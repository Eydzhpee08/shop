package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Eydzhpee08/first-api/pkg/customers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func (s *Server)GetcustomersAll(ctx *gin.Context) {
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

	list, err:=s.customerSvc.GetCustomers(ctx, requestData.Limit, requestData.Offset)
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


//Addcustomer ...
func (s *Server) AddCustomer(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData := customers.Customer{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(requestData)

    err = s.customerSvc.Addcustomer(ctx, requestData)
	if err != nil {
	
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("error on customer service, Addcustomer %s",err.Error())
		return
	}
	

	ctx.String(http.StatusOK, "%s", string("registered success!"))
}

//Editcustomer ...
func (s *Server) EditCustomer(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData := customers.Customer{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

    err = s.customerSvc.Editcustomer(ctx, requestData)
	if err != nil {
		if err==pgx.ErrNoRows{
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		if err.Error()=="conflict"{
			ctx.String(http.StatusBadRequest, "%s", string("this class is already have a customer!!"))
		}
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", string("edite success!"))
}

//Removecustomer ...
func (s *Server) RemoveCustomer(ctx *gin.Context) {
	 
	writer :=ctx.Writer
	customerID, err:=strconv.Atoi(ctx.Param("id"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
    err = s.customerSvc.Removecustomer(ctx, customerID)
	if err != nil {
		if err==pgx.ErrNoRows{
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", string("delete success!"))
}