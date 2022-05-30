package util

const (
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TRY:
		return true
	}
	return false
}
