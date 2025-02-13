package dto

type Search struct {
	Value string
}

func BuildSearch(value string) Search {
	return Search{Value: value}
}
