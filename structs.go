package main

type Payload struct {
	Attributes Attributes `json:"attributes"`
}

type Fields []struct {
	Aggregatable      bool     `json:"aggregatable,omitempty"`
	EsTypes           []string `json:"esTypes,omitempty"`
	Name              string   `json:"name"`
	Parent            string   `json:"parent,omitempty"`
	ReadFromDocValues bool     `json:"readFromDocValues"`
	Searchable        bool     `json:"searchable"`
	SubType           string   `json:"subType,omitempty"`
	Type              string   `json:"type"`
}

type Attributes struct {
	Title         string `json:"title"`
	TimeFieldName string `json:"timeFieldName"`
	Fields        string `json:"fields"`
}

type FieldList struct {
	Fields Fields `json:"fields"`
}
