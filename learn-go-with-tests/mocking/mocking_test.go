package main

import (
	"bytes"
	"testing"
)

type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleep{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Error en el sleep, se esperaba 4 se obtuvo %d", spySleeper.Calls)
	}
}
