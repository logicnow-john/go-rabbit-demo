package config

import "github.com/spf13/viper"

type Config struct {
	RabbitUrl  string `mapstructure:"RABBIT_URL"`
	QueueName  string `mapstructure:"QUEUE_NAME"`
	RoutingKey string `mapstructure:"ROUTING_KEY"`
	Exchange   string `mapstructure:"EXCHANGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
