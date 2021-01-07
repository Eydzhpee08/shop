//main.go
package main
import (
 "github.com/Eydzhpee08/first-api/Config"
 "github.com/Eydzhpee08/first-api/Models"
 "github.com/Eydzhpee08/first-api/Routes"
 "fmt"
"github.com/jinzhu/gorm"
)
var err error
func main() {
 Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
if err != nil {
  fmt.Println("Status:", err)
 }
defer Config.DB.Close()
 Config.DB.AutoMigrate(&Models.User{})
r := Routes.SetupRouter()
 //running
 r.Run()
}