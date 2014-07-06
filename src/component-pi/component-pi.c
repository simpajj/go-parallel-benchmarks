#include <stdio.h>
#include <stdlib.h>
#include <omp.h>

const int NITER = 100000;

void montecarlo(double, double, double, double, int, int);

int main(int argc, char* argv[])
{
  double x, y, z, pi;
  int i, count = 0;
  omp_set_dynamic(0);
  int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU);
  const int N = atoi(argv[1]);

  for(i = 0; i <= N; i++) {
      montecarlo(x, y, z, pi, count, iCPU);
  }

  return 0;
}

void montecarlo(double x, double y, double z, double pi, int count, int iCPU) {
  int i; 
#pragma omp parallel firstprivate(x, y, z, i) reduction(+:count) num_threads(iCPU)
  {
    srand48((int)time(NULL) ^ omp_get_thread_num());    //Give random() a seed value
    for (i = 0; i < NITER; ++i)
    {
      x = (double)drand48(); // Random x coordinate
      y = (double)drand48(); // Random y coordinate
      z = ((x*x)+(y*y));
      if (z<=1)
      {
        ++count; // Valid random point if number is inside unit circle  
      }       
    }
  } 
  pi = ((double)count/(double)(NITER*iCPU))*4.0;
  printf("Pi: %f\n", pi);
}