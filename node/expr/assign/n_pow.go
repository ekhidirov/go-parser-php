package assign

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Pow node
type Pow struct {
	Comments   []*comment.Comment
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewPow node constructor
func NewPow(Variable node.Node, Expression node.Node) *Pow {
	return &Pow{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *Pow) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Pow) GetPosition() *position.Position {
	return n.Position
}

func (n *Pow) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Pow) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Pow) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Pow) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expression != nil {
		v.EnterChildNode("Expression", n)
		n.Expression.Walk(v)
		v.LeaveChildNode("Expression", n)
	}

	v.LeaveNode(n)
}
