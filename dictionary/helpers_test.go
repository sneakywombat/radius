package dictionary_test

import (
	"reflect"
	"testing"

	"github.com/sneakywombat/radius/dictionary"
)

func TestMerge(t *testing.T) {
	files := []MemoryFile{
		{
			Filename: "dict1",
			Contents: `
VENDOR Test 32473
BEGIN-VENDOR Test
ATTRIBUTE Test-Vendor-Name 5 string
END-VENDOR Test`,
		},
		{
			Filename: "dict2",
			Contents: `
VENDOR Test 32473
BEGIN-VENDOR Test
ATTRIBUTE Test-Vendor-Int 10 integer
END-VENDOR Test`,
		},
	}

	parser := &dictionary.Parser{
		Opener: MemoryOpener(files),
	}
	d1, err := parser.ParseFile("dict1")
	if err != nil {
		t.Fatal(err)
	}

	d2, err := parser.ParseFile("dict2")
	if err != nil {
		t.Fatal(err)
	}

	merged, err := dictionary.Merge(d1, d2)
	if err != nil {
		t.Fatal(err)
	}

	expected := &dictionary.Dictionary{
		Vendors: []*dictionary.Vendor{
			{
				Name:   "Test",
				Number: 32473,
				Attributes: []*dictionary.Attribute{
					{
						Name: "Test-Vendor-Name",
						Type: dictionary.AttributeString,
						OID:  "5",
					},
					{
						Name: "Test-Vendor-Int",
						Type: dictionary.AttributeInteger,
						OID:  "10",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(merged, expected) {
		t.Fatalf("got:\n%#v\nexpected:\n%#v", merged, expected)
	}
}
