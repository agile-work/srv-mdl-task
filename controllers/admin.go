package controllers

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/feature"
	"github.com/agile-work/srv-mdl-shared/models/response"
)

// Features return all available module features
func Features(res http.ResponseWriter, req *http.Request) {
	resp := response.New()
	features, err := feature.Load("features.json")
	if err != nil {
		resp.NewError("features load json file", err)
		resp.Render(res, req)
		return
	}

	resp.Data = features
	resp.Render(res, req)
}

// Register create all structures necessary for this module
func Register(res http.ResponseWriter, req *http.Request) {
	resp := response.New()
	// TODO: execute job "register module" passing the json structure file
	resp.Render(res, req)
}
