package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	OneVNDong       float64 = 1
	OneThousandDong         = OneVNDong * 1000
	OneMillionDong          = OneThousandDong * 1000
	OneBillionDong          = OneMillionDong * 1000
)

// GetStringAmountVND format amount to string vnđ
func GetStringAmountVND(amount float64) string {
	return fmt.Sprintf("%s đ", NumberFormat(amount, 0))
}

// GetStringFormatAmount format amount
// 15345 -> 15.345đ , 15000 -> 15K
func GetStringFormatAmount(amount float64) string {
	if math.Mod(amount, OneThousandDong) != 0 {
		return fmt.Sprintf("%sđ", NumberFormat(amount, 0))
	}
	return FormatShortCurrencyNumberRoundDown(amount, 0)
}

func NumberFormat(val float64, precision int) string {
	var thousandSep = byte('.')
	var decSep = byte(',')
	// Parse the float as a string, with no exponent, and keeping precision
	// number of decimal places. Note that the precision passed in to FormatFloat
	// must be a positive number.
	usePrecision := precision
	if precision < 1 {
		usePrecision = 1
	}
	asString := strconv.FormatFloat(val, 'f', usePrecision, 64)
	// Split the string at the decimal point separator.
	separated := strings.Split(asString, ".")
	beforeDecimal := separated[0]
	// Our final string will need a total space of the original parsed string
	// plus space for an additional separator character every 3rd character
	// before the decimal point.
	withSeparator := make([]byte, 0, len(asString)+(len(beforeDecimal)/3))

	// Deal with a (possible) negative sign:
	if beforeDecimal[0] == '-' {
		withSeparator = append(withSeparator, '-')
		beforeDecimal = beforeDecimal[1:]
	}

	// Drain the initial characters that are "left over" after dividing the length
	// by 3. For example, if we had "12345", this would drain "12" from the string
	// append the separator character, and ensure we're left with something
	// that is exactly divisible by 3.
	initial := len(beforeDecimal) % 3
	if initial > 0 {
		withSeparator = append(withSeparator, beforeDecimal[0:initial]...)
		beforeDecimal = beforeDecimal[initial:]
		if len(beforeDecimal) >= 3 {
			withSeparator = append(withSeparator, thousandSep)
		}
	}

	// For each chunk of 3, append it and add a thousands separator,
	// slicing off the chunks of 3 as we go.
	for len(beforeDecimal) >= 3 {
		withSeparator = append(withSeparator, beforeDecimal[0:3]...)
		beforeDecimal = beforeDecimal[3:]
		if len(beforeDecimal) >= 3 {
			withSeparator = append(withSeparator, thousandSep)
		}
	}
	// Append everything after the '.', but only if we have positive precision.
	if precision > 0 {
		withSeparator = append(withSeparator, decSep)
		withSeparator = append(withSeparator, separated[1]...)
	}
	return string(withSeparator)
}

// ParseUint64 parse interface to uint64
func ParseUint64(i interface{}) uint64 {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(v)
}

// ParseInt64 parse interface to int64
func ParseInt64(i interface{}) int64 {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int64(v)
}

// ParseInt32 parse interface to int32
func ParseInt32(i interface{}) int32 {
	str := ToIntString(i)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int32(v)
}

// ToIntString convert an interface to integer string
func ToIntString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch i.(type) {
	case float32:
		return fmt.Sprintf("%.0f", i)
	case float64:
		return fmt.Sprintf("%.0f", i)
	case int:
		return fmt.Sprintf("%d", i)
	case int8:
		return fmt.Sprintf("%d", i)
	case int16:
		return fmt.Sprintf("%d", i)
	case int32:
		return fmt.Sprintf("%d", i)
	case int64:
		return fmt.Sprintf("%d", i)
	case uint:
		return fmt.Sprintf("%d", i)
	case uint8:
		return fmt.Sprintf("%d", i)
	case uint16:
		return fmt.Sprintf("%d", i)
	case uint32:
		return fmt.Sprintf("%d", i)
	case uint64:
		return fmt.Sprintf("%d", i)
	}
	return fmt.Sprint(i)
}

// FormatShortCurrencyNumberRoundDown format number to short string with round down floating-point value
// Example:
// 1. 123.567, precision = 0 -> 123đ
// 2. 123.567, precision = 1 -> 123,5đ
// 3. 123.567, precision = 2 -> 123,56đ
func FormatShortCurrencyNumberRoundDown(amount float64, precision int) string {
	var shortAmount float64
	var shortCurrency string
	if amount >= OneBillionDong {
		shortAmount = amount / OneBillionDong
		shortCurrency = " Tỷ"
	} else if amount >= OneMillionDong {
		shortAmount = amount / OneMillionDong
		shortCurrency = "Tr"
	} else if amount >= OneThousandDong {
		shortAmount = amount / OneThousandDong
		shortCurrency = "K"
	} else {
		shortAmount = amount
		shortCurrency = "đ"
	}
	usedPrecision := precision
	if usedPrecision < 0 {
		usedPrecision = 0
	}
	pow := math.Pow10(usedPrecision)
	roundDownAmount := math.Floor(shortAmount*pow) / pow
	usedPrecision = RemoveRedundantZeroNumber(roundDownAmount, usedPrecision)
	formattedAmount := NumberFormat(roundDownAmount, usedPrecision)
	return fmt.Sprintf("%s%s", formattedAmount, shortCurrency)
}

// RemoveRedundantZeroNumber remove zero number after floating-points
// Example: 1.234001, p = 5 ==> new precision = 3 (1.234)
func RemoveRedundantZeroNumber(value float64, maxPrecision int) int {
	if maxPrecision < 0 {
		return 0
	}
	precisionCount := maxPrecision
	for i := maxPrecision; i >= 1; i-- {
		v2 := int64(value * math.Pow10(i))
		v3 := int64(value*math.Pow10(i-1)) * 10
		if v2-v3 == 0.0 {
			precisionCount--
		} else {
			break
		}
	}
	return precisionCount
}
