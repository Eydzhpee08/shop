package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Eydzhpee08/shop/pkg/products"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)



//GetproductsAll testing
func (s *Server)GetProductAll(ctx *gin.Context) {
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

	list, err:=s.productSvc.GetProduct(ctx, requestData.Limit, requestData.Offset)
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

//Addproducts ...
func (s *Server) AddProducts(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData := products.Product{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

    err = s.productSvc.AddProducts(ctx, requestData)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", string("add success!"))
}

//Editproducts ...
func (s *Server) EditProducts(ctx *gin.Context) {
	request :=ctx.Request
	writer :=ctx.Writer
	
	body, err:=ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	requestData := products.Product{}
	err = json.Unmarshal(body, &requestData)
	if err!=nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

    err = s.productSvc.EditProducts(ctx, requestData)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", string("edite success!"))
}

//Removeproducts ...
func (s *Server) RemoveProducts(ctx *gin.Context) {
	
	writer :=ctx.Writer
	productID, err:=strconv.Atoi(ctx.Param("id"))
	if err != nil  {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
    err = s.productSvc.RemoveProducts(ctx, productID)
	if err != nil {
		if err==pgx.ErrNoRows{
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	

	ctx.String(http.StatusOK, "%s", "delete success!")
}