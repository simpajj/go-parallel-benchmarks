#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

void ping(int);

int main (int argc, char *argv[]) {
	int tid, i;
	omp_set_dynamic(0);
	omp_set_num_threads(atoi(argv[1]));
	const int N = atoi(argv[2]);

	for (i = 0; i <= N; i++)
		ping(tid);

}

void ping(int tid) {
#pragma omp parallel private(tid)
  {

	  tid = omp_get_thread_num();
	  printf("Ping from thread #%d\n", tid);

  }
}