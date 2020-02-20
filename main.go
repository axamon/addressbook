package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/axamon/addressbook/addressbookpb"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func main() {

	var file = "book.db"

	// Recupera dati da file.
	book := &addressbookpb.AddressBook{}
	data, err := ioutil.ReadFile(file)
	proto.Unmarshal(data, book)

	// Crea nuova persona.
	p := addressbookpb.Person{
		Id:   uuid.New().String(),
		Name: "Pluto",
		Phones: []*addressbookpb.Person_PhoneNumber{
			&addressbookpb.Person_PhoneNumber{
				Number: "555",
				Type:   addressbookpb.Person_HOME},
			&addressbookpb.Person_PhoneNumber{
				Number: "444",
				Type:   addressbookpb.Person_WORK},
		},
		Email: "aaaa",
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
