package v1

import (
	"gorm.io/gorm"

	metav1 "github.com/opsdata/common-base/pkg/meta/v1"

	"github.com/opsdata/common-base/pkg/util/idutil"
)

// Secret represents a secret restful resource. It is also used as gorm model.
type Secret struct {
	// Standard object metadata.
	metav1.ObjectMeta `       json:"metadata,omitempty"`

	Username    string `json:"username"    gorm:"column:username"    validate:"omitempty"`
	SecretID    string `json:"secret_id"   gorm:"column:secret_id"   validate:"omitempty"`
	SecretKey   string `json:"secret_key"  gorm:"column:secret_key"  validate:"omitempty"`
	Expires     int64  `json:"expires"     gorm:"column:expires"     validate:"omitempty"`
	Description string `json:"description" gorm:"column:description" validate:"description"`

	// Deprecated
	// SecretName string `json:"secret_name" gorm:"column:secret_name" validate:"omitempty"`
}

// SecretList is the whole list of all secrets which have been stored in stroage.
type SecretList struct {
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`

	Items []*Secret `json:"items"`
}

// TableName maps to postgresql table name.
func (s *Secret) TableName() string {
	return "public.secrets"
}

// AfterCreate run after create database record.
func (s *Secret) AfterCreate(tx *gorm.DB) error {
	s.InstanceID = idutil.GetInstanceID(s.ID, "secret-")
	return tx.Save(s).Error
}
