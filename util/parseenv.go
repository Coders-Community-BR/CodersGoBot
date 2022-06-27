package util

import (
	"os"
)

func GetEnv(name string) string {
	env := os.Getenv(name)
	if env == "foo" || env == "" {
		panic("Please set a value to $" + name + "!")
	}
	return env
}
