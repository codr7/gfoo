42 .. check: is(42)
check: is(42)

42 'foo _ check: is(42)

check: ("foo" type is(String))

{
  let: t (thread: () {pause: 1 pause: 2 3})
  t call check: is(1)
  t call check: is(2)
  t call check: is(3)
}