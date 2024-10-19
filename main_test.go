package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var EIU = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, EIU
	}

	return utf8.RuneCount(input), nil
}

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{input: []byte("Hello, 世界"), expected: 9, err: nil},

		{input: []byte{0x80}, expected: 0, err: EIU},

		{input: []byte{}, expected: 0, err: nil},

		{input: []byte("你"), expected: 1, err: nil},
	}

	for _, tt := range tests {
		got, err := GetUTFLength(tt.input)
		if got != tt.expected {
			t.Errorf("GetUTFLength() got = %v, expected = %v", got, tt.expected)
		}
		if err != tt.err {
			t.Errorf("GetUTFLength() error = %v, expected = %v", err, tt.err)
		}
	}
}
