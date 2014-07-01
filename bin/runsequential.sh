#$1 = number of iterations
{ time ./seq-bubblesort $1; } 2>> ../output/seq-bubblesort.txt
{ time ./seq-binarysearch $1; } 2>> ../output/seq-binarysearch.txt
{ time ./seq-gcd-lcm $1; } 2>> ../output/seq-gcd-lcm.txt
{ time ./seq-matrixmult $1; } 2>> ../output/seq-matrixmult.txt
{ time ./seq-bubblesort-c $1; } 2>> ../output/seq-bubblesort-c.txt
{ time ./seq-binarysearch-c $1; } 2>> ../output/seq-binarysearch-c.txt
{ time ./seq-gcd-lcm-c $1; } 2>> ../output/seq-gcd-lcm-c.txt
{ time ./seq-matrixmult-c $1; } 2>> ../output/seq-matrixmult-c.txt