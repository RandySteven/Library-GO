package configs

import "time"

type Config struct {
	Config struct {
		Env    string `yaml:"env"`
		Server struct {
			Host    string `yaml:"host"`
			Port    string `yaml:"port"`
			Timeout struct {
				Server time.Duration `yaml:"server"`
				Read   time.Duration `yaml:"read"`
				Write  time.Duration `yaml:"write"`
				Idle   time.Duration `yaml:"idle"`
			} `yaml:"timeout"`
		} `yaml:"server"`

		MySQL struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
			ConnPool struct {
				MaxIdle   int `yaml:"maxIdle"`
				ConnLimit int `yaml:"connLimit"`
				IdleTime  int `yaml:"idleTime"`
			} `yaml:"connPool"`
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

		Mailtrap struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
	} `yaml:"app"`
}
