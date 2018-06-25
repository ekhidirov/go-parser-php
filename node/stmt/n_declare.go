package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Declare node
type Declare struct {
	Comments []*comment.Comment
	Position *position.Position
	Consts   []node.Node
	Stmt     node.Node
}

// NewDeclare node constructor
func NewDeclare(Consts []node.Node, Stmt node.Node) *Declare {
	return &Declare{
		Consts: Consts,
		Stmt:   Stmt,
	}
}

// SetPosition sets node position
func (n *Declare) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Declare) GetPosition() *position.Position {
	return n.Position
}

func (n *Declare) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Declare) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Declare) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Declare) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
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

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
