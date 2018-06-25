package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// StaticVar node
type StaticVar struct {
	Comments []*comment.Comment
	Position *position.Position
	Variable node.Node
	Expr     node.Node
}

// NewStaticVar node constructor
func NewStaticVar(Variable node.Node, Expr node.Node) *StaticVar {
	return &StaticVar{
		Variable: Variable,
		Expr:     Expr,
	}
}

// SetPosition sets node position
func (n *StaticVar) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *StaticVar) GetPosition() *position.Position {
	return n.Position
}

func (n *StaticVar) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *StaticVar) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *StaticVar) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticVar) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
