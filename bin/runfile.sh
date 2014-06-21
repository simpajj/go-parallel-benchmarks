{ time ./micro-broadcast $1 $2 ; } 2>> ../output/micro-broadcast.txt
{ time ./micro-multiplex $1 $2; } 2>> ../output/micro-multiplex.txt
{ time ./micro-ping $1 $2; } 2>> ../output/micro-ping.txt
{ time ./micro-pingpong $1 $2; } 2>> ../output/micro-pingpong.txt
{ time ./seq-bubblesort-c; } 2>> ../output/seq-bubblesort-c.txt
{ time ./seq-binarysearch-c; } 2>> ../output/seq-binarysearch-c.txt
{ time ./seq-gcd-lcm-c; } 2>> ../output/seq-gcd-lcm-c.txt