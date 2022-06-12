package main

import (
	"fmt"

	"github.com/davidJeronimoG/tablefinder/settings"
	"github.com/davidJeronimoG/tablefinder/utilities"
)

func main() {
	FilterTables("atm 3.1", "./cases/atm3.1.txt")
	FilterTables("atm 3.2", "./cases/atm3.2.txt")
}

func FilterTables(caseName string, file string) {

	caseOfUse, _ := utilities.GetValuesFromFile(file)

	scenery := Scenery{
		Name:       caseName,
		Values:     caseOfUse,
		Tables:     settings.GetTables(),
		Conditions: settings.GetFunctionalConditions(),
	}

	tablesTodelete := make(map[string]bool)
	for _, condition := range scenery.Conditions {
		present, name := condition(caseOfUse)

		if present {
			for _, table := range scenery.Tables {
				if table.ConditionsValues[name] == 0 {
					tablesTodelete[table.Name] = true
					// for i := range scenery.Tables {
					// 	if scenery.Tables[i].Name == table.Name {
					// 		scenery.Tables = append(scenery.Tables[:i], scenery.Tables[i+1:]...)
					// 		break
					// 	}
					// }
				}
			}
		}
	}

	for key, _ := range tablesTodelete {
		for i := range scenery.Tables {
			if scenery.Tables[i].Name == key {
				scenery.Tables = append(scenery.Tables[:i], scenery.Tables[i+1:]...)
				break
			}
	}}


	result := ""
	for _, table := range scenery.Tables {
		result += table.Name + " "
	}

	fmt.Println("The case:", scenery.Name, "matches the tables:", result)
}

type Scenery struct {
	Name       string
	Values     map[string]string
	Tables     []settings.Table
	Conditions [](func(m map[string]string) (bool, string))
}
