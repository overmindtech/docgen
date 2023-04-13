package main

import (
	"go/ast"
	"sort"
	"strings"
)

// ParseFile Parses overmind documentation comments out of an already parsed go
// file. Go files can be parsed with `parser.parseFile()` from the "go/parser"
// package
func ParseFile(file *ast.File) SourceDoc {
	sd := SourceDoc{
		linksMap: make(map[string]struct{}),
	}

	for _, group := range file.Comments {
		sd.parseGroup(group)
	}

	// Combine the links map into a slice and sort
	links := make([]string, 0)

	// Combine the links map into a slice
	for link := range sd.linksMap {
		links = append(links, link)
	}

	// Sort links alphabetically
	sort.Strings(links)

	sd.Links = links

	return sd
}

type SourceDoc struct {
	// The type that the source returns e.g. `ec2-instance` Parsed from the
	// +overmind:type comment
	OvermindType string `json:"type"`

	// The desriptive type e.g. `EC2 Instance` Parsed from the
	// +overmind:descriptiveType
	DescriptiveType string `json:"descriptiveType"`

	// Description of the Get method for this source. Parsed from the
	// +overmind:get comment
	GetDescription string `json:"getDescription,omitempty"`

	// Description of the List method for this source. Parsed from the
	// +overmind:list comment
	ListDescription string `json:"listDescription,omitempty"`

	// Description of the Search method for this source. Parsed from the
	// +overmind:search comment
	SearchDescription string `json:"searchDescription,omitempty"`

	// The group that this source belongs to e.g. "AWS". Parsed from the
	// +overmind:group comment
	SourceGroup string `json:"group"`

	// Types of items that this can be linked to, parsed from many
	// +overmind:link comments
	Links []string `json:"links"`

	// Used for deduplication
	linksMap map[string]struct{} `json:"-"`
}

// parseGroup Parses a comment group and adds the details to the SourceDoc
// struct
func (s *SourceDoc) parseGroup(group *ast.CommentGroup) {
	var after string
	var found bool

	lines := group.Text()

	for _, line := range strings.Split(lines, "\n") {
		// Check for prefixes
		if after, found = strings.CutPrefix(line, "+overmind:type"); found {
			s.OvermindType = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:descriptiveType"); found {
			s.DescriptiveType = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:get"); found {
			s.GetDescription = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:list"); found {
			s.ListDescription = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:search"); found {
			s.SearchDescription = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:group"); found {
			s.SourceGroup = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:link"); found {
			s.linksMap[strings.Trim(after, " ")] = struct{}{}
		}
	}
}
