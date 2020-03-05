### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ go build -o gfoo main.go
$ ./gfoo
gfoo v0.11

Press Return on empty line to evaluate.

  42

[42]
```

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
Besides first class features described elsewhere in this document, the following types are provided. New types may be created using `type:`

```
  type: Foo Int
  35 as(Foo)

[35]
  .. type

[35 Foo]
```

New types are derived from their implementation and may be used as such without conversion.

```
  _ +(7)

[42]
  type

[Int]
```

#### Bool
All values have boolean representations; non-zero integers are true, empty strings and slices false etc.

```
  42 bool
  "" bool
  
[T F]
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

#### Pair
Pairs allow treating two values as one, and may be created using `,`.

```
  1 2,

[1 2,]
```

#### Record
Records are ordered mappings from identifiers to values.

```
  data.record: (foo 1 bar 2)

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

`set` may be used to insert/update fields.

```
  .. set('baz 3)

[Record(bar 2 baz 3 foo 1)]
```

`union` may be used to update several fields at once, the left value is kept for duplicates.

```
  data.record: (foo 1 bar 2) .. union(data.record: (foo 3 bar 4 baz 5))

[Record(bar 2 baz 5 foo 1)]
```

#### Slice
Slices may be created by enclosing code in brackets.

```
  ['foo 'bar 'baz]
  
[['foo 'bar 'baz]]
```

`;` may be used as shorthand for a nested slice.

```
  ['foo; 'bar 'baz]

[['foo ['bar 'baz]]]
```

#### Time
`now` may be used to get the current time.

```
  time.now
[2020-03-01T01:30:30.3399994Z]
```

### values
`is` may be used to check if two values share the same address.

```
  'foo is('foo)

[T]
```

Some values; including perhaps surprisingly, integers; are heap allocated.

```
  42 is(42)

[F]
```

`=` may be used to check if two values are equal.

```
  42 =(42)

[T]
```

Composite values support deep equality.

```
  [1 2 3] =([1 2 3])

[T]
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

Error in 'repl', line 1, column 5: Duplicate binding: foo
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

Error in 'repl', line 1, column 6: Attempt to override compile time binding: foo
```

### scopes
`scope:` may be used to create a new scope with specified bindings.

```
  scope: (foo 35)

[Scope(0xc0000447c0)]
```

`do:` may be used to evaluate code in an external scope.

```
  .. do: {
    let: bar (foo +(7))
  }

[Scope(0xc0000447c0)]
```

Identifiers starting with `.` get their scope from the stack.

```
  .bar

[42]
```

### branches
`?:` may be used to conditionally evaluate code.

```
  T ?: 'ok 'fail
  F ?: 'fail 'ok

['ok 'ok]
```

`if:` and `else:` are defined in the [abc](https://github.com/codr7/gfoo/tree/master/lib/abc.gf) module.

```
  include: "lib/abc.gf"
  T if: 'ok
  F else: 'ok

['ok 'ok]
```

### lambdas
Lambdas may be created using `\:`;

```
  \: (x y) {x y 3}

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
  method: foo (x Int; Id) {'int x,}
  method: foo (x String; Id) {'string x,}
  foo(42) foo("bar")

['int 42, 'string "bar",]
```

Methods belong to the containing scope.

```
  method: bar (;Id) {'outer}
  bar

  {
    method: bar (;Id) {'inner}
    bar
  }

  bar

['outer 'inner 'outer]
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

Error in 'repl', line 1, column 0: Duplicate binding: bar
```

Identifiers may be prefixed with `#` to avoid capturing bindings at the point of expansion.

```
  macro: foo () {
    '(let: #bar 42)
  }

[]
  foo bar

Error in 'repl', line 1, column 0: Unknown identifier: bar
[]
  #bar

Error in 'repl', line 1, column 0: Unknown identifier: #bar
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
Threads are implemented as Goroutines, which means they are preemptive yet cheaper than OS threads. New threads may be started using `thread:`, which takes an initial stack and body as arguments and starts running immediately. Calling a thread waits for it to stop executing and returns the contents of its stack.

```
  thread: (1 2 3) {4 5 6}

[Thread(0xc0000a2000)]

  call

[1 2 3 4 5 6]
```

Threads may be paused until next call, which then returns the specified argument.

```
  thread: () {pause: 1 pause: 2 3}

[Thread(0xc0000a2000)]
  .. call

[Thread(0xc0000a2000) 1]
  _ .. call

[Thread(0xc0000a2000) 2]
  _ .. call

[Thread(0xc0000a2000) 3]
  _ call

Error in 'repl', line 1, column 2: Thread is done
```

### tests
Conditions may be asserted using `check:`, which signals an error describing the condition and incoming stack on failure.

```
  T check: =(F)

Error in 'repl', line 1, column 2: Check failed: (F =) [T]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>