# 1 = # of iterations
{ time ./component-dotprod $1; } 2>> ../output/component-dotprod.txt
{ time ./component-factorial $1; } 2>> ../output/component-factorial.txt
{ time ./component-pi $1; } 2>> ../output/component-pi.txt
{ time ./component-primefilter $1; } 2>> ../output/component-primefilter.txt
{ time ./component-quicksort $1; } 2>> ../output/component-quicksort.txt
{ time ./component-dotprod-c $1; } 2>> ../output/component-dotprod-c.txt
{ time ./component-factorial-c $1; } 2>> ../output/component-factorial-c.txt
{ time ./component-pi-c $1; } 2>> ../output/component-pi-c.txt
{ time ./component-primefilter-c $1; } 2>> ../output/component-primefilter-c.txt
{ time ./component-quicksort-c $1; } 2>> ../output/component-quicksort-c.txt