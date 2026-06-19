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
