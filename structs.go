package main

// Payload : data sent to kibana api
type Payload struct {
	Attributes Attributes `json:"attributes"`
}

// Fields : kibana index fields
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

// Attributes : the attributes that make up the payload
type Attributes struct {
	Title         string `json:"title"`
	TimeFieldName string `json:"timeFieldName"`
	Fields        string `json:"fields"`
}

// FieldList : List of Fields
type FieldList struct {
	Fields Fields `json:"fields"`
}
