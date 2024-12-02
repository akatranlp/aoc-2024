package its_test

import (
	"aoc-lib/its"
	"bufio"
	"bytes"
	"log"
	"testing"
)

func TestFile(t *testing.T) {
	t.Run("SplitByBlocks", func(t *testing.T) {
		tests := []struct {
			input  string
			blocks []string
		}{{
			input:  "\nHallo\n\nTest\n",
			blocks: []string{"\nHallo", "Test\n"},
		}, {
			input:  "\n\n\n\n\n\n",
			blocks: []string{"", "", ""},
		}, {
			input:  "",
			blocks: []string{},
		}, {
			input:  " ",
			blocks: []string{" "},
		}, {
			input:  "\n\n ",
			blocks: []string{"", " "},
		}, {
			input:  "\n\n",
			blocks: []string{""},
		}}
		for _, test := range tests {
			buf := bytes.NewBufferString(test.input)
			scanner := bufio.NewScanner(buf)
			scanner.Split(its.SplitByBlocks)

			for i := range len(test.blocks) {
				if !scanner.Scan() {
					t.Fatalf("ERROR: expected a %d block\n", i+1)
				}
				if scanner.Err() != nil {
					t.Fatalf("ERROR: expected no error in block %d\n", i+1)
				}

				expected := test.blocks[i]
				actual := scanner.Text()
				if expected != actual {
					t.Fatalf("ERROR: expected %q(%d) actual %q(%d)\n", expected, len(expected), actual, len(actual))
				}
			}

			if scanner.Scan() {
				log.Fatal("don't expect a next block")
			}
			if scanner.Err() != nil {
				log.Fatal("expected no error")
			}
		}
	})
}
