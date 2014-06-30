{ time ./micro-broadcast $1 $2 ; } 2>> ../output/micro-broadcast.txt
{ time ./micro-multiplex $1 $2; } 2>> ../output/micro-multiplex.txt
{ time ./micro-ping $1 $2; } 2>> ../output/micro-ping.txt
{ time ./micro-pingpong $1 $2; } 2>> ../output/micro-pingpong.txt
{ time ./micro-broadcast-c $1 $2 ; } 2>> ../output/micro-broadcast-c.txt
{ time ./micro-multiplex-c $1 $2; } 2>> ../output/micro-multiplex-c.txt
{ time ./micro-ping-c $1 $2; } 2>> ../output/micro-ping-c.txt
{ time ./micro-pingpong-c $1 $2; } 2>> ../output/micro-pingpong-c.txt