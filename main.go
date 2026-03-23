package main

import (
  greetingv1 "hexlet-go/greeting"
  greetingv2 "hexlet-go/greeting/v2"
  "fmt"
)

func main() {
  fmt.Println("Первое приветствие: ", greetingv1.Get(), "\n", "Второе приветствие: ", greetingv2.Get())
}
