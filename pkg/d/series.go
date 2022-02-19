package d

import "street/ent"

type Series struct {
	*ent.Series
	Profile *Profile `json:"profile,omitempty"`
}

type SeriesForm struct {
	Title     string `json:"title" binding:"required"`
	Type      string `json:"type" binding:"required"`
	ProfileID string `json:"profileID" binding:"required,uuid"`
}

func SeriesFromEnt(series *ent.Series) *Series {
	return &Series{
		Series:  series,
		Profile: ProfileFromEnt(series.Edges.Owner),
	}
}

func SeriesManyFromEnt(manySeries []*ent.Series) []*Series {
	series := make([]*Series, len(manySeries))
	for i, s := range manySeries {
		series[i] = SeriesFromEnt(s)
	}

	return series
}
