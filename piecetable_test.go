package piecetable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiStringInsertion(t *testing.T) {
	text := `this is my testing
document i want to see how it fares
and all of that fun
stuff`
	output := `this is my piece foo table testing
document i want to see how it fares
and all of that fun
stuff`

	table := MakePieceTable(text)
	table.Insert("piece table ", 0, 11)
	table.Insert("foo ", 0, 17)

	fmt.Printf(table.String())

	assert.Equal(t, output, table.String())
}

func TestStringInsertion(t *testing.T) {
	text := `this is my testing
document i want to see how it fares
and all of that fun
stuff`
	output := `this is my piece table testing
document i want to see how it fares
and all of that fun
stuff`

	table := MakePieceTable(text)
	table.Insert("piece table ", 0, 11)

	fmt.Printf(table.String())

	assert.Equal(t, output, table.String())
}

func TestPrintDocument(t *testing.T) {
	text := `this is my testing
document i want to see how it fares
and all of that fun
stuff`

	table := MakePieceTable(text)

	fmt.Printf(table.String())

	assert.Equal(t, text, table.String(), "Un-modified piece table output doesn't match value expected")
}
