# This script plots series of benchmark results to pdf file.
# Usage: gnuplot -c plot.gp <input_data_file> <output_data_file>

set terminal pdf
set output ARG2
set title "ns/op from N"
set key left top
plot for [i=0:*] ARG1 index i with linespoints pointsize 0.5 title columnhead
