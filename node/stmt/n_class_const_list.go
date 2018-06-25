package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ClassConstList node
type ClassConstList struct {
	Comments  []*comment.Comment
	Position  *position.Position
	Modifiers []node.Node
	Consts    []node.Node
}

// NewClassConstList node constructor
func NewClassConstList(Modifiers []node.Node, Consts []node.Node) *ClassConstList {
	return &ClassConstList{
		Modifiers: Modifiers,
		Consts:    Consts,
	}
}

// SetPosition sets node position
func (n *ClassConstList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClassConstList) GetPosition() *position.Position {
	return n.Position
}

func (n *ClassConstList) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *ClassConstList) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *ClassConstList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassConstList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Modifiers != nil {
		v.EnterChildList("Modifiers", n)
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Modifiers", n)
	}

	if n.Consts != nil {
		v.EnterChildList("Consts", n)
		for _, nn := range n.Consts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Consts", n)
	}

	v.LeaveNode(n)
}
