package fs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldReadFileIntoString(t *testing.T) {
	Convey("Should read file into string", t, func() {
		if data, err := ReadFileString("../../examples/request.toml"); err != nil {
			t.Fail()
		} else if len(data) == 0 {
			t.Fail()
		}
	})

	Convey("Should not panic when file not exist", t, func() {
		if data, err := ReadFileString("../../examples/request1.toml"); err == nil {
			t.Fail()
		} else if len(data) > 0 {
			t.Fail()
		}
	})
}
