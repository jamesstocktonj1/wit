# wit

A Go library for programmatically generating [WIT (WebAssembly Interface Types)](https://component-model.bytecodealliance.org/design/wit.html) definitions. WIT is the interface description language for the WebAssembly Component Model — it describes the types, functions, and interfaces that components expose and consume.

This package gives you a type-safe, fluent API to construct WIT definitions in Go and serialize them to `.wit` files, making it straightforward to generate component interfaces as part of a build pipeline or code generation tool.

## Installation

```bash
go get github.com/jamesstocktonj1/wit
```

## Example

The following builds a WIT package with a `handler` interface and writes it to a file:

```go
package main

import (
    "log"
    "os"

    "github.com/jamesstocktonj1/wit"
)

func main() {
    f, err := os.Create("world.wit")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    w := wit.NewWit().
        WithPackage(
            wit.NewPackage("example", "http").WithVersion("0.1.0"),
        ).
        WithWorld(
            wit.NewWorld("proxy").
                WithImports("handler").
                WithExports("handler"),
        ).
        WithInterface(
            wit.NewInterface("handler",
                wit.NewRecord("request",
                    wit.NewField("method", wit.NewPrimitive(wit.String)),
                    wit.NewField("path", wit.NewPrimitive(wit.String)),
                ).WithDocs("An incoming HTTP request."),
                wit.NewRecord("response",
                    wit.NewField("status", wit.NewPrimitive(wit.Unsigned16)),
                    wit.NewField("body", wit.NewList(wit.NewPrimitive(wit.Unsigned8))),
                ).WithDocs("An HTTP response."),
            ).WithFunctions(
                wit.NewFunction("handle",
                    wit.NewParam("req", wit.NewReference("request")),
                ).WithResult(wit.NewParam("", wit.NewReference("response"))),
            ),
        )

    if err := wit.NewEncoder(f).Encode(w); err != nil {
        log.Fatal(err)
    }
}
```

This produces a `.wit` file equivalent to:

```wit
package example:http@0.1.0;

world proxy {
    import handler;
    export handler;
}

interface handler {
    /// An incoming HTTP request.
    record request {
        method: string,
        path: string,
    }

    /// An HTTP response.
    record response {
        status: u16,
        body: list<u8>,
    }

    handle: func(req: request) -> response;
}
```

### Type system

The package covers the full WIT type system:

| Go constructor | WIT type |
|---|---|
| `NewPrimitive(Bool\|String\|Char\|...)` | `bool`, `string`, `char`, `s8`–`s64`, `u8`–`u64`, `f32`, `f64` |
| `NewRecord(name, fields...)` | `record` |
| `NewEnum(name, cases...)` | `enum` |
| `NewVariant(name, cases...)` | `variant` |
| `NewFlags(name, cases...)` | `flags` |
| `NewList(elemType)` | `list<T>` |
| `NewOption(innerType)` | `option<T>` |
| `NewResult(ok, err)` | `result<T, E>` |
| `NewTuple(fields...)` | `tuple<T, ...>` |
| `NewAlias(name, kind)` | `type name = T` |
| `NewReference(name)` | reference to a named type |

## Contributing

Contributions are welcome. Please open an issue before starting significant work so we can discuss the approach.

1. Fork the repository and create a feature branch.
2. Make your changes and add tests where appropriate.
3. Run the test suite: `go test ./...`
4. Open a pull request with a clear description of what changed and why.

All submitted code should pass `go vet` and be formatted with `gofmt`.
