#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	int nthreads, tid, count; 
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1]));

#pragma omp master
	{
		count = 1; 
	} 

#pragma omp parallel private(nthreads, tid)
	{
		tid = omp_get_thread_num();
		printf("Count is: %d - Thread ID: %d\n", count, tid);
	}

	return 0;
}