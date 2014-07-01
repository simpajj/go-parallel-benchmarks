#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void pingpong(int, int);

int main(int argc, char *argv[])
{
	int tid,
			i, 
			var = 0; 
	omp_set_dynamic(0);
	omp_set_num_threads(2);
	const int N = atoi(argv[1]);

	for (i = 0; i <= N; i++)
		pingpong(tid, var);

	return 0;
}

void pingpong(int tid, int var) {
#pragma omp parallel shared(var)
	{
#pragma omp master
		{
			tid = omp_get_thread_num();
			var = 1;
		}
#pragma omp flush(var)
#pragma omp barrier

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
			// printf("Tid: %d - value: %d\n", tid, var);
		}
	}
}