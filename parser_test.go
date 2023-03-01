package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func testSD(sd SourceDoc, t *testing.T) {
	if sd.Type != "http" {
		t.Errorf("expected type to be \"http\", got %v", sd.Type)
	}

	if sd.Get != "Runs a HEAD request against a given URL" {
		t.Errorf("expected get to be \"Runs a HEAD request against a given URL\", got %v", sd.Get)
	}

	if sd.List != "**Not supported**" {
		t.Errorf("expected list to be \"**Not supported**\", got %v", sd.List)
	}

	if sd.Description() != "This type does HTTP requests e.g.\n\nhttps://www.google.com" {
		t.Errorf("expected description to be \"This type does HTTP requests e.g.\", got %v", sd.Description())
	}
}

func TestParseFile(t *testing.T) {
	testFiles := map[string]string{
		"all together": `package main

		//go:generate sourceDoc
		// +overmind:type http
		// +overmind:get Runs a HEAD request against a given URL
		// +overmind:list **Not supported**
		// +overmind:description
		// This type does HTTP requests e.g.
		//
		// https://www.google.com
		//
		//
		
		func Foo() bool {
		
		}`,
		"all separate": `package main

		//go:generate sourceDoc
		// +overmind:type http

		// +overmind:get Runs a HEAD request against a given URL

		// +overmind:list **Not supported**

		// +overmind:description
		// This type does HTTP requests e.g.
		//
		// https://www.google.com
		//
		//
		
		func Foo() bool {
		
		}`,
		"reverse order": `package main

		// +overmind:description
		// This type does HTTP requests e.g.
		//
		// https://www.google.com
		//
		//
		// +overmind:type http
		// +overmind:get Runs a HEAD request against a given URL
		// +overmind:list **Not supported**
		
		func Foo() bool {
		
		}`,
	}

	for name, file := range testFiles {
		t.Run(name, func(t *testing.T) {
			fset := token.NewFileSet()
			parsed, err := parser.ParseFile(fset, "", file, parser.ParseComments)

			if err != nil {
				t.Fatal(err)
			}

			sd := ParseFile(parsed)

			testSD(sd, t)

		})
	}
}
