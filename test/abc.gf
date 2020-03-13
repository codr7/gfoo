NIL check: is(NIL)

T and: 42 check: =(42)
F and: 42 check: =(F)

F or: 42 check: =(42)
42 or: F check: =(42)

42 .. check: =(42)
check: =(42)

42 'fail _ check: =(42)

35 +(7) check: =(42)
42 -(7) check: =(35)
6 *(7) check: =(42)

42 check: is(42)

1 check: <(2)
1 check: <=(2)
1 check: <=(1)

2 check: >(1)
2 check: >=(1)
1 check: >=(1)

42 bool check: =(T)
"" bool check: =(F)

Int check: isa(Number)

[3 ...] check: =([0 1 2])

length("abc") check: =(3)

["foo"...] check: =([\'f \'o \'o])

[,'foo 42...] check: =(['foo 42])

peek([]) check: is(NIL)

peek([1 2 3]..)
check: is(3)
check: =([1 2 3])

pop([]) check: is(NIL)

pop([1 2 3]..)
check: is(3)
check: =([1 2])

[1 2].. push(3) check: =([1 2 3])
length([1 2 3]) check: =(3)

[1 2 [3 4 5]...] check: =([1 2 3 4 5])

{
  let: foo 'bar

  {
    let: foo 'baz
    foo check: =('baz)
  }
  
  foo check: =('bar)
}

{
  42 let: foo ()
  foo check: =(42)
}
  
{
  define: foo 'bar
  foo check: =('bar)
}

{
  let: r (data.record: (foo 1 bar 2 baz 3))
  r length check: =(3)
  r .foo check: =(1)
  r .bar check: =(2)
  r .baz check: =(3)
 
  r .qux check: is(NIL)
  r set('qux 4)
  r .qux check: =(4)

  let: c clone(r)
  c set('qux 5)
  r .qux check: =(4)
}

data.record: (foo 1 bar 2) ..
union(data.record: (foo 3 bar 4 baz 5))
check: =(data.record: (foo 1 bar 2 baz 5))

scope: (foo 35)

.. do: {
  let: bar (foo +(7))
}

.bar check: =(42)

{
  method: foo (x Int; Pair) {,'int x}
  method: foo (x String; Pair) {,'string x}
  foo(42) check: =(,'int 42)
  foo("bar") check: =(,'string "bar")
}

{
  method: foo (_ Int; Id) (_ 'int)
  method: foo (_ String; Id) (_ 'string)

  foo(42) check: =('int)
  foo("bar") check: =('string)
}

{
  method: min (x Any y 0; 0) {x <=(y) ?: x y}
  min("foo" "bar") check: =("bar")
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

{
  macro: foo () {'(let: #bar 42)}
  foo foo
}

include: "../lib/abc.gf"
T if: 'ok check: =('ok)
F else: 'ok check: =('ok)