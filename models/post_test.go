package models

import (
	"testing"
)

func TestPost_MarshalBinary(t *testing.T) {
	type fields struct {
		Id     int
		UserId int
		Title  string
		Body   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "Normally initialized",
			fields: fields{Id: 1, UserId: 1, Title: "title 1", Body: "body 1"}, wantErr: false},
		{name: "Partly Initialized",
			fields: fields{Id: 1}, wantErr: false},
		{name: "Uninitialized", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Id:     tt.fields.Id,
				UserId: tt.fields.UserId,
				Title:  tt.fields.Title,
				Body:   tt.fields.Body,
			}
			_, err := p.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPost_UnmarshalBinary(t *testing.T) {
	type fields struct {
		Id     int
		UserId int
		Title  string
		Body   string
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Normally initialized",
			fields: fields{Id: 1, UserId: 1, Title: "title 1", Body: "body 1"}, wantErr: false,
			args: args{bytes: []byte(`{
	"id": 2,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}`)}},

		{name: "Partly Initialized",
			fields: fields{Id: 1}, wantErr: false, args: args{bytes: []byte(`{
	"id": 2,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}`)}},
		{name: "Uninitialized", wantErr: false, args: args{bytes: []byte(`{
	"id": 2,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}`)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Id:     tt.fields.Id,
				UserId: tt.fields.UserId,
				Title:  tt.fields.Title,
				Body:   tt.fields.Body,
			}
			if err := p.UnmarshalBinary(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("Post.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshalListBinary(t *testing.T) {
	type args struct {
		byte []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Empty curly brackets", wantErr: true, args: args{byte: []byte(`{}`)}},
		{name: "Empty square brackets", wantErr: false, args: args{byte: []byte(`[]`)}},
		{name: "One element without array", wantErr: true, args: args{byte: []byte(`{
	"id": 1,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}`)}},
		{name: "One element with array", wantErr: false, args: args{byte: []byte(`[{
	"id": 1,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}]`)}},
		{name: "Multiple element with array", wantErr: false, args: args{byte: []byte(`[{
	"id": 1,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
},{
	"id": 1,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}]`)}}, {name: "Multiple element with array, partial and complete mix", wantErr: false, args: args{byte: []byte(`[{
	"id": 1,
	"userId": 1
},{
	"id": 1,
	"userId": 1,
	"title": "title 1",
	"body": "body 1"
}]`)}}, {name: "Multiple element with array, partial mix", wantErr: false, args: args{byte: []byte(`[{
	"id": 1,
	"userId": 1
},{
	"title": "title 1",
	"body": "body 1"
}]`)}}, {name: "Unrelated array json should ignore", wantErr: false, args: args{byte: []byte(`[{
	"x": 1,
	"y": 1
},{
	"z": 1,
	"uaserId": 1,
	"tistle": "title 1",
	"bodsy": "body 1"
}]`)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UnmarshalListBinary(tt.args.byte)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalListBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
