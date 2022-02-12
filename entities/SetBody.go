package entities

type SetBody struct {
	Key   string
	Value string
}

func NewSetBody(key, value string) *SetBody {
	return &SetBody{key, value}
}
