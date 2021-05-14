# owoify-go
Turning your worst nightmare into a Go module.

[![Go](https://github.com/deadshot465/owoify-go/actions/workflows/go.yml/badge.svg)](https://github.com/deadshot465/owoify-go/actions/workflows/go.yml)

- [Go module](https://pkg.go.dev/github.com/deadshot465/owoify-go)

This is a Go port of [mohan-cao's owoify-js](https://github.com/mohan-cao/owoify-js), which will help you turn any string into nonsensical babyspeak similar to LeafySweet's infamous Chrome extension.

Just like my other Owoify ports, three levels of owoness are available:

1. **owo (default)**: The most vanilla one.
2. **uwu**: The moderate one.
3. **uvu**: Litewawwy unweadabwal.

Please refer to the original [owoify-js repository](https://github.com/mohan-cao/owoify-js) for more information.

Seriously, if you have seen my other ports, you probably have already known the details.

## Reason for development
Because [pkg.go.dev](https://pkg.go.dev/) doesn't have any owoify package. And since I have already ported it to several languages, I might as well just keep porting it.

## Install instructions
On command line, simply run:
```bash
go get github.com/deadshot465/owoify-go
```

## Usage
owoify-go is implemented as a function inside the package `owoify_go`. That means you don't need to create an instance of any class; instead, just call the method `owoify` with the package name.
```go
package main

import (
	"fmt"
	"github.com/deadshot465/owoify-go"
)

func main() {
	fmt.Println(owoify_go.Owoify("This is the string to owo! Kinda cute, isn't it?", "owo"))
	fmt.Println(owoify_go.Owoify("This is the string to owo! Kinda cute, isn't it?", "uvu"))
}

// Output:
// This is teh stwing two owo! Kinda cute, isn't it?
// fwis is teh stwing two owo(/ =ω=)/ Kinda cute(/ =ω=)/ isn't it?
```

## Disclaimer
As usual, I'm writing this package for both practicing and bots' needs. Performance is **NOT** guaranteed.

That being said, all regular expressions are declared beforehand rather than being declared inside a function. This module utilizes `sync.once` to make sure all initialization of regular expressions will only run once. That should be able to avoid the performance cost of building regular expressions every time the method is called.

## See also
- [owoify-js](https://github.com/mohan-cao/owoify-js) - The original owoify-js repository.
- [Owoify.Net](https://www.nuget.org/packages/Owoify.Net/1.0.1) - The C# port of Owoify written by me.
- [Owoify++](https://github.com/deadshot465/OwoifyCpp) - The C++ header-only port of Owoify written by me.
- [owoify_rs](https://crates.io/crates/owoify_rs) - The Rust port of Owoify written by me.
- [owoify-py](https://pypi.org/project/owoify-py/) - The Python port of Owoify written by me.
- [owoify_dart](https://pub.dev/packages/owoify_dart) - The Dart port of Owoify written by me.
- [owoify_rb](https://rubygems.org/gems/owoify_rb) - The Ruby port of Owoify written by me.