package data

type Metrics struct {
	Values         string `json:"values"`
	Dstypes        string `json:"dstypes"`
	Dsnames        string `json:"dsnames"`
	Time           string `json:"time"`
	Interval       string `json:"interval"`
	Host           string `json:"host"`
	Plugin         string `json:"plugin"`
	PluginInstance string `json:"plugin_instance"`
	Type           string `json:"type"`
	TypeInstance   string `json:"type_instance"`
	Meta           Meta
}

type Meta struct {
	Name string `json:"name"`
}

func NewMetric(nilAttributes ...string) *Metrics {
	m := &Metrics{}

	return m
}
