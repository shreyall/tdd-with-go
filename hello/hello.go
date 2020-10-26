package main

const spanish = "Spanish"
const french = "French"
const gujju = "Gujarati"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const gujjuHelloPrefix = "Kemche, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case gujju:
		prefix = gujjuHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	//	fmt.Println(Hello("world", "English"))
}
