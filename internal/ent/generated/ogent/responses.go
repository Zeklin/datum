// Code generated by ent, DO NOT EDIT.

package ogent

import "github.com/datumforge/datum/internal/ent/generated"

func NewIntegrationCreate(e *generated.Integration) *IntegrationCreate {
	if e == nil {
		return nil
	}
	var ret IntegrationCreate
	ret.ID = e.ID
	ret.Kind = e.Kind
	ret.Description = NewOptString(e.Description)
	ret.SecretName = e.SecretName
	ret.CreatedAt = e.CreatedAt
	ret.DeletedAt = NewOptDateTime(e.DeletedAt)
	return &ret
}

func NewIntegrationCreates(es []*generated.Integration) []IntegrationCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]IntegrationCreate, len(es))
	for i, e := range es {
		r[i] = NewIntegrationCreate(e).Elem()
	}
	return r
}

func (i *IntegrationCreate) Elem() IntegrationCreate {
	if i == nil {
		return IntegrationCreate{}
	}
	return *i
}

func NewIntegrationList(e *generated.Integration) *IntegrationList {
	if e == nil {
		return nil
	}
	var ret IntegrationList
	ret.ID = e.ID
	ret.Kind = e.Kind
	ret.Description = NewOptString(e.Description)
	ret.SecretName = e.SecretName
	ret.CreatedAt = e.CreatedAt
	ret.DeletedAt = NewOptDateTime(e.DeletedAt)
	return &ret
}

func NewIntegrationLists(es []*generated.Integration) []IntegrationList {
	if len(es) == 0 {
		return nil
	}
	r := make([]IntegrationList, len(es))
	for i, e := range es {
		r[i] = NewIntegrationList(e).Elem()
	}
	return r
}

func (i *IntegrationList) Elem() IntegrationList {
	if i == nil {
		return IntegrationList{}
	}
	return *i
}

func NewIntegrationRead(e *generated.Integration) *IntegrationRead {
	if e == nil {
		return nil
	}
	var ret IntegrationRead
	ret.ID = e.ID
	ret.Kind = e.Kind
	ret.Description = NewOptString(e.Description)
	ret.SecretName = e.SecretName
	ret.CreatedAt = e.CreatedAt
	ret.DeletedAt = NewOptDateTime(e.DeletedAt)
	return &ret
}

func NewIntegrationReads(es []*generated.Integration) []IntegrationRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]IntegrationRead, len(es))
	for i, e := range es {
		r[i] = NewIntegrationRead(e).Elem()
	}
	return r
}

func (i *IntegrationRead) Elem() IntegrationRead {
	if i == nil {
		return IntegrationRead{}
	}
	return *i
}

func NewIntegrationUpdate(e *generated.Integration) *IntegrationUpdate {
	if e == nil {
		return nil
	}
	var ret IntegrationUpdate
	ret.ID = e.ID
	ret.Kind = e.Kind
	ret.Description = NewOptString(e.Description)
	ret.SecretName = e.SecretName
	ret.CreatedAt = e.CreatedAt
	ret.DeletedAt = NewOptDateTime(e.DeletedAt)
	return &ret
}

func NewIntegrationUpdates(es []*generated.Integration) []IntegrationUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]IntegrationUpdate, len(es))
	for i, e := range es {
		r[i] = NewIntegrationUpdate(e).Elem()
	}
	return r
}

func (i *IntegrationUpdate) Elem() IntegrationUpdate {
	if i == nil {
		return IntegrationUpdate{}
	}
	return *i
}

func NewIntegrationOrganizationRead(e *generated.Organization) *IntegrationOrganizationRead {
	if e == nil {
		return nil
	}
	var ret IntegrationOrganizationRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewIntegrationOrganizationReads(es []*generated.Organization) []IntegrationOrganizationRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]IntegrationOrganizationRead, len(es))
	for i, e := range es {
		r[i] = NewIntegrationOrganizationRead(e).Elem()
	}
	return r
}

func (o *IntegrationOrganizationRead) Elem() IntegrationOrganizationRead {
	if o == nil {
		return IntegrationOrganizationRead{}
	}
	return *o
}

func NewMembershipCreate(e *generated.Membership) *MembershipCreate {
	if e == nil {
		return nil
	}
	var ret MembershipCreate
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewMembershipCreates(es []*generated.Membership) []MembershipCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipCreate, len(es))
	for i, e := range es {
		r[i] = NewMembershipCreate(e).Elem()
	}
	return r
}

func (m *MembershipCreate) Elem() MembershipCreate {
	if m == nil {
		return MembershipCreate{}
	}
	return *m
}

func NewMembershipList(e *generated.Membership) *MembershipList {
	if e == nil {
		return nil
	}
	var ret MembershipList
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewMembershipLists(es []*generated.Membership) []MembershipList {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipList, len(es))
	for i, e := range es {
		r[i] = NewMembershipList(e).Elem()
	}
	return r
}

func (m *MembershipList) Elem() MembershipList {
	if m == nil {
		return MembershipList{}
	}
	return *m
}

func NewMembershipRead(e *generated.Membership) *MembershipRead {
	if e == nil {
		return nil
	}
	var ret MembershipRead
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewMembershipReads(es []*generated.Membership) []MembershipRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipRead, len(es))
	for i, e := range es {
		r[i] = NewMembershipRead(e).Elem()
	}
	return r
}

func (m *MembershipRead) Elem() MembershipRead {
	if m == nil {
		return MembershipRead{}
	}
	return *m
}

func NewMembershipUpdate(e *generated.Membership) *MembershipUpdate {
	if e == nil {
		return nil
	}
	var ret MembershipUpdate
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewMembershipUpdates(es []*generated.Membership) []MembershipUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipUpdate, len(es))
	for i, e := range es {
		r[i] = NewMembershipUpdate(e).Elem()
	}
	return r
}

func (m *MembershipUpdate) Elem() MembershipUpdate {
	if m == nil {
		return MembershipUpdate{}
	}
	return *m
}

func NewMembershipOrganizationRead(e *generated.Organization) *MembershipOrganizationRead {
	if e == nil {
		return nil
	}
	var ret MembershipOrganizationRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewMembershipOrganizationReads(es []*generated.Organization) []MembershipOrganizationRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipOrganizationRead, len(es))
	for i, e := range es {
		r[i] = NewMembershipOrganizationRead(e).Elem()
	}
	return r
}

func (o *MembershipOrganizationRead) Elem() MembershipOrganizationRead {
	if o == nil {
		return MembershipOrganizationRead{}
	}
	return *o
}

func NewMembershipUserRead(e *generated.User) *MembershipUserRead {
	if e == nil {
		return nil
	}
	var ret MembershipUserRead
	ret.ID = e.ID
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewMembershipUserReads(es []*generated.User) []MembershipUserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]MembershipUserRead, len(es))
	for i, e := range es {
		r[i] = NewMembershipUserRead(e).Elem()
	}
	return r
}

func (u *MembershipUserRead) Elem() MembershipUserRead {
	if u == nil {
		return MembershipUserRead{}
	}
	return *u
}

func NewOrganizationCreate(e *generated.Organization) *OrganizationCreate {
	if e == nil {
		return nil
	}
	var ret OrganizationCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewOrganizationCreates(es []*generated.Organization) []OrganizationCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationCreate, len(es))
	for i, e := range es {
		r[i] = NewOrganizationCreate(e).Elem()
	}
	return r
}

func (o *OrganizationCreate) Elem() OrganizationCreate {
	if o == nil {
		return OrganizationCreate{}
	}
	return *o
}

func NewOrganizationList(e *generated.Organization) *OrganizationList {
	if e == nil {
		return nil
	}
	var ret OrganizationList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewOrganizationLists(es []*generated.Organization) []OrganizationList {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationList, len(es))
	for i, e := range es {
		r[i] = NewOrganizationList(e).Elem()
	}
	return r
}

func (o *OrganizationList) Elem() OrganizationList {
	if o == nil {
		return OrganizationList{}
	}
	return *o
}

func NewOrganizationRead(e *generated.Organization) *OrganizationRead {
	if e == nil {
		return nil
	}
	var ret OrganizationRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewOrganizationReads(es []*generated.Organization) []OrganizationRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationRead, len(es))
	for i, e := range es {
		r[i] = NewOrganizationRead(e).Elem()
	}
	return r
}

func (o *OrganizationRead) Elem() OrganizationRead {
	if o == nil {
		return OrganizationRead{}
	}
	return *o
}

func NewOrganizationUpdate(e *generated.Organization) *OrganizationUpdate {
	if e == nil {
		return nil
	}
	var ret OrganizationUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewOrganizationUpdates(es []*generated.Organization) []OrganizationUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationUpdate, len(es))
	for i, e := range es {
		r[i] = NewOrganizationUpdate(e).Elem()
	}
	return r
}

func (o *OrganizationUpdate) Elem() OrganizationUpdate {
	if o == nil {
		return OrganizationUpdate{}
	}
	return *o
}

func NewOrganizationIntegrationsList(e *generated.Integration) *OrganizationIntegrationsList {
	if e == nil {
		return nil
	}
	var ret OrganizationIntegrationsList
	ret.ID = e.ID
	ret.Kind = e.Kind
	ret.Description = NewOptString(e.Description)
	ret.SecretName = e.SecretName
	ret.CreatedAt = e.CreatedAt
	ret.DeletedAt = NewOptDateTime(e.DeletedAt)
	return &ret
}

func NewOrganizationIntegrationsLists(es []*generated.Integration) []OrganizationIntegrationsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationIntegrationsList, len(es))
	for i, e := range es {
		r[i] = NewOrganizationIntegrationsList(e).Elem()
	}
	return r
}

func (i *OrganizationIntegrationsList) Elem() OrganizationIntegrationsList {
	if i == nil {
		return OrganizationIntegrationsList{}
	}
	return *i
}

func NewOrganizationMembershipsList(e *generated.Membership) *OrganizationMembershipsList {
	if e == nil {
		return nil
	}
	var ret OrganizationMembershipsList
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewOrganizationMembershipsLists(es []*generated.Membership) []OrganizationMembershipsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]OrganizationMembershipsList, len(es))
	for i, e := range es {
		r[i] = NewOrganizationMembershipsList(e).Elem()
	}
	return r
}

func (m *OrganizationMembershipsList) Elem() OrganizationMembershipsList {
	if m == nil {
		return OrganizationMembershipsList{}
	}
	return *m
}

func NewUserCreate(e *generated.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewUserCreates(es []*generated.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *generated.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewUserLists(es []*generated.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *generated.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewUserReads(es []*generated.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *generated.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Email = e.Email
	ret.CreatedAt = e.CreatedAt
	return &ret
}

func NewUserUpdates(es []*generated.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserMembershipsList(e *generated.Membership) *UserMembershipsList {
	if e == nil {
		return nil
	}
	var ret UserMembershipsList
	ret.ID = e.ID
	ret.Current = e.Current
	ret.CreatedAt = e.CreatedAt
	ret.UpdatedAt = e.UpdatedAt
	return &ret
}

func NewUserMembershipsLists(es []*generated.Membership) []UserMembershipsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserMembershipsList, len(es))
	for i, e := range es {
		r[i] = NewUserMembershipsList(e).Elem()
	}
	return r
}

func (m *UserMembershipsList) Elem() UserMembershipsList {
	if m == nil {
		return UserMembershipsList{}
	}
	return *m
}
