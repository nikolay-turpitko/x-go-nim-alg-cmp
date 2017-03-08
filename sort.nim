import algorithm

proc std_sort(a: var openArray[cint]) {.exportc.} =
  sort(a, system.cmp[cint])
