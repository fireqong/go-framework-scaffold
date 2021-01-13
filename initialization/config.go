package initialization

import "github.com/spf13/viper"

func Config() {
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err.Error())
	}
}
