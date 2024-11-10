package wcAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mAuth/rqAuth"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// FeedbacksMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcAuth__ORM.GEN.go
type FeedbacksMutator struct {
	rqAuth.Feedbacks
	mutations []A.X
	logs      []A.X
}

// NewFeedbacksMutator create new ORM writer/command object
func NewFeedbacksMutator(adapter *Tt.Adapter) (res *FeedbacksMutator) {
	res = &FeedbacksMutator{Feedbacks: rqAuth.Feedbacks{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (f *FeedbacksMutator) Logs() []A.X { //nolint:dupl false positive
	return f.logs
}

// HaveMutation check whether Set* methods ever called
func (f *FeedbacksMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(f.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (f *FeedbacksMutator) ClearMutations() { //nolint:dupl false positive
	f.mutations = []A.X{}
	f.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (f *FeedbacksMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.ToUpdateArray())
	return !L.IsError(err, `Feedbacks.DoOverwriteById failed: `+f.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (f *FeedbacksMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !f.HaveMutation() {
		return true
	}
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.mutations)
	return !L.IsError(err, `Feedbacks.DoUpdateById failed: `+f.SpaceName())
}

// DoDeletePermanentById permanent delete
func (f *FeedbacksMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Delete(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id})
	return !L.IsError(err, `Feedbacks.DoDeletePermanentById failed: `+f.SpaceName())
}

// func (f *FeedbacksMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := f.ToArray()
//	_, err := f.Adapter.Upsert(f.SpaceName(), arr, A.X{
//		A.X{`=`, 0, f.Id},
//		A.X{`=`, 1, f.CreatedBy},
//		A.X{`=`, 2, f.CreatedAt},
//		A.X{`=`, 3, f.UpdatedBy},
//		A.X{`=`, 4, f.UpdatedAt},
//		A.X{`=`, 5, f.DeletedAt},
//		A.X{`=`, 6, f.UserMessage},
//		A.X{`=`, 7, f.AdminReply},
//	})
//	return !L.IsError(err, `Feedbacks.DoUpsert failed: `+f.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (f *FeedbacksMutator) DoInsert() bool { //nolint:dupl false positive
	arr := f.ToArray()
	row, err := f.Adapter.Insert(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Feedbacks.DoInsert failed: `+f.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (f *FeedbacksMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := f.ToArray()
	row, err := f.Adapter.Replace(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Feedbacks.DoUpsert failed: `+f.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (f *FeedbacksMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != f.Id {
		f.mutations = append(f.mutations, A.X{`=`, 0, val})
		f.logs = append(f.logs, A.X{`id`, f.Id, val})
		f.Id = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (f *FeedbacksMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.CreatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 1, val})
		f.logs = append(f.logs, A.X{`createdBy`, f.CreatedBy, val})
		f.CreatedBy = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (f *FeedbacksMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.CreatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 2, val})
		f.logs = append(f.logs, A.X{`createdAt`, f.CreatedAt, val})
		f.CreatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (f *FeedbacksMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.UpdatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 3, val})
		f.logs = append(f.logs, A.X{`updatedBy`, f.UpdatedBy, val})
		f.UpdatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (f *FeedbacksMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.UpdatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 4, val})
		f.logs = append(f.logs, A.X{`updatedAt`, f.UpdatedAt, val})
		f.UpdatedAt = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (f *FeedbacksMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != f.DeletedAt {
		f.mutations = append(f.mutations, A.X{`=`, 5, val})
		f.logs = append(f.logs, A.X{`deletedAt`, f.DeletedAt, val})
		f.DeletedAt = val
		return true
	}
	return false
}

// SetUserMessage create mutations, should not duplicate
func (f *FeedbacksMutator) SetUserMessage(val string) bool { //nolint:dupl false positive
	if val != f.UserMessage {
		f.mutations = append(f.mutations, A.X{`=`, 6, val})
		f.logs = append(f.logs, A.X{`userMessage`, f.UserMessage, val})
		f.UserMessage = val
		return true
	}
	return false
}

// SetAdminReply create mutations, should not duplicate
func (f *FeedbacksMutator) SetAdminReply(val string) bool { //nolint:dupl false positive
	if val != f.AdminReply {
		f.mutations = append(f.mutations, A.X{`=`, 7, val})
		f.logs = append(f.logs, A.X{`adminReply`, f.AdminReply, val})
		f.AdminReply = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (f *FeedbacksMutator) SetAll(from rqAuth.Feedbacks, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		f.Id = from.Id
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		f.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		f.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		f.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		f.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		f.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`userMessage`] && (forceMap[`userMessage`] || from.UserMessage != ``) {
		f.UserMessage = S.Trim(from.UserMessage)
		changed = true
	}
	if !excludeMap[`adminReply`] && (forceMap[`adminReply`] || from.AdminReply != ``) {
		f.AdminReply = S.Trim(from.AdminReply)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// SessionsMutator DAO writer/command struct
type SessionsMutator struct {
	rqAuth.Sessions
	mutations []A.X
	logs      []A.X
}

// NewSessionsMutator create new ORM writer/command object
func NewSessionsMutator(adapter *Tt.Adapter) (res *SessionsMutator) {
	res = &SessionsMutator{Sessions: rqAuth.Sessions{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *SessionsMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *SessionsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *SessionsMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.SessionToken},
//		A.X{`=`, 1, s.UserId},
//		A.X{`=`, 2, s.ExpiredAt},
//		A.X{`=`, 3, s.Device},
//		A.X{`=`, 4, s.LoginAt},
//		A.X{`=`, 5, s.LoginIPs},
//	})
//	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteBySessionToken update all columns, error if not exists, not using mutations/Set*
func (s *SessionsMutator) DoOverwriteBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.ToUpdateArray())
	return !L.IsError(err, `Sessions.DoOverwriteBySessionToken failed: `+s.SpaceName())
}

// DoUpdateBySessionToken update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SessionsMutator) DoUpdateBySessionToken() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.mutations)
	return !L.IsError(err, `Sessions.DoUpdateBySessionToken failed: `+s.SpaceName())
}

// DoDeletePermanentBySessionToken permanent delete
func (s *SessionsMutator) DoDeletePermanentBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken})
	return !L.IsError(err, `Sessions.DoDeletePermanentBySessionToken failed: `+s.SpaceName())
}

// DoInsert insert, error if already exists
func (s *SessionsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Insert(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Replace(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetSessionToken create mutations, should not duplicate
func (s *SessionsMutator) SetSessionToken(val string) bool { //nolint:dupl false positive
	if val != s.SessionToken {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`sessionToken`, s.SessionToken, val})
		s.SessionToken = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (s *SessionsMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != s.UserId {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`userId`, s.UserId, val})
		s.UserId = val
		return true
	}
	return false
}

// SetExpiredAt create mutations, should not duplicate
func (s *SessionsMutator) SetExpiredAt(val int64) bool { //nolint:dupl false positive
	if val != s.ExpiredAt {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`expiredAt`, s.ExpiredAt, val})
		s.ExpiredAt = val
		return true
	}
	return false
}

// SetDevice create mutations, should not duplicate
func (s *SessionsMutator) SetDevice(val string) bool { //nolint:dupl false positive
	if val != s.Device {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`device`, s.Device, val})
		s.Device = val
		return true
	}
	return false
}

// SetLoginAt create mutations, should not duplicate
func (s *SessionsMutator) SetLoginAt(val int64) bool { //nolint:dupl false positive
	if val != s.LoginAt {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`loginAt`, s.LoginAt, val})
		s.LoginAt = val
		return true
	}
	return false
}

// SetLoginIPs create mutations, should not duplicate
func (s *SessionsMutator) SetLoginIPs(val string) bool { //nolint:dupl false positive
	if val != s.LoginIPs {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`loginIPs`, s.LoginIPs, val})
		s.LoginIPs = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *SessionsMutator) SetAll(from rqAuth.Sessions, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`sessionToken`] && (forceMap[`sessionToken`] || from.SessionToken != ``) {
		s.SessionToken = S.Trim(from.SessionToken)
		changed = true
	}
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		s.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`expiredAt`] && (forceMap[`expiredAt`] || from.ExpiredAt != 0) {
		s.ExpiredAt = from.ExpiredAt
		changed = true
	}
	if !excludeMap[`device`] && (forceMap[`device`] || from.Device != ``) {
		s.Device = S.Trim(from.Device)
		changed = true
	}
	if !excludeMap[`loginAt`] && (forceMap[`loginAt`] || from.LoginAt != 0) {
		s.LoginAt = from.LoginAt
		changed = true
	}
	if !excludeMap[`loginIPs`] && (forceMap[`loginIPs`] || from.LoginIPs != ``) {
		s.LoginIPs = S.Trim(from.LoginIPs)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// UsersMutator DAO writer/command struct
type UsersMutator struct {
	rqAuth.Users
	mutations []A.X
	logs      []A.X
}

// NewUsersMutator create new ORM writer/command object
func NewUsersMutator(adapter *Tt.Adapter) (res *UsersMutator) {
	res = &UsersMutator{Users: rqAuth.Users{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (u *UsersMutator) Logs() []A.X { //nolint:dupl false positive
	return u.logs
}

// HaveMutation check whether Set* methods ever called
func (u *UsersMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(u.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (u *UsersMutator) ClearMutations() { //nolint:dupl false positive
	u.mutations = []A.X{}
	u.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteById failed: `+u.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateById failed: `+u.SpaceName())
}

// DoDeletePermanentById permanent delete
func (u *UsersMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id})
	return !L.IsError(err, `Users.DoDeletePermanentById failed: `+u.SpaceName())
}

// func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := u.ToArray()
//	_, err := u.Adapter.Upsert(u.SpaceName(), arr, A.X{
//		A.X{`=`, 0, u.Id},
//		A.X{`=`, 1, u.Email},
//		A.X{`=`, 2, u.Password},
//		A.X{`=`, 3, u.CreatedAt},
//		A.X{`=`, 4, u.CreatedBy},
//		A.X{`=`, 5, u.UpdatedAt},
//		A.X{`=`, 6, u.UpdatedBy},
//		A.X{`=`, 7, u.DeletedAt},
//		A.X{`=`, 8, u.PasswordSetAt},
//		A.X{`=`, 9, u.SecretCode},
//		A.X{`=`, 10, u.SecretCodeAt},
//		A.X{`=`, 11, u.VerificationSentAt},
//		A.X{`=`, 12, u.VerifiedAt},
//		A.X{`=`, 13, u.LastLoginAt},
//		A.X{`=`, 14, u.FullName},
//		A.X{`=`, 15, u.UserName},
//		A.X{`=`, 16, u.Country},
//		A.X{`=`, 17, u.Language},
//		A.X{`=`, 18, u.PropertyCount},
//		A.X{`=`, 19, u.PropertyBought},
//	})
//	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByEmail update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteByEmail failed: `+u.SpaceName())
}

// DoUpdateByEmail update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateByEmail() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateByEmail failed: `+u.SpaceName())
}

// DoDeletePermanentByEmail permanent delete
func (u *UsersMutator) DoDeletePermanentByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email})
	return !L.IsError(err, `Users.DoDeletePermanentByEmail failed: `+u.SpaceName())
}

// DoOverwriteByUserName update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteByUserName() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteByUserName failed: `+u.SpaceName())
}

// DoUpdateByUserName update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateByUserName() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateByUserName failed: `+u.SpaceName())
}

// DoDeletePermanentByUserName permanent delete
func (u *UsersMutator) DoDeletePermanentByUserName() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName})
	return !L.IsError(err, `Users.DoDeletePermanentByUserName failed: `+u.SpaceName())
}

// DoInsert insert, error if already exists
func (u *UsersMutator) DoInsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Insert(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoInsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Replace(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (u *UsersMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != u.Id {
		u.mutations = append(u.mutations, A.X{`=`, 0, val})
		u.logs = append(u.logs, A.X{`id`, u.Id, val})
		u.Id = val
		return true
	}
	return false
}

// SetEmail create mutations, should not duplicate
func (u *UsersMutator) SetEmail(val string) bool { //nolint:dupl false positive
	if val != u.Email {
		u.mutations = append(u.mutations, A.X{`=`, 1, val})
		u.logs = append(u.logs, A.X{`email`, u.Email, val})
		u.Email = val
		return true
	}
	return false
}

// SetPassword create mutations, should not duplicate
func (u *UsersMutator) SetPassword(val string) bool { //nolint:dupl false positive
	if val != u.Password {
		u.mutations = append(u.mutations, A.X{`=`, 2, val})
		u.Password = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (u *UsersMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.CreatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 3, val})
		u.logs = append(u.logs, A.X{`createdAt`, u.CreatedAt, val})
		u.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (u *UsersMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.CreatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 4, val})
		u.logs = append(u.logs, A.X{`createdBy`, u.CreatedBy, val})
		u.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.UpdatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 5, val})
		u.logs = append(u.logs, A.X{`updatedAt`, u.UpdatedAt, val})
		u.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.UpdatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 6, val})
		u.logs = append(u.logs, A.X{`updatedBy`, u.UpdatedBy, val})
		u.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (u *UsersMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != u.DeletedAt {
		u.mutations = append(u.mutations, A.X{`=`, 7, val})
		u.logs = append(u.logs, A.X{`deletedAt`, u.DeletedAt, val})
		u.DeletedAt = val
		return true
	}
	return false
}

// SetPasswordSetAt create mutations, should not duplicate
func (u *UsersMutator) SetPasswordSetAt(val int64) bool { //nolint:dupl false positive
	if val != u.PasswordSetAt {
		u.mutations = append(u.mutations, A.X{`=`, 8, val})
		u.logs = append(u.logs, A.X{`passwordSetAt`, u.PasswordSetAt, val})
		u.PasswordSetAt = val
		return true
	}
	return false
}

// SetSecretCode create mutations, should not duplicate
func (u *UsersMutator) SetSecretCode(val string) bool { //nolint:dupl false positive
	if val != u.SecretCode {
		u.mutations = append(u.mutations, A.X{`=`, 9, val})
		u.SecretCode = val
		return true
	}
	return false
}

// SetSecretCodeAt create mutations, should not duplicate
func (u *UsersMutator) SetSecretCodeAt(val int64) bool { //nolint:dupl false positive
	if val != u.SecretCodeAt {
		u.mutations = append(u.mutations, A.X{`=`, 10, val})
		u.SecretCodeAt = val
		return true
	}
	return false
}

// SetVerificationSentAt create mutations, should not duplicate
func (u *UsersMutator) SetVerificationSentAt(val int64) bool { //nolint:dupl false positive
	if val != u.VerificationSentAt {
		u.mutations = append(u.mutations, A.X{`=`, 11, val})
		u.logs = append(u.logs, A.X{`verificationSentAt`, u.VerificationSentAt, val})
		u.VerificationSentAt = val
		return true
	}
	return false
}

// SetVerifiedAt create mutations, should not duplicate
func (u *UsersMutator) SetVerifiedAt(val int64) bool { //nolint:dupl false positive
	if val != u.VerifiedAt {
		u.mutations = append(u.mutations, A.X{`=`, 12, val})
		u.logs = append(u.logs, A.X{`verifiedAt`, u.VerifiedAt, val})
		u.VerifiedAt = val
		return true
	}
	return false
}

// SetLastLoginAt create mutations, should not duplicate
func (u *UsersMutator) SetLastLoginAt(val int64) bool { //nolint:dupl false positive
	if val != u.LastLoginAt {
		u.mutations = append(u.mutations, A.X{`=`, 13, val})
		u.logs = append(u.logs, A.X{`lastLoginAt`, u.LastLoginAt, val})
		u.LastLoginAt = val
		return true
	}
	return false
}

// SetFullName create mutations, should not duplicate
func (u *UsersMutator) SetFullName(val string) bool { //nolint:dupl false positive
	if val != u.FullName {
		u.mutations = append(u.mutations, A.X{`=`, 14, val})
		u.logs = append(u.logs, A.X{`fullName`, u.FullName, val})
		u.FullName = val
		return true
	}
	return false
}

// SetUserName create mutations, should not duplicate
func (u *UsersMutator) SetUserName(val string) bool { //nolint:dupl false positive
	if val != u.UserName {
		u.mutations = append(u.mutations, A.X{`=`, 15, val})
		u.logs = append(u.logs, A.X{`userName`, u.UserName, val})
		u.UserName = val
		return true
	}
	return false
}

// SetCountry create mutations, should not duplicate
func (u *UsersMutator) SetCountry(val string) bool { //nolint:dupl false positive
	if val != u.Country {
		u.mutations = append(u.mutations, A.X{`=`, 16, val})
		u.logs = append(u.logs, A.X{`country`, u.Country, val})
		u.Country = val
		return true
	}
	return false
}

// SetLanguage create mutations, should not duplicate
func (u *UsersMutator) SetLanguage(val string) bool { //nolint:dupl false positive
	if val != u.Language {
		u.mutations = append(u.mutations, A.X{`=`, 17, val})
		u.logs = append(u.logs, A.X{`language`, u.Language, val})
		u.Language = val
		return true
	}
	return false
}

// SetPropertyCount create mutations, should not duplicate
func (u *UsersMutator) SetPropertyCount(val int64) bool { //nolint:dupl false positive
	if val != u.PropertyCount {
		u.mutations = append(u.mutations, A.X{`=`, 18, val})
		u.logs = append(u.logs, A.X{`propertyCount`, u.PropertyCount, val})
		u.PropertyCount = val
		return true
	}
	return false
}

// SetPropertyBought create mutations, should not duplicate
func (u *UsersMutator) SetPropertyBought(val int64) bool { //nolint:dupl false positive
	if val != u.PropertyBought {
		u.mutations = append(u.mutations, A.X{`=`, 19, val})
		u.logs = append(u.logs, A.X{`propertyBought`, u.PropertyBought, val})
		u.PropertyBought = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (u *UsersMutator) SetAll(from rqAuth.Users, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		u.Id = from.Id
		changed = true
	}
	if !excludeMap[`email`] && (forceMap[`email`] || from.Email != ``) {
		u.Email = S.Trim(from.Email)
		changed = true
	}
	if !excludeMap[`password`] && (forceMap[`password`] || from.Password != ``) {
		u.Password = S.Trim(from.Password)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		u.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		u.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		u.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		u.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		u.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`passwordSetAt`] && (forceMap[`passwordSetAt`] || from.PasswordSetAt != 0) {
		u.PasswordSetAt = from.PasswordSetAt
		changed = true
	}
	if !excludeMap[`secretCode`] && (forceMap[`secretCode`] || from.SecretCode != ``) {
		u.SecretCode = S.Trim(from.SecretCode)
		changed = true
	}
	if !excludeMap[`secretCodeAt`] && (forceMap[`secretCodeAt`] || from.SecretCodeAt != 0) {
		u.SecretCodeAt = from.SecretCodeAt
		changed = true
	}
	if !excludeMap[`verificationSentAt`] && (forceMap[`verificationSentAt`] || from.VerificationSentAt != 0) {
		u.VerificationSentAt = from.VerificationSentAt
		changed = true
	}
	if !excludeMap[`verifiedAt`] && (forceMap[`verifiedAt`] || from.VerifiedAt != 0) {
		u.VerifiedAt = from.VerifiedAt
		changed = true
	}
	if !excludeMap[`lastLoginAt`] && (forceMap[`lastLoginAt`] || from.LastLoginAt != 0) {
		u.LastLoginAt = from.LastLoginAt
		changed = true
	}
	if !excludeMap[`fullName`] && (forceMap[`fullName`] || from.FullName != ``) {
		u.FullName = S.Trim(from.FullName)
		changed = true
	}
	if !excludeMap[`userName`] && (forceMap[`userName`] || from.UserName != ``) {
		u.UserName = S.Trim(from.UserName)
		changed = true
	}
	if !excludeMap[`country`] && (forceMap[`country`] || from.Country != ``) {
		u.Country = S.Trim(from.Country)
		changed = true
	}
	if !excludeMap[`language`] && (forceMap[`language`] || from.Language != ``) {
		u.Language = S.Trim(from.Language)
		changed = true
	}
	if !excludeMap[`propertyCount`] && (forceMap[`propertyCount`] || from.PropertyCount != 0) {
		u.PropertyCount = from.PropertyCount
		changed = true
	}
	if !excludeMap[`propertyBought`] && (forceMap[`propertyBought`] || from.PropertyBought != 0) {
		u.PropertyBought = from.PropertyBought
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
