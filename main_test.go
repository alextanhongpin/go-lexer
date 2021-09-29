package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	tests := []struct {
		input    string
		expected *SelectStatement
	}{
		{

			input: `SELECT * FROM mytable`,
			expected: &SelectStatement{
				TableName: "mytable",
				Fields:    []string{"*"},
			},
		},
		{
			input: `SELECT name, age FROM mytable`,
			expected: &SelectStatement{
				TableName: "mytable",
				Fields:    []string{"name", "age"},
			},
		},
	}

	for _, test := range tests {
		p := NewParser(bytes.NewReader([]byte(test.input)))
		stmt, err := p.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(test.expected, stmt) {
			t.Fatalf("expected %v, got %v", test.expected, stmt)
		}
	}
}
