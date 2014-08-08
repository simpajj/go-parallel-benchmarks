#include <stdio.h>
#include <time.h>
#include <stdlib.h>

#define numberOfElements 100

void bubble_sort(long [], long);

int main(int argc, char *argv[]) {
	srand(time(NULL));
	const int N = atoi(argv[1]);
	long array[numberOfElements];
	int i, c;

	for (i = 0; i <= N; i++) {
		for (c = 0; c < numberOfElements; c++) 
			array[c] = rand() % numberOfElements;

		bubble_sort(array, numberOfElements);
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