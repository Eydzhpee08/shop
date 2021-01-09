//main.go
package main
import (
 "github.com/Eydzhpee08/first-api/pkg/config"
 "github.com/Eydzhpee08/first-api/cmd/app"
 "github.com/Eydzhpee08/first-api/pkg/bill"
 "github.com/Eydzhpee08/first-api/pkg/customers"
 "github.com/Eydzhpee08/first-api/pkg/products"
 "github.com/Eydzhpee08/first-api/pkg/database"
 "fmt"
"github.com/jinzhu/gorm"
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
	customersSvc:=customerss.NewCustomersService(db)
	billSvc:=billes.NewBillService(db)
	router := gin.Default()

	appServer:=app.NewServer(router, productSvc, customersSvc, billSvc)
	appServer.InitRoutes()
	
	panic(appServer.Run(config.Server.Host, config.Server.Port))
}