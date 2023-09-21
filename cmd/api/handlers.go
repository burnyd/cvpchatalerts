package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cvpchatalert/pkg/gpt"
	"cvpchatalert/pkg/jsonlog"
)

type AlertData struct {
	Alerts map[string]interface{}
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := Envelope{
		"status": "available",
		"system_info": map[string]string{
			"version": "v1",
		},
		"Status": "Running",
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) CreateAlert(w http.ResponseWriter, r *http.Request) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	datamodel := AlertResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &datamodel)
	if err != nil {
		fmt.Println(err)
	}

	re := datamodel.Alerts

	var QuestionString string

	for _, i := range re {
		QuestionString = i.Annotations.Description
	}

	logger.PrintInfo("received question "+QuestionString, nil)
	logger.PrintInfo("Checking with chatgpt", nil)
	gpt.AskGpt(QuestionString)
}
