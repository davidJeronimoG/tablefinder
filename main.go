package main

import (
	// "encoding/json"
	"fmt"

	"github.com/davidJeronimoG/tablefinder/service"
	// "github.com/davidJeronimoG/tablefinder/settings"
)

func main() {
	fmt.Println("***************************** Casos ATM *****************************")
	service.FilterFolder("./cases/casosATM-TXT")
	fmt.Println("***************************** Casos Retail *****************************")
	service.FilterFolder("./cases/casosRetail-TXT")
	fmt.Println("***************************** Casos CNP *****************************")
	service.FilterFolder("./cases/CNP-TXT")
	fmt.Println("***************************** Casos T&E *****************************")
	service.FilterFolder("./cases/TandE-TXT")
// 	fmt.Println("***************************** Caso individual *****************************")
// 	service.FilterTables("", "./cases/casosATM-TXT/atm-issuing-qVSDC-v0s-9.1a.txt")

// for _, table := range settings.GetTables() {
// 	fmt.Println(table.Name, len(table.ConditionsValues))
// }

// tablas := settings.GetTables()

// convertir tablas a json
// jsonTablas, _ := json.Marshal(tablas)	
// fmt.Println(string(jsonTablas))

}