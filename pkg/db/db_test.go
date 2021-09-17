package db

import (
	"fmt"
	"github.com/sadasystems/gcsb/pkg/config"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDB(t *testing.T) {
	Convey("DB", t, func() {
		Convey("NewDB", func() {
			db, _ := NewDB(config.GCSBConfig{
				Database: "gcsb-test-db-1",
				Project:  "gcsb",
				Instance: "gcsb-test-1",
			})
			fmt.Println(db)
			err := db.GetDatabase()

			fmt.Println(err)
		})

	})
}
