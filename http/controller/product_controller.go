package controller

import (
	"bytes"
	"dimall.id/standard-template/model"
	"dimall.id/standard-template/repo"
	"encoding/json"
	"github.com/dimall-id/lumos/data"
	"github.com/dimall-id/lumos/event"
	web "github.com/dimall-id/lumos/http"
	"net/http"
)

type ProductController struct {
	Repo repo.ProductRepo
}

func (c *ProductController) Index (r *http.Request) (interface{}, web.HttpError) {

	tx := c.Repo.DB
	tx.Model(model.Product{})
	Q := data.New(c.Repo.DB)

	var data []model.Product
	response := Q.BuildResponse(r, &data)

	jsonRes, err := json.Marshal(response)
	var res bytes.Buffer
	err = json.Compact(&res, jsonRes)
	if err != nil {
		return nil, web.InternalServerError()
	}

	msg := event.LumosMessage{
		Topic: "COMMAND_EVENT",
		Key: "DATA",
		Value: res.String(),
		Headers: map[string]string{},
	}

	err = event.GenerateOutbox(c.Repo.DB, msg)
	if err != nil {
		return nil, web.InternalServerError()
	}

	return response, web.HttpError{}
}