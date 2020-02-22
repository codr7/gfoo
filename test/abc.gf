42 .. check: is(42)
check: is(42)

42 'fail _ check: is(42)

42 type check: is(Integer)

35 +(7) check: is(42)
42 -(7) check: is(35)
6 *(7) check: is(42)

"foo" type check: is(String)

{
  let: t (thread: () {pause: 1 pause: 2 3})
  t call check: is(1)
  t call check: is(2)
  t call check: is(3)
}