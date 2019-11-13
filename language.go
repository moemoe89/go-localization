// Package go-localization will converted some string message to another language that defined before.
package language

import (
	"encoding/json"
	"io/ioutil"
)

// map variable for store the language messages.
var langs = make(map[string](map[string]string))

// Config has two property: path for save the config file json path and mainLocale for default language code.
type Config struct {
	path       string
	mainLocale string
}

// Create a new Config with default value.
func New() *Config {
	cfg := new(Config)
	return cfg
}

// Store the config file json path value.
func (c *Config) BindPath(path string) {
	c.path = path
}

// Store the default language code value.
func (c *Config) BindMainLocale(mainLocale string) {
	c.mainLocale = mainLocale
}

// Store the message to map variable.
func (c Config) Init() (*Config, error) {

	file := c.path
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &langs)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Lookup the destination message based on language code & message.
func (c Config) Lookup(locale, msg string) string {

	if len(langs[locale]) < 1 {
		locale = c.mainLocale
	}

	if len(langs[locale][msg]) < 1 {
		return msg
	}

	return langs[locale][msg]
}
