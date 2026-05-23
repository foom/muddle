package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Name    string            `yaml:"name"`
	Host    string            `yaml:"host"`
	Port    string            `yaml:"port"`
	Aliases map[string]string `yaml:"aliases"`
}

func loadProfile(path string) (Profile, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := createDefaultProfile(path)
		if err != nil {
			return Profile{}, err
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Profile{}, err
	}

	var profile Profile

	err = yaml.Unmarshal(data, &profile)
	if err != nil {
		return Profile{}, err
	}

	if profile.Aliases == nil {
		profile.Aliases = make(map[string]string)
	}

	return profile, nil
}

func createDefaultProfile(path string) error {
	err := os.MkdirAll("profiles", 0755)
	if err != nil {
		return err
	}

	defaultProfile := Profile{
		Name: "Default",
		Host: "",
		Port: "",
		Aliases: map[string]string{
			"muddle": "emote muddles around!",
		},
	}

	data, err := yaml.Marshal(defaultProfile)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func saveProfile(path string, profile Profile) error {
	data, err := yaml.Marshal(profile)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
