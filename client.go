package jsonPlaceHolderClient

import (
	bytes2 "bytes"
	"errors"
	"io/ioutil"
	"jsonPlaceHolderClient/models"
	"log"
	"net/http"
)

const (
	BASE  = "https://jsonplaceholder.typicode.com/"
	POSTS = "posts/"
)

type jsonPlaceHolderClient struct {
	Client *http.Client
}

func NewJsonPlaceHolderClient() (*jsonPlaceHolderClient, error) {
	client := http.DefaultClient
	return &jsonPlaceHolderClient{Client: client}, nil

}
func (j *jsonPlaceHolderClient) GetPost(id string) (*models.Post, error) {
	resp, err := j.Client.Get(BASE + POSTS + id)
	if err != nil {
		log.Println("error encountered while getting post_id= ", id, ": ", err)
		return nil, err
	}
	//Close body after function finish processing
	defer resp.Body.Close()
	log.Println(resp)
	if resp.StatusCode == http.StatusOK {
		post := &models.Post{}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("error encountered while unmarshaling post_id= ", id, ": ", err)
			return nil, err
		}
		err = post.UnmarshalBinary(bytes)
		if err != nil {
			log.Println("error encountered while unmarshaling post_id= ", id, ": ", err)
			return nil, err
		}
		return post, nil
	} else {
		return nil, errors.New(resp.Status)
	}
}

func (j *jsonPlaceHolderClient) CreatePost(userId, id int, title, body string) error {
	post := &models.Post{UserId: userId, Id: id, Title: title, Body: body}
	bytes, err := post.MarshalBinary()
	if err != nil {
		log.Println("error encountered while marshalling post data", err)
	}
	resp, err := j.Client.Post(BASE+POSTS, "application/json", bytes2.NewBuffer(bytes))
	if err != nil {
		log.Println("error encountered while creating post: ", err)
		return err
	}
	//Close body after function finish processing
	defer resp.Body.Close()
	log.Println(resp)
	if resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		return errors.New(resp.Status)
	}
}

func (j *jsonPlaceHolderClient) ListPost() ([]models.Post, error) {
	resp, err := j.Client.Get(BASE + POSTS)
	if err != nil {
		log.Println("error encountered while listing posts : ", err)
		return nil, err
	}
	//Close body after function finish processing
	defer resp.Body.Close()
	log.Println(resp)
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("error encountered while reading list posts response : ", err)
		}
		return models.UnmarshalListBinary(bytes)
	} else {
		return nil, errors.New(resp.Status)
	}
}
