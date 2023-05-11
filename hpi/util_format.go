package hpi

import (
	"fmt"
	"strings"
)

// FormatNumber is function to format number and return string format
func (c *Transform) FormatNumber(num float64, intDigit int64, padding string, decDigit int64, thousandSeparator string, decimalSeparator string) string {

	// init result
	result := ""

	if decDigit < 0 {
		decDigit = 0
	}

	// float value
	decDigitStr := fmt.Sprintf("%d", decDigit)
	fFormat := "%0." + decDigitStr + "f"
	fValue := fmt.Sprintf(fFormat, num)
	dValue := "0"
	arr := strings.Split(fValue, ".")

	// validate value
	if len(arr) > 0 {
		dValue = arr[0]
		if len(arr) > 1 {
			fValue = arr[1]
		} else {
			fValue = "0"
		}
	}

	// Add Padding
	dDiff := int(intDigit) - len(dValue)
	var sbD strings.Builder
	for i := 0; i < dDiff; i++ {
		sbD.WriteString(padding)
	}
	sbD.WriteString(dValue)
	dValue = sbD.String()

	// Add thousandSeparator
	dResult := ""
	count := 0
	for i := len(dValue) - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			dResult = thousandSeparator + dResult
		}
		s := dValue[i]
		dResult = string(s) + dResult
		count++
	}

	result = dResult + decimalSeparator + fValue

	return result
}

// FormatNaming is function to format naming
func (c *Transform) FormatNaming(str string, format string) string {

	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	format = strings.TrimSpace(format)
	if format == "" {
		format = "LowerCase"
	}
	format = strings.ReplaceAll(format, " ", "")

	// Formating...
	if strings.EqualFold(format, "LowerCase") {
		return strings.ToLower(str)
	} else if strings.EqualFold(format, "UpperCase") {
		return strings.ToUpper(str)
	} else if strings.EqualFold(format, "FirstCharLowerCase") {
		fChar := str[:1]
		return strings.ToLower(fChar) + str[1:]
	} else if strings.EqualFold(format, "FirstCharUpperCase") {
		fChar := str[:1]
		return strings.ToUpper(fChar) + str[1:]
	} else if strings.EqualFold(format, "RemoveSpace") {
		str = strings.ReplaceAll(str, " ", "")
		return str
	} else if strings.EqualFold(format, "RemoveSpecialChar") {
		var sb strings.Builder
		for _, rune := range str {

			if rune >= 48 && rune <= 57 {
				sb.WriteRune(rune)
			} else if rune >= 65 && rune <= 90 {
				sb.WriteRune(rune)
			} else if rune >= 97 && rune <= 122 {
				sb.WriteRune(rune)
			} else if rune == 95 {
				sb.WriteRune(rune)
			} else if rune == 36 {
				sb.WriteRune(rune)
			}
		}
		return sb.String()
	} else if strings.EqualFold(format, "RemoveFirstCharNumber") {

		var sb strings.Builder
		canAdd := false
		for _, rune := range str {

			if rune >= 48 && rune <= 57 {
				// Ignore
			} else {
				canAdd = true
			}

			if canAdd {
				sb.WriteRune(rune)
			}
		}
		return sb.String()

	}

	return str
}
