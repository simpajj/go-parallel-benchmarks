#include <stdio.h>
#include <time.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	srand(time(NULL));
	const int N = atoi(argv[1]);
	const int numberOfElements = atoi(argv[2]);
	int i, c, first, last, middle, search = 13, array[numberOfElements];

	for (i = 0; i <= N; i++) {
		for (c = 0; c < numberOfElements; c++)
			array[c] = rand() % 101;

		first = 0;
		last = numberOfElements - 1;
		middle = (first + last)/2;

		while (first <= last) {
			if (array[middle] < search)
				first = middle + 1;
			else if (array[middle] == search) {
				return 0;
			}
			else
				last = middle - 1;

			middle = (first + last)/2;
		}
		if (first > last)
			return 0;
	}
	return 0;
}