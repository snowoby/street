package d

type NoEdges struct {
	Edges *struct{} `json:"edges,omitempty"`
}

type NoPath struct {
	Path *struct{} `json:"path,omitempty"`
}

type ValueType struct {
	ValueType string `json:"valueType"`
}
