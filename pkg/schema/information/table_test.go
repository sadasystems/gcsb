package information

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTable(t *testing.T) {
	Convey("ListTablesQuery", t, func() {
		So(ListTablesQuery().SQL, ShouldEqual, `SELECT TABLE_CATALOG, TABLE_SCHEMA, TABLE_NAME, TABLE_TYPE, PARENT_TABLE_NAME, ON_DELETE_ACTION, SPANNER_STATE FROM information_schema.tables WHERE table_catalog = "" AND table_schema = "" ORDER BY table_catalog, table_schema, table_name`)
	})

	Convey("GetTableQuery", t, func() {
		st := GetTableQuery("foo")
		So(st, ShouldNotBeNil)
		So(st.Params["table_name"], ShouldEqual, "foo")

		// fmt.Println(st)
	})
}
