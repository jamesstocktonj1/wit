package integration

import "github.com/jamesstocktonj1/wit"

func basic() wit.Wit {
	ref := wit.Reference("message")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "wasi",
			Package:   "test",
		},
		Interfaces: []wit.Interface{
			{
				Name: "greeting",
				TypeDefs: []wit.Type{
					&wit.Record{
						Name: "message",
						Fields: []wit.Field{
							{
								Name: "name",
								Kind: wit.NewPrimitive(wit.String),
							},
							{
								Name: "message",
								Kind: wit.NewPrimitive(wit.String),
							},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "greet",
						Params: []wit.Param{
							{
								Name: "name",
								Kind: wit.NewPrimitive(wit.String),
							},
						},
						Results: &wit.Param{
							Kind: &ref,
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name: "greeter",
				Exports: []string{
					"greeting",
				},
			},
		},
	}
}

func listTypes() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "lists",
		},
		Interfaces: []wit.Interface{
			{
				Name: "collection",
				TypeDefs: []wit.Type{
					&wit.Record{
						Name: "point",
						Fields: []wit.Field{
							{Name: "x", Kind: wit.NewPrimitive(wit.Float32)},
							{Name: "y", Kind: wit.NewPrimitive(wit.Float32)},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "sum",
						Params: []wit.Param{
							{Name: "values", Kind: &wit.List{Elem: wit.NewPrimitive(wit.Signed32)}},
						},
						Results: &wit.Param{Kind: wit.NewPrimitive(wit.Signed32)},
					},
					{
						Name: "join",
						Params: []wit.Param{
							{Name: "parts", Kind: &wit.List{Elem: wit.NewPrimitive(wit.String)}},
						},
						Results: &wit.Param{Kind: wit.NewPrimitive(wit.String)},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "collection-world",
				Exports: []string{"collection"},
			},
		},
	}
}

func optionResult() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "option-result",
		},
		Interfaces: []wit.Interface{
			{
				Name: "maybe",
				Functions: []wit.Function{
					{
						Name: "find",
						Params: []wit.Param{
							{Name: "key", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{
							Kind: &wit.Option{Inner: wit.NewPrimitive(wit.String)},
						},
					},
					{
						Name: "divide",
						Params: []wit.Param{
							{Name: "numerator", Kind: wit.NewPrimitive(wit.Float64)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  wit.NewPrimitive(wit.Float64),
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "maybe-world",
				Exports: []string{"maybe"},
			},
		},
	}
}

func enumFlags() wit.Wit {
	dirRef := wit.Reference("direction")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "enum-flags",
		},
		Interfaces: []wit.Interface{
			{
				Name: "permissions",
				TypeDefs: []wit.Type{
					&wit.Flags{
						Name:  "access",
						Cases: []wit.Case{{Name: "read"}, {Name: "write"}, {Name: "execute"}},
					},
					&wit.Enum{
						Name: "direction",
						Cases: []wit.Case{
							{Name: "north"},
							{Name: "south"},
							{Name: "east"},
							{Name: "west"},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "can-access",
						Params: []wit.Param{
							{Name: "dir", Kind: &dirRef},
						},
						Results: &wit.Param{Kind: wit.NewPrimitive(wit.Bool)},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "permissions-world",
				Exports: []string{"permissions"},
			},
		},
	}
}

func multiInterface() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "wasi",
			Package:   "io",
			Version:   "0.2.0",
		},
		Interfaces: []wit.Interface{
			{
				Name: "input",
				Functions: []wit.Function{
					{
						Name: "read",
						Params: []wit.Param{
							{Name: "len", Kind: wit.NewPrimitive(wit.Unsigned32)},
						},
						Results: &wit.Param{
							Kind: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)},
						},
					},
				},
			},
			{
				Name: "output",
				Functions: []wit.Function{
					{
						Name: "write",
						Params: []wit.Param{
							{Name: "buf", Kind: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
						},
						Results: &wit.Param{Kind: wit.NewPrimitive(wit.Unsigned32)},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "streams",
				Imports: []string{"input"},
				Exports: []string{"output"},
			},
		},
	}
}

func multiWorld() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "worlds",
		},
		Interfaces: []wit.Interface{
			{
				Name: "logger",
				Functions: []wit.Function{
					{
						Name: "log",
						Params: []wit.Param{
							{Name: "msg", Kind: wit.NewPrimitive(wit.String)},
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "server",
				Exports: []string{"logger"},
			},
			{
				Name:    "client",
				Imports: []string{"logger"},
			},
		},
	}
}

func nestedTypes() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "nested",
		},
		Interfaces: []wit.Interface{
			{
				Name: "deep",
				Functions: []wit.Function{
					{
						Name: "flatten",
						Params: []wit.Param{
							{
								Name: "matrix",
								Kind: &wit.List{Elem: &wit.List{Elem: wit.NewPrimitive(wit.Signed32)}},
							},
						},
						Results: &wit.Param{
							Kind: &wit.List{Elem: wit.NewPrimitive(wit.Signed32)},
						},
					},
					{
						Name: "lookup",
						Params: []wit.Param{
							{Name: "key", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{
							Kind: &wit.Option{Inner: &wit.List{Elem: wit.NewPrimitive(wit.String)}},
						},
					},
					{
						Name: "parse",
						Params: []wit.Param{
							{Name: "input", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  &wit.Option{Inner: wit.NewPrimitive(wit.Unsigned64)},
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "nested-world",
				Exports: []string{"deep"},
			},
		},
	}
}

func tupleTypes() wit.Wit {
	boundsRef := wit.Reference("bounds")
	pair := &wit.Tuple{Fields: []wit.Type{wit.NewPrimitive(wit.Float64), wit.NewPrimitive(wit.Float64)}}
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "tuples",
		},
		Interfaces: []wit.Interface{
			{
				Name: "coords",
				TypeDefs: []wit.Type{
					&wit.Record{
						Name: "bounds",
						Fields: []wit.Field{
							{Name: "min", Kind: pair},
							{Name: "max", Kind: pair},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "midpoint",
						Params: []wit.Param{
							{Name: "b", Kind: &boundsRef},
						},
						Results: &wit.Param{Kind: pair},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "coords-world",
				Exports: []string{"coords"},
			},
		},
	}
}

func aliasTypes() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "alias",
		},
		Interfaces: []wit.Interface{
			{
				Name: "types",
				TypeDefs: []wit.Type{
					&wit.Alias{Name: "name", Kind: wit.NewPrimitive(wit.String)},
					&wit.Alias{Name: "age", Kind: wit.NewPrimitive(wit.Unsigned32)},
					&wit.Alias{Name: "names", Kind: &wit.List{Elem: wit.NewReference("name")}},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "alias-world",
				Exports: []string{"types"},
			},
		},
	}
}

func complexRecord() wit.Wit {
	responseRef := wit.Reference("response")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "complex-record",
		},
		Interfaces: []wit.Interface{
			{
				Name: "data",
				TypeDefs: []wit.Type{
					&wit.Record{
						Name: "response",
						Fields: []wit.Field{
							{Name: "status", Kind: wit.NewPrimitive(wit.Unsigned32)},
							{
								Name: "body",
								Kind: &wit.Option{Inner: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
							},
							{Name: "error", Kind: &wit.Option{Inner: wit.NewPrimitive(wit.String)}},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "fetch",
						Params: []wit.Param{
							{Name: "url", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{Kind: &responseRef},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "data-world",
				Exports: []string{"data"},
			},
		},
	}
}

func wasiCli() wit.Wit {
	streamErrRef := wit.Reference("stream-error")
	exitCodeRef := wit.Reference("exit-code")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "wasi",
			Package:   "cli",
			Version:   "0.2.0",
		},
		Interfaces: []wit.Interface{
			{
				Name: "stdin",
				TypeDefs: []wit.Type{
					&wit.Record{
						Name: "stream-error",
						Fields: []wit.Field{
							{Name: "code", Kind: wit.NewPrimitive(wit.Unsigned32)},
							{Name: "message", Kind: wit.NewPrimitive(wit.String)},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "read",
						Params: []wit.Param{
							{Name: "len", Kind: wit.NewPrimitive(wit.Unsigned64)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)},
								Err: &streamErrRef,
							},
						},
					},
				},
			},
			{
				Name: "stdout",
				Functions: []wit.Function{
					{
						Name: "write",
						Params: []wit.Param{
							{Name: "buf", Kind: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  wit.NewPrimitive(wit.Unsigned64),
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
					{
						Name: "flush",
					},
				},
			},
			{
				Name: "environment",
				Functions: []wit.Function{
					{
						Name: "get-environment",
						Results: &wit.Param{
							Kind: &wit.List{
								Elem: &wit.Tuple{
									Fields: []wit.Type{
										wit.NewPrimitive(wit.String),
										wit.NewPrimitive(wit.String),
									},
								},
							},
						},
					},
					{
						Name: "get-arguments",
						Results: &wit.Param{
							Kind: &wit.List{Elem: wit.NewPrimitive(wit.String)},
						},
					},
					{
						Name: "initial-cwd",
						Results: &wit.Param{
							Kind: &wit.Option{Inner: wit.NewPrimitive(wit.String)},
						},
					},
				},
			},
			{
				Name: "exit",
				TypeDefs: []wit.Type{
					&wit.Enum{
						Name: "exit-code",
						Cases: []wit.Case{
							{Name: "success"},
							{Name: "failure"},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "exit",
						Params: []wit.Param{
							{Name: "code", Kind: &exitCodeRef},
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "command",
				Imports: []string{"stdin", "environment"},
				Exports: []string{"stdout", "exit"},
			},
			{
				Name:    "reactor",
				Imports: []string{"stdin", "stdout", "environment"},
			},
		},
	}
}

func kvStore() wit.Wit {
	errorKindRef := wit.Reference("error-kind")
	keyRef := wit.Reference("key")
	valueRef := wit.Reference("value")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "kv-store",
			Version:   "2.1.0",
		},
		Interfaces: []wit.Interface{
			{
				Name: "types",
				TypeDefs: []wit.Type{
					&wit.Alias{Name: "key", Kind: wit.NewPrimitive(wit.String)},
					&wit.Alias{Name: "value", Kind: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
					&wit.Flags{
						Name:  "consistency",
						Cases: []wit.Case{{Name: "strong"}, {Name: "eventual"}},
					},
					&wit.Enum{
						Name: "error-kind",
						Cases: []wit.Case{
							{Name: "not-found"},
							{Name: "permission-denied"},
							{Name: "internal"},
						},
					},
					&wit.Record{
						Name: "error",
						Fields: []wit.Field{
							{Name: "kind", Kind: &errorKindRef},
							{Name: "message", Kind: wit.NewPrimitive(wit.String)},
						},
					},
					&wit.Record{
						Name: "entry",
						Fields: []wit.Field{
							{Name: "key", Kind: &keyRef},
							{Name: "value", Kind: &valueRef},
							{Name: "version", Kind: wit.NewPrimitive(wit.Unsigned64)},
						},
					},
				},
			},
			{
				Name: "store",
				Functions: []wit.Function{
					{
						Name: "get",
						Params: []wit.Param{
							{Name: "key", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  &wit.Option{Inner: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
					{
						Name: "set",
						Params: []wit.Param{
							{Name: "key", Kind: wit.NewPrimitive(wit.String)},
							{Name: "value", Kind: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  wit.NewPrimitive(wit.Unsigned64),
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
					{
						Name: "delete",
						Params: []wit.Param{
							{Name: "key", Kind: wit.NewPrimitive(wit.String)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  wit.NewPrimitive(wit.Bool),
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
					{
						Name: "list-keys",
						Params: []wit.Param{
							{Name: "prefix", Kind: wit.NewPrimitive(wit.String)},
							{Name: "limit", Kind: wit.NewPrimitive(wit.Unsigned32)},
						},
						Results: &wit.Param{
							Kind: &wit.Result{
								Ok:  &wit.List{Elem: wit.NewPrimitive(wit.String)},
								Err: wit.NewPrimitive(wit.String),
							},
						},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "kv-client",
				Imports: []string{"store"},
			},
			{
				Name:    "kv-server",
				Exports: []string{"store"},
			},
		},
	}
}

func empty() wit.Wit {
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "foo",
			Package:   "foo",
		},
		Worlds: []wit.World{
			{
				Name: "empty",
			},
		},
	}
}

func docs() wit.Wit {
	responseRef := wit.Reference("response")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "docs",
		},
		Interfaces: []wit.Interface{
			{
				Name: "data",
				TypeDefs: []wit.Type{
					&wit.Enum{
						Name: "method",
						Cases: []wit.Case{
							{
								Name: "GET",
								Docs: wit.Docs{Content: "GET - get request"},
							},
							{
								Name: "POST",
								Docs: wit.Docs{Content: "POST - post request"},
							},
							{
								Name: "PUT",
								Docs: wit.Docs{Content: "PUT - put request"},
							},
							{
								Name: "DELETE",
								Docs: wit.Docs{Content: "DELETE - delete request"},
							},
						},
						Docs: wit.Docs{Content: "method is an enum for the request method"},
					},
					&wit.Record{
						Name: "response",
						Docs: wit.Docs{Content: "response is the response structure"},
						Fields: []wit.Field{
							{
								Name: "status",
								Kind: wit.NewPrimitive(wit.Unsigned32),
								Docs: wit.Docs{Content: "status is the http status code"},
							},
							{
								Name: "body",
								Kind: &wit.Option{Inner: &wit.List{Elem: wit.NewPrimitive(wit.Unsigned8)}},
								Docs: wit.Docs{Content: "body is an optional byte array"},
							},
							{
								Name: "error",
								Kind: &wit.Option{Inner: wit.NewPrimitive(wit.String)},
								Docs: wit.Docs{Content: "error is an optional error message"},
							},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name: "fetch",
						Params: []wit.Param{
							{Name: "url", Kind: wit.NewPrimitive(wit.String)},
							{Name: "method", Kind: wit.NewReference("method")},
						},
						Results: &wit.Param{Kind: &responseRef},
						Docs:    wit.Docs{Content: "fetch calls the url and responds with the response record"},
					},
				},
				Docs: wit.Docs{Content: "data is the main interface in the package `example:docs`\nthis is also a multi-line doc"},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "data-world",
				Exports: []string{"data"},
				Docs:    wit.Docs{Content: "data-world exports the data interface"},
			},
		},
	}
}

func variantTypes() wit.Wit {
	shapeRef := wit.Reference("shape")
	return wit.Wit{
		Package: &wit.Package{
			Namespace: "example",
			Package:   "variant",
		},
		Interfaces: []wit.Interface{
			{
				Name: "shapes",
				TypeDefs: []wit.Type{
					&wit.Variant{
						Name: "shape",
						Cases: []wit.Field{
							{Name: "circle", Kind: wit.NewPrimitive(wit.Float32)},
							{Name: "rectangle", Kind: &wit.Tuple{Fields: []wit.Type{wit.NewPrimitive(wit.Float32), wit.NewPrimitive(wit.Float32)}}},
							{Name: "point"},
						},
					},
				},
				Functions: []wit.Function{
					{
						Name:    "area",
						Params:  []wit.Param{{Name: "s", Kind: &shapeRef}},
						Results: &wit.Param{Kind: wit.NewPrimitive(wit.Float32)},
					},
				},
			},
		},
		Worlds: []wit.World{
			{
				Name:    "shapes-world",
				Exports: []string{"shapes"},
			},
		},
	}
}
