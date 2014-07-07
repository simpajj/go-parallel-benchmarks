#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void ping(int, int, int);

int main (int argc, char *argv[]) {
	int tid, i, var, local;
	omp_set_dynamic(0);
	const int N = atoi(argv[1]); // Iterations
	int iCPU = omp_get_num_procs();
	omp_set_num_threads(iCPU); // Threads

	for (i = 0; i <= N; i++) 
		ping(tid, var, local);

}

void ping(int tid, int var, int local) {
#pragma omp parallel private(tid, local) shared(var)
  {
#pragma omp master
  	{
  		var = 1; 
  	}
#pragma omp barrier
#pragma omp critical
  	{
  		tid = omp_get_thread_num();
  		local = var;
	  	// printf("Ping from thread #%d: %d\n", tid, local);
  	}
  }
}