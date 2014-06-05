package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const (
      Host = "localhost:27017,localhost:27018,localhost:27019"
)

type Person struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	Weight int    `bson:"weight"`
}


func main() {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}
	c := session.DB("mydb").C("testCollection")

	if err = c.Insert(&Person{Name: "Bill DeRose", Age: 20}); err != nil {
		panic(err)
	}
	
	var pp Person
	if err = c.Find(bson.M{"name": "Bill DeRose", "age": 20}).One(&pp); err != nil {
		panic(err)
	}
	fmt.Println("Hello, my name is", pp.Name, "and I am", pp.Age, "years old")
}


