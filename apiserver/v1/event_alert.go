package v1

import (
	"time"
)

// Alert represents an alert restful resource. It is also used as gorm model.
type Alert struct {
	// Basic info
	Serial          int       `json:"serial"               gorm:"primary_key;auto_increment;column:serial"`
	EventId         string    `json:"event_id"             gorm:"column:event_id"`
	Identifier      string    `json:"identifier"           gorm:"column:identifier"`
	Node            string    `json:"node"                 gorm:"column:node"`
	NodeAlias       string    `json:"node_alias"           gorm:"column:node_alias"`
	Severity        int       `json:"severity"             gorm:"column:severity"`         // 告警当前级别
	OldSev          int       `json:"old_sev,omitempty"    gorm:"column:old_sev"`          // 告警初始级别
	Summary         string    `json:"summary"              gorm:"column:summary"`          // 告警内容信息
	Additional      string    `json:"additional,omitempty" gorm:"column:additional"`       // 告警辅助信息
	FirstOccurrence time.Time `json:"first_occurrence"     gorm:"column:first_occurrence"` // 告警发生时间
	LastOccurrence  time.Time `json:"last_occurrence"      gorm:"column:last_occurrence"`  // 最近发生时间
	Tally           int       `json:"tally"                gorm:"column:tally"`

	// Source info
	Manager    string `json:"manager,omitempty"     gorm:"column:manager"`
	AlertGroup string `json:"alert_group,omitempty" gorm:"column:alert_group"`
	AlertKey   string `json:"alert_key,omitempty"   gorm:"column:alert_key"`

	// Status info
	Type         int  `json:"type"                   gorm:"column:type"`         // 事件类型:1-故障类,2-恢复类
	EventStatus  int  `json:"event_status"           gorm:"column:event_status"` // 事件状态:0-已恢复(自动),1-处理中,2-已关闭(手动)
	ClearedType  int  `json:"cleared_type,omitempty" gorm:"column:cleared_type"` // 清除方式:0-默认,1-自动,2-手动
	Acknowledged bool `json:"acknowledged,omitempty" gorm:"column:acknowledged"` // 是否确认:false-未确认,true-已确认
	Suppressed   bool `json:"suppressed,omitempty"   gorm:"column:suppressed"`   // 是否抑制:false-未抑制,true-已抑制

	// Action info
	AckedTime   time.Time `json:"acked_time,omitempty"   gorm:"column:acked_time"`   // 告警确认时间
	ClearedTime time.Time `json:"cleared_time,omitempty" gorm:"column:cleared_time"` // 告警清除时间

	// Others
	Tags     string `json:"tags,omitempty"     gorm:"column:tags"` // Tags []Tag `json:"tags,omitempty"`
	AppName  string `json:"app_name,omitempty" gorm:"column:app_name"`
	PcsName  string `json:"pcs_name,omitempty" gorm:"column:pcs_name"`
	OrgName  string `json:"org_name,omitempty" gorm:"column:org_name"`
	Location string `json:"location,omitempty" gorm:"column:location"`
}

type AlertList struct {
	Alerts []*Alert `json:"alerts"`
}

// Internal use only.
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TableName maps to sqlite table name.
func (a *Alert) TableName() string {
	return "alerts"
}
