package config

import "strconv"

type Config struct {
	Port int    `default:"development" env:"WAFFLR_PORT"`
	Addr string `default:"development" env:"WAFFLR_ADDR"`
	Db   struct {
		Url  string `default:"mongodb://localhost:27017" env:"WAFFLR_DB_URL"`
		Name string `default:"wafflr_db" env:"WAFFLR_DB_NAME"`
	}
}

func (c *Config) Validate() error {
	return nil
}

func (c *Config) ListenAddr() string {
	return c.Addr + ":" + strconv.Itoa(c.Port)
}
