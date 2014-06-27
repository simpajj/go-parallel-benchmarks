#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void broadcast(int, int); 

int main(int argc, char *argv[])
{
	int tid, count, i; 
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1]));
	const int N = atoi(argv[2]);

	for (i = 0; i <= N; i++)
		broadcast(count, tid);

	return 0;
}

void broadcast(int count, int tid) {
#pragma omp master
	{
		count = 1; 
	} 

#pragma omp parallel private(tid)
	{
		tid = omp_get_thread_num();
		printf("Count is: %d - Thread ID: %d\n", count, tid);
	}
}