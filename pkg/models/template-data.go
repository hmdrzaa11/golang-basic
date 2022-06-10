package models

// TemplateData a general struct represents the data that handlers will send into templates
type TemplateData struct {
	//because we do not know how many string we might send to a template from start we use a map
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	OtherData map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
