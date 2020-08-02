Insertion Sort Rotation
Description
You are given an array A(0-indexed) having N elements. Consider the following pseudo-code of insertion sort algorithm for sorting it:

for (int i = 1; i < N; i = i + 1) {

   int j = i;

   while (j > 0 && a[j] < a[j - 1]) {

       swap(a[j], a[j - 1]);

       j = j - 1;

   }

}

Since, sorting the array takes too many swaps, you are allowed to rotate the array any number of times. Rotating an array involves removing the last element and putting it at the start of the array. Calculate the minimum number of swaps required in the above algorithm after rotating the array as many times as needed modulo 10^9+7.