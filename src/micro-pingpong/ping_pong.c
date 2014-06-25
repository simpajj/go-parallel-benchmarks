#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	int numthreads, 
			tid,
			i, 
			var = 0; 
	omp_set_dynamic(0);
	omp_set_num_threads(2);

#pragma omp parallel shared(var)
	{
#pragma omp master
		{
			tid = omp_get_thread_num();
			var = 1;
		}

#pragma omp critical
		{
			tid = omp_get_thread_num();
			if (tid == 1) {
				var = 2;
			}
		}

#pragma omp flush(var)
#pragma omp barrier
#pragma omp master
		{
			tid = omp_get_thread_num();
			printf("Tid: %d - value: %d\n", tid, var);
		}
	}

	return 0;
}