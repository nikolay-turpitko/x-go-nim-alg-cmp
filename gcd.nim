import common

proc iter_gcd(a, b: cint): cint {.exportc.} =
  var b = b
  result = a
  while b != 0:
    var t = b
    b = result mod b
    result = t

proc iter_sub_gcd(a, b: cint): cint {.exportc.} =
  var b = b
  result = a
  while result != b:
    if result > b:
       result = result - b
    else:
       b = b - result

proc rec_gcd(a, b: cint): cint {.exportc.} =
  result = a
  if b != 0:
     result = rec_gcd(b, a mod b)

let mem_gcd_internal = memoize(rec_gcd)

proc mem_gcd(a, b: cint): cint {.exportc.} =
  result = mem_gcd_internal(a, b)
