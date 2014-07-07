# Create executable for all source files
# --------------------------------------

# If needed
#export GOOS=linux
#export GOARCH=amd64
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
go install src/component-factorial/component-factorial.go
go install src/component-dotprod/component-dotprod.go
go install src/component-pi/component-pi.go
go install src/component-primefilter/component-primefilter.go
go install src/component-quicksort/component-quicksort.go