package model

import (
	"strings"
)

var (
	Sharp = map[string]string{
		"spot": "C",
		"swap": "S",
		"future": "F",
		"option": "O",
		"aoption": "A",
		"isolated": "I",
		"margin": "M",
	}
)

func GetSharp(name string) string {
	return Sharp[strings.ToLower(name)]
}

