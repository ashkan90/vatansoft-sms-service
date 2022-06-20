package consumer

type Configs struct {
	Consumer Application `mapstructure:"application"`
}

type Application struct {
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
