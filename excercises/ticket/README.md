Buy Ticket
Description


Your purse has N distinct denominations for buying things. You are given the denominations in the array A. You have an infinite amount of each denomination available with you because you are very rich. You went to a theatre to buy a movie ticket. You don't know the exact price of the ticket, just that it can cost anything from 1 to X inclusive. Since you hate change, you want to make sure that you can pay for the ticket without having to haggle for change. How many ticket prices from 1 to X can you buy without having to haggle for change?

 

Function Description:

Complete the ticket function in the editor below. It has the following parameter(s):

﻿



Constraints:

1<=N<=10^3

1<=A[i]<=10^5

1<=X<=10^9

 

Input Format:

The first line contains an integer, N, denoting the number of elements in A.

Each line i of the N subsequent lines (where 0 ≤ i < N) contains an integer describing A[i].  

The next line contains an integer, X, denoting the upper limit of ticket prices.

 

Output Format:

An integer denoting the number of tickets you can buy.

 

Example

Sample Input 1:

2

3

5

10

 

Sample Output 1:

6

 

Sample Output 1 Explanation:

You can't buy ticket prices of price = {1,2,4,7}. Note you can buy a ticket price of 6 with 2 notes of 3 each, and a ticket price of 8 with one note of 3 and 5 each.

 

Sample Input 2:

3

1

100

1000

1000

 

Sample Output 2:

1000

 

Sample Output 2 Explanation:

Since you have a note of value 1, you can buy any ticket price with using just multiple notes of value 1.

Execution Time Limit