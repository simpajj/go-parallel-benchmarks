#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	int numthreads = omp_get_num_threads(); 
	int i, tid;
	int local, sum = 0; 
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1]));

#pragma omp parallel shared(sum) private(local)
	{
		tid = omp_get_thread_num();

#pragma omp for 
		for (i = 0; i <= numthreads; i++)
			local = tid + 1;

#pragma omp critical
		{
			sum += local;
		}

	}
	printf("Sum is: %d\n", sum);

	return 0;
}