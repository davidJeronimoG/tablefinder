package main

import (
	"fmt"
	"github.com/davidJeronimoG/tablefinder/service"
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
	fmt.Println("***************************** Caso individual *****************************")
	service.FilterTables("", "./cases/casosATM-TXT/atm-issuing-qVSDC-v0s-9.1a.txt")
}
