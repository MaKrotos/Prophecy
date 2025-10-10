package names

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {
	// Инициализируем генератор случайных чисел с источником, основанным на времени
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateRandomName генерирует случайное двухсловное имя
func GenerateRandomName() string {
	if len(Adjectives) == 0 || len(Nouns) == 0 {
		return "Аноним"
	}

	adjective := Adjectives[rng.Intn(len(Adjectives))]
	noun := Nouns[rng.Intn(len(Nouns))]

	return adjective + " " + noun
}
