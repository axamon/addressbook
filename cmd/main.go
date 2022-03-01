package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"addressbook/proto/pb"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func main() {

	var file = "book.db"

	// Recupera dati da file.
	book := &pb.AddressBook{}
	data, err := ioutil.ReadFile(file)
	proto.Unmarshal(data, book)

	// Crea nuova persona.
	p := pb.Person{
		Id:   uuid.New().String(),
		Name: randomdata.FullName(1),
		Phones: []*pb.Person_PhoneNumber{
			&pb.Person_PhoneNumber{
				Number: randomdata.PhoneNumber(),
				Type:   pb.Person_HOME},
			&pb.Person_PhoneNumber{
				Number: randomdata.PhoneNumber(),
				Type:   pb.Person_WORK},
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
