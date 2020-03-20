### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ go build -o gfoo main.go
$ ./gfoo
gfoo v0.19

Press Return on empty line to evaluate.

  42

[42]
```

### status
While there is much left to do, all examples in this document should work as advertised and I have a couple of projects in progress.

* [bookng](https://github.com/codr7/bookng)

* [nojs](https://github.com/codr7/nojs)

### syntax
By default, arguments are expected before operations.

```
  35 7 +

[42]
```

Trailing arguments may be enclosed in parens to get prefix/infix notation.

```
  35 +(7)

[42]
```

### stacks
Literals, values of bindings and results from operations are pushed on the stack.

```
  1 2 3

[1 2 3]
```

The top value may be duplicated using `..`,

```
  ..
  
[1 2 3 3]
```

and dropped using `_`.

```
  _
  
[1 2 3]
```

Evaluating nothing in the REPL clears the stack.

```

[]
```

### types
New types may be created using `type:`

```
  trait: Foo ()
  type: Bar (Foo) Int
  35 as-bar

[35]
  typeof

[Bar]
```

New types are derived from their implementation type and may be used as such without conversion.

```
  _ +(7)

[42]
  type

[Int]
```

`isa` returns the direct parent, or `NIL` if none exists.

```
  Bar isa(Foo)

[Foo]
  isa(Bar)

[NIL]
```

Besides first class features described elsewhere in this document, the following types are provided.

#### Bool
All values have boolean representations; non-zero integers are true, empty strings and slices false etc.

```
  42 bool
  "" bool
  
[T F]
```

Booleans may be negated using `!`.

```
  !T

[F]
```

`and:` returns the value on top of the stack if it's false, otherwise the right operand is evaluated.

```
  T and: 42

[42]

  _ F and: say("evaluate!")

[F]
```

`or:` returns the value on top of the stack if it's true, otherwise the right operand is evaluated.

```
  F or: 42

[42]

  _ 42 or: say("evaluate!")

[F]
```

#### Char
Characters are prefixed with `\`.

```
  \n

[\n]
```

Literal characters are quoted.

```
  \'n

[\'n]
```

#### Id
Identifiers may be quoted and used as values.

```
  'foo

['foo]
  type

[Id]
```

#### Int
Integers may be specified using decimal, hexadecimal or binary notation.

```
  42 0x2a 0b101010

[42 42 42]
```

Integers may be negated using `!`,

```
  !42

[-42]
```

and spread using `...`.

```
  3...

[0 1 2]
```

#### Pair
Pairs allow treating two values as one, and may be created using `,`.

```
  ,1 2

[,1 2]
```

Pairs may be negated using `!`,

```
  !,1 2

[.-1 -2]
```

#### Record
Records are ordered mappings from identifiers to values.

```
  use: data (record: merge)
  record: (foo 1 bar 2)

[Record(bar 2 foo 1)]
```

Fields may be accessed directly without quoting.

```
  .. .foo

[Record(bar 2 foo 1) 1]
```

Missing fields return `NIL`.

```
  _ .baz

[NIL]
```

`set` may be used to insert/update fields;

```
  .. set('baz 3)

[Record(bar 2 baz 3 foo 1)]
```

and `merge` to update several fields at once, keeping the left value for duplicates.

```
  record: (foo 1 bar 2) .. merge(record: (foo 3 bar 4 baz 5))

[Record(bar 2 baz 5 foo 1)]
```

Records may be negated using `!`.

```
  !record: (foo 1 bar 2)

[Record(bar -2 foo -1)]
```

#### Slice

```
  ['foo 'bar 'baz]
  
[['foo 'bar 'baz]]
```

Slices support basic stack operations,

```
  [1 2] push(3)

[1 2 3]
  .. peek

[[1 2 3] 3]
  _ pop

[[1 2] 3]

  _ length

[2]
```

negation using `!`,

```
  ![1 2 3]

[-1 -2 -3]
```

and spreading using `...`.

```
  [1 2 3]...

[1 2 3]
```

`;` may be used as shorthand for a nested slice.

```
  ['foo; 'bar 'baz]

[['foo ['bar 'baz]]]
```

#### String

```
  "foo"

["foo"]
```

Strings may be spread using `...`.

```
  ...

[\'f \'o \'o]
```

#### Time
`now` may be used to get the current time,

```
  time.now
[2020-03-01T01:30:30.3399994Z]
```

while `today` truncates to whole days.

```
  time.today

[2020-03-01T00:00:00.00Z]
```

#### TimeDelta
Time deltas may be used to perform date arithmetics.

```
  time.today
  10 time.days

[2020-03-01T00:00:00Z TimeDelta(0 0 10)]
  +

[2020-03-10T00:00:00Z]
```

### modules
Definitions belong to modules and have to be imported into the current scope to be used without prefix.

```
  time.today

[]
  use: time (today)
  today

[]
```

Fundamental definitions are defined in the `abc` module, which is fully imported by default in the REPL. The same thing may be accomplished in scripts by using `abc...`.

test.bar
```
use: abc...
say(35 +(7))
```

```
$ gfoo test.bar
42
```

### values
`is` may be used to check if two values share the same identity.

```
  'foo is('foo)

[T]
```

```
  [1 2 3] is([1 2 3])

[F]
```

`=` may be used instead to check if two values are equal.

```
  [1 2 3] =([1 2 3])

[T]
```

### interactions
Values may be printed as is using `dump`,

```
  dump("foo")

"foo"
[]
```

or pretty-printed using `say`.

```
  say("foo")

foo
[]
```

Slices may be used to print several values at once.

```
  say([1 \n 2 \n 3])

1
2
3
[]
```

### bindings
Bindings come in two flavors, compile time and runtime.

Identifiers may be bound to values at runtime in the current scope using `let:`.
```
  let: foo 42
  foo

[42]
```

Rebinding in the same scope results in a compile time error,

```
  let: foo "bar"

Error in 'n/a', line 1, column 5: Duplicate binding: foo
```

while child scopes are allowed to override inherited bindings.

```
  {let: foo "bar" foo}

["bar"]

  foo

["bar" 42]
```

Specifying the empty group as value pops the stack.

```
  "baz"
  
["baz"]

  let: bar ()

[]

  bar

["baz"]
```

`define:` may be used to create compile time bindings.

```
  define: foo 42
  foo
```

Overriding compile time bindings is not allowed.

```
  {let: foo 42}

Error in 'n/a', line 1, column 6: Attempt to override compile time binding: foo
```

### branches
`?:` may be used to conditionally evaluate code.

```
  T ?: 'ok 'fail
  F ?: 'fail 'ok

['ok 'ok]
```

`$` is bound to the condition while evaluating branches.

```
  42 ?= $ 'fail

[42]
```

`if:` and `else:` may be used instead for single branch conditions.

```
  T if: 'ok
  F else: 'ok

['ok 'ok]
```

### sequences
`for:` may be used to execute code once for each item in a sequence, `$` is bound to the current item.

```
  3 for: ($ *(2))

[0 2 4]
```

`map:` may be used to lazily transform sequences, `$` is bound to the current item.

```
  3 map: ($ *(2))

[Iter(0xc000090130)]
  ...

[0 2 4]
```

Iterators may be manually consumed using `next`,

```
  3 count

[Iter(0xc000006148)]
  .. iter.next

[Iter(0xc000006148) 0]
  _ ...

[1 2]
```

and chained using `~`.

```
  3 count iter.~([3 4] items)

[Iter(0xc000090278)]
  ...

[0 1 2 3 4]
```


### lambdas
Lambdas may be created using `/:`;

```
  /: (x y) {x y 3}

[Lambda(0xc0000483c0)]
```

and evaluated using `call`, or `call:` which pushes specified arguments after the target is popped.

```
  call: (1 2)

[1 2 3]
```

### methods
Metods allow dispatching on argument types.

```
  method: foo(x Int; Pair) {,'int x}
  method: foo(x String; Pair) {,'string x}
  foo(42) foo("bar")

['int 42, 'string "bar",]
```

Arguments may be grouped per type;

```
  method: foo((x y) Int; Int) {x +(y)}
  foo(35 7)

[42]
```

or anonymous, which leaves the value on the stack.

```
  method: foo(_ Int; Id) (_ 'int)
  method: foo(_ String; Id) (_ 'string)
  foo(42) foo("bar")

['int 'string]
```

Indexes may be used instead of types to match any type compatible with the specified argument.

```
  method: min(x Any y 0; 0) {x <=(y) ?: x y}
  min("foo" "bar")

["bar"]
```

Methods belong to the containing scope.

```
  method: bar(;Id) 'outer
  bar

  {
    method: bar(;Id) 'inner
    bar
  }

  bar

['outer 'inner 'outer]
```

Calls may be negated using `!`.

```
  method: foo(;Int) 42
  !foo

[-42]
```

### macros
Macros are called during compilation and expand to the unquoted contents of their stacks.

```
  macro: foo () {
    '(let: bar 42)
  }

[]
  foo bar

[42]
  foo

Error in 'n/a', line 1, column 0: Duplicate binding: bar
```

Identifiers may be prefixed with `#` to avoid capturing bindings at the point of expansion.

```
  macro: foo () {
    '(let: #bar 42)
  }

[]
  foo bar

Error in 'n/a', line 1, column 0: Unknown identifier: bar
[]
  #bar

Error in 'n/a', line 1, column 0: Unknown identifier: #bar
[]
  foo

[]
```

Macro arguments are bound to forms following the call in specified order. By convention, macros that take compile time arguments have names ending with `:`. Values may be spliced into quoted forms using `@`.

```
  macro: if: (body) {
    '(?: @body ())
  }

  macro: else: (body) {
    '(?: () @body)
  }
```

### threads
Threads are implemented as Goroutines, which means they are preemptive yet cheaper than OS threads. New threads may be started using `thread:`, which takes an initial stack and body as arguments and starts running immediately.

```
  thread: (1 2 3) {4 5 6}

[Thread(0xc0000a2000)]

  wait

[1 2 3 4 5 6]
```

Threads may be paused until next call, which then returns the specified argument.

```
  thread: () {pause: 1 pause: 2 3}

[Thread(0xc0000a2000)]
  .. wait

[Thread(0xc0000a2000) 1]
  _ .. wait

[Thread(0xc0000a2000) 2]
  _ .. wait

[Thread(0xc0000a2000) 3]
  _ call

Error in 'n/a', line 1, column 2: Thread is done
```

### tests
Conditions may be asserted using `check:`, which signals an error describing the condition and incoming stack on failure.

```
  T check: =(F)

Error in 'n/a', line 1, column 2: Check failed: (F =) [T]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>