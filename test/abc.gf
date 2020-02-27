42 .. check: =(42)
check: =(42)

42 'fail _ check: =(42)

35 +(7) check: =(42)
42 -(7) check: =(35)
6 *(7) check: =(42)

1 check: <(2)
1 check: <=(2)
1 check: <=(1)

2 check: >(1)
2 check: >=(1)
1 check: >=(1)

42 bool check: =(T)
"" bool check: =(F)

Number check: <(Int)
Number check: <=(Int)
Number check: <=(Number)

Number check: >(Any)
Number check: >=(Any)
Number check: >=(Number)

{
  method: foo (x Int; Id) {'int x,}
  method: foo (x String; Id) {'string x,}
  foo(42) check: =('int 42,)
  foo("bar") check: =('string "bar",)
}

{
  method: bar (;Id) {'outer}

  {
    method: bar (;Id) {'inner}
    bar check: =('inner)
  }

  bar check: =('outer)
}

thread: (35) {+(7)}
call check: =(42)

{
  let: t (thread: () {pause: 1 pause: 2 3})
  t call check: =(1)
  t call check: =(2)
  t call check: =(3)
}

include: "../lib/abc.gf"
T if: 'ok check: =('ok)
F else: 'ok check: =('ok)