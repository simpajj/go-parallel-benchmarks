{ time ./micro-broadcast $1 $2 ; } 2>> ../output/micro-broadcast.txt
{ time ./micro-multiplex $1 $2; } 2>> ../output/micro-multiplex.txt
{ time ./micro-ping $1 $2; } 2>> ../output/micro-ping.txt
{ time ./micro-pingpong $1 $2; } 2>> ../output/micro-pingpong.txt