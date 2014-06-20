#include <stdio.h>
#include <time.h>
#include <stdlib.h>

void bubble_sort(long [], long);
const int numberOfElements = 100;

int main() {
	srand(time(NULL));
	long array[numberOfElements], n, c, d, swap;
	n = 63;

	for (c = 0; c < numberOfElements; c++) 
		array[c] = rand() % 101;

	bubble_sort(array, n);
	printf("Sorted list:\n");

	for(c = 0; c < n; c++) 
		printf("%ld\n", array[c]);

	return 0;
}

void bubble_sort(long list[], long n) {
	long c, d, t;

	for(c = 0; c < (n - 1); c++) {
		for(d = 0; d < (n - 1); d++) {
			if (list[d] > list[d+1]) {
				t = list[d];
				list[d] = list[d+1];
				list[d+1] = t; 
			}
		}
	}
}