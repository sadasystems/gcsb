package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDB(t *testing.T) {
	Convey("DB", t, func() {
		// Do not connect to databases in unit tests
		//
		// Convey("NewDB", func() {
		// 	db, _ := NewDB(config.GCSBConfig{
		// 		Database: "gcsb-test-db-1",
		// 		Project:  "gcsb",
		// 		Instance: "gcsb-test-1",
		// 	})
		// 	fmt.Println(db)
		// 	err := db.GetDatabase()

		// 	fmt.Println(err)
		// })

	})
}
