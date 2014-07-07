#include <stdio.h>
#include <time.h>
#include <stdlib.h>

long gcd(long, long);
const int range = 1000;

int main(int argc, char *argv[]) {
	srand(time(NULL));
	int i, first, second, hcf, lcm;
	const int N = atoi(argv[1]);

	for (i = 0; i <= N; i++) {
		first = rand() % range;
		second = rand() % range;

		hcf = gcd(first, second);
		lcm = (first * second)/hcf;
	}
	return 0;
}

long gcd(long first, long second) {
	if (second == 0)
		return first;
	else 
		return gcd(second, first % second);
}