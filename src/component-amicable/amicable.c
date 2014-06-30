#include <stdio.h>
#include <stdlib.h>
#include <omp.h>

unsigned int sum_divisors(unsigned int n) {
	unsigned int s = 0, i = 1;
	while(i < n) {
		if(n % i == 0) s = s + i;
		i++;
	}
	return s;
}

int main(int argc, char const *argv[])
{
	unsigned int a, b, result = 0;
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1])); // threads
	const int N = atoi(argv[2]); // iterations
	const int numbers = atoi(argv[3]);

#pragma omp parallel for reduction(+:result) private(b)
		for(a = 0; a < numbers; a++) {
			b = sum_divisors(a);
			if(b > a && sum_divisors(b) == a) result = result + a + b; 
		}

	return 0;
}