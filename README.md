# Monkey Programming Language

Welcome to the Monkey programming language! This project is an implementation based on the concepts taught in the books [Writing An Interpreter In Go](https://interpreterbook.com/) and [Writing A Compiler In Go](https://compilerbook.com/) by Thorsten Ball. 

## About Monkey

Monkey is a programming language that supports various data types, functions, conditionals, recursion, and closures. Here's an overview of its features:

### Integers & Arithmetic Expressions

```javascript
// Integers & arithmetic expressions...
let version = 1 + (50 / 2) - (8 * 3);
```

### Strings

```javascript
// ... and strings
let name = "The Monkey programming language";
```

### Booleans

```javascript
// ... booleans
let isMonkeyFastNow = true;
```

### Arrays & Hash Maps

```javascript
// ... arrays & hash maps
let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];
```

It also has functions!

### User-defined Functions

```javascript
// User-defined functions...
let getName = fn(person) { person["name"]; };
getName(people[0]); // => "Anna"
getName(people[1]); // => "Bob"
```

## Built-in Functions

```javascript
// and built-in functions
puts(len(people));  // prints: 2
```

### Conditionals and Recursion

Monkey supports conditionals, implicit and explicit returns, and recursive functions. Here's an example of a recursive function to compute Fibonacci numbers:

```javascript
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
```

### Closures

The crown jewel in every Monkey implementation is closures:

```javascript
// `newAdder` returns a closure that makes use of the free variables `a` and `b`:
let newAdder = fn(a, b) {
    fn(c) { a + b + c };
};
// This constructs a new `adder` function:
let adder = newAdder(1, 2);

adder(8); // => 11
```
