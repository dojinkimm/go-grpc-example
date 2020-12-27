package data

import (
	userpb "github.com/dojinkimm/go-grpc-example/protos/user"
)

var Users = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "Henry",
		PhoneNumber: "01012341234",
		Age:         22,
	},
	{
		UserId:      "2",
		Name:        "Michael",
		PhoneNumber: "01098128734",
		Age:         55,
	},
	{
		UserId:      "3",
		Name:        "Jessie",
		PhoneNumber: "01056785678",
		Age:         15,
	},
	{
		UserId:      "4",
		Name:        "Max",
		PhoneNumber: "01099999999",
		Age:         37,
	},
	{
		UserId:      "5",
		Name:        "Tony",
		PhoneNumber: "01012344321",
		Age:         25,
	},
}
