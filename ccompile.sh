gcc -O3 src/seq-binarysearch/seq-binarysearch.c -o bin/seq-binarysearch-c
gcc -O3 src/seq-bubblesort/seq-bubblesort.c -o bin/seq-bubblesort-c
gcc -O3 src/seq-gcd-lcm/seq-gcd-lcm.c -o bin/seq-gcd-lcm-c
gcc -O3 src/seq-matrixmult/seq-matrixmult.c -o bin/seq-matrixmult-c
gcc -fopenmp -O3 src/micro-broadcast/micro-broadcast.c -o bin/micro-broadcast-c
gcc -fopenmp -O3 src/micro-multiplex/micro-multiplex.c -o bin/micro-multiplex-c
gcc -fopenmp -O3 src/micro-ping/micro-ping.c -o bin/micro-ping-c
gcc -fopenmp -O3 src/micro-pingpong/micro-pingpong.c -o bin/micro-pingpong-c
gcc -fopenmp -O3 src/component-dotprod/component-dotprod.c -o bin/component-dotprod-c
gcc -fopenmp -O3 src/component-factorial/component-factorial.c -o bin/component-factorial-c
gcc -fopenmp -O3 src/component-pi/component-pi.c -o bin/component-pi-c
gcc -fopenmp -O3 src/component-primefilter/component-primefilter.c -o bin/component-primefilter-c
gcc -fopenmp -O3 src/component-quicksort/component-quicksort.c -o bin/component-quicksort-c