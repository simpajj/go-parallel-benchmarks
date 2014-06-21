#include <stdio.h>
#include <time.h>
#include <stdlib.h>

const int numberOfElements = 100;

int main() {
	srand(time(NULL));
	int c, first, last, middle, search, array[100];
	search = 13;

	for (c = 0; c < numberOfElements; c++)
		array[c] = rand() % 101;

	first = 0;
	last = numberOfElements - 1;
	middle = (first + last)/2;

	while (first <= last) {
		if (array[middle] < search)
			first = middle + 1;
		else if (array[middle] == search) {
			printf("%d found at location %d \n", search, middle + 1);
			break;
		}
		else
			last = middle - 1;

		middle = (first + last)/2;
	}
	if (first > last)
		printf("The number is not in the list!");

	return 0;
}