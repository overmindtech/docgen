package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fileName := os.Getenv("GOFILE")

	if fileName == "" {
		log.Print("GOFILE cannot be blank")
	}

	log.Printf("Generating Overmind docs for %v", fileName)

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)

	if err != nil {
		log.Printf("Error parsing file: %v", err)
	}

	doc := SourceDoc{}

	for _, group := range parsed.Comments {
		doc.ParseGroup(group)
	}

}
