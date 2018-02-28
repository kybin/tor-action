package action

import "strings"

type Text struct {
	lines []string
}

func NewText(str string) *Text {
	return &Text{strings.Split(str, "\n")}
}

func (t *Text) Lines() []string {
	return t.lines
}

func (t *Text) Line(i int) string {
	return t.lines[i]
}

func (t *Text) String() string {
	return strings.Join(t.lines, "\n")
}
