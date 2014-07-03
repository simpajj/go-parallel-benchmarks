#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
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
	unsigned int a = 284, b = 220, i;
	bool result = false;
	omp_set_dynamic(0);
	const int N = atoi(argv[1]); // iterations
  int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU)

#pragma omp parallel for reduction(+:result) private(b)
		for(i = 0; i < N; i++) {
			b = sum_divisors(a);
			if(sum_divisors(b) == a) result = true; 
		}

	printf("Final result = %d\n", result);
	return 0;
}