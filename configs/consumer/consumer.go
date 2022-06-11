package consumer

type Config struct {
	URL          string `mapstructure:"server-url"`
	QueueName    string `mapstructure:"sms-queue"`
	EventBusName string `mapstructure:"event-bus"`
}
