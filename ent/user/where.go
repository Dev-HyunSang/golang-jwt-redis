// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserUUID), v))
	})
}

// UserEmail applies equality check predicate on the "user_email" field. It's identical to UserEmailEQ.
func UserEmail(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserEmail), v))
	})
}

// UserPassword applies equality check predicate on the "user_password" field. It's identical to UserPasswordEQ.
func UserPassword(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserPassword), v))
	})
}

// UserNickname applies equality check predicate on the "user_nickname" field. It's identical to UserNicknameEQ.
func UserNickname(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserNickname), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserUUID), v))
	})
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserUUID), v))
	})
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...uuid.UUID) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserUUID), v...))
	})
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...uuid.UUID) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserUUID), v...))
	})
}

// UserUUIDGT applies the GT predicate on the "user_uuid" field.
func UserUUIDGT(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserUUID), v))
	})
}

// UserUUIDGTE applies the GTE predicate on the "user_uuid" field.
func UserUUIDGTE(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserUUID), v))
	})
}

// UserUUIDLT applies the LT predicate on the "user_uuid" field.
func UserUUIDLT(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserUUID), v))
	})
}

// UserUUIDLTE applies the LTE predicate on the "user_uuid" field.
func UserUUIDLTE(v uuid.UUID) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserUUID), v))
	})
}

// UserEmailEQ applies the EQ predicate on the "user_email" field.
func UserEmailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserEmail), v))
	})
}

// UserEmailNEQ applies the NEQ predicate on the "user_email" field.
func UserEmailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserEmail), v))
	})
}

// UserEmailIn applies the In predicate on the "user_email" field.
func UserEmailIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserEmail), v...))
	})
}

// UserEmailNotIn applies the NotIn predicate on the "user_email" field.
func UserEmailNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserEmail), v...))
	})
}

// UserEmailGT applies the GT predicate on the "user_email" field.
func UserEmailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserEmail), v))
	})
}

// UserEmailGTE applies the GTE predicate on the "user_email" field.
func UserEmailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserEmail), v))
	})
}

// UserEmailLT applies the LT predicate on the "user_email" field.
func UserEmailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserEmail), v))
	})
}

// UserEmailLTE applies the LTE predicate on the "user_email" field.
func UserEmailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserEmail), v))
	})
}

// UserEmailContains applies the Contains predicate on the "user_email" field.
func UserEmailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserEmail), v))
	})
}

// UserEmailHasPrefix applies the HasPrefix predicate on the "user_email" field.
func UserEmailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserEmail), v))
	})
}

// UserEmailHasSuffix applies the HasSuffix predicate on the "user_email" field.
func UserEmailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserEmail), v))
	})
}

// UserEmailEqualFold applies the EqualFold predicate on the "user_email" field.
func UserEmailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserEmail), v))
	})
}

// UserEmailContainsFold applies the ContainsFold predicate on the "user_email" field.
func UserEmailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserEmail), v))
	})
}

// UserPasswordEQ applies the EQ predicate on the "user_password" field.
func UserPasswordEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserPassword), v))
	})
}

// UserPasswordNEQ applies the NEQ predicate on the "user_password" field.
func UserPasswordNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserPassword), v))
	})
}

// UserPasswordIn applies the In predicate on the "user_password" field.
func UserPasswordIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserPassword), v...))
	})
}

// UserPasswordNotIn applies the NotIn predicate on the "user_password" field.
func UserPasswordNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserPassword), v...))
	})
}

// UserPasswordGT applies the GT predicate on the "user_password" field.
func UserPasswordGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserPassword), v))
	})
}

// UserPasswordGTE applies the GTE predicate on the "user_password" field.
func UserPasswordGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserPassword), v))
	})
}

// UserPasswordLT applies the LT predicate on the "user_password" field.
func UserPasswordLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserPassword), v))
	})
}

// UserPasswordLTE applies the LTE predicate on the "user_password" field.
func UserPasswordLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserPassword), v))
	})
}

// UserPasswordContains applies the Contains predicate on the "user_password" field.
func UserPasswordContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserPassword), v))
	})
}

// UserPasswordHasPrefix applies the HasPrefix predicate on the "user_password" field.
func UserPasswordHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserPassword), v))
	})
}

// UserPasswordHasSuffix applies the HasSuffix predicate on the "user_password" field.
func UserPasswordHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserPassword), v))
	})
}

// UserPasswordEqualFold applies the EqualFold predicate on the "user_password" field.
func UserPasswordEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserPassword), v))
	})
}

// UserPasswordContainsFold applies the ContainsFold predicate on the "user_password" field.
func UserPasswordContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserPassword), v))
	})
}

// UserNicknameEQ applies the EQ predicate on the "user_nickname" field.
func UserNicknameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserNickname), v))
	})
}

// UserNicknameNEQ applies the NEQ predicate on the "user_nickname" field.
func UserNicknameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserNickname), v))
	})
}

// UserNicknameIn applies the In predicate on the "user_nickname" field.
func UserNicknameIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserNickname), v...))
	})
}

// UserNicknameNotIn applies the NotIn predicate on the "user_nickname" field.
func UserNicknameNotIn(vs ...string) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserNickname), v...))
	})
}

// UserNicknameGT applies the GT predicate on the "user_nickname" field.
func UserNicknameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserNickname), v))
	})
}

// UserNicknameGTE applies the GTE predicate on the "user_nickname" field.
func UserNicknameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserNickname), v))
	})
}

// UserNicknameLT applies the LT predicate on the "user_nickname" field.
func UserNicknameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserNickname), v))
	})
}

// UserNicknameLTE applies the LTE predicate on the "user_nickname" field.
func UserNicknameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserNickname), v))
	})
}

// UserNicknameContains applies the Contains predicate on the "user_nickname" field.
func UserNicknameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserNickname), v))
	})
}

// UserNicknameHasPrefix applies the HasPrefix predicate on the "user_nickname" field.
func UserNicknameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserNickname), v))
	})
}

// UserNicknameHasSuffix applies the HasSuffix predicate on the "user_nickname" field.
func UserNicknameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserNickname), v))
	})
}

// UserNicknameEqualFold applies the EqualFold predicate on the "user_nickname" field.
func UserNicknameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserNickname), v))
	})
}

// UserNicknameContainsFold applies the ContainsFold predicate on the "user_nickname" field.
func UserNicknameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserNickname), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
