package models

import (
	"encoding/json"
	"log"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) MarshalBinary() ([]byte, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *Post) UnmarshalBinary(bytes []byte) error {
	return json.Unmarshal(bytes, p)
}

func UnmarshalListBinary(byte []byte) ([]Post, error) {
	var posts []Post
	err := json.Unmarshal(byte, &posts)
	if err != nil {
		log.Println("error occurred while unmarshal list of posts: ", err)
		return nil, err
	}
	return posts, nil
}
