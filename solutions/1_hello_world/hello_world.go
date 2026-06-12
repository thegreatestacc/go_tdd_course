package main

func main() {
}

const (
	english            = "English"
	englishHelloPrefix = "Hello "
	spanish            = "Spanish"
	spanishHelloPerf   = "Hola "
	french             = "French"
	frenchHelloPrefix  = "Bonjour "
)

func Hello(language, name string) string {
	name = validateName(name)

	prefix := gettingPrefix(language)

	return prefix + name
}

func validateName(name string) string {
	if name == "" {
		return "World"
	}
	return name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPerf
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix
}
