#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void multiplex(int, int, int, int, int);

int main(int argc, char *argv[])
{
	int i, tid;
	int local, sum = 0; 
	omp_set_dynamic(0);
	const int N = atoi(argv[1]); // Iterations
	int iCPU = omp_get_num_procs();
	omp_set_num_threads(iCPU); // Threads

	for (i = 0; i <= N; i++)
		multiplex(tid, local, sum, iCPU, i);

	return 0;
}

void multiplex(int tid, int local, int sum, int iCPU, int i) {
#pragma omp parallel shared(sum) private(local)
	{
		tid = omp_get_thread_num();

#pragma omp for 
		for (i = 0; i <= iCPU; i++)
			local = tid + 1;

#pragma omp critical
		{
			sum += local;
		}

	}
	// printf("Sum is: %d\n", sum);
}