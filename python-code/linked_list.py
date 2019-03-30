from random import randint

class Node:
    def __init__(self, dataval=None):
        self.dataval = dataval
        self.nextval = None

class SLinkedList:
    def __init__(self, node):
        self.headval = node
        self.tail = node
        self.sorted = False

    def add(self, node):
        self.tail.nextval = node
        self.tail = node

    def find(self, value):
        current = self.headval
        count = 0
        while current is not None:
            if current.dataval is value:
                return count
            count = count + 1
            current = current.nextval
        return None

    def bubble_sort(self):
        current_node = self.headval
        count = 0
        while current_node is not None:
            print(f'Count: {count}')
            current = self.headval
            while current is not None and current.nextval is not None:
                temp = current.nextval.dataval
                if temp < current.dataval:
                    # print(f'Swap: {temp} for {current.dataval}')
                    current.nextval.dataval = current.dataval
                    current.dataval = temp
                current = current.nextval
            current_node = current_node.nextval
            count = count + 1
        self.sorted = True

    def count(self):
        current = self.headval
        count = 0
        while current is not None:
            count = count + 1
            current = current.nextval
        return count
    
    def print_nodes(self):
        current = self.headval
        while current is not None:
            print(current.dataval)
            current = current.nextval

num_of_nums = 1000
nums = [randint(0,num_of_nums) for _ in range(num_of_nums)]

list1 = SLinkedList(Node(nums[0]))

for num in nums[1:]:
    list1.add(Node(num))

print('sorted')
list1.bubble_sort()
list1.print_nodes()

current = list1.headval