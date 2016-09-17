package sententia

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type maker struct{}

func (m *maker) Noun() string {
	return nouns[rand.Intn(len(nouns))]
}

func (m *maker) ANoun() string {
	return articlize(m.Noun())
}

func (m *maker) Nouns() string {
	return pluralize(m.Noun())
}

func (m *maker) Adjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func (m *maker) AnAdjective() string {
	return articlize(m.Adjective())
}

func articlize(word string) string {
	var article = "a"
	switch word[0] {
	case 'a', 'e', 'i', 'o', 'u':
		article = "an"
	}
	return fmt.Sprintf("%s %s", article, word)
}

func pluralize(word string) string {
	var suffix = "s"
	switch word[len(word)-1] {
	case 's', 'h':
		suffix = "es"
	}
	return fmt.Sprintf("%s%s", word, suffix)
}
