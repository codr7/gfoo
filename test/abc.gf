42 .. check: =(42)
check: =(42)

42 'fail _ check: =(42)

42 type check: is(Integer)

35 +(7) check: =(42)
42 -(7) check: =(35)
6 *(7) check: =(42)

"foo" type check: is(String)

{
  let: t (thread: () {pause: 1 pause: 2 3})
  t call check: =(1)
  t call check: =(2)
  t call check: =(3)
}