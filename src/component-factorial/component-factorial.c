#include <stdio.h>
#include <stdlib.h>
#include <omp.h>

int factorial(int);

#define RANGE 10

int main(int argc, char const *argv[])
{
	int i;
	const int N = atoi(argv[1]);
	int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU);
	omp_set_dynamic(0);
	srand(time(NULL));
	int a[RANGE];

	for(i=0; i < RANGE; i++)
		a[i] = rand() % RANGE - 1;

	for (i = 0; i <= N; i++)
		for(i=0; i < RANGE; i++)
			factorial(a[i]);
	
	return 0;
}

int factorial(int num)
 {
   int fac = 1, n;
   #pragma omp parallel for reduction(*:fac)
   for(n = 2; n <= num; ++n)
     fac *= n;
   // printf("%d\n", fac);
   return fac;
 }