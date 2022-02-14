package d

import "street/ent"

type Episode struct {
	*ent.Episode
	Profile *Profile
}

type EpisodeForm struct {
	ProfileID string `json:"profileID" binding:"uuid,required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content"`
	Cover     string `json:"cover"`
}

func EpisodeFromEnt(episode *ent.Episode) *Episode {
	return &Episode{
		Episode: episode,
		Profile: ProfileFromEnt(episode.Edges.Profile),
	}
}

func EpisodesFromEnt(episodes []*ent.Episode) []*Episode {
	result := make([]*Episode, len(episodes))
	for i, episode := range episodes {
		result[i] = EpisodeFromEnt(episode)
	}
	return result
}
