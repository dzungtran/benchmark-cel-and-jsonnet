package jsonnet_test

import (
	_ "embed"
	"fmt"
	"log"
	"testing"

	"github.com/google/go-jsonnet"
)

var (
	//go:embed data.json
	dataJson []byte

	//go:embed template.jsonnet
	templateJsonnet string

	//go:embed strings.libsonnet
	stringsJsonnet []byte
)

func TestLoadJsonnetLib(t *testing.T) {
	vm := jsonnet.MakeVM()
	// vm.TLAVar("__params", "import 'data.json'")
	vm.Importer(&jsonnet.MemoryImporter{
		Data: map[string]jsonnet.Contents{
			"data.json":         jsonnet.MakeContentsRaw(dataJson),
			"strings.libsonnet": jsonnet.MakeContentsRaw(stringsJsonnet),
		},
	})

	jsonStr, err := vm.EvaluateAnonymousSnippet("example1.jsonnet", templateJsonnet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonStr)
}
