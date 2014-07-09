#include <stdio.h>
#include <stdlib.h>
#include <omp.h>

int factorial(int);

#define RANGE 1000

int main(int argc, char const *argv[])
{
	int i, j;
	const int N = atoi(argv[1]);
	int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU);
	omp_set_dynamic(0);

	for (i = 0; i <= N; i++)
		for(j = 0; j < RANGE; j++)
			factorial(j);
	
	return 0;
}

int factorial(int num)
 {
   int fac = 1, n;
   #pragma omp parallel for reduction(*:fac)
   for(n = 2; n <= num; ++n)
     fac *= n;
   return fac;
 }