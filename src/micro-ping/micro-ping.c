#include <omp.h>
#include <stdio.h>
#include <stdlib.h>

int main (int argc, char *argv[]) 
{
int nthreads, tid;
omp_set_dynamic(0);
omp_set_num_threads(atoi(argv[1]));

#pragma omp parallel private(nthreads, tid)
  {

  tid = omp_get_thread_num();
  printf("Ping from thread #%d\n", tid);

  }

}