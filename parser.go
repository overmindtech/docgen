package main

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"
)

// ParseFile Parses overmind documentation comments out of an already parsed go
// file. Go files can be parsed with `parser.parseFile()` from the "go/parser"
// package
func ParseFile(file *ast.File) (SourceDoc, error) {
	sd := SourceDoc{
		linksMap: make(map[string]struct{}),
	}

	var err error

	for _, group := range file.Comments {
		err = sd.parseGroup(group)

		if err != nil {
			return sd, err
		}
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

	// Set Terraform defaults if required
	if len(sd.TerraformQueryMaps) > 0 {
		if sd.TerraformMethod == "" {
			sd.TerraformMethod = "GET"
		}

		if sd.TerraformScope == "" {
			sd.TerraformScope = "*"
		}
	}

	return sd, nil
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

	// Where the `query` data should come from when converting a `terraform
	// plan` to an Overmind query. This is in the format
	// `{resource_type}.{attribute_name}` and comes from the:
	// +overmind:terraform:query comment. Multiple comments are supported
	TerraformQueryMaps []string `json:"terraformQuery,omitempty"`

	// The method the query should have when converting from a `terraform plan`
	// to an overmind query. Defaults to `GET`. Valid values: `GET`, `LIST`,
	// `SEARCH`. Defined by the comment:
	// +overmind:terraform:method SEARCH
	TerraformMethod string `json:"terraformMethod,omitempty"`

	// The scope that the query should have when converting from a `terraform
	// plan` to an Overmind query, defaults to `*`. Defined by the comment:
	// +overmind:terraform:scope global
	TerraformScope string `json:"terraformScope,omitempty"`

	// Types of items that this can be linked to, parsed from many
	// +overmind:link comments
	Links []string `json:"links"`

	// Used for deduplication
	linksMap map[string]struct{} `json:"-"`
}

// parseGroup Parses a comment group and adds the details to the SourceDoc
// struct
func (s *SourceDoc) parseGroup(group *ast.CommentGroup) error {
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
		} else if after, found = strings.CutPrefix(line, "+overmind:terraform:query"); found {
			s.TerraformQueryMaps = append(s.TerraformQueryMaps, strings.Trim(after, " "))
		} else if after, found = strings.CutPrefix(line, "+overmind:terraform:method"); found {
			method := strings.Trim(after, " ")
			switch method {
			case "GET", "LIST", "SEARCH":
				s.TerraformMethod = method
			default:
				return fmt.Errorf("unsupported value for +overmind:terraform:method: %v Must be either GET, LIST or SEARCH", method)
			}
			s.TerraformMethod = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:terraform:scope"); found {
			s.TerraformScope = strings.Trim(after, " ")
		} else if after, found = strings.CutPrefix(line, "+overmind:link"); found {
			s.linksMap[strings.Trim(after, " ")] = struct{}{}
		}
	}

	// Sort query maps for deterministic output
	if len(s.TerraformQueryMaps) > 0 {
		sort.Strings(s.TerraformQueryMaps)
	}

	return nil
}
