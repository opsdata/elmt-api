package v1

import (
	"time"
)

type Journal struct {
	KeyField   string    `json:"key_field"   gorm:"column:key_field"`
	Serial     int       `json:"serial"      gorm:"column:serial"`      // 事件编号
	ActionType int       `json:"action_type" gorm:"column:action_type"` // 操作类型
	ActionTime time.Time `json:"action_time" gorm:"column:action_time"` // 操作时间
	UserId     int       `json:"user_id"     gorm:"column:user_id"`     // 操作用户
	UserName   string    `json:"user_name"   gorm:"column:user_name"`   // 用户名称
}

type JournalList struct {
	Journals []*Journal `json:"journals"`
}

func (j *Journal) TableName() string {
	return "journals"
}
