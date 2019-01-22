package jsonPlaceHolderClient

import (
	"jsonPlaceHolderClient/models"
	"net/http"
	"reflect"
	"testing"
)

func TestNewJsonPlaceHolderClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *jsonPlaceHolderClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJsonPlaceHolderClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJsonPlaceHolderClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJsonPlaceHolderClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonPlaceHolderClient_GetPost(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonPlaceHolderClient{
				Client: tt.fields.Client,
			}
			got, err := j.GetPost(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonPlaceHolderClient.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonPlaceHolderClient.GetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonPlaceHolderClient_CreatePost(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		userId int
		id     int
		title  string
		body   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonPlaceHolderClient{
				Client: tt.fields.Client,
			}
			if err := j.CreatePost(tt.args.userId, tt.args.id, tt.args.title, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("jsonPlaceHolderClient.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jsonPlaceHolderClient_ListPost(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonPlaceHolderClient{
				Client: tt.fields.Client,
			}
			got, err := j.ListPost()
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonPlaceHolderClient.ListPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonPlaceHolderClient.ListPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
