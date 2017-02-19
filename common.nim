import tables

# Function memoize returns wrapped function which caches results on first
# execution and returns cached result afterwards.
# Note: we do not interested in general implementation here.

proc memoize*(f: proc(i: cint): cint): proc(i: cint): cint =
  var cache = initTable[cint, cint]()
  result = proc(i: cint): cint =
    if cache.hasKey(i):
      result = cache[i]
    else:
      result = f(i)
      cache[i] = result

proc memoize*(f: proc(a, b: cint): cint): proc(a, b: cint): cint =
  var cache = initTable[clonglong, cint]()
  result = proc(a, b: cint): cint =
    var k = a shl 32 or b
    if cache.hasKey(k):
      result = cache[k]
    else:
      result = f(a, b)
      cache[k] = result
