package owoify_go

import (
	"io/ioutil"
	"strings"
	"testing"
)

const source = "This is the string to owo! Kinda cute, isn't it?"
const pokemonNameListPath = "assets/pokemons.txt"
const warAndPeacePath = "assets/war_and_peace_chapter01-20.txt"

func TestOwo(t *testing.T) {
	if got := Owoify(source, Owo); len(got) == 0 {
		t.Error("Failed to owoify the string with level of OwO.")
	}
}

func TestUwu(t *testing.T) {
	if got := Owoify(source, Uwu); len(got) == 0 {
		t.Error("Failed to owoify the string with level of UwU.")
	}
}

func TestUvu(t *testing.T) {
	if got := Owoify(source, Uvu); len(got) == 0 {
		t.Error("Failed to owoify the string with level of UvU.")
	}
}

func TestOwoify(t *testing.T) {
	if got := Owoify(source, Owo); got == source {
		t.Error("Failed to owoify the string with level of UwU.")
	}
}

func TestInvalidOwoness(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The invalid owoness level didn't panic.")
		}
	}()

	Owoify(source, 123)
}

func TestOwoNotEqualToUwu(t *testing.T) {
	if got := Owoify(source, Owo); got == Owoify(source, Uwu) {
		t.Error("A string with level of owo should not be equal to a string with level of uwu.")
	}
}

func TestOwoNotEqualToUvu(t *testing.T) {
	if got := Owoify(source, Owo); got == Owoify(source, Uvu) {
		t.Error("A string with level of owo should not be equal to a string with level of uvu.")
	}
}

func TestUwuNotEqualToUvu(t *testing.T) {
	if got := Owoify(source, Uwu); got == Owoify(source, Uvu) {
		t.Error("A string with level of uwu should not be equal to a string with level of uvu.")
	}
}

func TestPokemonNames(t *testing.T) {
	pokemonNamesBytes, _ := ioutil.ReadFile(pokemonNameListPath)
	pokemonNames := strings.Split(string(pokemonNamesBytes), "\n")
	for _, name := range pokemonNames {
		nameWithOwo := Owoify(name, Owo)
		nameWithUwu := Owoify(name, Uwu)
		nameWithUvu := Owoify(name, Uvu)
		if len(nameWithOwo) == 0 {
			t.Error("Pokemon name didn't get owoified correctly.")
		}
		if len(nameWithUwu) == 0 {
			t.Error("Pokemon name didn't get owoified correctly.")
		}
		if len(nameWithUvu) == 0 {
			t.Error("Pokemon name didn't get owoified correctly.")
		}
	}
}

func TestLongText(t *testing.T) {
	longTextBytes, _ := ioutil.ReadFile(warAndPeacePath)
	text := string(longTextBytes)
	textWithOwo := Owoify(text, Owo)
	textWithUwu := Owoify(text, Uwu)
	textWithUvu := Owoify(text, Uvu)
	if len(textWithOwo) == 0 {
		t.Error("Long text didn't get owoified correctly.")
	}
	if len(textWithUwu) == 0 {
		t.Error("Long text didn't get owoified correctly.")
	}
	if len(textWithUvu) == 0 {
		t.Error("Long text didn't get owoified correctly.")
	}
}
