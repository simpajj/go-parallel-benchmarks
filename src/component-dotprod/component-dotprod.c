#include <stdio.h>
#include <stdlib.h>
#include <omp.h>

int main(int argc, char const *argv[])
{
  int c, i, n = 100, chunk = 10;
  omp_set_dynamic(0);
  const int N = atoi(argv[1]); // iterations
  int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU);
  float a[100], b[100], result;
  
  for (c = 0; c < N; c++) {
    result = 0.0;
    for (i=0; i < n; i++) {
        a[i] = i * 1.0;
        b[i] = i * 2.0;
    }
    #pragma omp parallel for      \
        default(shared) private(i)  \
        schedule(static,chunk)      \
        reduction(+:result)

      for (i=0; i < n; i++)
          result += (a[i] * b[i]);

    // printf("Final result = %f\n", result);
  }
}