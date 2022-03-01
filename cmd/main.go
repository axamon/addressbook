package main

import (
	"addressbook"
	"fmt"
	"io/ioutil"
	"log"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func main() {

	var file = "book.db"

	// Recupera dati da file.
	book := &addressbook.pb.AddressBook{}
	data, err := ioutil.ReadFile(file)
	proto.Unmarshal(data, book)

	// Crea nuova persona.
	p := addressbook.pb.Person{
		Id:   uuid.New().String(),
		Name: randomdata.FullName(1),
		Phones: []*addressbook.pb.Person_PhoneNumber{
			&addressbook.pb.Person_PhoneNumber{
				Number: randomdata.PhoneNumber(),
				Type:   addressbook.pb.Person_HOME},
			&addressbook.pb.Person_PhoneNumber{
				Number: randomdata.PhoneNumber(),
				Type:   addressbook.pb.Person_WORK},
		},
		Email: randomdata.Email(),
	}

	// Aggiunge persona.
	book.People = append(book.People, &p)

	// Stampa tutte le persone.
	for _, person := range book.People {
		fmt.Println(person)
	}

	// Salva dati su file.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(file, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
