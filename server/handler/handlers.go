package handler

import (
	"awesomeProject1/server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var hb HumansBuffer

type HumansBuffer struct {
	humans []model.Human
}

func (hb *HumansBuffer) isExist(firstName string) (int, bool) {
	for i, value := range hb.humans {
		if firstName != value.FirstName {
			continue
		} else {
			return i, true
		}
	}
	return 0, false
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		h := &model.Human{}
		r.Body.Close()
		err = json.Unmarshal(res, h)
		if err != nil {
			log.Fatal(err)
		}
		if _,ok:=hb.isExist(h.FirstName);ok{
			w.WriteHeader(http.StatusCreated)
			break
		}
		hb.humans = append(hb.humans, *h)
		w.WriteHeader(http.StatusCreated)
		fmt.Println("saved human is - ", hb.humans[len(hb.humans)-1])
		fmt.Println("after saved :",hb.humans)
	case "GET":
		if len(hb.humans) == 0 {
			break
		}
		i, ok := hb.isExist(r.URL.Query().Get("first_name"))
		if ok {
			humanJSON, err := json.Marshal(hb.humans[i])
			if err != nil {
				log.Fatal(err)
			}
			w.Write(humanJSON)
			fmt.Println("after get :",hb.humans)
		}
	case "PUT":
		if len(hb.humans) == 0 {
			break
		}
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		h := &model.Human{}
		r.Body.Close()
		err = json.Unmarshal(res, h)
		if err != nil {
			log.Fatal(err)
		}
		var i int
		var ok bool
		if i,ok=hb.isExist(h.FirstName);ok   {
			if hb.humans[i]==*h {
				w.WriteHeader(http.StatusCreated)
				break
			}
			hb.humans[i]=*h
			w.WriteHeader(http.StatusOK)
			fmt.Println("after update :",hb.humans)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	case "DELETE":
		if len(hb.humans) == 0 {
			break
		}
		i, ok := hb.isExist(r.URL.Query().Get("first_name"))
		if ok {
			hb.humans=append(hb.humans[:i],hb.humans[i+1:]...)
			w.WriteHeader((http.StatusOK))
		}else {
			w.WriteHeader((http.StatusNoContent))
		}
		fmt.Println("after delete :",hb.humans)
		default:
		fmt.Println("not REST request")

	}
}
func StartServer() {
	http.HandleFunc("/", queryHandler)

	fmt.Println("Server starts")
	http.ListenAndServe(":8080", nil)
}
