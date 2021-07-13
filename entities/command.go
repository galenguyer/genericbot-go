package entities

type Command struct {
	Name    string
	Execute func(Context) error
}
