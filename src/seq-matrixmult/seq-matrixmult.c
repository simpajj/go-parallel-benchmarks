#include <stdio.h>
#include <stdlib.h>
#include <time.h>

const int range = 5;

int main(int argc, char *argv[]) {
	srand(time(NULL));
	enum { ROWS = 3, COLS = 3 };
	int i, c, d, k, sum = 0;
	int first[10][10], second[10][10], result[10][10];
	const int N = atoi(argv[1]);

	for (i = 0; i <= N; i++) {
		// Populate first matrix
		for (c = 0; c < ROWS; c++) 
			for (d = 0; d < COLS; d++)
				first[c][d] = rand() % range;

		// Populate second matrix
		for (c = 0; c < ROWS; c++)
			for (d = 0; d < COLS; d++)
				second[c][d] = rand() % range;

		// Multiply
		for (c = 0; c < ROWS; c++) {
			for (d = 0; d < COLS; d++) {
				for (k = 0; k < ROWS; k++) {
					sum = sum + first[c][k] * second[k][d];
				}

				result[c][d] = sum;
				sum = 0;
			}
		}

		// printf("Product: \n");

		// for (c = 0; c < ROWS; c++) {
		// 	for (d = 0; d < COLS; d++) {
		// 		printf("%d\t", result[c][d]);
		// 	}
		// 	printf("\n");
		// } 
	}
	return 0;
}