package catalog

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ensures that the User struct satisfies the core.RecordProxy interface
var _ core.RecordProxy = (*User)(nil)

// UserRole restricts what actions an authenticated user can perform.
type UserRole string

const (
	UserRoleAdmin  UserRole = "admin"
	UserRoleEditor UserRole = "editor"
)

type User struct {
	core.BaseRecordProxy
}

func (u *User) Email() string {
	return u.GetString("email")
}

func (u *User) SetEmail(email string) {
	u.Set("email", email)
}

func (u *User) DisplayName() string {
	return u.GetString("display_name")
}

func (u *User) SetDisplayName(name string) {
	u.Set("display_name", name)
}

func (u *User) Role() UserRole {
	return UserRole(u.GetString("role"))
}

func (u *User) SetRole(role UserRole) {
	u.Set("role", string(role))
}

func (u *User) CreatedAt() types.DateTime {
	return u.GetDateTime("created")
}

// IsAdmin returns true if the user has full administrative privileges.
func (u *User) IsAdmin() bool {
	return u.Role() == UserRoleAdmin
}

// IsEditor returns true if the user can manage products but not admin settings.
func (u *User) IsEditor() bool {
	return u.Role() == UserRoleEditor
}

// CanManageCatalog returns true if the user is allowed to create/edit/delete
// products and categories.
func (u *User) CanManageCatalog() bool {
	return u.IsAdmin() || u.IsEditor()
}

// FindUserByEmail finds a user by email address (case insensitive).
func FindUserByEmail(app core.App, email string) (*User, error) {
	user := &User{}
	err := app.RecordQuery("users").
		AndWhere(dbx.NewExp("LOWER(email)={:email}", dbx.Params{
			"email": strings.ToLower(email),
		})).
		Limit(1).
		One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
