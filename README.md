### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ go build -o gfoo main.go
$ ./gfoo
gfoo v0.4

Press Return on empty line to evaluate.

  42

[42]
```

### syntax
By default, arguments are expected before operations.

```
  42 type

[Int64]
```

Trailing arguments may be enclosed in parens to get prefix/infix notation.

```
  type("foo")

[String]
```

### stacks
Literals, values of bindings and results from operations are pushed on a stack.

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

`|` may be used to drop all items.

```
  1 2 3 | 4 5 6

[4 5 6]
```

### bindings
Identifiers may be bound to values in the current scope using `let:`.

```
  let: foo 42

[]
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

### numbers
Numeric literals may be specified using decimal, hexadecimal or binary notation.

```
  42 0x2a 0b101010

[42 42 42]
```

### slices
Slices may be created by enclosing code in brackets,

```
  ['foo 'bar 'baz]
  
[['foo 'bar 'baz]]
```

or by quoting groups.

```
  '(foo bar baz)
  
[['foo 'bar 'baz]]
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

### branches
`if:` and `else:` may be used to conditionally evaluate code.

```
  T if: 'is-true
  F else: 'is-false

['is-true 'is-false]
```

All values have boolean representations; non-zero integers are true; empty strings and slices false etc.

```
  42 if: 'is-true
  "" else: 'is-false
  
['is-true 'is-false]
```

### macros
Macro arguments are bound to forms following the call in specified order. By convention, macros that take compile time arguments have identifiers ending with `:`. Macros expand to the unquoted contents of the stack.

```
  macro: foo () {'(let: bar 42)}
  foo

[]
  bar

[42]
  foo

Error in 'repl', line 1, column 0: Duplicate binding: bar
```

Identifiers may be prefixed with `$` to avoid capturing bindings at the point of expansion.

```
  macro: foo () {'(let: $bar 42)}
  foo

[]

  bar

Error in 'repl', line 1, column 0: Unknown identifier: bar
[]
  $bar

Error in 'repl', line 1, column 0: Unknown identifier: $bar
[]
  foo

[]
```

### threads
Threads are implemented as Goroutines, which means they are preemptive yet more efficient than OS threads. New threads may be started using `thread:`, which takes an initial stack and body as arguments and starts the thread immediately. Calling a thread waits for it to stop executing and returns the result.

```
  thread: (1 2 3) {4 5 6}

[Thread(0xc0000a2000)]

  call

[1 2 3 4 5 6]
```

Threads may be paused until next call, which then returns the specified argument.

```
  thread: () {1 pause: 2 3}

[Thread(0xc0000a2000)]
  .. call

[Thread(0xc0000a2000) 2]
  _ call

[1 3]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project, every contribution helps.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>