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

42 type check: is(Int)

Number check: <(Int)
Number check: <=(Int)
Number check: <=(Number)

Number check: >(Any)
Number check: >=(Any)
Number check: >=(Number)

thread: (35) {+(7)}
call check: =(42)

{
  let: t (thread: () {pause: 1 pause: 2 3})
  t call check: =(1)
  t call check: =(2)
  t call check: =(3)
}