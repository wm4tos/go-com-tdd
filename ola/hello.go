package main

import "fmt"

const prefixEnglish = "Hello, "
const prefixPortuguese = "Olá, "
const prefixFrench = "Bonjour, "

const portuguese = "portuguese"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == portuguese {
		return prefixPortuguese + name + "!"
	}

	if language == french {
		return prefixFrench + name + "!"
	}

	return prefixEnglish + name + "!"
}

func main() {
	fmt.Println(Hello("Wes", ""))
}
