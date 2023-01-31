package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// Define a custom struct to hold Album data.
type Album struct {
	Title  string
	Artist string
	Price  float64
	Likes  int
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Fetch all album fields with the HGETALL command. Because HGETALL
	// returns an array reply, and because the underlying data structure
	// in redis is a hash, it makes sense to use the Map() helper
	// function to convert the reply to a map[string]string.
	reply, err := redis.StringMap(conn.Do("HGETALL", "album:1"))
	if err != nil {
		log.Fatal(err)
	}

	// Use the populateAlbum helper function to create a new Album
	// object from the map[string]string.
	album, err := populateAlbum(reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", album)
}

// Create, populate and return a pointer to a new Album struct, based
// on data from a map[string]string.
func populateAlbum(reply map[string]string) (*Album, error) {
	var err error
	album := new(Album)
	album.Title = reply["title"]
	album.Artist = reply["artist"]
	// We need to use the strconv package to convert the 'price' value
	// from a string to a float64 before assigning it.
	album.Price, err = strconv.ParseFloat(reply["price"], 64)
	if err != nil {
		return nil, err
	}
	// Similarly, we need to convert the 'likes' value from a string to
	// an integer.
	album.Likes, err = strconv.Atoi(reply["likes"])
	if err != nil {
		return nil, err
	}
	return album, nil
}
