package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText,
		make([]bool, len(plainText))}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return (position >= t.Start && position <= t.End)
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}

	for position, character := range b.plainText {
		c := character
		for _, r := range b.formatting {
			if r.Covers(position) && r.Capitalize {
				c = unicode.ToUpper(rune(c))
			}
		}
		sb.WriteRune(rune(c))
	}

	return sb.String()
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

func (ft *FormattedText) String() string {
	sb := strings.Builder{}

	for position, letter := range ft.plainText {
		if ft.capitalize[position] {
			sb.WriteRune(unicode.ToUpper(letter))

			continue
		}

		sb.WriteRune(letter)
	}

	return sb.String()
}

func (ft *FormattedText) Capitalize(start, end int) {
	for ; start <= end; start++ {
		ft.capitalize[start] = true
	}
}

func main() {
	text := "This is a brave new world"

	ft := NewFormattedText(text)
	ft.Capitalize(0, len(text)-1)
	fmt.Println(ft.String())

	bft := NewBetterFormattedText(text)
	bft.Range(0, len(text) - 1).Capitalize = true
	fmt.Println(bft.String())
}
