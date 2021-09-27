package config

import (
	"runtime"

	"github.com/spf13/viper"
)

// SetDefaults takes a viper instance and sets default configuration values
func SetDefaults(v *viper.Viper) {
	// Defaults
	v.SetDefault("num_conns", runtime.GOMAXPROCS(0))

	// Pool Defaults
	v.SetDefault("pool.healthcheck_workers", 10)
	v.SetDefault("pool.healthcheck_interval", "50m")
	v.SetDefault("pool.track_session_handles", false)
}
