package config

import "github.com/spf13/viper"

// SetDefaults takes a viper instance and sets default configuration values
func SetDefaults(v *viper.Viper) {
	// Default Connection
	v.SetDefault("connection.num_conns", 10)
	// v.SetDefault("")
}
