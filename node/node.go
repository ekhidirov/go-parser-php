package node

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Node interface
type Node interface {
	walker.Walkable
	Attributes() map[string]interface{} // Attributes returns node attributes as map
	SetPosition(p *position.Position)
	GetPosition() *position.Position
	AddComments(c []*comment.Comment, t comment.TokenName)
	GetComments() []*comment.Comment
}
