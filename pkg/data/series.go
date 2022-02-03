package data

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	ee "street/ent/episode"
	ep "street/ent/profile"
	es "street/ent/series"
)

type series struct {
	client *ent.SeriesClient
}

func (s *series) Create(ctx context.Context, title, cleanContent string, profileID uuid.UUID) (*ent.Series, error) {
	return s.client.Create().SetTitle(title).SetContent(cleanContent).SetProfileID(profileID).Save(ctx)
}

func (s *series) FindByCallSign(ctx context.Context, callSign string) (*ent.Series, error) {
	return s.client.Query().Where(es.Call(callSign)).WithProfile().Only(ctx)
}
func (s *series) FindByID(ctx context.Context, id uuid.UUID) (*ent.Series, error) {
	return s.client.Query().Where(es.ID(id)).WithProfile().Only(ctx)
}

func (s *series) Update(ctx context.Context, seriesID uuid.UUID, title, cleanContent string) (*ent.Series, error) {
	return s.client.UpdateOneID(seriesID).SetTitle(title).SetContent(cleanContent).Save(ctx)
}

func (s *series) IsOwner(ctx context.Context, profileID uuid.UUID, seriesID uuid.UUID) (bool, error) {
	return s.client.Query().
		Where(
			es.And(
				es.ID(seriesID),
				es.HasProfileWith(ep.ID(profileID))),
		).Exist(ctx)

}

func (s *series) HasEpisode(ctx context.Context, episodeID uuid.UUID, seriesID uuid.UUID) (bool, error) {
	return s.client.Query().
		Where(
			es.And(
				es.ID(seriesID),
				es.HasEpisodeWith(ee.ID(episodeID))),
		).Exist(ctx)

}

func (s *series) Delete(ctx context.Context, seriesID uuid.UUID) error {
	return s.client.DeleteOneID(seriesID).Exec(ctx)
}

func (s *series) SetCall(ctx context.Context, seriesID uuid.UUID, callSign string) (*ent.Series, error) {
	return s.client.UpdateOneID(seriesID).SetCall(callSign).Save(ctx)
}
