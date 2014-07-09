#include <stdio.h>
#include <omp.h>
#include <stdio.h>
#include <time.h>

#define RANGE 10000000

void qsort(int, int);
void qsort_parallel(int, int);

int a[RANGE];

int main(int argc, char const *argv[])
{
	int i, j;
	const int N = atoi(argv[1]); // iterations
  int iCPU = omp_get_num_procs();
  omp_set_num_threads(iCPU);
	omp_set_dynamic(0);
	srand(time(NULL));

	for(j = 0; j <= N; j++) {
		for(i = 0; i < RANGE; i++)
			a[i] = rand() % RANGE;

#pragma omp parallel	
		#pragma omp single
		qsort_parallel(0, RANGE-1);
	}

	return 0;
}

void qsort(int l, int r){
	if(r > l) {
		int pivot = a[r], tmp;
		int less = l-1, more;
		for(more = l; more <= r; more++) {
			if(a[more] <= pivot) {
				less++;
				 tmp = a[less];
				a[less] = a[more];
				a[more] = tmp;
			}
		}
		qsort(l, less-1);
		qsort(less+1, r);
	}
}

void qsort_parallel(int l, int r){
	int i;
	if(r > l) {
		int pivot = a[r], tmp;
		int less = l-1, more;
		for(more = l; more <= r; more++) {
			if(a[more] <= pivot) {
				less++;
				tmp = a[less];
				a[less] = a[more];
				a[more] = tmp;
			}
		}
		if((r - l) < 1000) {
			qsort(l, less-1);
			qsort(less+1, r);
		} else {
			#pragma omp task
			qsort_parallel(l, less-1);
			#pragma omp task
			qsort_parallel(less+1, r);
			#pragma omp taskwait
		}
	}
}