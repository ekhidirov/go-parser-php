package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Continue node
type Continue struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewContinue node constructor
func NewContinue(Expr node.Node) *Continue {
	return &Continue{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Continue) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Continue) GetPosition() *position.Position {
	return n.Position
}

func (n *Continue) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Continue) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Continue) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Continue) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
