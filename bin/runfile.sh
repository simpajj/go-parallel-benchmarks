# Sequential benchmarks
{ time ./seq-bubblesort; } 2>> ../output/seq-bubblesort.txt
{ time ./seq-binarysearch; } 2>> ../output/seq-binarysearch.txt
{ time ./seq-gcd-lcm; } 2>> ../output/seq-gcd-lcm.txt
{ time ./seq-matrixmult; } 2>> ../output/seq-matrixmult.txt
{ time ./seq-bubblesort-c; } 2>> ../output/seq-bubblesort-c.txt
{ time ./seq-binarysearch-c; } 2>> ../output/seq-binarysearch-c.txt
{ time ./seq-gcd-lcm-c; } 2>> ../output/seq-gcd-lcm-c.txt
{ time ./seq-matrixmult-c; } 2>> ../output/seq-matrixmult-c.txt

# Microbenchmarks
{ time ./micro-broadcast $1 $2 ; } 2>> ../output/micro-broadcast.txt
{ time ./micro-multiplex $1 $2; } 2>> ../output/micro-multiplex.txt
{ time ./micro-ping $1 $2; } 2>> ../output/micro-ping.txt
{ time ./micro-pingpong $1 $2; } 2>> ../output/micro-pingpong.txt
{ time ./micro-broadcast-c $1 $2 ; } 2>> ../output/micro-broadcast-c.txt
{ time ./micro-multiplex-c $1 $2; } 2>> ../output/micro-multiplex-c.txt
{ time ./micro-ping-c $1 $2; } 2>> ../output/micro-ping-c.txt
{ time ./micro-pingpong-c $1 $2; } 2>> ../output/micro-pingpong-c.txt