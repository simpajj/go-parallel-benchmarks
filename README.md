Directory structure
--------------------
The source code can be found in the /src directory. Binaries are in the /bin directory. All the output is piped to files corresponding to the executable file names in the /output directory.

Compiling the source code
-------------------------
If you wish to compile the source code yourself you can do so using one of the compilation scripts, e.g: 
> 	./gocompile.sh

The pre-compiled Go binaries are compiled to run on linux/amd64 architectures. Due to cross-compilation issues, the C source code has been left uncompiled and will have to be compiled on the target machine. 

Running the executables
-----------------------
All the executables can be run from the /bin folder using one of the runscripts, depending on which type of benchmark, e.g:  
> 	./runsequential.sh {# of iterations}

The number of iterations can be specified via the command line. The number of threads or goroutines is automatically set to equal the number of processors available on the system.

File types
----------
- **SC**: Sanity checks, not used for benchmarks
- **Seq**: Sequential baseline measurements
- **Micro**: Microbenchmarks, simple parallel operations
- **Component**: Non-trivial parallel operations