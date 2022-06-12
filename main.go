package main

import (
	"github.com/davidJeronimoG/tablefinder/service"
)

func main() {
	service.FilterTables("atm 3.1", "./cases/atm3.1.txt")
	service.FilterTables("atm 3.2", "./cases/atm3.2.txt")
	service.FilterTables("atm 4.1", "./cases/atm4.1.txt")
	service.FilterTables("atm 5.1", "./cases/atm5.1.txt")
	service.FilterTables("TyE 1.2 - incremental", "./cases/tye1.2incremental.txt")
}