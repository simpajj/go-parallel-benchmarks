#include <stdio.h>
#include <math.h>
#include <omp.h>

#define MAX_N 100000000
#define MAX_THREADS 100
#define RANGE 50000

int prime[MAX_N+1],  
    nextbase;  // next sieve multiplier to be used

// "crosses out" all odd multiples of k, from k*k on, k odd
void crossout(int k)
{  
  int i;

   for (i = k; i*k <= RANGE; i += 2)  {
      prime[i*k] = 0;
   }
}

int worker()  
{  
  int lim, base;
  lim = sqrt(RANGE);  
  do  {
    #pragma omp critical
    {  
      base = nextbase += 2; 
    }
    if (base <= lim)  {
      if (prime[base])  
        crossout(base);
    }
    else break; 
  } while (1);
}

int main(int argc, char **argv)
{  
  const int N = atoi(argv[1]); // Iterations
  int iCPU = omp_get_num_procs();
  omp_set_dynamic(0);
  omp_set_num_threads(iCPU); // Threads
  int nprimes,  // number of primes found 
      i,
      j;
  for (j = 0; j <= N; j++) {
    for (i = 2; i <= RANGE; i++) prime[i] = 1;
    for (i = 2; 2*i <= RANGE; i++) prime[2*i] = 0;
    nextbase = 1;

    #pragma omp parallel 
    {  
      worker(); 
    }
    nprimes = 0;
    for (i = 2; i <= RANGE; i++)  
      if (prime[i]) nprimes++;
    // printf("number of primes found: %d\n", nprimes);
  }
}
