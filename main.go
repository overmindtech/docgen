package main

import (
	"encoding/json"
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

	log.Printf("Generating docs data for %v", fileName)

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)

	if err != nil {
		log.Printf("Error parsing file: %v", err)
	}

	doc, err := ParseFile(parsed)

	if err != nil {
		log.Fatalf("Error parsing file %v: %v", fileName, err)
	}

	// Format as JSON
	b, err := json.MarshalIndent(doc, "", "	")

	if err != nil {
		log.Fatalf("Error generating JSON: %v", err)
	}

	outFile := path.Join(outPath, fmt.Sprintf("%v.json", doc.OvermindType))

	err = os.WriteFile(outFile, b, 0644)

	if err != nil {
		log.Fatalf("Error writing file %v: %v", outFile, err)
	}

	log.Printf("JSON written: %v", outFile)
}
