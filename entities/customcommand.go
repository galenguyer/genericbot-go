package entities

type CustomCommand struct {
	Name     string `bson:"_id"`
	Delete   bool
	Response string
}
