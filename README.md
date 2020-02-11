### setup

```
$ go get https://github.com/codr7/gfoo.git
$ cd ~/go/src/gfoo
$ go build -o gfoo main.go
$ ./gfoo
gfoo v0.1

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

### the stack
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

`|` may be used to reset the stack.

```
  1 2 3 | 4 5 6

[4 5 6]
```

### bindings
Identifiers may be bound to values in the current scope using `let:`.

```
  let: foo 42
  foo

[42]
```

Rebinding identifiers within the same scope signals compile time errors.

```
  let: foo "bar"

Error in 'repl', line 1, column 5: Duplicate binding: foo
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

or by quoting group forms.

```
  '(foo bar baz)
  
[['foo 'bar 'baz]]
```

### license
[MIT](https://github.com/codr7/gfoo/blob/master/LICENSE.txt)

### support
Please consider a donation if you would like to support the project, every contribution helps.

<a href="https://liberapay.com/codr7/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a>