package selfbot

type Config struct {
	Token string
}

func NewConfigDefault(token string) Config {
	config := Config{
		Token: token,
	}

	return config
}
