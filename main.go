package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
)

func main() {
	fileName := os.Getenv("GOFILE")

	if fileName == "" {
		log.Print("GOFILE cannot be blank")
	}

	if len(os.Args) != 2 {
		log.Fatal("Must pass the desired output path as the only argument")
	}

	outPath := os.Args[1]

	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		log.Printf("Creating output folder: %v", outPath)
		err := os.Mkdir(outPath, os.ModePerm)

		if err != nil {
			log.Fatalf("Error creating folder: %v", err)
		}
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

	md, err := doc.FormatMarkdown()

	if err != nil {
		log.Fatalf("Error generating markdown: %v", err)
	}

	outFile := path.Join(outPath, fmt.Sprintf("%v.md", doc.Type))

	err = os.WriteFile(outFile, []byte(md), 0644)

	if err != nil {
		log.Fatalf("Error writing file %v: %v", outFile, err)
	}

	log.Printf("Markdown written: %v", outFile)
}
