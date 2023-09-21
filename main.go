package main

import (
	"cvpchatalert/cmd/api"
)

func main() {
	MakeApi := api.Application{
		Port: 8086,
	}
	//Start the API
	MakeApi.StartApi()
}
