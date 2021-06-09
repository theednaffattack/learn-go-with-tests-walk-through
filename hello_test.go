package main

import "testing"



func TestHello(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello(HelloParams{name: "Eddie", lang: "English"})
		want := "Hello, Eddie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello(HelloParams{name: ""})
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("respond in Spanish if specifying 'Spanish' in the function call", func(t *testing.T) {
		got := Hello(HelloParams{name: "Elodie", lang: "Spanish"})
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}