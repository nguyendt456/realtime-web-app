package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID
	Name string
}

func (u User) Print() {
	fmt.Printf("\nUser ID: %d\nName: %s\n\n", u.Id, u.Name)
}
