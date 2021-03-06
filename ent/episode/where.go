// Code generated by entc, DO NOT EDIT.

package episode

import (
	"street/ent/predicate"
	"street/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Sid applies equality check predicate on the "sid" field. It's identical to SidEQ.
func Sid(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSid), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Cover applies equality check predicate on the "cover" field. It's identical to CoverEQ.
func Cover(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCover), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Files applies equality check predicate on the "files" field. It's identical to FilesEQ.
func Files(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiles), v))
	})
}

// SidEQ applies the EQ predicate on the "sid" field.
func SidEQ(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSid), v))
	})
}

// SidNEQ applies the NEQ predicate on the "sid" field.
func SidNEQ(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSid), v))
	})
}

// SidIn applies the In predicate on the "sid" field.
func SidIn(vs ...schema.ID) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSid), v...))
	})
}

// SidNotIn applies the NotIn predicate on the "sid" field.
func SidNotIn(vs ...schema.ID) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSid), v...))
	})
}

// SidGT applies the GT predicate on the "sid" field.
func SidGT(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSid), v))
	})
}

// SidGTE applies the GTE predicate on the "sid" field.
func SidGTE(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSid), v))
	})
}

// SidLT applies the LT predicate on the "sid" field.
func SidLT(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSid), v))
	})
}

// SidLTE applies the LTE predicate on the "sid" field.
func SidLTE(v schema.ID) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSid), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// CoverEQ applies the EQ predicate on the "cover" field.
func CoverEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCover), v))
	})
}

// CoverNEQ applies the NEQ predicate on the "cover" field.
func CoverNEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCover), v))
	})
}

// CoverIn applies the In predicate on the "cover" field.
func CoverIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCover), v...))
	})
}

// CoverNotIn applies the NotIn predicate on the "cover" field.
func CoverNotIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCover), v...))
	})
}

// CoverGT applies the GT predicate on the "cover" field.
func CoverGT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCover), v))
	})
}

// CoverGTE applies the GTE predicate on the "cover" field.
func CoverGTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCover), v))
	})
}

// CoverLT applies the LT predicate on the "cover" field.
func CoverLT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCover), v))
	})
}

// CoverLTE applies the LTE predicate on the "cover" field.
func CoverLTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCover), v))
	})
}

// CoverContains applies the Contains predicate on the "cover" field.
func CoverContains(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCover), v))
	})
}

// CoverHasPrefix applies the HasPrefix predicate on the "cover" field.
func CoverHasPrefix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCover), v))
	})
}

// CoverHasSuffix applies the HasSuffix predicate on the "cover" field.
func CoverHasSuffix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCover), v))
	})
}

// CoverIsNil applies the IsNil predicate on the "cover" field.
func CoverIsNil() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCover)))
	})
}

// CoverNotNil applies the NotNil predicate on the "cover" field.
func CoverNotNil() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCover)))
	})
}

// CoverEqualFold applies the EqualFold predicate on the "cover" field.
func CoverEqualFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCover), v))
	})
}

// CoverContainsFold applies the ContainsFold predicate on the "cover" field.
func CoverContainsFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCover), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// FilesEQ applies the EQ predicate on the "files" field.
func FilesEQ(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiles), v))
	})
}

// FilesNEQ applies the NEQ predicate on the "files" field.
func FilesNEQ(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFiles), v))
	})
}

// FilesIn applies the In predicate on the "files" field.
func FilesIn(vs ...schema.Medias) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFiles), v...))
	})
}

// FilesNotIn applies the NotIn predicate on the "files" field.
func FilesNotIn(vs ...schema.Medias) predicate.Episode {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Episode(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFiles), v...))
	})
}

// FilesGT applies the GT predicate on the "files" field.
func FilesGT(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFiles), v))
	})
}

// FilesGTE applies the GTE predicate on the "files" field.
func FilesGTE(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFiles), v))
	})
}

// FilesLT applies the LT predicate on the "files" field.
func FilesLT(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFiles), v))
	})
}

// FilesLTE applies the LTE predicate on the "files" field.
func FilesLTE(v schema.Medias) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFiles), v))
	})
}

// HasProfile applies the HasEdge predicate on the "profile" edge.
func HasProfile() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfileTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfileWith applies the HasEdge predicate on the "profile" edge with a given conditions (other predicates).
func HasProfileWith(preds ...predicate.Profile) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSeries applies the HasEdge predicate on the "series" edge.
func HasSeries() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SeriesTable, SeriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeriesWith applies the HasEdge predicate on the "series" edge with a given conditions (other predicates).
func HasSeriesWith(preds ...predicate.Series) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SeriesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SeriesTable, SeriesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Episode) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		p(s.Not())
	})
}
