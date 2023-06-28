# 👩‍⚕️🤖 docgen

*Documentation generation for Overmind sources*

This tool generates JSON that can then be used to generate docs for Overmind sources. The format of the comments is as follows:

```go
//go:generate docgen ./docs
// +overmind:type ec2-instance
// +overmind:descriptiveType EC2 Instance
// +overmind:get Get an EC2 instance by ID
// +overmind:list List all EC2 instances
// +overmind:search Search for EC2 instances by name
// +overmind:group AWS
// +overmind:link ip
// +overmind:link ec2-security-group
// +overmind:terraform:method SEARCH
// +overmind:terraform:queryMap aws_autoscaling_group.name
// +overmind:terraform:scope global
```

Produces this markdown:

```json
{
	"type": "ec2-instance",
	"descriptiveType": "EC2 Instance",
	"getDescription": "Get an EC2 instance by ID",
	"listDescription": "List all EC2 instances",
	"searchDescription": "Search for EC2 instances by name",
	"group": "AWS",
	"links": [
		"ip",
		"ec2-security-group"
	]
}
```

Note that the format of the `go generate` comment is:

```
//go:generate docgen {destination_folder}
```

Where `destination_folder` is relative to the current file

## Installation

Intel:

```shell
mkdir -p ~/.local/bin/ && curl -Lo ~/.local/bin/docgen https://github.com/overmindtech/docgen/releases/latest/download/docgen-amd64 && chmod +x ~/.local/bin/docgen
```

ARM:

```shell
mkdir -p ~/.local/bin/ && curl -Lo ~/.local/bin/docgen https://github.com/overmindtech/docgen/releases/latest/download/docgen-arm64 && chmod +x ~/.local/bin/docgen
```

## Comment Tags

All tags gor a given type should exist within the same file, however they can live anywhere withing that file.

### Single-Line Tags

These tags must include the desired documentation following the tag on a single line

* `+overmind:type`: The type that the source returns e.g. `ec2-instance`
* `+overmind:descriptiveType: The desriptive type e.g. `EC2 Instance`
* `+overmind:get`: Description of the Get method for this source
* `+overmind:list`: Description of the List method for this source
* `+overmind:search`: Description of the Search method for this source
* `+overmind:group`: The group that this source belongs to e.g. "AWS"
* `+overmind:link`: Types of items that this can be linked to
* `+overmind:terraform:method`: The method the query should have when converting from a `terraform plan` to an overmind query. Defaults to `GET`. Valid values: `GET`, `LIST`, `SEARCH`
* `+overmind:terraform:queryMap`:  Where the `query` data should come from when converting a `terraform plan` to an Overmind query. This is in the format `{resource_type}.{attribute_name}`. This can be specified multiple times if there are many terrafrom resources taht could affect this overmind type
* `+overmind:terraform:scope`: The scope that the query should have when converting from a `terraform plan` to an Overmind query, defaults to `*`
