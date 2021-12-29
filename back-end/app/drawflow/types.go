package drawflow

type DrawflowData struct {
	Uid   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	Data  string   `json:"data,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type CodeDto struct {
	Code   string `json:"code,omitempty"`
	Result string `json:"result,omitempty"`
}

type Query struct {
	Results []DrawflowData `json:"results"`
}
