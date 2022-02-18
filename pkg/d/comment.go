package d

import "street/ent"

type CommentForm struct {
	To      string `json:"to" binding:"required"`
	From    string `json:"from" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Comment struct {
	*ent.Comment
}

func CommentFromEnt(c *ent.Comment) *Comment {
	return &Comment{c}
}

func CommentsFromEnt(cs []*ent.Comment) []*Comment {
	result := make([]*Comment, len(cs))
	for i, c := range cs {
		result[i] = CommentFromEnt(c)
	}
	return result
}
