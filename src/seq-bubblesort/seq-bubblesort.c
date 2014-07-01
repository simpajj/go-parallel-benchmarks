#include <stdio.h>
#include <time.h>
#include <stdlib.h>

void bubble_sort(long [], long);
const int numberOfElements = 100;

int main(int argc, char *argv[]) {
	srand(time(NULL));
	long array[numberOfElements];
	int i, c;
	const int N = atoi(argv[1]);

	for (i = 0; i <= N; i++) {
		for (c = 0; c < numberOfElements; c++) 
			array[c] = rand() % numberOfElements;

		bubble_sort(array, numberOfElements);
		// printf("Sorted list:\n");

		// for(c = 0; c < numberOfElements; c++) 
		// 	printf("%ld\n", array[c]);
	}

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