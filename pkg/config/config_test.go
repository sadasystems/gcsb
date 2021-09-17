package config

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	Convey("Config", t, func() {
		Convey("NewConfig", func() {
			Convey("Valid", func() {
				v := viper.New()
				v.SetConfigType("yaml")

				// Read the config
				err := v.ReadConfig(bytes.NewBuffer(
					[]byte(`
connection:
  project:  test-123 # GCP Project ID
  instance: gcsb-test-1 # Spanner Instance ID
  database: gcsb-test-db-1 # Spanner Database Name
`)))
				So(err, ShouldBeNil)

				c, err := NewConfig(v)
				So(err, ShouldBeNil)
				So(c, ShouldNotBeNil)
			})

			Convey("Invalid Configuration", func() {
				v := viper.New()
				v.SetConfigType("yaml")

				// Read the config
				err := v.ReadConfig(bytes.NewBuffer([]byte(``)))
				So(err, ShouldBeNil)

				c, err := NewConfig(v)
				So(err, ShouldNotBeNil)
				So(c, ShouldBeNil)
			})
		})
	})
}
