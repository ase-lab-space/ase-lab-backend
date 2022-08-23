package config

func NewConfigMock() *Config {
	return &Config{
		PORT:                                  1234,
		GIN_MODE:                              "test",
		CONTACT_FORM_NOTIFICATOR_ACCESS_TOKEN: "ACCESS_TOKEN",
	}
}
