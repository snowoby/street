package data

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	"street/ent/episode"
	"street/ent/profile"
	es "street/ent/series"
)

type series struct {
	*ent.SeriesClient
}

func (s *series) Create(ctx context.Context, title, cleanContent string, profileID uuid.UUID) (*ent.Series, error) {
	return s.SeriesClient.Create().SetTitle(title).SetContent(cleanContent).SetProfileID(profileID).Save(ctx)
}

func (s *series) FindByCallSign(ctx context.Context, callSign string) (*ent.Series, error) {
	return s.Query().Where(es.CallSign(callSign)).WithProfile().Only(ctx)
}
func (s *series) FindByID(ctx context.Context, id uuid.UUID) (*ent.Series, error) {
	return s.Query().Where(es.ID(id)).WithProfile().Only(ctx)
}

func (s *series) Update(ctx context.Context, seriesID uuid.UUID, title, cleanContent string) (*ent.Series, error) {
	return s.UpdateOneID(seriesID).SetTitle(title).SetContent(cleanContent).Save(ctx)
}

func (s *series) IsOwner(ctx context.Context, profileID uuid.UUID, seriesID uuid.UUID) (bool, error) {
	return s.Query().
		Where(
			es.And(
				es.ID(seriesID),
				es.HasProfileWith(profile.ID(profileID))),
		).Exist(ctx)

}

func (s *series) HasEpisode(ctx context.Context, episodeID uuid.UUID, seriesID uuid.UUID) (bool, error) {
	return s.Query().
		Where(
			es.And(
				es.ID(seriesID),
				es.HasEpisodeWith(episode.ID(episodeID))),
		).Exist(ctx)

}

func (s *series) Delete(ctx context.Context, seriesID uuid.UUID) error {
	return s.DeleteOneID(seriesID).Exec(ctx)
}

func (s *series) SetCallSign(ctx context.Context, seriesID uuid.UUID, callSign string) (*ent.Series, error) {
	return s.UpdateOneID(seriesID).SetCallSign(callSign).Save(ctx)
}

//
//func (s *db) CreateSeries(ctx context.Context, title, cleanContent string, profileID uuid.UUID) (*ent.Series, error) {
//	return s.client.Series.Create().SetTitle(title).SetContent(cleanContent).SetProfileID(profileID).Save(ctx)
//}
//
//func (s *db) FindSeries(ctx context.Context, callSign string) (*ent.Series, error) {
//	return s.client.Series.Query().Where(series.CallSign(callSign)).WithProfile().Only(ctx)
//}
//func (s *db) FindSeriesByID(ctx context.Context, id uuid.UUID) (*ent.Series, error) {
//	return s.client.Series.Query().Where(series.ID(id)).WithProfile().Only(ctx)
//}
//
//func (s *db) UpdateSeries(ctx context.Context, seriesID uuid.UUID, title, cleanContent string) (*ent.Series, error) {
//	return s.client.Series.UpdateOneID(seriesID).SetTitle(title).SetContent(cleanContent).Save(ctx)
//}
//
//func (s *db) SeriesBelongs(ctx context.Context, profileID uuid.UUID, seriesID uuid.UUID) (bool, error) {
//	return s.client.Series.Query().
//		Where(
//			series.And(
//				series.ID(seriesID),
//				series.HasProfileWith(profile.ID(profileID))),
//		).Exist(ctx)
//
//}
//
//func (s *db) EpisodeBelongsSeries(ctx context.Context, episodeID uuid.UUID, seriesID uuid.UUID) (bool, error) {
//	return s.client.Series.Query().
//		Where(
//			series.And(
//				series.ID(seriesID),
//				series.HasEpisodeWith(episode.ID(episodeID))),
//		).Exist(ctx)
//
//}
//
//func (s *db) DeleteSeries(ctx context.Context, seriesID uuid.UUID) error {
//	return s.client.Series.DeleteOneID(seriesID).Exec(ctx)
//}
//
//func (s *db) SetSeriesCallSign(ctx context.Context, seriesID uuid.UUID, callSign string) (*ent.Series, error) {
//	return s.client.Series.UpdateOneID(seriesID).SetCallSign(callSign).Save(ctx)
//}
