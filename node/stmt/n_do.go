package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Do node
type Do struct {
	Comments []*comment.Comment
	Position *position.Position
	Stmt     node.Node
	Cond     node.Node
}

// NewDo node constructor
func NewDo(Stmt node.Node, Cond node.Node) *Do {
	return &Do{
		Stmt: Stmt,
		Cond: Cond,
	}
}

// SetPosition sets node position
func (n *Do) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Do) GetPosition() *position.Position {
	return n.Position
}

func (n *Do) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Do) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Do) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Do) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	v.LeaveNode(n)
}
