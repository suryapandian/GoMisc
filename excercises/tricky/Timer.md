### Problem statement
Create a clock application that will print the following values at the following intervals to console:
- "tick" every second
- "tock" every minute
- "bong" every hour
Only one value should be printed in a given second, i.e. when printing "bong" on the hour, the "tick" and "tock" values should not be printed. Similarly while printing "tock" , "tick" should not be printed.
A mechanism should exist for the user to alter any of the printed values while the program is running, i.e. after the clock has run for 10 minutes I should, without stopping the program, be able to change it so that it stops printing "tick" every second and starts printing "quack" instead.
Also, provide a mechanism to input for how long the clock needs to run from the user.
Please provide appropriate test coverage.Note:
1. Please use Golang as the development language (ver 1.7 and above)
2. Try to use goroutine, waitGroup, sync package and channel wherever required
Also provide appropriate steps in a readme file needed to execute the program.