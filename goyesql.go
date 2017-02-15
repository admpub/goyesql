// Package goyesql is a Go port of Yesql
//
// It allows you to write SQL queries in separate files.
//
// See rationale at https://github.com/krisajenkins/yesql#rationale
package goyesql

import (
	"bytes"
	"os"
)

// Some helpers to read files

// ParseFile reads a file and return Queries or an error
func ParseFile(path string, preprocessors ...func(string) string) (Queries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ParseReader(file, preprocessors...)
}

// MustParseFile calls ParseFile but panic if an error occurs
func MustParseFile(path string, preprocessors ...func(string) string) Queries {
	queries, err := ParseFile(path, preprocessors...)
	if err != nil {
		panic(err)
	}

	return queries
}

// ParseBytes parses bytes and returns Queries or an error.
func ParseBytes(b []byte, preprocessors ...func(string) string) (Queries, error) {
	return ParseReader(bytes.NewReader(b), preprocessors...)
}

// MustParseBytes parses bytes but panics if an error occurs.
func MustParseBytes(b []byte) Queries {
	queries, err := ParseBytes(b)
	if err != nil {
		panic(err)
	}

	return queries
}
