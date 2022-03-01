package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"addressbook/proto/pb"

	randomdata "github.com/Pallinder/go-randomdata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	var file = "book.db"

	// Recupera dati da file.
	book := &pb.AddressBook{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print(err)
	}
	proto.Unmarshal(data, book)

	rand.Seed(time.Now().UnixNano())
	gender := rand.Intn(2)
	profile := randomdata.GenerateProfile(gender)

	fmt.Printf("%+v\n", profile)

	// Crea nuova persona.
	p := pb.Person{
		Name:        profile.Name.First + " " + profile.Name.Last,
		Id:          profile.Login.Sha256,
		Email:       profile.Email,
		Gender:      profile.Gender,
		Nationality: profile.Nat,
		Picture:     profile.Picture.Large,
		Dob:         profile.Dob,
		Phones:      []*pb.Person_PhoneNumber{{Number: profile.Phone, Type: pb.Person_HOME}, {Number: profile.Cell, Type: pb.Person_MOBILE}},
		LastUpdated: &timestamppb.Timestamp{Seconds: time.Now().UnixNano()},
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
