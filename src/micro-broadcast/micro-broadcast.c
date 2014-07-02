#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void broadcast(int, int, int); 

int main(int argc, char *argv[])
{
	int tid, count, i, local; 
	omp_set_dynamic(0);
	const int N = atoi(argv[1]); // Iterations
	int iCPU = omp_get_num_procs();
	omp_set_num_threads(iCPU); // Threads

	for (i = 0; i <= N; i++)
		broadcast(count, tid, local);

	return 0;
}

void broadcast(int count, int tid, int local) {
#pragma omp parallel private(tid, local) shared(count)
	{
#pragma omp master
		{
			local = 1;
			count = local; 
		} 
#pragma omp barrier

#pragma omp critical
		{
			tid = omp_get_thread_num();
			local = count; 
			// printf("Count is: %d - Thread ID: %d\n", local, tid);
		}
	}
}