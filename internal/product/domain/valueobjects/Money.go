package valueobjects

import "fmt"

type (
	CurrencyTokenPosition int

	Currency struct {
		Name          string
		Token         string
		TokenPosition CurrencyTokenPosition
		DecimalToken  string
		ThousandToken string
	}

	Money struct {
		Value    int
		Currency Currency
	}
)

const (
	CurrencyTokenOnStart CurrencyTokenPosition = iota
	CurrencyTokenOnEnd
)

func (currency *Currency) FormatString(integerPart int, decimalPart int) string {
	intStr := fmt.Sprintf("%d", integerPart)
	n := len(intStr)
	if n > 3 {
		var result string
		for i, c := range intStr {
			if i != 0 && (n-i)%3 == 0 {
				result += currency.ThousandToken
			}
			result += string(c)
		}
		intStr = result
	}

	// Format fractional part
	fracStr := fmt.Sprintf("%02d", decimalPart)

	// Combine with decimal token
	formatted := intStr + currency.DecimalToken + fracStr

	// Add currency token
	if currency.TokenPosition == CurrencyTokenOnStart {
		formatted = currency.Token + formatted
	} else {
		formatted = formatted + " " + currency.Token
	}

	return formatted
}

func (money *Money) FormatString() string {
	absValue := money.Value
	integerPart := absValue / 100
	fractionalPart := absValue % 100

	return money.Currency.FormatString(integerPart, fractionalPart)
}

func (money *Money) Validate() error {
	return nil
}
