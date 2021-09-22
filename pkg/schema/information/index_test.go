package information

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIndex(t *testing.T) {
	Convey("GetIndexesQuery", t, func() {
		st := GetIndexesQuery("foo")
		So(st, ShouldNotBeNil)
		So(st.Params["table_name"], ShouldEqual, "foo")

		// fmt.Println(st)
	})
}
