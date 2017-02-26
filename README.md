# x-go-nim-alg-cmp
Some exercises to feel Nim vs Go, to invoke Nim code from Go, scripts to plot benchmark results.

This repository is my experiments with Nim and Go.

If you ever stumble on it, feel free to use any code within it as you like (hereby it's in the public domain), but keep in mind, that this code is completely unstable and unsupported (also maybe incorrect and probably usles anyway).

This project illustrates:

1. Several simple algorithms in both Nim and Go, not of any good quality, just to feel both languages vs each other.
2. Linking static library, compiled with Nim, into the Go program.
3. Several awk and gnuplot scripts to draw golang benchmark output.

__Note__: to use plot scripts benchmarks should be coded with simple conventions:
* All benchmarks for different N (size of problem) should be a subtests within one BenchmarkXxx function.
* Name of subtest should be in form `fmt.Sprintf("%s#%d", nameOfAlgImplementation, N)`. Example `go test` output:

        BenchmarkFib/SimpleFibNim#1-8             	10000000	       185 ns/op
        BenchmarkFib/SimpleFibGo#1-8              	300000000	         3.55 ns/op
        BenchmarkFib/SimpleFibNim#2-8             	10000000	       184 ns/op
        BenchmarkFib/SimpleFibGo#2-8              	300000000	         3.63 ns/op
        BenchmarkFib/SimpleFibNim#4-8             	10000000	       192 ns/op
        BenchmarkFib/SimpleFibGo#4-8              	100000000	        12.5 ns/op

__Note__: file type for plot (default is `svg`) could be passed via `make` argument like `make plot-type=png`.
Possible values are `svg`, `pdf`, `jpeg`, `png` (see gnuplot manual for full list). If you have `Inkscape` installed, you may use command `inkview *.svg` to preview `svg` files. Also, you don't need to run all benchmarks again to draw plots in another format, because data files saved in `bench-out` folder. Just use command `./bench/plot-all ./bench-out jpeg`.


TODO:
- sorting algorithms, benchmark vs standard (try to use functional style and list comprehancion in Nim)
- parallel or external sorting
- graph traversal algorithm (use adj matrix and adj list)
- path finding algorithm
- https://blog.gopheracademy.com/advent-2015/glow-map-reduce-for-golang/
- parallel matrix alg - multiplication, min path in graph, ...
- 8 queens problem (recursive and las vegas)
- max flow
- traveling salesman (using ants approach)
