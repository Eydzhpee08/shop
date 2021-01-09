//main.go
package main
import (
 "github.com/Eydzhpee08/shop/pkg/config"
 "github.com/Eydzhpee08/shop/cmd/app"
 "github.com/Eydzhpee08/shop/pkg/bills"
 "github.com/Eydzhpee08/shop/pkg/customers"
 "github.com/Eydzhpee08/shop/pkg/products"
 "github.com/Eydzhpee08/shop/pkg/database"
 "log"
)
var err error
func main() {
	config:=config.InitConfigs()
	db:=database.NewDBConnection(config.DSN)
	err:=database.Init(db)
	if err!=nil{
		log.Fatal(err)
		return
  }
  
  
	productSvc:=products.NewProductService(db)
	customersSvc:=customers.NewCustomersService(db)
	billSvc:=bill.NewBillService(db)
	router := gin.Default()

	appServer:=app.NewServer(router, productSvc, customersSvc, billSvc)
	appServer.InitRoutes()
	
	panic(appServer.Run(config.Server.Host, config.Server.Port))
}