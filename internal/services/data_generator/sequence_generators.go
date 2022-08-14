package data_generator

import (
	"math/rand"
	"strconv"
)

const lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
const digitChars = "0123456789"
const alphaNumeric = lowercaseChars + digitChars

func generateAlphaNumeric(length int) string {
	byteStr := make([]byte, length)
	for i := range byteStr {
		byteStr[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(byteStr)
}

func generateWord(length int) string {
	byteStr := make([]byte, length)
	for i := range byteStr {
		byteStr[i] = lowercaseChars[rand.Intn(len(lowercaseChars))]
	}
	return string(byteStr)
}

func generateNumber(length int) string {
	byteStr := make([]byte, length)

	// do not generate a leading zero
	byteStr[0] = digitChars[1:][rand.Intn(len(digitChars)-1)]

	for i := 1; i < length; i++ {
		byteStr[i] = digitChars[rand.Intn(len(digitChars))]
	}
	return string(byteStr)
}

func generateDate() string {
	year := strconv.Itoa(2007 + rand.Intn(15))
	month := strconv.Itoa(rand.Intn(13))
	day := strconv.Itoa(rand.Intn(31))

	hours := strconv.Itoa(rand.Intn(24))
	if numDigits := len(hours); numDigits == 1 {
		hours = "0" + hours
	}
	minutes := strconv.Itoa(rand.Intn(60))
	seconds := strconv.Itoa(rand.Intn(60))

	return year + "-" + month + "-" + day + "T" + hours + ":" + minutes + ":" + seconds + "Z"
}
