Directory structure
--------------------
The source code can be found in the /src directory. Executable files are in the /bin directory. All the output is piped to files corresponding to the executable file names in the /output directory.

Compiling the source code
-------------------------
If you wish to compile the source code yourself you can do so using the install shell script: 
> 	./install.sh

It will compile all the code and place all the executables in the /bin directory. The current binaries are compiled to run on linux/amd64.

Running the executables
-----------------------
All the executables can be run from the /bin folder using one of the runscripts, depending on which type of benchmark, e.g:  
> 	./runsequential.sh {# of iterations}

Using the CLI, you can specify the number of iterations of the scripts. The number of threads or goroutines is automatically set to equal the number of processors available on the system. 

File types
----------
- **SC**: Sanity checks, not used for benchmarks
- **Seq**: Sequential baseline measurements
- **Micro**: Microbenchmarks, simple parallel operations
- **Component**: Non-trivial parallel operations, minimal I/O
- **Suite**: Benchmarks taken from the Rodinia Parallel Benchmarking Suite and implemented in Go