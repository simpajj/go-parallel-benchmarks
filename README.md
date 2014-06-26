Directory structure
--------------------
The source code can be found in the /src directory. Executable files are in the /bin directory. All the output is piped to files corresponding to the executable file names in the /output directory.

Compiling the source code
-------------------------
If you wish to compile the source code yourself you can do so using the install shell script: 
> 	./install.sh

It will compile all the code and place all the executables in the /bin directory.

Running the executables
-----------------------
All the executables can be run from the /bin folder using the following command: 
> 	./runscript.sh {argument 1} {argument 2}

Argument 1 specifies the number of threads (or goroutines) and argument two the number of cores to be used. These arguments do not apply for all programs.

File types
----------
- **SC**: Sanity checks, not used for benchmarks
- **Seq**: Sequential baseline measurements
- **Micro**: Microbenchmarks, simple parallel operations
- **Component**: Non-trivial parallel operations, minimal I/O
- **Suite**: Benchmarks taken from the Rodinia Parallel Benchmarking Suite and implemented in Go