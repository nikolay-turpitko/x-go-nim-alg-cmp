# This script plots series of benchmark results to pdf file.
# Usage: gnuplot -c plot.gp <input_output_data_file_without_ext> <output_format>
# Output format could be "pdf", "svg", "jpeg", "png".

set terminal ARG2
set output sprintf("%s.%s", ARG1, ARG2)
set title "ns/op from N"
set key left top
plot for [i=0:*] sprintf("%s.dat", ARG1) index i with linespoints pointsize 0.5 title columnhead
