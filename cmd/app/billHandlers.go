package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/Eydzhpee08/shop/pkg/bills"
	"github.com/gin-gonic/gin"
)

type RequestData struct{
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

//GetbillsAll testing
func (s *Server)GetBillAll(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData:=RequestData{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	list, err:=s.billSvc.GetBill(ctx, requestData.Limit, requestData.Offset)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	
	if len(list)==0{
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	response, err:=json.Marshal(list)

	ctx.String(http.StatusOK, "%s", string(response))
	
}

//Addbills ...
func (s *Server) AddBill(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData := bill.Bill{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

    err = s.billSvc.AddBill(ctx, requestData)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", string(" Add success!"))
}
