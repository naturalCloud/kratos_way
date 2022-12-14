// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"
	"user/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickName), v))
	})
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthday), v))
	})
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// AddTime applies equality check predicate on the "add_time" field. It's identical to AddTimeEQ.
func AddTime(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// IsDeletedAt applies equality check predicate on the "is_deleted_at" field. It's identical to IsDeletedAtEQ.
func IsDeletedAt(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeletedAt), v))
	})
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMobile), v))
	})
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMobile), v...))
	})
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMobile), v...))
	})
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMobile), v))
	})
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMobile), v))
	})
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMobile), v))
	})
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMobile), v))
	})
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMobile), v))
	})
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMobile), v))
	})
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMobile), v))
	})
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMobile), v))
	})
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMobile), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickName), v))
	})
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickName), v))
	})
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNickName), v...))
	})
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNickName), v...))
	})
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickName), v))
	})
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickName), v))
	})
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickName), v))
	})
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickName), v))
	})
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickName), v))
	})
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickName), v))
	})
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickName), v))
	})
}

// NickNameIsNil applies the IsNil predicate on the "nick_name" field.
func NickNameIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNickName)))
	})
}

// NickNameNotNil applies the NotNil predicate on the "nick_name" field.
func NickNameNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNickName)))
	})
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickName), v))
	})
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickName), v))
	})
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBirthday), v))
	})
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBirthday), v))
	})
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBirthday), v...))
	})
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBirthday), v...))
	})
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBirthday), v))
	})
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBirthday), v))
	})
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBirthday), v))
	})
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBirthday), v))
	})
}

// BirthdayIsNil applies the IsNil predicate on the "birthday" field.
func BirthdayIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBirthday)))
	})
}

// BirthdayNotNil applies the NotNil predicate on the "birthday" field.
func BirthdayNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBirthday)))
	})
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGender), v))
	})
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGender), v...))
	})
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGender), v...))
	})
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGender), v))
	})
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGender), v))
	})
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGender), v))
	})
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGender), v))
	})
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGender), v))
	})
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGender), v))
	})
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGender), v))
	})
}

// GenderIsNil applies the IsNil predicate on the "gender" field.
func GenderIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGender)))
	})
}

// GenderNotNil applies the NotNil predicate on the "gender" field.
func GenderNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGender)))
	})
}

// GenderEqualFold applies the EqualFold predicate on the "gender" field.
func GenderEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGender), v))
	})
}

// GenderContainsFold applies the ContainsFold predicate on the "gender" field.
func GenderContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGender), v))
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRole), v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRole), v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...int32) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRole), v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...int32) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRole), v...))
	})
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRole), v))
	})
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRole), v))
	})
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRole), v))
	})
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v int32) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRole), v))
	})
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRole)))
	})
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRole)))
	})
}

// AddTimeEQ applies the EQ predicate on the "add_time" field.
func AddTimeEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddTime), v))
	})
}

// AddTimeNEQ applies the NEQ predicate on the "add_time" field.
func AddTimeNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddTime), v))
	})
}

// AddTimeIn applies the In predicate on the "add_time" field.
func AddTimeIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddTime), v...))
	})
}

// AddTimeNotIn applies the NotIn predicate on the "add_time" field.
func AddTimeNotIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddTime), v...))
	})
}

// AddTimeGT applies the GT predicate on the "add_time" field.
func AddTimeGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddTime), v))
	})
}

// AddTimeGTE applies the GTE predicate on the "add_time" field.
func AddTimeGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddTime), v))
	})
}

// AddTimeLT applies the LT predicate on the "add_time" field.
func AddTimeLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddTime), v))
	})
}

// AddTimeLTE applies the LTE predicate on the "add_time" field.
func AddTimeLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddTime), v))
	})
}

// AddTimeIsNil applies the IsNil predicate on the "add_time" field.
func AddTimeIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAddTime)))
	})
}

// AddTimeNotNil applies the NotNil predicate on the "add_time" field.
func AddTimeNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAddTime)))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateTime)))
	})
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateTime)))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// IsDeletedAtEQ applies the EQ predicate on the "is_deleted_at" field.
func IsDeletedAtEQ(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtNEQ applies the NEQ predicate on the "is_deleted_at" field.
func IsDeletedAtNEQ(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtIn applies the In predicate on the "is_deleted_at" field.
func IsDeletedAtIn(vs ...int8) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIsDeletedAt), v...))
	})
}

// IsDeletedAtNotIn applies the NotIn predicate on the "is_deleted_at" field.
func IsDeletedAtNotIn(vs ...int8) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIsDeletedAt), v...))
	})
}

// IsDeletedAtGT applies the GT predicate on the "is_deleted_at" field.
func IsDeletedAtGT(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtGTE applies the GTE predicate on the "is_deleted_at" field.
func IsDeletedAtGTE(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtLT applies the LT predicate on the "is_deleted_at" field.
func IsDeletedAtLT(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtLTE applies the LTE predicate on the "is_deleted_at" field.
func IsDeletedAtLTE(v int8) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeletedAt), v))
	})
}

// IsDeletedAtIsNil applies the IsNil predicate on the "is_deleted_at" field.
func IsDeletedAtIsNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsDeletedAt)))
	})
}

// IsDeletedAtNotNil applies the NotNil predicate on the "is_deleted_at" field.
func IsDeletedAtNotNil() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsDeletedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
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
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		p(s.Not())
	})
}
