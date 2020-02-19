package main

import (
	"fmt"

	"github.com/axamon/addressbook/addressbookpb"
	"github.com/google/uuid"
)

func main() {
	p := addressbookpb.Person{}
	p.Id = uuid.New().String()
	p.Name = "pippo"

	fmt.Println(p)
}
