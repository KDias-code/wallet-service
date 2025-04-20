package configs

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
	Db   string `mapstructure:"DB"`
}

func LoadConfigs() (*Config, error) {
	conf := new(Config)

	v := viper.New()
	v.AutomaticEnv()

	err := v.BindEnv("PORT")
	if err != nil {
		return nil, err
	}
	err = v.BindEnv("DB")
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
