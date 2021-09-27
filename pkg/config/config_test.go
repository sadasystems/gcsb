package config

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var (
	cfgExample = []byte(`
project:  test-123 # GCP Project ID
instance: gcsb-test-1 # Spanner Instance ID
database: gcsb-test-db-1 # Spanner Database Name
num_conns: 10
`)
)

func readConfig(c []byte) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	// Read the config
	err := v.ReadConfig(bytes.NewBuffer(c))

	return v, err
}

func TestConfig(t *testing.T) {
	Convey("Config", t, func() {
		Convey("NewConfig", func() {
			Convey("Valid", func() {
				v := viper.New()
				v.SetConfigType("yaml")
				// Read the config
				err := v.ReadConfig(bytes.NewBuffer(cfgExample))
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
				So(err, ShouldBeNil)
				So(c, ShouldNotBeNil)

				vErr := c.Validate()
				So(vErr, ShouldNotBeNil)
			})

			Convey("DSN", func() {
				v, err := readConfig(cfgExample)
				So(err, ShouldBeNil)
				So(v, ShouldNotBeNil)

				// Unmarshal the config
				var c Config

				err = v.Unmarshal(&c)
				So(err, ShouldBeNil)

				So(c.DB(), ShouldEqual, fmt.Sprintf("projects/%s/instances/%s/database/%s", c.Project, c.Instance, c.Database))
			})
		})
	})
}
