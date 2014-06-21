#include <stdio.h>
#include <time.h>
#include <stdlib.h>

long gcd(long, long);
const int range = 1000000;

int main() {
	srand(time(NULL));
	long first, second, hcf, lcm;

	first = rand() % range;
	second = rand() % range;

	hcf = gcd(first, second);
	lcm = (first * second)/hcf;

	printf("The greatest common divisor of %ld and %ld = %ld\n", first, second, hcf);
	printf("The least common multiple of %ld and %ld = %ld\n", first, second, lcm);

	return 0;
}

long gcd(long first, long second) {
	if (second == 0)
		return first;
	else 
		return gcd(second, first % second);
}