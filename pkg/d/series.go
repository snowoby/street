package d

import "street/ent"

type Series struct {
	*ent.Series
	Profile *Profile `json:"profile,omitempty"`
	NoEdges
	ValueType
}

type SeriesForm struct {
	Title     string `json:"title" binding:"required"`
	Type      string `json:"type" binding:"required"`
	ProfileID string `json:"profileID" binding:"required,uuid"`
}

func SeriesFromEnt(series *ent.Series) *Series {
	if series == nil {
		return nil
	}
	return &Series{
		Series:    series,
		Profile:   ProfileFromEnt(series.Edges.Owner),
		ValueType: ValueType{"series"},
	}
}

func SeriesManyFromEnt(manySeries []*ent.Series) []*Series {
	series := make([]*Series, len(manySeries))
	for i, s := range manySeries {
		series[i] = SeriesFromEnt(s)
	}

	return series
}
