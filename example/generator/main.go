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
		WithPackage(wit.NewPackage("foo", "bar").WithVersion("0.1.2-rc1")).
		WithWorld(
			wit.NewWorld("foo").
				WithImports("handler").
				WithExports("handler"),
		).
		WithInterface(
			wit.NewInterface("handler",
				wit.NewRecord("foo",
					wit.NewField("foo", wit.NewPrimitive(wit.Char)),
				).WithDocs("foo is a record with a nested record"),
			).WithFunctions(
				wit.NewFunction("handler", wit.NewParam("name", wit.NewPrimitive(wit.String))).
					WithResult(wit.NewParam("", wit.NewPrimitive(wit.String))),
			),
		)

	err = wit.NewEncoder(f).Encode(w)
	if err != nil {
		log.Fatal(err)
	}
}
