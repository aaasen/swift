package swift

// Abstract Syntax Tree

import (
	"fmt"
	"strings"
)

type AST struct {
	Children []*AST
	Value    interface{}
}

func (self *AST) String() string {
	return self.levelString(0)
}

func (self *AST) levelString(level int) string {
	childStrings := make([]string, len(self.Children))

	if self.Children != nil {
		for i, child := range self.Children {
			childStrings[i] = child.levelString(level + 1)
		}
	}

	return fmt.Sprintf("%s- %v\n%s", strings.Repeat("    ", level),
		self.Value,
		strings.Join(childStrings, ""))
}
