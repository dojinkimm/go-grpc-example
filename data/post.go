package data

import (
	postpb "github.com/dojinkimm/go-grpc-example/protos/post"
)

type PostData struct {
	UserID string
	Posts  []*postpb.PostMessage
}

var UserPosts = []*PostData{
	{
		UserID: "1",
		Posts: []*postpb.PostMessage{
			{
				PostId: "1",
				Author: "",
				Title:  "gRPC 구축하기 (1)",
				Body:   "gRPC 구축하려면 이렇게 하면 된다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "2",
				Author: "",
				Title:  "gRPC 구축하기 (2)",
				Body:   "gRPC server란 이렇다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "3",
				Author: "",
				Title:  "Golang context",
				Body:   "Golang context를 공부하려면 공식 docs를 봐라",
				Tags:   []string{"Golang", "context"},
			},
			{
				PostId: "4",
				Author: "",
				Title:  "스타트업 취업 후기",
				Body:   "나는 이렇게 취업했다",
				Tags:   []string{"IT", "interview", "startup", "developer"},
			},
		},
	},
	{
		UserID: "3",
		Posts: []*postpb.PostMessage{
			{
				PostId: "5",
				Author: "",
				Title:  "2020.10.25 일기",
				Body:   "오늘의 일기 끝",
				Tags:   []string{"daily", "homework"},
			},
		},
	},
	{
		UserID: "4",
		Posts: []*postpb.PostMessage{
			{
				PostId: "6",
				Author: "",
				Title:  "은퇴하는 방법",
				Body:   "은퇴하고 싶니?",
				Tags:   []string{"retirement", "know-how"},
			},
			{
				PostId: "7",
				Author: "",
				Title:  "저축하는 방법",
				Body:   "저축하고 싶니?",
				Tags:   []string{"money", "asset management", "know-how"},
			},
		},
	},
	{
		UserID: "5",
		Posts: []*postpb.PostMessage{
			{
				PostId: "8",
				Author: "",
				Title:  "회고하는 방법",
				Body:   "좋은 리더는 이렇게 회고하더라",
				Tags:   []string{"retrospect", "company", "teamwork", "leader"},
			},
		},
	},
}
