### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ mkdir dist
$ ./build
$ dist/gfoo
gfoo v0.1

Press Return in empty row to evaluate.

  42

[42]
```

### syntax
By default, arguments are expected to appear before function calls.

```
  42 type

[Int64]
```

Trailing arguments may be enclosed in parens to get prefix/infix notation.

```
  type('foo)

[Id]
```

### the stack
Literals, values of bindings and results of macros and methods are pushed on the stack.

```
  1 2 3

[1 2 3]
```

The top value may be dropped using `_`.

```
  1 2 3 _

[1 2]
```

### bindings
Identifiers may be bound to values in the current scope using `let:`.

```
  let: foo 42
  foo

[42]
```

Rebinding identifiers within the same scope is not allowed.

```
  let: foo 42

Error in 'repl', line 1, column 0: Duplicate binding: foo
```

Specifying `_` as value pops it from the stack.

```
  "baz"
  
["baz"]

  let: bar _

[]

  bar

["baz"]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project, every contribution helps.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>