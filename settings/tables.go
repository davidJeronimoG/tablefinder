package settings

type Table struct {
	Name string
	// Map of conditions where 0 = false, 1 = true, 2 = unknown. It is not necessary to map the unknowns.
	ConditionsValues map[string]int
}

func GetTables() []Table {

	table1 := Table{
		Name: "213",
		ConditionsValues: map[string]int{
			"ATM":                 1,
			"CardnNotPresent":     1,
			"CardPresent":         1,
			"VSDC":                1,
			"NotVSDC":             1,
			"TandE":               1,
			"CPS":                 1,
			"AccountVerification": 1,
			"CreditVoucher":       1,
			"BalanceInquiry":      1,
			"CardAccountTransfer": 1,
			"isDeferredAuthorization": 1,
			"WithPIN":             1,
			"PINChangeUnblock":      1,
			"IncrementalAuthorization": 1,
			"Reversal":             1,
			"Authorization": 	  1,
			"Advice" :              1,
		},
	}

	table2 := Table{
		Name: "214",
		ConditionsValues: map[string]int{
			"ATM":                 1,
			"CardnNotPresent":     1,
			"CardPresent":         1,
			"VSDC":                1,
			"NotVSDC":             1,
			"TandE":               1,
			"CPS":                 1,
			"AccountVerification": 1,
			"CreditVoucher":       1,
			"BalanceInquiry":      1,
			"CardAccountTransfer": 1,
			"isDeferredAuthorization": 1,
			"WithPIN":             1,
			"PINChangeUnblock":      1,
			"IncrementalAuthorization": 1,
			"Reversal":             1,
			"Authorization": 	  1,
			"Advice" :              1,
		},
	}

	table3 := Table{
		Name: "215",
		ConditionsValues: map[string]int{
			"ATM":                 1,
			"CardnNotPresent":     1,
			"CardPresent":         1,
			"VSDC":                1,
			"NotVSDC":             1,
			"TandE":               1,
			"CPS":                 1,
			"AccountVerification": 1,
			"CreditVoucher":       1,
			"BalanceInquiry":      1,
			"CardAccountTransfer": 1,
			"isDeferredAuthorization": 1,
			"WithPIN":             1,
			"PINChangeUnblock":      1,
			"IncrementalAuthorization": 1,
			"Reversal":             1,
			"Authorization": 	  1,
			"Advice" :              1,
		},
	}

	return []Table{table1, table2, table3}
}
