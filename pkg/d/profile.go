package d

import "street/ent"

type Profile struct {
	*ent.Profile
	NoEdges
	ValueType
}

type ProfileForm struct {
	Title    string `json:"title" binding:"required"`
	Call     string `json:"call" binding:"required"`
	Category string `json:"category" binding:"required"`
	Avatar   string `json:"avatar"`
}

func ProfileFromEnt(profile *ent.Profile) *Profile {
	return &Profile{
		Profile:   profile,
		ValueType: ValueType{"profile"},
	}
}

func ProfilesFromEnt(profiles []*ent.Profile) []*Profile {
	result := make([]*Profile, len(profiles))
	for i, profile := range profiles {
		result[i] = ProfileFromEnt(profile)
	}
	return result
}
