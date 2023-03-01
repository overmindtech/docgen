package main

import (
	"go/ast"
	"strings"
)

// ParseFile Parses overmind documentation comments out of an already parsed go
// file. Go files can be parsed with `parser.parseFile()` from the "go/parser"
// package
func ParseFile(file *ast.File) SourceDoc {
	sd := SourceDoc{}

	for _, group := range file.Comments {
		sd.ParseGroup(group)
	}

	return sd
}

type SourceDoc struct {
	Type             string // Parsed from a +overmind:type comment
	Get              string // Parsed from a +overmind:get comment
	List             string // Parsed from a +overmind:list comment
	Search           string // Parsed from a +overmind:search comment
	descriptionLines []string
}

// Description Returns the description of the source, parsed from a
// +overmind:description comment over multipe lines
func (s *SourceDoc) Description() string {
	return strings.Trim(strings.Join(s.descriptionLines, "\n"), "\n")
}

// ParseGroup Parses a comment group and adds the details to the SourceDoc
// struct
func (s *SourceDoc) ParseGroup(group *ast.CommentGroup) {
	var after string
	var found bool
	var writeDescription bool

	lines := group.Text()

	for _, line := range strings.Split(lines, "\n") {
		// Check for prefixes
		if after, found = strings.CutPrefix(line, "+overmind:type"); found {
			s.Type = strings.Trim(after, " ")
			writeDescription = false
		} else if after, found = strings.CutPrefix(line, "+overmind:get"); found {
			s.Get = strings.Trim(after, " ")
			writeDescription = false
		} else if after, found = strings.CutPrefix(line, "+overmind:list"); found {
			s.List = strings.Trim(after, " ")
			writeDescription = false
		} else if after, found = strings.CutPrefix(line, "+overmind:search"); found {
			s.Search = strings.Trim(after, " ")
			writeDescription = false
		} else if after, found = strings.CutPrefix(line, "+overmind:description"); found {
			writeDescription = true
			line = strings.Trim(after, " ")
		}

		// If we are within the description block, collect it
		if writeDescription {
			s.descriptionLines = append(s.descriptionLines, line)
		}
	}
}
