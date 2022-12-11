package models

type TemplateData struct {
	StrMap   map[string]string
	IntMap   map[string]int
	FloatMap map[string]float32
	Data     map[string]interface{}
	Error    string
	Warning  string
	Flash    string
	CSRToken string
}
