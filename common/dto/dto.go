package dto

type SearchData struct {
	Target string `json:"target,omitempty"`
}

type SearchDetailData struct {
	Data string `json:"data,omitempty"`
	ID   string `json:"id,omitempty"`
}

type SearchResult struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}
