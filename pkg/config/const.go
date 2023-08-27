package config

type ConfgField string

const (
	Key      ConfgField = "key"
	Switches ConfgField = "switches"
	EventOn  ConfgField = "on"
	EventOff ConfgField = "off"
)

var StaticConfigFields = []ConfgField{
	Key,
	Switches,
}
