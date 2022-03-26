package d

import "street/ent"

type CommentForm struct {
	From    string `json:"from" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Comment struct {
	*ent.Comment
	Author *Profile `json:"author"`
	ValueType
	NoEdges
}

func CommentFromEnt(c *ent.Comment) *Comment {
	return &Comment{
		Comment:   c,
		Author:    ProfileFromEnt(c.Edges.Author),
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
