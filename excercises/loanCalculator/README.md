Loan Calculator
Description
A net loan calculator is to be created at a vehicle showroom. The applicant details are stored in a class named Source which contains the following details

A string s denoting the name of applicant,

An integer n denoting type of vehicle; where 0 represents a car and 1 represents a bike.

An integer x denoting the net price of the vehicle,

An integer y denoting the duration for which loan is provided.



Given the details of n such applicants as user input, write a function that lists the total amount payable for all candidates along with their names.

The interest charged on bike loans is 15 % per annum and interest on car loans is 20% per annum for total amount payable below 50 lakhs and 10 % per annum for bikes and 15% per annum for cars if the total amount payable is above 50 lakhs.



For a vehicle type that doesn’t exist, output “Invalid vehicle type specified”, the corresponding name of the applicant and total amount as -1.



Note 1- Interest is calculated using the formula shown below:

Interest = (Net price of vehicle*Rate of interest*Duration in years)/100



Note 2- The higher rate of interest needs to be used for calculating the slab for total amount payable by applicant. 



Input Format
The first line contains an integer n denoting the number of applicants for whom the total amount payable needs to be calculated.

The next subsequent i lines contain the details of name of applicant, type of vehicle, net price and duration for which loan is provided for the ith applicant in the specified format.

Output format:
An array containing the names of the n applicants and the total amount payable by each candidate.



Example

Sample Input 1:

3

Arvind

0

150000

3

Sumit

0

200000

3

Jacky

1

2000000

4

 

Sample Output 1:

Arvind

2,40,000

Sumit

3,20,000

Jacky

32,00,000

Explanation:

As per the formula given above 150000+Interest = 2,40,000 for Arvind,

200000+Interest = 3,20,000 for Sumit and 2000000+Interest= 32,00,000 for Jacky

 



Sample Input 2:

2

James

0

20000

1

Jane

0

30000

1

 Sample Output 2:

James

24000

Jane

36000

 



Sample Input 3:

3

Jane

0

100000

3

Amit

0

50000

3

Sumit

0

20000

2

Sample Output 3:

Jane

1,60000

Amit

80,000

Sumit

28,000