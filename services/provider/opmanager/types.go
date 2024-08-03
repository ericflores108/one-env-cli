package opmanager

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Items []Item

type Field struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Purpose   string `json:"purpose"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Reference string `json:"reference"`
}

type ItemResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Fields []Field
}
