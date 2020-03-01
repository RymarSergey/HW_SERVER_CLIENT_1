package main

import (
	"awesomeProject1/client"
	"awesomeProject1/server/handler"
	"awesomeProject1/server/model"
	"fmt"
	"net/http"
	"time"
)

func main() {
	h := &model.Human{
		FirstName:  "Sergey",
		SecondName: "Rymar",
		Age:        "34",
	}

	go handler.StartServer()
	//Try to save human
	if client.SaveHumanByPOSTRequest(h.ForSendJSON()) {
		fmt.Println("created human")
	}
	time.Sleep(time.Second * 2)
	//Try to get human
	if h1, ok := client.ReadHumanByGetRequest(h.FirstName); ok {
		fmt.Println("readed human is - ", *h1)
	}
	time.Sleep(time.Second * 2)
	//Try to update
	h.Age = "35"
	switch client.UpdateByPUTRequest(h.ForSendJSON()) {
	case http.StatusCreated:
		fmt.Println("human exist ")
	case http.StatusOK:
		fmt.Println("human was updated ")
	case http.StatusNoContent:
		fmt.Println("human was't updated ")
	}
	//Try to delete
	switch client.DeleteHumanByRequest(h.FirstName) {
	case http.StatusOK:
		fmt.Println("human was updated ")
	case http.StatusNoContent:
		fmt.Println("human was't updated ")
	}

}
