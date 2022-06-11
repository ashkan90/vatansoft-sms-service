package app

type Configs struct {
	Application ApplicationConfigs `mapstructure:"application"`
}

type ApplicationConfigs struct {
	Mobilisim MobilisimConfig `mapstructure:"mobilisim"`
	Rabbit    RabbitConfig    `mapstructure:"rabbitmq"`
}

type MobilisimConfig struct {
	URL    string `mapstructure:"url"`
	ApiKey string `mapstructure:"api-key"`
}

type RabbitConfig struct {
	URL          string `mapstructure:"server-url"`
	QueueName    string `mapstructure:"sms-queue"`
	EventBusName string `mapstructure:"event-bus"`
}
