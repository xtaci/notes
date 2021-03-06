1.1 Implement an algorithm to determine if a string has all unique characters. What if you can not use additional data structures?
_
_________________________________________________________________pg 95
1.2 Write code to reverse a C-Style String. (C-String means that “abcd” is represented as five characters, including the null character.)
_
_________________________________________________________________pg 96
1.3 Design an algorithm and write code to remove the duplicate characters in a string without using any additional buffer. NOTE: One or two additional variables are fine. An extra copy of the array is not.
FOLLOW UP
Write the test cases for this method.
_
_________________________________________________________________pg 97
1.4 Write a method to decide if two strings are anagrams or not.
_
_________________________________________________________________pg 99
1.5 Write a method to replace all spaces in a string with ‘%20’.
_
________________________________________________________________pg 100
1.6 Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
_
________________________________________________________________pg 101
1.7 Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column is set to 0.
_
________________________________________________________________pg 102
1.8 Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring (i.e., “waterbottle” is a rotation of “erbottlewat”).
_
________________________________________________________________pg 103








2.1 Write code to remove duplicates from an unsorted linked list.
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
_
________________________________________________________________pg 105
2.2 Implement an algorithm to find the nth to last element of a singly linked list.
_
________________________________________________________________pg 106
2.3 Implement an algorithm to delete a node in the middle of a single linked list, given only access to that node.
EXAMPLE
Input: the node ‘c’ from the linked list a->b->c->d->e
Result: nothing is returned, but the new linked list looks like a->b->d->e
_
________________________________________________________________pg 107
2.4 You have two numbers represented by a linked list, where each node contains a single digit. The digits are stored in reverse order, such that the 1’s digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
Output: 8 -> 0 -> 8
_
________________________________________________________________pg 108
2.5 Given a circular linked list, implement an algorithm which returns node at the beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node’s next pointer points to an earlier node, so as to make a loop in the linked list.
EXAMPLE
input: A -> B -> C -> D -> E -> C [the same C as earlier]
output: C
_
________________________________________________________________pg 109















3.1 Describe how you could use a single array to implement three stacks.
_
________________________________________________________________pg 111
3.2 How would you design a stack which, in addition to push and pop, also has a function min which returns the minimum element? Push, pop and min should all operate in O(1) time.
_
________________________________________________________________pg 113
3.3 Imagine a (literal) stack of plates. If the stack gets too high, it might topple. Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be composed of several stacks, and should create a new stack once the previous one exceeds capacity. SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack).
FOLLOW UP
Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.
_
________________________________________________________________pg 115
3.4 In the classic problem of the Towers of Hanoi, you have 3 rods and N disks of different sizes which can slide onto any tower. The puzzle starts with disks sorted in ascending order of size from top to bottom (e.g., each disk sits on top of an even larger one). You have the following constraints:
(A) Only one disk can be moved at a time.
(B) A disk is slid off the top of one rod onto the next rod.
(C) A disk can only be placed on top of a larger disk.
Write a program to move the disks from the first rod to the last using Stacks.
_
________________________________________________________________pg 118
3.5 Implement a MyQueue class which implements a queue using two stacks.
_
________________________________________________________________pg 120
3.6 Write a program to sort a stack in ascending order. You should not make any assumptions about how the stack is implemented. The following are the only functions that should be used to write this program: push | pop | peek | isEmpty.
























4.1 Implement a function to check if a tree is balanced. For the purposes of this question, a balanced tree is defined to be a tree such that no two leaf nodes differ in distance from the root by more than one.
_
________________________________________________________________pg 123
4.2 Given a directed graph, design an algorithm to find out whether there is a route between two nodes.
_
________________________________________________________________pg 124
4.3 Given a sorted (increasing order) array, write an algorithm to create a binary tree with minimal height.
_
________________________________________________________________pg 125
4.4 Given a binary search tree, design an algorithm which creates a linked list of all the nodes at each depth (i.e., if you have a tree with depth D, you’ll have D linked lists).
_
________________________________________________________________pg 126
4.5 Write an algorithm to find the ‘next’ node (i.e., in-order successor) of a given node in a binary search tree where each node has a link to its parent.
_
________________________________________________________________pg 127
4.6 Design an algorithm and write code to find the first common ancestor of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not necessarily a binary search tree.
_
________________________________________________________________pg 128
4.7 You have two very large binary trees: T1, with millions of nodes, and T2, with hundreds of nodes. Create an algorithm to decide if T2 is a subtree of T1.
_
________________________________________________________________pg 130
4.8 You are given a binary tree in which each node contains a value. Design an algorithm to print all paths which sum up to that value. Note that it can be any path in the tree - it does not have to start at the root.
_
________________________________________________________________



































5.1 You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a method to set all bits between i and j in N equal to M (e.g., M becomes a substring of N located at i and starting at j).
EXAMPLE:
Input: N = 10000000000, M = 10101, i = 2, j = 6
Output: N = 10001010100
_
________________________________________________________________pg 133
5.2 Given a (decimal - e.g. 3.72) number that is passed in as a string, print the binary representation. If the number can not be represented accurately in binary, print “ERROR”
_
________________________________________________________________pg 134
5.3 Given an integer, print the next smallest and next largest number that have the same number of 1 bits in their binary representation.
_
________________________________________________________________pg 135
5.4 Explain what the following code does: ((n & (n-1)) == 0).
_
________________________________________________________________pg 138
5.5 Write a function to determine the number of bits required to convert integer A to integer B.
Input: 31, 14
Output: 2
_
________________________________________________________________pg 139
5.6 Write a program to swap odd and even bits in an integer with as few instructions as possible (e.g., bit 0 and bit 1 are swapped, bit 2 and bit 3 are swapped, etc).
_
________________________________________________________________pg 140
5.7 An array A[1... n] contains all the integers from 0 to n except for one number which is missing. In this problem, we cannot access an entire integer in A with a single operation. The elements of A are represented in binary, and the only operation we can use to access them is “fetch the jth bit of A[i]”, which takes constant time. Write code to find the missing integer. Can you do it in O(n) time?






























































8.1 Write a method to generate the nth Fibonacci number.
_
________________________________________________________________pg 169
8.2 Imagine a robot sitting on the upper left hand corner of an NxN grid. The robot can only move in two directions: right and down. How many possible paths are there for the robot?
FOLLOW UP
Imagine certain squares are “off limits”, such that the robot can not step on them. Design an algorithm to get all possible paths for the robot.
_
________________________________________________________________pg 170
8.3 Write a method that returns all subsets of a set.
_
________________________________________________________________pg 171
8.4 Write a method to compute all permutations of a string.
_
________________________________________________________________pg 173
8.5 Implement an algorithm to print all valid (e.g., properly opened and closed) combinations of n-pairs of parentheses.
EXAMPLE:
input: 3 (e.g., 3 pairs of parentheses)
output: ()()(), ()(()), (())(), ((()))
_
________________________________________________________________pg 174
8.6 Implement the “paint fill” function that one might see on many image editing programs. That is, given a screen (represented by a 2 dimensional array of Colors), a point, and a new color, fill in the surrounding area until you hit a border of that color.’
_
________________________________________________________________pg 175
8.7 Given an infinite number of quarters (25 cents), dimes (10 cents), nickels (5 cents) and pennies (1 cent), write code to calculate the number of ways of representing n cents.
_
________________________________________________________________pg 176
8.8 Write an algorithm to print all ways of arranging eight queens on a chess board so that none of them share the same row, column or diagonal.



























































9.1 You are given two sorted arrays, A and B, and A has a large enough buffer at the end to hold B. Write a method to merge B into A in sorted order.
_
________________________________________________________________pg 179
9.2 Write a method to sort an array of strings so that all the anagrams are next to each other.
_
________________________________________________________________pg 180
9.3 Given a sorted array of n integers that has been rotated an unknown number of times, give an O(log n) algorithm that finds an element in the array. You may assume that the array was originally sorted in increasing order.
EXAMPLE:
Input: find 5 in array (15 16 19 20 25 1 3 4 5 7 10 14)
Output: 8 (the index of 5 in the array)
_
________________________________________________________________pg 181
9.4 If you have a 2 GB file with one string per line, which sorting algorithm would you use to sort the file and why?
_
________________________________________________________________pg 182
9.5 Given a sorted array of strings which is interspersed with empty strings, write a method to find the location of a given string.
Example: find “ball” in [“at”, “”, “”, “”, “ball”, “”, “”, “car”, “”, “”, “dad”, “”, “”] will return 4
Example: find “ballcar” in [“at”, “”, “”, “”, “”, “ball”, “car”, “”, “”, “dad”, “”, “”] will return -1
_
________________________________________________________________pg 183
9.6 Given a matrix in which each row and each column is sorted, write a method to find an element in it.
_
________________________________________________________________pg 184
9.7 A circus is designing a tower routine consisting of people standing atop one another’s shoulders. For practical and aesthetic reasons, each person must be both shorter and lighter than the person below him or her. Given the heights and weights of each person in the circus, write a method to compute the largest possible number of people in such a tower.
EXAMPLE:
Input (ht, wt): (65, 100) (70, 150) (56, 90) (75, 190) (60, 95) (68, 110)
Output: The longest tower is length 6 and includes from top to bottom: (56, 90) (60,95) (65,100) (68,110) (70,150) (75,190)
_
________________________________________________________________
