package information

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColumn(t *testing.T) {
	Convey("GetColumnsQuery", t, func() {
		st := GetColumnsQuery("foo")
		So(st, ShouldNotBeNil)
		So(st.Params["table_name"], ShouldEqual, "foo")
		// fmt.Println(st)
	})
}
