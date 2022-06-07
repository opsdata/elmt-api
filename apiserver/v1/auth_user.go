package v1

import (
	"time"

	"gorm.io/gorm"

	metav1 "github.com/opsdata/common-base/pkg/meta/v1"

	"github.com/opsdata/common-base/pkg/auth"
	"github.com/opsdata/common-base/pkg/util/idutil"
)

// User represents a user restful resource. It is also used as gorm model.
type User struct {
	// Standard object metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UID         int       `json:"uid"                  gorm:"column:uid"         validate:"required"`
	Password    string    `json:"password"             gorm:"column:password"    validate:"required"`
	Activated   int       `json:"activated"            gorm:"column:activated"   validate:"omitempty"`
	IsAdmin     int       `json:"is_admin"             gorm:"column:is_admin"    validate:"omitempty"`
	LoginedAt   time.Time `json:"logined_at,omitempty" gorm:"column:logined_at"  validate:"omitempty"`
	TotalPolicy int64     `json:"totalPolicy"          gorm:"-"                  validate:"omitempty"`

	// Deprecated
	// Username string `json:"username" gorm:"column:username"    validate:"required,min=1,max=30"`
}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}

// TableName maps to postgresql table name.
func (u *User) TableName() string {
	return "public.users"
}

// Compare with the plain text password. Returns true if it is the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// AfterCreate run after create database record.
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "user-")
	return tx.Save(u).Error
}

// TODO:
// Legacy userInfo struct for ELMT frontent.
type UserInfo struct {
	ID               uint32   `json:"id"`
	Login            string   `json:"login"`
	Email            string   `json:"email"`
	Activated        bool     `json:"activated"`
	CreatedBy        string   `json:"createdBy"`
	LastModifiedBy   string   `json:"lastModifiedBy"`
	LastModifiedDate string   `json:"lastModifiedDate"`
	Authorities      []string `json:"authorities"`
	Apps             []string `json:"apps"`
	Orgs             []string `json:"orgs"`
	Pcss             []string `json:"pcss"`
	Isvs             []string `json:"isvs"`
}
