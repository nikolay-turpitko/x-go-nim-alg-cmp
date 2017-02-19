import common

proc rec_fib(n: cint): cint {.exportc.} =
  if n <= 2: result = 1
  else: result = rec_fib(n-1)+rec_fib(n-2)

proc tail_rec_fib(n: cint): cint {.exportc.} =
  proc f(term, val, prev: cint): cint =
    if term == 0: result = prev
    elif term == 1: result = val
    else: result = f(term-1,prev+val,val)
  result = f(n, 1, 0)

proc iter_fib(n: cint): cint {.exportc.} =
  if n <= 2:
    result = 1
  else:
    var prev1: cint = 0
    var prev2: cint = 1
    for i in 2..n:
      let tmp = prev1
      prev1 = prev2
      prev2 = tmp+prev2
    result = prev2

let mem_fib_internal = memoize(rec_fib)

proc mem_fib(n: cint): cint {.exportc.} =
  result = mem_fib_internal(n)

discard """
# Type tramp defines a function which returns result or next trampoline.
type
  trampres = object
    res: cint
    next: tramp
  tramp = proc(): trampres

# Function trampExec executes trampoline loop.
proc trampoline(t: tramp): cint =
  var r = t()
  while r.next != nil:
    r = r.next()
  result = r.res

proc tramp_tail_rec_fib_internal(term, val, prev: cint): tramp =
  if term == 0:
    result = proc(): trampres =
      result = trampres(res: prev, next: nil)
  elif term == 1:
    result = proc(): trampres =
      result = trampres(res: val, next: nil)
  else:
    result = proc(): trampres =
      var next = tramp_tail_rec_fib_internal(term-1, prev+val, val)
      result = trampres(res: cint(0), next: next)

proc tramp_tail_rec_fib(n: cint): cint {.exportc.} =
  result = trampoline(tramp_tail_rec_fib_internal(n, 1, 0))

echo tramp_tail_rec_fib(8)
"""

# This is a stab till code above fixed.
proc tramp_tail_rec_fib(n: cint): cint {.exportc.} =
  result = 42
