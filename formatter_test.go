package main

import "testing"

func TestFormatMarkdown(t *testing.T) {
	t.Run("with all populated", func(t *testing.T) {
		sd := SourceDoc{
			Type:   "type",
			Get:    "get",
			List:   "list",
			Search: "search",
			descriptionLines: []string{
				"description",
			},
		}

		expected := `# type

description

## Supported Methods

* **Get:** get
* **List:** list
* **Search:** search
`

		out, err := sd.FormatMarkdown()

		if err != nil {
			t.Fatal(err)
		}

		if out != expected {
			t.Errorf("expected out to be:\n%v\nGot:\n%v", expected, out)
		}
	})

	t.Run("with none populated", func(t *testing.T) {
		sd := SourceDoc{}

		expected := `# 



## Supported Methods

* **Get**
* **List**
`

		out, err := sd.FormatMarkdown()

		if err != nil {
			t.Fatal(err)
		}

		if out != expected {
			t.Errorf("expected out to be:\n%v\nGot:\n%v", expected, out)
		}
	})

	t.Run("with no search", func(t *testing.T) {
		sd := SourceDoc{
			Type: "type",
			Get:  "get",
			List: "list",
			descriptionLines: []string{
				"description",
			},
		}

		expected := `# type

description

## Supported Methods

* **Get:** get
* **List:** list
`

		out, err := sd.FormatMarkdown()

		if err != nil {
			t.Fatal(err)
		}

		if out != expected {
			t.Errorf("expected out to be:\n%v\nGot:\n%v", expected, out)
		}
	})
}
