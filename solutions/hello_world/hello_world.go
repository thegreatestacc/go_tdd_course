package main

const (
	english            = "English"
	englishHelloPrefix = "Hello "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola "
	french             = "French"
	frenchHelloPrefix  = "Bonjour "
)

func main() {
}

func Hello(language, name string) string {
	name = validateName(name)

	prefix := greetingPrefix(language)

	return prefix + name
}

func validateName(name string) string {
	if name == "" {
		return "World"
	}
	return name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
