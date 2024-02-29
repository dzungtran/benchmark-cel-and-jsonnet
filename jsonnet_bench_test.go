package jsonnet_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/tidwall/gjson"
)

type JsonnetCase struct {
	Expr string
	In   any
	Out  string
}

var (
	ReferenceJsonnetCases = []JsonnetCase{
		{
			Expr: `inp.string_value == "value"`,
			In: map[string]any{
				"string_value": "value",
			},
			Out: "true",
		},
		{
			Expr: `std.member(["a","b","c","value"], inp.string_value)`,
			In: map[string]any{
				"string_value": "value",
			},
			Out: "true",
		},
		{
			Expr: `std.member(inp.array_value, inp.string_value)`,
			In: map[string]any{
				"string_value": "value",
				"array_value":  []string{"a", "b", "c", "value"},
			},
			Out: "true",
		},
		{
			Expr: `'The value of inp.string_value is %s.' % inp.string_value`,
			In: map[string]any{
				"string_value": "value",
				"array_value":  []string{"a", "b", "c", "value"},
			},
			Out: `The value of inp.string_value is value.`,
		},
	}

	tmplExprDoc = `
	local inp = std.parseJson('%s');
	{
		result: %s
	}`
)

func BenchmarkJsonnet(b *testing.B) {
	vm := jsonnet.MakeVM()

	for _, rc := range ReferenceJsonnetCases {
		b.Run(fmt.Sprintf("%s", rc.Expr), func(b *testing.B) {
			templateJsonnet := fmt.Sprintf(tmplExprDoc, ToJsonBytes(rc.In), rc.Expr)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				jsonStr, _ := vm.EvaluateAnonymousSnippet("result1.jsonnet", templateJsonnet)
				res := gjson.GetBytes([]byte(jsonStr), "result")
				if res.String() != rc.Out {
					b.Fatalf("not equal: %v - %v", rc.Out, res.String())
				}
			}
		})
	}
}

func ToJsonBytes(data any) []byte {
	b, _ := json.Marshal(data)
	return b
}
