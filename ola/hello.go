package ola

import "fmt"

const portuguese = "portuguese"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getPreffix(language) + name + "!"
}

func getPreffix(language string) (preffix string) {
	const preffixEnglish = "Hello, "
	const preffixPortuguese = "Olá, "
	const preffixFrench = "Bonjour, "

	switch language {
	case french:
		preffix = preffixFrench
	case portuguese:
		preffix = preffixPortuguese
	default:
		preffix = preffixEnglish
	}

	return
}

func main() {
	fmt.Println(Hello("Wes", ""))
}
