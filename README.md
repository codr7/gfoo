### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ go build -o gfoo main.go
$ ./gfoo
gfoo v0.6

Press Return on empty line to evaluate.

  42

[42]
```

### syntax
By default, arguments are expected before operations.

```
  42 type

[Int]
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

### pairs
Pairs allow treating two values as one, and may be created using `,`.

```
  1 2,

[1 2,]
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
`?:` may be used to conditionally evaluate code.

```
  T ?: 'ok 'fail
  F ?: 'fail 'ok

['ok 'ok]
```

`if:` and `else:` are defined in the [abc](https://github.com/codr7/gfoo/tree/master/lib/abc.gf) module, and may be used to when there is only one branch.

```
  load("lib/abc.gf")
  T if: 'ok
  F else: 'ok

['ok 'ok]
```

All values have Bool representations; non-zero Ints are true; empty Strings and Slices false etc.

```
  42 if: 'ok
  "" else: 'ok
  
['ok 'ok]
```

### macros
Macros are called before compilation and expand to the unquoted contents of their stacks.

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

Identifiers may be prefixed with `$` to avoid capturing bindings at the point of expansion.

```
  macro: foo () {
    '(let: $bar 42)
  }

[]
  foo bar

Error in 'repl', line 1, column 0: Unknown identifier: bar
[]
  $bar

Error in 'repl', line 1, column 0: Unknown identifier: $bar
[]
  foo

[]
```

Macro arguments are bound to forms following the call in specified order. By convention, macros that take compile time arguments have names ending with `:`. Values may be spliced into quoted forms using `@`

```
  macro: while: (cond body) {
    '(loop: (@cond else: break @body))
  }

[]

  3 while: () {
    say(..) --
  }

3
2
1
[]
```

### threads
Threads are implemented as Goroutines, which means they are preemptive yet more efficient than OS threads. New threads may be started using `thread:`, which takes an initial stack and body as arguments and starts the thread immediately. Calling a thread waits for it to stop executing and returns the contents of its stack.

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
  T check: is(F)

Error in 'repl', line 1, column 2: Check failed: (F is) [T]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project, every contribution helps.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>