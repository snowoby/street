package d

import "street/ent"

type CommentForm struct {
	From    string  `json:"from" binding:"uuid"`
	Content string  `json:"content" binding:"required"`
	ReplyTo *string `json:"replyTo" binding:"omitempty,uuid"`
}

type Comment struct {
	*ent.Comment
	ReplyTo *Comment `json:"replyTo,omitempty"`
	Author  *Profile `json:"author"`
	ValueType
	NoEdges
	NoPath
}

func CommentFromEnt(c *ent.Comment) *Comment {
	if c == nil {
		return nil
	}
	return &Comment{
		Comment:   c,
		ReplyTo:   CommentFromEnt(c.Edges.ReplyTo),
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
