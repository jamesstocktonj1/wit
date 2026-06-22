package integration

import "github.com/jamesstocktonj1/wit"

func basic() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("wasi", "test")).
		WithInterface(
			wit.NewInterface("greeting",
				wit.NewRecord("message",
					wit.NewField("name", wit.String),
					wit.NewField("message", wit.String),
				),
			).WithFunctions(
				wit.NewFunction("greet", wit.NewParam("name", wit.String)).
					WithResult(wit.NewReference("message")),
			),
		).
		WithWorld(
			wit.NewWorld("greeter").WithExports(wit.NewInterfaceReference("greeting")),
		)
}

func listTypes() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "lists")).
		WithInterface(
			wit.NewInterface("collection",
				wit.NewRecord("point",
					wit.NewField("x", wit.Float32),
					wit.NewField("y", wit.Float32),
				),
			).WithFunctions(
				wit.NewFunction("sum", wit.NewParam("values", wit.NewList(wit.Signed32))).
					WithResult(wit.Signed32),
				wit.NewFunction("join", wit.NewParam("parts", wit.NewList(wit.String))).
					WithResult(wit.String),
			),
		).
		WithWorld(
			wit.NewWorld("collection-world").WithExports(wit.NewInterfaceReference("collection")),
		)
}

func optionResult() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "option-result")).
		WithInterface(
			wit.NewInterface("maybe").WithFunctions(
				wit.NewFunction("find", wit.NewParam("key", wit.String)).
					WithResult(wit.NewOption(wit.String)),
				wit.NewFunction("divide", wit.NewParam("numerator", wit.Float64)).
					WithResult(wit.NewResult(wit.Float64, wit.String)),
			),
		).
		WithWorld(
			wit.NewWorld("maybe-world").WithExports(wit.NewInterfaceReference("maybe")),
		)
}

func enumFlags() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "enum-flags")).
		WithInterface(
			wit.NewInterface("permissions",
				wit.NewFlags("access",
					wit.NewCase("read"),
					wit.NewCase("write"),
					wit.NewCase("execute"),
				),
				wit.NewEnum("direction",
					wit.NewCase("north"),
					wit.NewCase("south"),
					wit.NewCase("east"),
					wit.NewCase("west"),
				),
			).WithFunctions(
				wit.NewFunction("can-access", wit.NewParam("dir", wit.NewReference("direction"))).
					WithResult(wit.Bool),
			),
		).
		WithWorld(
			wit.NewWorld("permissions-world").WithExports(wit.NewInterfaceReference("permissions")),
		)
}

func multiInterface() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("wasi", "io").WithVersion("0.2.0")).
		WithInterface(
			wit.NewInterface("input").WithFunctions(
				wit.NewFunction("read", wit.NewParam("len", wit.Unsigned32)).
					WithResult(wit.NewList(wit.Unsigned8)),
			),
			wit.NewInterface("output").WithFunctions(
				wit.NewFunction("write", wit.NewParam("buf", wit.NewList(wit.Unsigned8))).
					WithResult(wit.Unsigned32),
			),
		).
		WithWorld(
			wit.NewWorld("streams").WithImports(wit.NewInterfaceReference("input")).WithExports(wit.NewInterfaceReference("output")),
		)
}

func multiWorld() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "worlds")).
		WithInterface(
			wit.NewInterface("logger").WithFunctions(
				wit.NewFunction("log", wit.NewParam("msg", wit.String)),
			),
		).
		WithWorld(
			wit.NewWorld("server").WithExports(wit.NewInterfaceReference("logger")),
			wit.NewWorld("client").WithImports(wit.NewInterfaceReference("logger")),
		)
}

func nestedTypes() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "nested")).
		WithInterface(
			wit.NewInterface("deep").WithFunctions(
				wit.NewFunction("flatten", wit.NewParam("matrix", wit.NewList(wit.NewList(wit.Signed32)))).
					WithResult(wit.NewList(wit.Signed32)),
				wit.NewFunction("lookup", wit.NewParam("key", wit.String)).
					WithResult(wit.NewOption(wit.NewList(wit.String))),
				wit.NewFunction("parse", wit.NewParam("input", wit.String)).
					WithResult(wit.NewResult(
						wit.NewOption(wit.Unsigned64),
						wit.String,
					)),
			),
		).
		WithWorld(
			wit.NewWorld("nested-world").WithExports(wit.NewInterfaceReference("deep")),
		)
}

func tupleTypes() wit.Wit {
	boundsRef := wit.NewReference("bounds")
	pair := wit.NewTuple(wit.Float64, wit.Float64)
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "tuples")).
		WithInterface(
			wit.NewInterface("coords",
				wit.NewRecord("bounds",
					wit.NewField("min", pair),
					wit.NewField("max", pair),
				),
			).WithFunctions(
				wit.NewFunction("midpoint", wit.NewParam("b", boundsRef)).
					WithResult(pair),
			),
		).
		WithWorld(
			wit.NewWorld("coords-world").WithExports(wit.NewInterfaceReference("coords")),
		)
}

func aliasTypes() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "alias")).
		WithInterface(
			wit.NewInterface("types",
				wit.NewAlias("name", wit.String),
				wit.NewAlias("age", wit.Unsigned32),
				wit.NewAlias("names", wit.NewList(wit.NewReference("name"))),
			),
		).
		WithWorld(
			wit.NewWorld("alias-world").WithExports(wit.NewInterfaceReference("types")),
		)
}

func complexRecord() wit.Wit {
	responseRef := wit.NewReference("response")
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "complex-record")).
		WithInterface(
			wit.NewInterface("data",
				wit.NewRecord("response",
					wit.NewField("status", wit.Unsigned32),
					wit.NewField("body", wit.NewOption(wit.NewList(wit.Unsigned8))),
					wit.NewField("error", wit.NewOption(wit.String)),
				),
			).WithFunctions(
				wit.NewFunction("fetch", wit.NewParam("url", wit.String)).
					WithResult(responseRef),
			),
		).
		WithWorld(
			wit.NewWorld("data-world").WithExports(wit.NewInterfaceReference("data")),
		)
}

func wasiCli() wit.Wit {
	streamErrRef := wit.NewReference("stream-error")
	exitCodeRef := wit.NewReference("exit-code")
	return wit.NewWit().
		WithPackage(wit.NewPackage("wasi", "cli").WithVersion("0.2.0")).
		WithInterface(
			wit.NewInterface("stdin",
				wit.NewRecord("stream-error",
					wit.NewField("code", wit.Unsigned32),
					wit.NewField("message", wit.String),
				),
			).WithFunctions(
				wit.NewFunction("read", wit.NewParam("len", wit.Unsigned64)).
					WithResult(wit.NewResult(
						wit.NewList(wit.Unsigned8),
						streamErrRef,
					)),
			),
			wit.NewInterface("stdout").WithFunctions(
				wit.NewFunction("write", wit.NewParam("buf", wit.NewList(wit.Unsigned8))).
					WithResult(wit.NewResult(
						wit.Unsigned64,
						wit.String,
					)),
				wit.NewFunction("flush"),
			),
			wit.NewInterface("environment").WithFunctions(
				wit.NewFunction("get-environment").
					WithResult(wit.NewList(wit.NewTuple(
						wit.String,
						wit.String,
					))),
				wit.NewFunction("get-arguments").
					WithResult(wit.NewList(wit.String)),
				wit.NewFunction("initial-cwd").
					WithResult(wit.NewOption(wit.String)),
			),
			wit.NewInterface("exit",
				wit.NewEnum("exit-code",
					wit.NewCase("success"),
					wit.NewCase("failure"),
				),
			).WithFunctions(
				wit.NewFunction("exit", wit.NewParam("code", exitCodeRef)),
			),
		).
		WithWorld(
			wit.NewWorld("command").WithImports(wit.NewInterfaceReference("stdin"), wit.NewInterfaceReference("environment")).WithExports(wit.NewInterfaceReference("stdout"), wit.NewInterfaceReference("exit")),
			wit.NewWorld("reactor").WithImports(wit.NewInterfaceReference("stdin"), wit.NewInterfaceReference("stdout"), wit.NewInterfaceReference("environment")),
		)
}

func kvStore() wit.Wit {
	errorKindRef := wit.NewReference("error-kind")
	keyRef := wit.NewReference("key")
	valueRef := wit.NewReference("value")
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "kv-store").WithVersion("2.1.0")).
		WithInterface(
			wit.NewInterface("types",
				wit.NewAlias("key", wit.String),
				wit.NewAlias("value", wit.NewList(wit.Unsigned8)),
				wit.NewFlags("consistency",
					wit.NewCase("strong"),
					wit.NewCase("eventual"),
				),
				wit.NewEnum("error-kind",
					wit.NewCase("not-found"),
					wit.NewCase("permission-denied"),
					wit.NewCase("internal"),
				),
				wit.NewRecord("error",
					wit.NewField("kind", errorKindRef),
					wit.NewField("message", wit.String),
				),
				wit.NewRecord("entry",
					wit.NewField("key", keyRef),
					wit.NewField("value", valueRef),
					wit.NewField("version", wit.Unsigned64),
				),
			),
			wit.NewInterface("store").WithFunctions(
				wit.NewFunction("get", wit.NewParam("key", wit.String)).
					WithResult(wit.NewResult(
						wit.NewOption(wit.NewList(wit.Unsigned8)),
						wit.String,
					)),
				wit.NewFunction("set",
					wit.NewParam("key", wit.String),
					wit.NewParam("value", wit.NewList(wit.Unsigned8)),
				).WithResult(wit.NewResult(
					wit.Unsigned64,
					wit.String,
				)),
				wit.NewFunction("delete", wit.NewParam("key", wit.String)).
					WithResult(wit.NewResult(
						wit.Bool,
						wit.String,
					)),
				wit.NewFunction("list-keys",
					wit.NewParam("prefix", wit.String),
					wit.NewParam("limit", wit.Unsigned32),
				).WithResult(wit.NewResult(
					wit.NewList(wit.String),
					wit.String,
				)),
			),
		).
		WithWorld(
			wit.NewWorld("kv-client").WithImports(wit.NewInterfaceReference("store")),
			wit.NewWorld("kv-server").WithExports(wit.NewInterfaceReference("store")),
		)
}

func empty() wit.Wit {
	return wit.NewWit().
		WithPackage(wit.NewPackage("foo", "foo")).
		WithWorld(
			wit.NewWorld("empty"),
		)
}

func docs() wit.Wit {
	responseRef := wit.NewReference("response")
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "docs")).
		WithInterface(
			wit.NewInterface("data",
				wit.NewEnum("method",
					wit.NewCase("GET").WithDocs("GET - get request"),
					wit.NewCase("POST").WithDocs("POST - post request"),
					wit.NewCase("PUT").WithDocs("PUT - put request"),
					wit.NewCase("DELETE").WithDocs("DELETE - delete request"),
				).WithDocs("method is an enum for the request method"),
				wit.NewRecord("response",
					wit.NewField("status", wit.Unsigned32).WithDocs("status is the http status code"),
					wit.NewField("body", wit.NewOption(wit.NewList(wit.Unsigned8))).WithDocs("body is an optional byte array"),
					wit.NewField("error", wit.NewOption(wit.String)).WithDocs("error is an optional error message"),
				).WithDocs("response is the response structure"),
			).WithFunctions(
				wit.NewFunction("fetch",
					wit.NewParam("url", wit.String),
					wit.NewParam("method", wit.NewReference("method")),
				).WithResult(responseRef).
					WithDocs("fetch calls the url and responds with the response record"),
			).WithDocs("data is the main interface in the package `example:docs`\nthis is also a multi-line doc"),
		).
		WithWorld(
			wit.NewWorld("data-world").WithExports(wit.NewInterfaceReference("data")).WithDocs("data-world exports the data interface"),
		)
}

func variantTypes() wit.Wit {
	shapeRef := wit.NewReference("shape")
	return wit.NewWit().
		WithPackage(wit.NewPackage("example", "variant")).
		WithInterface(
			wit.NewInterface("shapes",
				wit.NewVariant("shape",
					wit.NewField("circle", wit.Float32),
					wit.NewField("rectangle", wit.NewTuple(wit.Float32, wit.Float32)),
					wit.NewField("point", nil),
				),
			).WithFunctions(
				wit.NewFunction("area", wit.NewParam("s", shapeRef)).
					WithResult(wit.Float32),
			),
		).
		WithWorld(
			wit.NewWorld("shapes-world").WithExports(wit.NewInterfaceReference("shapes")),
		)
}
