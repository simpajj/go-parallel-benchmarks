# Create executable for all source files
# --------------------------------------

# Go files
#export GOPATH=$GOPATH
#export GOBIN=$GOPATH/bin
go install src/seq-binarysearch/seq-binarysearch.go
go install src/seq-bubblesort/seq-bubblesort.go
go install src/seq-gcd-lcm/seq-gcd-lcm.go
go install src/seq-matrixmult/seq-matrixmult.go
go install src/micro-broadcast/micro-broadcast.go
go install src/micro-multiplex/micro-multiplex.go
go install src/micro-ping/micro-ping.go
go install src/micro-pingpong/micro-pingpong.go

# C files
gcc src/seq-binarysearch/seq-binarysearch.c -o bin/seq-binarysearch-c
gcc src/seq-bubblesort/seq-bubblesort.c -o bin/seq-bubblesort-c
gcc src/seq-gcd-lcm/seq-gcd-lcm.c -o bin/seq-gcd-lcm-c
gcc src/seq-matrixmult/seq-matrixmult.c -o bin/seq-matrixmult-c
gcc-4.9 -fopenmp src/micro-broadcast/micro-broadcast.c -o bin/micro-broadcast-c
gcc-4.9 -fopenmp src/micro-multiplex/micro-multiplex.c -o bin/micro-multiplex-c
gcc-4.9 -fopenmp src/micro-ping/micro-ping.c -o bin/micro-ping-c
gcc-4.9 -fopenmp src/micro-pingpong/micro-pingpong.c -o bin/micro-pingpong-c