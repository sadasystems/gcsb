package config

import (
	"runtime"

	"github.com/spf13/viper"
)

// SetDefaults takes a viper instance and sets default configuration values
func SetDefaults(v *viper.Viper) {
	// Defaults
	v.SetDefault("num_conns", runtime.GOMAXPROCS(0))
	v.SetDefault("threads", 10)

	// Operations defualts
	v.SetDefault("operations.total", 10000)
	v.SetDefault("operations.read", 50)
	v.SetDefault("operations.write", 50)
	v.SetDefault("operations.sample_size", 50)
	v.SetDefault("operations.read_stale", false)

	// Pool Defaults
	v.SetDefault("pool.max_opened", 1000)
	v.SetDefault("pool.min_opened", 100)
	v.SetDefault("pool.max_idle", 0)
	v.SetDefault("pool.write_sessions", 0.2)
	v.SetDefault("pool.healthcheck_workers", 10)
	v.SetDefault("pool.healthcheck_interval", "50m")
	v.SetDefault("pool.track_session_handles", false)
}
