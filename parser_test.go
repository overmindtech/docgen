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

	if sd.TerraformMethod != "GET" {
		t.Errorf("expected terraform method to be GET, got %v", sd.TerraformMethod)
	}

	if len(sd.TerraformQueryMaps) != 2 {
		t.Errorf("expected 2 terraform query maps, got %v", len(sd.TerraformQueryMaps))
	}

	if sd.TerraformQueryMaps[0] != "resource_type.id" {
		t.Errorf("expected terraform query map to be resource_type.id, got %v", sd.TerraformQueryMaps[0])
	}

	if sd.TerraformQueryMaps[1] != "resource_type2.id" {
		t.Errorf("expected terraform query map to be resource_type2.id, got %v", sd.TerraformQueryMaps[1])
	}

	if sd.TerraformScope != "*" {
		t.Errorf("expected terraform scope to be *, got %v", sd.TerraformScope)
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
		// +overmind:terraform:query resource_type.id	
		// +overmind:terraform:query resource_type2.id	
		
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
		
		// +overmind:terraform:query resource_type2.id	

		// +overmind:terraform:query resource_type.id	
		
		// +overmind:terraform:method GET

		// +overmind:terraform:scope *
		
		func Foo() bool {
		
		}`,
		"reverse order": `package main
		
		// +overmind:terraform:query resource_type2.id	
		// +overmind:terraform:query resource_type.id	
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

			sd, err := ParseFile(parsed)

			if err != nil {
				t.Error(err)
			}

			testSD(sd, t)

		})
	}
}
