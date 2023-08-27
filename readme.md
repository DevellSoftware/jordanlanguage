# Jordan programming language

## Introductio

Jordan is a programming language that emphasises DSL, type autodetection, code readability, and other features that are not common in most of other languages.

Jordan is written in Go lang.

## Contribution

I warm welcome everyone that wants to contribute in this project.

## Features examples

### Function definition

```
fun helloworld {
  print "Hello, world!"
}

fun hello name {
  print "Hello $name"
}

helloworld
hello "Filip"
```

### Types aliases and autodetection

```
class Account {
  created at
  name
}

-- Equavilent to

class Account {
  created at: Date
  name: string
}
```

### Built in DSL functions

```
print name of accounts where active is true
```

* `name` is a property of accounts
* `active` is a property of accounts
* `accounts` is equavilent to loop over the `account`, english plurals are autodetected.
* `where` is a filtering functions
* `is` is a comparision operator, equavilent to `==`
* `of` is a keyword to specify the object to be used in a functions

### Other features examples

* last line of the function is returning value
* pattern matching
* type inference
* no semilicons needed
* objective and functional programming paradigms are supported

## Roadmap

Currently a proof of concept code was in the implementation.
The next step is to either modify it or rewrite from scratch,
basing on similar concept.

* [ ] Lexer
* [ ] Parser
* [ ] AST
* [ ] Runtime
* [ ] Webassembly compiler
* [ ] Webassembly runtime
* [ ] CLI
* [ ] VSCode extension
* [ ] Neovim extension


