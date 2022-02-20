package d

import "street/ent"

type CommentForm struct {
	To      string `json:"to" binding:"required"`
	From    string `json:"from" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Comment struct {
	*ent.Comment
	ValueType
	NoEdges
}

func CommentFromEnt(c *ent.Comment) *Comment {
	return &Comment{
		Comment:   c,
		ValueType: ValueType{"comment"},
	}
}

func CommentsFromEnt(cs []*ent.Comment) []*Comment {
	result := make([]*Comment, len(cs))
	for i, c := range cs {
		result[i] = CommentFromEnt(c)
	}
	return result
}
