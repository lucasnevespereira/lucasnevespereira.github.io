package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type SiteConfig struct {
	Name    string  `yaml:"name"`
	Bio     string  `yaml:"bio"`
	Picture string  `yaml:"picture"`
	Meta    Meta    `yaml:"meta"`
	Links   []Link  `yaml:"links"`
	Contact Contact `yaml:"contact"`
	Theme   string  `yaml:"theme"`
}

type Contact struct {
	Email        string `yaml:"email"`
	Phone        string `yaml:"phone"`
	BuyMeACoffee string `yaml:"buymeacoffee"`
	CV           string `yaml:"cv"`
	Socials      []Link `yaml:"socials"`
}

type Meta struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Lang        string `yaml:"lang"`
	Author      string `yaml:"author"`
	SiteUrl     string `yaml:"siteUrl"`
}

type Link struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func LoadSiteConfig(path string) (*SiteConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config SiteConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
