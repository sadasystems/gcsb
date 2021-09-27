package config

import "github.com/spf13/viper"

func Bind(v *viper.Viper) {
	v.BindEnv("connection.project", "PROJECT")
}
