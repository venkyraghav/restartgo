package hello

/* Constants */
const (
	spanish            = "Spanish"
	english            = "English"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

/* getHelloPrefix returns the hello prefix based on the language */
func getHelloPrefix(lang string) string {
	switch lang {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	}
	return englishHelloPrefix
}

/* Hello returns a greeting */
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return getHelloPrefix(lang) + name
}
