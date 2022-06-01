package v1

type ZbxHost struct {
	HostName string `json:"host_name"`
	HostId   int    `json:"host_id"`
	AppAlias string `json:"app_alias"`
}

func (z *ZbxHost) TableName() string {
	return "mon.zbx_hosts"
}

type ZbxItem struct {
	ItemKey  string `json:"item_key"`
	ItemName string `json:"item_name"`
	ItemId   int    `json:"item_id"`
	HostName string `json:"host_name"`
	HostId   int    `json:"host_id"`
}

func (z *ZbxItem) TableName() string {
	return "mon.zbx_items"
}
