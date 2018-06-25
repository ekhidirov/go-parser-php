package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Case node
type Case struct {
	Comments []*comment.Comment
	Position *position.Position
	Cond     node.Node
	Stmts    []node.Node
}

// NewCase node constructor
func NewCase(Cond node.Node, Stmts []node.Node) *Case {
	return &Case{
		Cond:  Cond,
		Stmts: Stmts,
	}
}

// SetPosition sets node position
func (n *Case) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Case) GetPosition() *position.Position {
	return n.Position
}

func (n *Case) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Case) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Case) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Case) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
