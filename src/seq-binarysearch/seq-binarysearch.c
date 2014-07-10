#include <stdio.h>
#include <time.h>
#include <stdlib.h>

#define numberOfElements 100

int binarysearch(int []);

int binarysearch(int array[]) {
	int c, 
			first = 0, 
			last = numberOfElements - 1, 
			middle = (first + last)/2, 
			search = 13;

	for (c = 0; c < numberOfElements; c++)
		array[c] = rand() % 101;

	while (first <= last) {
		if (array[middle] < search)
			first = middle + 1;
		else if (array[middle] == search) {
			search = array[middle];
			return search;
		}
		else
			last = middle - 1;

		middle = (first + last)/2;
	}
	if (first > last)
		return 0;
	return 0;
}

int main(int argc, char *argv[]) {
	srand(time(NULL));
	const int N = atoi(argv[1]);
	int i, c, array[numberOfElements];

	for (i = 0; i <= N; i++)
		binarysearch(array);
}