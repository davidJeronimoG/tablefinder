package service

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/davidJeronimoG/tablefinder/settings"
	"github.com/davidJeronimoG/tablefinder/utilities"
)

func FilterTables(caseName string, file string) {

	caseOfUse, _ := utilities.GetValuesFromFile(file)

	scenery := Scenery{
		Name:       caseName,
		Values:     caseOfUse,
		Tables:     settings.GetTables(),
		Conditions: settings.GetFunctionalConditions(),
	}

	// fmt.Println(scenery.Values["F3.1"] != "")

	tablesTodelete := make(map[string]bool)
	for _, condition := range scenery.Conditions {
		present, name := condition(caseOfUse)

		if present {
			for _, table := range scenery.Tables {
				if table.ConditionsValues[name] == 0 {
					tablesTodelete[table.Name] = true
				}
			}
		}
		if !present {
			for _, table := range scenery.Tables {
				if table.ConditionsValues[name] == 1 {
					tablesTodelete[table.Name] = true
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
		}
	}

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

func FilterFolder(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		nombre := f.Name()
		FilterTables(nombre, folder+"/"+nombre)
	}

}

// MapOfCase is a map key-value of a case
func GetCaseMap(file string) {
	caseOfUse, _ := utilities.GetValuesFromFile(file)

	scenery := Scenery{
		Values: caseOfUse,
	}
	fmt.Println(scenery.Values)
}
