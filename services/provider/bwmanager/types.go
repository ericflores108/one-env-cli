package bwmanager

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Items []Item

type Field struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ItemResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
	Login  Login   `json:"login"`
}
