package main

import (
	"go/parser"
	"go/token"
	"testing"
)

//go:generate ./docgen ./docs
// +overmind:type ec2-instance
// +overmind:descriptiveType EC2 Instance
// +overmind:get Get an EC2 instance by ID
// +overmind:list List all EC2 instances
// +overmind:search Search for EC2 instances by name
// +overmind:group AWS
// +overmind:link ip
// +overmind:link ec2-security-group

func testSD(sd SourceDoc, t *testing.T) {
	if sd.OvermindType != "ec2-instance" {
		t.Errorf("expected type to be ec2-instance, got %v", sd.OvermindType)
	}

	if sd.DescriptiveType != "EC2 Instance" {
		t.Errorf("expected descriptive type to be EC2 Instance, got %v", sd.DescriptiveType)
	}

	if sd.GetDescription != "Get an EC2 instance by ID" {
		t.Errorf("expected get description to be Get an EC2 instance by ID, got %v", sd.GetDescription)
	}

	if sd.ListDescription != "List all EC2 instances" {
		t.Errorf("expected list description to be List all EC2 instances, got %v", sd.ListDescription)
	}

	if sd.SearchDescription != "Search for EC2 instances by name" {
		t.Errorf("expected search description to be Search for EC2 instances by name, got %v", sd.SearchDescription)
	}

	if sd.SourceGroup != "AWS" {
		t.Errorf("expected source group to be AWS, got %v", sd.SourceGroup)
	}

	if len(sd.Links) != 2 {
		t.Errorf("expected 2 links, got %v", len(sd.Links))
	}

	if sd.Links[0] != "ec2-security-group" {
		t.Errorf("expected second link to be ec2-security-group, got %v", sd.Links[1])
	}

	if sd.Links[1] != "ip" {
		t.Errorf("expected first link to be ip, got %v", sd.Links[0])
	}
}

func TestParseFile(t *testing.T) {
	testFiles := map[string]string{
		"all together": `package main

		// +overmind:type ec2-instance
		// +overmind:descriptiveType EC2 Instance
		// +overmind:get Get an EC2 instance by ID
		// +overmind:list List all EC2 instances
		// +overmind:search Search for EC2 instances by name
		// +overmind:group AWS
		// +overmind:link ip
		// +overmind:link ec2-security-group		
		
		func Foo() bool {
		
		}`,
		"all separate": `package main

		// +overmind:type ec2-instance

		// +overmind:descriptiveType EC2 Instance

		// +overmind:get Get an EC2 instance by ID

		// +overmind:list List all EC2 instances

		// +overmind:search Search for EC2 instances by name
		// +overmind:group AWS

		// +overmind:link ip
		// +overmind:link ec2-security-group
		
		
		func Foo() bool {
		
		}`,
		"reverse order": `package main

		// +overmind:group AWS
		// +overmind:list List all EC2 instances
		// +overmind:type ec2-instance
		// +overmind:link ip
		// +overmind:link ip
		// +overmind:search Search for EC2 instances by name
		// +overmind:descriptiveType EC2 Instance
		// +overmind:link ip
		// +overmind:link ec2-security-group
		// +overmind:get Get an EC2 instance by ID
		
		
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
