package configs

type Config struct {
	Config struct {
		Server struct {
			Host    string `yaml:"host"`
			Port    string `yaml:"port"`
			Timeout struct {
				Server int `yaml:"server"`
				Read   int `yaml:"read"`
				Write  int `yaml:"write"`
				Idle   int `yaml:"idle"`
			} `yaml:"timeout"`
		} `yaml:"server"`

		MySQL struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`

		AWS struct {
			AccessKeyID     string `yaml:"accessKeyID"`
			SecretAccessKey string `yaml:"secretAccessKey"`
			Region          string `yaml:"region"`
		} `yaml:"aws"`

		Redis struct {
			Host          string `yaml:"host"`
			Port          string `yaml:"port"`
			Password      string `yaml:"password"`
			MinIddleConns int    `yaml:"minIddleConns"`
			MaxIddleConns int    `yaml:"maxIddleConns"`
			PoolSize      int    `yaml:"poolSize"`
			PoolTimeout   int    `yaml:"poolTimeout"`
			DB            int    `yaml:"db"`
		} `yaml:"redis"`

		Algolia struct {
			AppID  string `yaml:"appID"`
			ApiKey string `yaml:"apiKey"`
		} `yaml:"algolia"`
	} `yaml:"app"`
}
