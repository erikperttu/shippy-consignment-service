package main

import "gopkg.in/mgo.v2"

// Implement all methods from the protobuf def
type service struct {
	session *mgo.Session
}