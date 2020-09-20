Function Description:

Complete the minSwaps function in the editor below. It has the following parameter(s):

﻿

Constraints:

1<=N<=2*10^5
1<=A[i]<=2*10^5
 

Input Format:

The first line contains an integer, N, denoting the number of elements in A.
Each line i of the N subsequent lines (where 0 ≤ i < N) contains an integer describing A[i]. 
 

Output Format:

An integer denoting the minimum number of swaps required.
 

Example

Sample Input 1:

6

2

8

4

1

1

5

Sample Output 1:

3

Sample Output 1 Explanation:

The number of swaps required for insertion sort in the original array is: 8.

The minimum swaps required for insertion sort is 3 when array is rotated to [1,1,5,2,8,4]



Sample Input 2:

4

4

3

2

1

Sample Output 2:

2

Sample Output 2 Explanation:

The rotation is: [2,1,4,3] when the minimum number of swaps required for insertion sort is 2.

Execution Time Limit
Default.