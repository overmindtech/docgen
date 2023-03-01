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

	if len(os.Args) != 2 {
		log.Fatal("Must pass the desired output file as the only argument")
	}

	outFile := os.Args[1]

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

	md, err := doc.FormatMarkdown()

	if err != nil {
		log.Fatalf("Error generating markdown: %v", err)
	}

	err = os.WriteFile(outFile, []byte(md), 0644)

	if err != nil {
		log.Fatalf("Error writing file %v: %v", outFile, err)
	}

	log.Printf("Markdown written: %v", outFile)
}
