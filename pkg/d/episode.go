package d

import (
	"street/ent"
	"street/ent/schema"
)

type Episode struct {
	*ent.Episode
	Profile  *Profile   `json:"profile"`
	Series   *Series    `json:"series"`
	Comments []*Comment `json:"comments"`
	NoEdges
	ValueType
}

type EpisodeWithCommentCount struct {
	*Episode
	CommentCount int `json:"commentCount" binding:"omitempty"`
}

type EpisodeForm struct {
	ProfileID string        `json:"profileID" binding:"uuid,required"`
	SeriesID  *string       `json:"seriesID" binding:"omitempty,uuid"`
	Title     *string       `json:"title" binding:"omitempty"`
	Content   string        `json:"content"`
	Cover     *string       `json:"cover" binding:"omitempty"`
	Files     schema.Medias `json:"files" binding:"omitempty,dive"`
}

func EpisodeWithCommentCountFromEnt(episode *ent.Episode, count int) *EpisodeWithCommentCount {
	ep := EpisodeFromEnt(episode)
	return &EpisodeWithCommentCount{
		Episode:      ep,
		CommentCount: count,
	}
}

func EpisodeWithCommentCountsFromEnt(episodes []*ent.Episode, counts []*IDCount) []*EpisodeWithCommentCount {
	result := make([]*EpisodeWithCommentCount, len(episodes))
	flagCount := 0
	for i, episode := range episodes {

		if flagCount < len(counts) && episode.ID == counts[flagCount].ID {
			result[i] = EpisodeWithCommentCountFromEnt(episode, counts[flagCount].Count)
			flagCount++
		} else {
			result[i] = EpisodeWithCommentCountFromEnt(episode, 0)
		}
	}
	return result
}

func EpisodeFromEnt(episode *ent.Episode) *Episode {
	return &Episode{
		Episode:   episode,
		Profile:   ProfileFromEnt(episode.Edges.Profile),
		Series:    SeriesFromEnt(episode.Edges.Series),
		Comments:  CommentsFromEnt(episode.Edges.Comments),
		ValueType: ValueType{"episode"},
	}
}

func EpisodesFromEnt(episodes []*ent.Episode) []*Episode {
	result := make([]*Episode, len(episodes))
	for i, episode := range episodes {
		result[i] = EpisodeFromEnt(episode)
	}
	return result
}
