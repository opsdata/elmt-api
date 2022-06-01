package v1

type Zone struct {
	ZoneId   int    `json:"zone_id"    gorm:"primary_key;column:zone_id"`
	ZoneName string `json:"zone_name"  gorm:"column:zone_name"`
}

func (z *Zone) TableName() string {
	return "cmdb.zones"
}

type Cloud struct {
	PcsId    int    `json:"pcs_id"    gorm:"primary_key;column:pcs_id"`
	PcsName  string `json:"pcs_name"  gorm:"column:pcs_name"`
	ZoneName string `json:"zone_name" gorm:"column:zone_name"`
}

func (c *Cloud) TableName() string {
	return "cmdb.clouds"
}

type Location struct {
	LocationId   int    `json:"location_id"    gorm:"primary_key;column:location_id"`
	LocationName string `json:"location_name"  gorm:"column:location_name"`
}

func (l *Location) TableName() string {
	return "cmdb.locations"
}

type Isv struct {
	IsvId   int    `json:"isv_id"    gorm:"primary_key;column:isv_id"`
	IsvName string `json:"isv_name"  gorm:"column:isv_name"`
	IsvType string `json:"isv_type"  gorm:"column:isv_type"`
}

func (i *Isv) TableName() string {
	return "cmdb.isvs"
}

type Org struct {
	OrgId   int    `json:"org_id"    gorm:"primary_key;column:org_id"`
	OrgName string `json:"org_name"  gorm:"column:org_name"`
}

func (o *Org) TableName() string {
	return "cmdb.orgs"
}
