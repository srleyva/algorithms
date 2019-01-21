package algorithms

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewNode(t *testing.T) {
	newNode := NewNode(10)
	if newNode.Value != 10 {
		t.Errorf("%d expected but %d returned", 10, newNode.Value)
	}
}

func TestNodeInsert(t *testing.T) {
	newNode := NewNode(10)

	newNode.Insert(20)
	newNode.Insert(5)
	newNode.Insert(8)
	newNode.Insert(30)
	newNode.Insert(15)

	if newNode.Value != 10 {
		t.Errorf("Root node: %d expected but %d returned", 10, newNode.Value)
	}

	if newNode.Left.Value != 5 {
		t.Errorf("Wrong node position for root node left expected 5 got %d", newNode.Left.Value)
	}

	if newNode.Left.Right.Value != 8 {
		t.Errorf("Wrong node position for second node left expected 8 got %d", newNode.Left.Left.Value)
	}

	if newNode.Right.Value != 20 {
		t.Errorf("Wrong node position for root node right expected 20 got %d", newNode.Right.Value)
	}

	if newNode.Right.Left.Value != 15 {
		t.Errorf("Wrong node position for second node left expected 25 got %d", newNode.Right.Value)
	}

}

func TestNodePrintOrder(t *testing.T) {
	newNode := NewNode(10)
	newNode.Insert(20)
	newNode.Insert(5)
	newNode.Insert(8)
	newNode.Insert(30)
	newNode.Insert(15)

	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	newNode.OrderPrint()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	expected := "5 8 10 15 20 30 "
	if out != expected {
		t.Errorf("Expected: %s Actual: %s", expected, out)
	}
}
