package v1

import (
	"github.com/ory/ladon"
	"gorm.io/gorm"

	"github.com/opsdata/common-base/pkg/json"
	metav1 "github.com/opsdata/common-base/pkg/meta/v1"
	"github.com/opsdata/common-base/pkg/util/idutil"
)

// AuthzPolicy defines iam policy type.
type AuthzPolicy struct {
	ladon.DefaultPolicy
}

// Policy represents a policy restful resource, include a ladon policy.
// It is also used as gorm model.
type Policy struct {
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	PolicyName string `json:"policy_name" gorm:"column:policy_name" validate:"omitempty"`

	// The user of the policy.
	Username string `json:"username" gorm:"column:username" validate:"omitempty"`

	// AuthzPolicy policy, will not be stored in db.
	Policy AuthzPolicy `json:"policy,omitempty" gorm:"-" validate:"omitempty"`

	// The ladon policy content, just a string format of ladon.DefaultPolicy. DO NOT modify directly.
	PolicyShadow string `json:"-" gorm:"column:policy_shadow" validate:"omitempty"`
}

// PolicyList is the whole list of all policies which have been stored in stroage.
type PolicyList struct {
	// Standard list metadata.
	metav1.ListMeta `json:",inline"`

	// List of policies.
	Items []*Policy `json:"items"`
}

// TableName maps to mysql table name.
func (p *Policy) TableName() string {
	return "public.policies"
}

// String returns the string format of Policy.
func (ap AuthzPolicy) String() string {
	data, _ := json.Marshal(ap)

	return string(data)
}

// BeforeCreate run before create database record.
func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeCreate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.Policy.String()

	return nil
}

// AfterCreate run after create database record.
func (p *Policy) AfterCreate(tx *gorm.DB) error {
	p.InstanceID = idutil.GetInstanceID(p.ID, "policy-")

	return tx.Save(p).Error
}

// BeforeUpdate run before update database record.
func (p *Policy) BeforeUpdate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeUpdate(tx); err != nil {
		return err
	}

	p.PolicyShadow = p.Policy.String()

	return nil
}

// AfterFind run after find to unmarshal a policy string into ladon.DefaultPolicy struct.
func (p *Policy) AfterFind(tx *gorm.DB) (err error) {
	if err := p.ObjectMeta.AfterFind(tx); err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(p.PolicyShadow), &p.Policy); err != nil {
		return err
	}

	return nil
}
