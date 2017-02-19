.PHONY: clean build bench all

plot-type := svg

all: | clean build bench

clean:
	@rm -rf nimcache bench-out
	@rm -f libnimalgs.a

build:
	@nim c \
		-d:release \
		--app:staticLib \
		--noMain \
		--header \
		nimalgs.nim

bench:
	@mkdir -p bench-out
	@go test -bench=. | tee /dev/tty | ./bench/to-gnuplot bench-out/
	@cat ./bench-out/Fib.dat | ./bench/remove-plots GoRec NimRec > ./bench-out/FibScaled-1.dat
	@cat ./bench-out/FibScaled-1.dat | ./bench/remove-plots GoTramp NimTramp > ./bench-out/FibScaled-2.dat
	@./bench/plot-all ./bench-out $(plot-type)
