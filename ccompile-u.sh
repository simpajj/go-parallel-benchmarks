gcc src/seq-binarysearch/seq-binarysearch.c -o bin/seq-binarysearch-c
gcc src/seq-bubblesort/seq-bubblesort.c -o bin/seq-bubblesort-c
gcc src/seq-gcd-lcm/seq-gcd-lcm.c -o bin/seq-gcd-lcm-c
gcc src/seq-matrixmult/seq-matrixmult.c -o bin/seq-matrixmult-c
gcc -fopenmp src/micro-broadcast/micro-broadcast.c -o bin/micro-broadcast-c
gcc -fopenmp src/micro-multiplex/micro-multiplex.c -o bin/micro-multiplex-c
gcc -fopenmp src/micro-ping/micro-ping.c -o bin/micro-ping-c
gcc -fopenmp src/micro-pingpong/micro-pingpong.c -o bin/micro-pingpong-c
gcc -fopenmp src/component-dotprod/component-dotprod.c -o bin/component-dotprod-c
gcc -fopenmp src/component-factorial/component-factorial.c -o bin/component-factorial-c
gcc -fopenmp src/component-pi/component-pi.c -o bin/component-pi-c
gcc -fopenmp src/component-primefilter/component-primefilter.c -o bin/component-primefilter-c
gcc -fopenmp src/component-quicksort/component-quicksort.c -o bin/component-quicksort-c