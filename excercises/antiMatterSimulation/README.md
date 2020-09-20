Antimatter Simulation
Description
A Physics simulator simulates the interactions between hypothetical antimatter types. If any piece of matter in the simulation combines with another, it results in a piece of matter whose size is the difference in sizes of the original two pieces. Combination is the only type of interaction possible. Any two pieces in the simulation can interact in this way, in any order.



Given the sizes of all pieces of matter that are being simulated, find the smallest possible piece of matter that can be created by the interaction of the pieces.

 

Note: Size of particles is represented by an integer value.



﻿Constraints﻿

●    1 ≤ n ≤ 10^3

●    1 ≤ A[i] ≤ 10^9

 

Input Format

The first line contains an integer, n, denoting the number of pieces being simulated.
Each line i of the n subsequent lines (where 0 ≤ i < n) contains an integer that represents the size of the piece of matter i.


Output Format

Output an INTEGER denoting the size of the smallest possible piece of matter that can be created .


Sample Case 0

Sample Input

3

30

10

8

Sample Output

2

Explanation

The smallest possible piece that can result from an interaction is formed if the piece of size 10 interacts (and combines) with the one of size 8 and results in a piece of size 2.

 

Sample Case 1

Sample Input

4

1

2

4

8

Sample Output

1

Explanation:

The smallest possible piece that can result from an interaction is formed if the piece of size 2 interacts (and combines) with the piece of size 1 and results in a piece of size 1.