package toml

import (
	"fmt"
	"testing"

	"github.com/PMoneda/pruu/infra/fs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldParseTomlFileIntoModelsRequest(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Should parse toml file into Request structure", t, func() {
		content, err := fs.ReadFileString("../../examples/request.toml")
		if err != nil {
			t.Fail()
		}
		parser := new(Parser)
		if req, err := parser.Parse(content); err != nil {
			t.Fail()
		} else {
			if req.Name != "post_data" {
				t.Fail()
			}
			if req.Method != "post" {
				t.Fail()
			}
			if value, ok := req.Headers["Content-Type"]; !ok {
				t.Fail()
			} else if value != "application/json" {
				t.Fail()
			}

			if req.Body.Content == "" {
				t.Fail()
			}

			if req.Response.Template == "" {
				t.Fail()
			}

			if req.Output.Stdout == "" {
				t.Fail()
			}
			fmt.Println(req)
		}

	})
}
