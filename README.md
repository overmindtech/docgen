# 👩‍⚕️🤖 docgen

*Documentation generation for Overmind sources*

This tool generates markdown docs for Overmind sources based on comments in the code. For example these comments:

```go
//go:generate docgen ./docs
// +overmind:type http
// +overmind:get Runs a HEAD request against a given URL
// +overmind:list **Not supported**
// +overmind:search By ARN
// +overmind:description
// The HTTP source runs HEAD requests to get the details of an HTTP endpoint.
// All items are returned in the "global" scope and container links to
// certificates and DNS entries
```

Produces this markdown:

```markdown
# http

The HTTP source runs HEAD requests to get the details of an HTTP endpoint.
All items are returned in the "global" scope and container links to
certificates and DNS entries

## Supported Methods

* **Get:** Runs a HEAD request against a given URL
* **List:** **Not supported**
* **Search:** By ARN
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

* `+overmind:type`: The type of item returned by the source
* `+overmind:get`: Documentation for the get method
* `+overmind:list`: Documentation for the get method
* `+overmind:search`: Documentation for the Search method, if this is omitted the search method will not be documented

## Multi-Line Tags

These tags will consume documentation from the line containing the tag, and all subsequent lines until either the end of the comment or another tag is found.

* `+overmind:description`: Detailed description of the source in this file
