package v1

type App struct {
	AppId    int    `json:"app_id"    gorm:"primary_key;column:app_id"`
	AppName  string `json:"app_name"  gorm:"column:app_name"`
	AppAlias string `json:"app_alias" gorm:"column:app_alias"`
	AppLevel string `json:"app_level" gorm:"column:app_level"`
	IsvName  string `json:"isv_name"  gorm:"column:isv_name"`
}

type AppStatus struct {
	App
	Status int `json:"status"`
}

func (a *App) TableName() string {
	return "cmdb.apps"
}

type ExtraInfo struct {
	AppName  string `json:"app_name" gorm:"column:app_name"`
	PcsName  string `json:"pcs_name" gorm:"column:pcs_name"`
	Location string `json:"location" gorm:"column:location"`
}

type Host struct {
	Node      string `json:"node"`
	NodeAlias string `json:"node_alias"`
	OsType    string `json:"os_type"`
	ExtraInfo
}

type HostStatus struct {
	Host
	Status int `json:"status"`
	HostId int `json:"host_id"`
}

func (h *Host) TableName() string {
	return "cmdb.hosts"
}

type Device struct {
	Node      string `json:"node"`
	NodeAlias string `json:"node_alias"`
	DevType   string `json:"dev_type"`
	DevVendor string `json:"dev_vendor"`
	Rack      string `json:"rack"`
	RackU     string `json:"rack_u"`
	ExtraInfo
}

func (d *Device) TableName() string {
	return "cmdb.devices"
}

type Indicator struct {
	ItemKey    string `json:"item_key"`
	ItemName   string `json:"item_name"`
	ItemType   int    `json:"item_type"`
	IsWeb      bool   `json:"is_web"`
	IsMulti    bool   `json:"is_multi"`
	Component  string `json:"component"`
	ObjectType string `json:"object_type"`
	Source     string `json:"source"`
}

func (i *Indicator) TableName() string {
	return "cmdb.indicators"
}
