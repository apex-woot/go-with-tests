package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

type OptionFunc func(*GreetingOpts)

func WithLanguage(lang string) OptionFunc {
	return func(opts *GreetingOpts) {
		opts.language = lang
	}
}

type GreetingOpts struct {
	language string
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, options ...OptionFunc) string {
	if name == "" {
		name = "World"
	}

	opts := &GreetingOpts{
		language: "",
	}

	for _, opt := range options {
		opt(opts)
	}

	return greetingPrefix(opts.language) + name
}

func main() {
	fmt.Println(Hello("world"))
}
