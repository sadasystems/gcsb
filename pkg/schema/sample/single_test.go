package sample

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleSampleSchema(t *testing.T) {
	Convey("singleTableSingleKey", t, func() {
		want := `CREATE TABLE Singers (
  SingerId INT64 NOT NULL,
  FirstName STRING(1024),
  LastName STRING(1024),
  BirthDate DATE,
  ByteField BYTES(1024),
  FloatField FLOAT64,
  ArrayField ARRAY<INT64>,
  TSField TIMESTAMP,
  NumericField NUMERIC,
) PRIMARY KEY(SingerId)`
		So(singleTableSingleKey.SQL(), ShouldEqual, want)
	})
}
