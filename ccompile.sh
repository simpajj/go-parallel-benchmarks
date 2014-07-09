gcc src/seq-binarysearch/seq-binarysearch.c -O3 bin/seq-binarysearch-c
gcc src/seq-bubblesort/seq-bubblesort.c -O3 bin/seq-bubblesort-c
gcc src/seq-gcd-lcm/seq-gcd-lcm.c -O3 bin/seq-gcd-lcm-c
gcc src/seq-matrixmult/seq-matrixmult.c -O3 bin/seq-matrixmult-c
gcc -fopenmp src/micro-broadcast/micro-broadcast.c -O3 bin/micro-broadcast-c
gcc -fopenmp src/micro-multiplex/micro-multiplex.c -O3 bin/micro-multiplex-c
gcc -fopenmp src/micro-ping/micro-ping.c -O3 bin/micro-ping-c
gcc -fopenmp src/micro-pingpong/micro-pingpong.c -O3 bin/micro-pingpong-c
gcc -fopenmp src/component-dotprod/component-dotprod.c -O3 bin/component-dotprod-c
gcc -fopenmp src/component-factorial/component-factorial.c -O3 bin/component-factorial-c
gcc -fopenmp src/component-pi/component-pi.c -O3 bin/component-pi-c
gcc -fopenmp src/component-primefilter/component-primefilter.c -O3 bin/component-primefilter-c
gcc -fopenmp src/component-quicksort/component-quicksort.c -O3 bin/component-quicksort-c