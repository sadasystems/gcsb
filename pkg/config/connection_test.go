package config

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hashicorp/go-multierror"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

var (
	dsnOnly = []byte(`
connection:
  project:  test-123 # GCP Project ID
  instance: gcsb-test-1 # Spanner Instance ID
  database: gcsb-test-db-1 # Spanner Database Name
`)
	cfgExample = []byte(`
connection:
  project:  test-123 # GCP Project ID
  instance: gcsb-test-1 # Spanner Instance ID
  database: gcsb-test-db-1 # Spanner Database Name
  num_conns: 10
  pool:
    max_opened: 1000
    min_opened: 100
    max_idle: 0
    write_sessions: 0.2
    healthcheck_workers: 10
    healthcheck_interval: 50m
    track_session_handles: false
`)
)

func readConfig(c []byte) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	// Read the config
	err := v.ReadConfig(bytes.NewBuffer(c))

	return v, err
}

func TestConnection(t *testing.T) {
	Convey("Connection", t, func() {
		Convey("Unmarshal", func() {
			v, err := readConfig(cfgExample)
			So(err, ShouldBeNil)
			So(v, ShouldNotBeNil)

			// Unmarshal the config
			var c Config

			err = v.Unmarshal(&c)
			So(err, ShouldBeNil)

			So(c.Connection, ShouldNotBeNil)
			So(c.Connection.Project, ShouldEqual, "test-123")
			So(c.Connection.Instance, ShouldEqual, "gcsb-test-1")
			So(c.Connection.Database, ShouldEqual, "gcsb-test-db-1")
			So(c.Connection.NumConns, ShouldEqual, 10)
			So(c.Connection.Pool, ShouldNotBeNil)
			So(c.Connection.Pool.MinOpened, ShouldEqual, 100)
		})

		Convey("Viper Defaults work on Unmarshal", func() {
			// Ensure that Unmarshal adequately factors viper.SetDefault
			v, err := readConfig(dsnOnly)
			So(err, ShouldBeNil)
			So(v, ShouldNotBeNil)

			// Unmarshal the config
			var c Config

			// Unmarshal with no defaults
			err = v.Unmarshal(&c)
			So(err, ShouldBeNil)
			So(c.Connection.NumConns, ShouldEqual, 0)

			// Set some defaults
			v.SetDefault("connection.num_conns", 100)

			// Unmarshal the config
			var c2 Config

			// Unmarshal again
			err = v.Unmarshal(&c2)
			So(err, ShouldBeNil)
			So(c2.Connection.NumConns, ShouldEqual, 100)
		})

		Convey("Validate", func() {
			v, err := readConfig([]byte(``))
			So(err, ShouldBeNil)
			So(v, ShouldNotBeNil)

			// Unmarshal the config
			var c Config

			err = v.Unmarshal(&c)
			So(err, ShouldBeNil)
			So(c.Connection.Project, ShouldEqual, "")

			errs := c.Validate()
			So(errs, ShouldNotBeNil)

			merr, ok := errs.(*multierror.Error)
			So(ok, ShouldBeTrue)
			So(merr, ShouldNotBeNil)
			So(merr.Len(), ShouldEqual, 3) // we expect to see 3 errors
		})

		Convey("DSN", func() {
			v, err := readConfig(cfgExample)
			So(err, ShouldBeNil)
			So(v, ShouldNotBeNil)

			// Unmarshal the config
			var c Config

			err = v.Unmarshal(&c)
			So(err, ShouldBeNil)

			So(c.Connection.DB(), ShouldEqual, fmt.Sprintf("projects/%s/instances/%s/database/%s", c.Connection.Project, c.Connection.Instance, c.Connection.Database))
		})
	})
}
