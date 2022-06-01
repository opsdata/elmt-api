package v1

type MetricData struct {
	HostName  string       `json:"host_name"` // 主机名称
	Component string       `json:"component"` // 组件名称
	Items     []MetricItem `json:"items"`     // 指标数据
}

type MetricItem struct {
	ItemKey   string `json:"item_key"`  // 示例: vfs.dev.read.rate
	ItemDesc  string `json:"item_desc"` // 示例: Disk read rate
	ItemName  string `json:"item_name"` // 示例: sda
	ItemType  int    `json:"item_type"` // 示例: 0(数值),3(文本)
	LastCheck string `json:"last_check"`
	LastValue string `json:"last_value"`

	// 显示用监控指标描述(ItemDesc: ItemName)
	// - 示例1: CPU utilization
	// - 示例2: Disk read rate: sda
	Name string `json:"name"`
}

// 单对象指标示例:
// key = 2021-05-17T03:59:43Z
// value = 38416
//
type MetricHistory struct {
	Key   string `json:"key"`   // 日期时间
	Value string `json:"value"` // 指标数值
}

// 多对象指标示例:
// item_names  = []{"eth0", "eth1"}
// item_values = [
// {"Timestamp": "2021-05-17T03:59:43Z", "eth0": "38416", "eth1": "0"},
// {"Timestamp": "2021-05-17T03:59:44Z", "eth0": "42624", "eth1": "0"}]
//
type MultiMetricHistory struct {
	ItemNames  []string            `json:"item_names"`
	ItemValues []map[string]string `json:"item_values"`
}
