package settings

import "strings"

// returns a slice of functions that can be used to compare values
func GetFunctionalConditions() [](func(m map[string]string) (bool, string)) {

	isAtm := func(m map[string]string) (bool, string) {
		return m["F18"] == "6011", "ATM"
	}

	isCardnNotPresent := func(m map[string]string) (bool, string) {
		return m["F22.1"] == "01" || m["F22.1"] == "10", "CardnNotPresent"
	}

	isCardPresent := func(m map[string]string) (bool, string) {
		return m["F22.1"] == "02" || m["F22.1"] == "03" || m["F22.1"] == "90" || m["F22.1"] == "91" || m["F22.1"] == "05" || m["F22.1"] == "07" || m["F22.1"] == "95", "CardPresent"
	}

	isVSDC := func(m map[string]string) (bool, string) {
		return m["F22.1"] == "05" || m["F22.1"] == "07" || m["F22.1"] == "95", "VSDC"
	}

	isNotVSDC := func(m map[string]string) (bool, string) {
		return m["F22.1"] == "02" || m["F22.1"] == "03" || m["F22.1"] == "90" || m["F22.1"] == "91", "NotVSDC"
	}

	isTandE := func(m map[string]string) (bool, string) {
		return m["F62.4"] == "H" || m["F62,4"] == "A", "TandE"
	}

	isCPS := func(m map[string]string) (bool, string) {
		return m["F62.3"] != "", "CPS"
	}

	isAccountVerification := func(m map[string]string) (bool, string) {
		return m["F4"] == "000000000000" && m["F25"] == "51", "AccountVerification"
	}

	isCreditVoucher := func(m map[string]string) (bool, string) {
		return m["F3.1"] == "20", "CreditVoucher"
	}

	isBalanceInquiry := func(m map[string]string) (bool, string) {
		return m["F3.1"] == "30", "BalanceInquiry"
	}

	isCardAccountTransfer  := func(m map[string]string) (bool, string) {
		return m["F3.1"] == "40", "CardAccountTransfer"
	}

	isDeferredAuthorization := func(m map[string]string) (bool, string) {
		return m["F63.3"] == "5206", "DeferredAuthorization"
	}

	isWithPIN := func(m map[string]string) (bool, string) {
		return m["F52"] != "", "WithPIN"
	}

	isPINChangeUnblock := func(m map[string]string) (bool, string) {
		return m["F3.1"] == "70" || m["F3.1"] == "72", "PINChangeUnblock"
	}

	isIncrementalAuthorization := func(m map[string]string) (bool, string) {
		return m["F62.1"] == "I", "IncrementalAuthorization"
	}

	isReversal := func(m map[string]string) (bool, string) {
		return strings.Split(m["MTI"], "")[0] == "0" && strings.Split(m["MTI"], "")[1] == "4", "Reversal"
	}

	isAuthorization := func(m map[string]string) (bool, string) {
		return (strings.Split(m["MTI"], "")[0] == "0" && (strings.Split(m["MTI"], "")[1] == "1") && !(strings.Split(m["MTI"], "")[2] == "2" && strings.Split(m["MTI"], "")[3] == "0") || strings.Split(m["MTI"], "")[1] == "2"), "Authorization"
	}

	isAdvice := func(m map[string]string) (bool, string) {
		return (strings.Split(m["MTI"], "")[2] == "2" && strings.Split(m["MTI"], "")[3] == "0"), "Advice"
	}

	// isPOS := func(m map[string]string) (bool, string) {
	// 	return m["F18"] == "6011", "POS"
	// }


	return [](func(m map[string]string) (bool, string)){isAtm, 
		isCardnNotPresent, isCardPresent, 
		isVSDC, isNotVSDC, isTandE, isCPS, 
		isAccountVerification, isCreditVoucher,
		isBalanceInquiry, isCardAccountTransfer, 
		isDeferredAuthorization, isWithPIN,
		isPINChangeUnblock, isIncrementalAuthorization,
		isReversal, isAuthorization, isAdvice}

}
