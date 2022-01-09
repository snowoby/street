package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	"street/ent/episode"
	"street/ent/profile"
)

func (s *db) CreateEpisode(ctx context.Context, title, cleanContent string, profileID uuid.UUID) (*ent.Episode, error) {
	return s.client.Episode.Create().SetTitle(title).SetContent(cleanContent).SetProfileID(profileID).Save(ctx)
}

func (s *db) FindEpisode(ctx context.Context, episodeID uuid.UUID) (*ent.Episode, error) {
	return s.client.Episode.Query().Where(episode.ID(episodeID)).WithProfile().Only(ctx)
}

func (s *db) UpdateEpisode(ctx context.Context, episodeID uuid.UUID, title, cleanContent string) (*ent.Episode, error) {
	return s.client.Episode.UpdateOneID(episodeID).SetTitle(title).SetContent(cleanContent).Save(ctx)
}

func (s *db) EpisodeBelongs(ctx context.Context, profileID uuid.UUID, episodeID uuid.UUID) (bool, error) {
	return s.client.Episode.Query().
		Where(
			episode.And(
				episode.ID(episodeID),
				episode.HasProfileWith(profile.ID(profileID))),
		).Exist(ctx)

}

func (s *db) DeleteEpisode(ctx context.Context, episodeID uuid.UUID) error {
	return s.client.Episode.DeleteOneID(episodeID).Exec(ctx)
}
