# Leetcode 150 DSA sheet - Golang

## Index

1. [Merge Sorted Array](#01-merge-sorted-array)
2. [Remove Element](#02-remove-element)
3. [Problem Name Placeholder](#03-problem-name-placeholder)

## Solutions

### 01. Merge Sorted Array

[Problem Link](https://leetcode.com/problems/remove-element/)

**2 pointer approach:** Start 2 pointers from both ends of each num1 and num2 array. Add a new pointer from the end of new array to be written(extended end of num1 here) that will write the result. Based on comparision of elements from 2 sorted array from the end start filling the auxilary array with biggest element from it's end. Decrement the pointer until either of the array gets exhausted out of elements to be compared(num2 in this case).

[My Submission](https://leetcode.com/submissions/detail/1461678292/)

[solutions/01-Merge-Sorted-Array](https://github.com/singhxayush/dsa-150-sheet-in-golang/blob/master/solutions/01-Merge-Sorted-Array.go)

**Dry run**

```txt
num1[] = {1, 2, 3, 4}
num2[] = {0, 1, 3, 8}

1 2 3 .4 - - - |-
0 1 3 .8

4 < 8
1 2 3 .4 - - |- 8
0 1 .3 8

4 > 3
1 2 .3 4 - |- 4 8
0 1 .3

3 == 3
1 .2 3 4 |- 3 4 8
0 1 .3

2 < 3
1 .2 3 |4 3 3 4 8
0 .1

2 > 1
.1 2 |3 2 3 3 4 8
0 .1

1 == 1
1 |2 1 2 3 3 4 8
0 .1

no comparision here - simply write
|1 1 1 2 3 3 4 8
.0

no comparision here - simply write
0 1 1 2 3 3 4 8
num2 empty
```

---

### 02. Remove Element

[Problem Link](https://leetcode.com/problems/merge-sorted-array/description/)

**Constructive Approach:** Iterate through all the elements and maintain a counter (also as an index) for all the elements that will count all the elements that are not equal to the given value. Now write back all the non equal elements at the index of counter. The Value written will always be ahead of all the non equal values previously written and 

[My Submission](https://leetcode.com/submissions/detail/1461662052/)

[solutions/02-Remove-Element](https://github.com/singhxayush/dsa-150-sheet-in-golang/blob/master/solutions/02-Remove-Element.go)

**Dry run**

```txt
arr[] = {1, 2, 2, 4, 3, 2, 7}
val = 2

cnt = 0

1 2 2 4 3 2 7
^ 1 != val, cnt=0, arr[cnt++] = 1

1 2 2 4 3 2 7
  ^ 2 == val, do nothing

1 2 2 4 3 2 7
    ^ 2 == val, do nothing

1 4 2 4 3 2 7
      ^ 4 != val, cnt = 1, arr[cnt++] = 4

1 4 3 4 3 2 7
        ^ 3 != val, cnt = 2, arr[cnt++] = 3

1 4 3 4 3 2 7
          ^ 2 == val, do nothing

1 4 3 7 3 2 7
            ^ 7 != val, cnt = 3, arr[cnt] = 7
```