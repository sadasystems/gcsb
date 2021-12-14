package schema

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSchema(t *testing.T) {
	Convey("Schema", t, func() {
		Convey("Getters/Setters", func() {
			s := NewSchema()
			So(s, ShouldNotBeNil)

			t1 := NewTable()
			t1.SetName("single_table")
			s.SetTable(t1)
			So(s.Table().Name(), ShouldEqual, "single_table")

			for i := 1; i <= 3; i++ {
				t := NewTable()
				t.SetName(fmt.Sprintf("table_%d", i))
				s.AddTable(t)
			}

			So(s.Tables(), ShouldNotBeNil)
			So(s.Tables().Len(), ShouldEqual, 3)
			So(s.GetTable("table_1"), ShouldNotBeNil)
			So(s.GetTable("table_1").Name(), ShouldEqual, "table_1")

			So(s.GetTable("table_1").IsApex(), ShouldBeTrue)

			t1 = s.GetTable("table_1")
			t2 := s.GetTable("table_2")
			t3 := s.GetTable("table_3")
			So(t1, ShouldNotBeNil)
			So(t2, ShouldNotBeNil)
			So(t3, ShouldNotBeNil)

			t3.SetParentName(t2.Name())
			t2.SetParentName(t1.Name())

			err := s.Traverse()
			So(err, ShouldBeNil)
			So(t3.IsApex(), ShouldBeFalse)
			So(t2.IsApex(), ShouldBeFalse)
			So(t1.IsApex(), ShouldBeTrue)

			// Child links should be set as well
			So(t1.HasChild(), ShouldBeTrue)
			So(t1.Child().Name(), ShouldEqual, t2.Name())
			So(t2.HasChild(), ShouldBeTrue)
			So(t2.Child().Name(), ShouldEqual, t3.Name())
			So(t3.HasChild(), ShouldBeFalse)

			So(t3.IsBottom(), ShouldBeTrue)

			// Missing parent tables should return error
			// I actually don't even know if this is possible but it's handled
			t4 := NewTable()
			t4.SetName("table_4")
			t4.SetParentName("NOT_EXIST")
			s.AddTable(t4)

			err = s.Traverse()
			So(err, ShouldNotBeNil)

			// Non-interleaved
			t5 := NewTable()
			t5.SetName("table_5")
			So(t5.IsInterleaved(), ShouldBeFalse)
		})
	})
}