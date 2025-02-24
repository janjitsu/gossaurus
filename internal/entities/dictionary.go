package entities

type Dictionary struct {
	Id     string
	Title  string
	From   string
	To     string
	Author string
	Cover  string
	Pages  []Page
}
