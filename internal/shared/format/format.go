package format

import "strings"

func FormatPrice(price string) string {
	parts := strings.Split(price, ".")
	integer := parts[0]
	decimal := ""
	if len(parts) > 1 {
		decimal = "." + parts[1]
	}

	formattedInteger := ""
	for i, r := range integer {
		if i > 0 && (len(integer)-i)%3 == 0 {
			formattedInteger += " "
		}
		formattedInteger += string(r)
	}
	return formattedInteger + decimal
}
