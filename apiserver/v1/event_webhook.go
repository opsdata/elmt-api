package v1

import "time"

type ZabbixEvent struct {
	Node         string `json:"node"`
	NodeAlias    string `json:"node_alias"`
	AlertGroup   string `json:"alert_group"`
	AlertKey     string `json:"alert_key"`
	Summary      string `json:"summary"`
	Additional   string `json:"additional"`
	Severity     string `json:"severity"`
	EventValue   string `json:"event_value"`
	EventId      string `json:"event_id"`
	EventDate    string `json:"event_date"`
	EventTime    string `json:"event_time"`
	EventTags    string `json:"event_tags"`
	RecoveryId   string `json:"r_event_id"`
	RecoveryDate string `json:"r_event_date"`
	RecoveryTime string `json:"r_event_time"`
	RecoveryTags string `json:"r_event_tags"`
}

type AlertmanagerEvent struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`        // key identifying the group of alerts (e.g. to deduplicate)
	TruncatedAlerts   int               `json:"truncatedAlerts"` // how many alerts have been truncated due to "max_alerts"
	Status            string            `json:"status"`          // resolved|firing
	ExternalURL       string            `json:"externalURL"`     // backlink to the Alertmanager
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	Alerts            []AlertInfo       `json:"alerts"`
}

type AlertInfo struct {
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	Status       string            `json:"status"`
	GeneratorURL string            `json:"generatorURL"` // the entity that caused the alert
	Fingerprint  string            `json:"fingerprint"`  // identify the alert
}
