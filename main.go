package main

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	text := color.RedString("Hello, Hexlet!")
	logrus.Println(text)
}