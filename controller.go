package main

import (
	"log"
    "encoding/json"
	"net/http"
	srv "tamboonapi/service"
)

func (handler *TamboonHandler) GetCharityList(resp http.ResponseWriter, req *http.Request) {

	charities := srv.CharityList()

	if e := json.NewEncoder(resp).Encode(charities); e != nil {
		http.Error(resp, e.Error(), 500)
		return
	}
}

func (handler *TamboonHandler) PostDonate(resp http.ResponseWriter, req *http.Request) {		
	donationReq := &srv.DonationReq{}

	defer req.Body.Close()

	if e := json.NewDecoder(req.Body).Decode(donationReq); e != nil {
		log.Printf("decode error")
		log.Fatal(e)
		http.Error(resp, e.Error(), 400)
		return
	}
	
	charge, operation := donationReq.DoDonate(handler.client)

	if e := handler.client.Do(charge, operation); e != nil {
		log.Printf("charge error")
		log.Fatal(e)
		http.Error(resp, e.Error(), 400)
		return
	}

	if e := json.NewEncoder(resp).Encode(&srv.Result{Success: true}); e != nil {
		log.Printf("encode error")
		log.Fatal(e)
		http.Error(resp, e.Error(), 500)
		return
	}
}