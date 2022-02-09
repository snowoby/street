package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	ee "street/ent/episode"
	ep "street/ent/profile"
	"street/ent/schema"
)

type episode struct {
	client *ent.EpisodeClient
}

func (e *episode) All(ctx context.Context) ([]*ent.Episode, error) {
	return e.client.Query().WithProfile().WithSeries().All(ctx)
}

func (e *episode) Create(ctx context.Context, title, cleanContent string, profileID uuid.UUID, navPicture string) (*ent.Episode, error) {
	return e.client.Create().SetTitle(title).SetContent(cleanContent).SetProfileID(profileID).SetExtra(schema.EpisodeExtra{
		NavPicture: navPicture,
	}).Save(ctx)
}

func (e *episode) FindByIDWithProfile(ctx context.Context, episodeID uuid.UUID) (*ent.Episode, error) {
	return e.client.Query().Where(ee.ID(episodeID)).WithProfile().Only(ctx)
}

func (e *episode) Update(ctx context.Context, episodeID uuid.UUID, title, cleanContent string) (*ent.Episode, error) {
	return e.client.UpdateOneID(episodeID).SetTitle(title).SetContent(cleanContent).Save(ctx)
}

func (e *episode) Delete(ctx context.Context, episodeID uuid.UUID) error {
	return e.client.DeleteOneID(episodeID).Exec(ctx)
}

func (e *episode) IsOwner(ctx context.Context, profileID uuid.UUID, episodeID uuid.UUID) (bool, error) {
	return e.client.Query().
		Where(
			ee.And(
				ee.ID(episodeID),
				ee.HasProfileWith(ep.ID(profileID))),
		).Exist(ctx)
}

func (e *episode) OwnerID(ctx context.Context, episodeID uuid.UUID) (string, error) {
	episode, err := e.FindByIDWithProfile(ctx, episodeID)
	if err != nil {
		return "", err
	}
	return episode.Edges.Profile.ID.String(), nil
}
