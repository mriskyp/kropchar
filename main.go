package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"

	"kropchar/infrastructure"
)

const (
	bulkLowerLetter  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bulkUpperletter  = "abcdefghijklmnopqrstuvwxyz"
	bulkNumberLetter = "0123456789"
	bulkLetter       = bulkLowerLetter + bulkUpperletter + bulkNumberLetter

	/**
	-6 bits to represent a letter index
	-All 1-bits, as many as letterIndexBits
	-# of letter indices fitting in 63 bits
	*/
	letterIndexBits = 6
	letterIndexMask = 1<<letterIndexBits - 1
	letterIndexMax  = 63 / letterIndexBits
)

func main() {
	// define size of random string
	lengthNumber := 32

	// example
	str := GenerateRandomRuneString(lengthNumber)
	fmt.Println(str)

	str = GenerateRandomByteString(lengthNumber)
	fmt.Println(str)

	str = GenerateRandomRemainderString(lengthNumber)
	fmt.Println(str)

	str = GenerateRandomStringBytesMask(lengthNumber)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrc(lengthNumber)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrcSB(lengthNumber)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrcUnsafe(lengthNumber)
	fmt.Println(str)

	str = GenerateRandSeq(lengthNumber)
	fmt.Println(str)
}

func GenerateRandSeq(n int) string {
	now := time.Now()

	defer infrastructure.EvaluateProcessTime("[GenerateRandSeq]", now)

	rand.Seed(time.Now().UnixNano())
	letters := []rune(bulkLetter)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
func GenerateRandStringBytesMaskImprSrcUnsafe(n int) string {

	now := time.Now()

	defer infrastructure.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrcUnsafe]", now)

	src := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIndexMax
		}
		if idx := int(cache & letterIndexMask); idx < len(bulkLetter) {
			b[i] = bulkLetter[idx]
			i--
		}
		cache >>= letterIndexBits
		remain--
	}

	stringTemp := *(*string)(unsafe.Pointer(&b))

	return stringTemp
}

func GenerateRandStringBytesMaskImprSrcSB(n int) string {
	now := time.Now()

	defer infrastructure.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrcSB]", now)

	src := rand.NewSource(time.Now().UnixNano())

	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIndexMax
		}
		if idx := int(cache & letterIndexMask); idx < len(bulkLetter) {
			sb.WriteByte(bulkLetter[idx])
			i--
		}
		cache >>= letterIndexBits
		remain--
	}

	return sb.String()
}

func GenerateRandomStringBytesMask(n int) string {
	now := time.Now()

	defer infrastructure.EvaluateProcessTime("[GenerateRandomStringBytesMask]", now)

	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIndexMask); idx < len(bulkLetter) {
			b[i] = bulkLetter[idx]
			i++
		}
	}

	return string(b)
}

func GenerateRandomByteString(length int) string {
	now := time.Now()
	defer infrastructure.EvaluateProcessTime("[GenerateRandomByteString]", now)

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func GenerateRandomRemainderString(n int) string {
	now := time.Now()
	defer infrastructure.EvaluateProcessTime("[GenerateRandomRemainderString]", now)

	b := make([]byte, n)
	for i := range b {
		b[i] = bulkLetter[rand.Int63()%int64(len(bulkLetter))]
	}

	return string(b)
}

func GenerateRandStringBytesMaskImprSrc(n int) string {
	now := time.Now()
	defer infrastructure.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrc]", now)

	src := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIndexMax
		}
		if idx := int(cache & letterIndexMask); idx < len(bulkLetter) {
			b[i] = bulkLetter[idx]
			i--
		}
		cache >>= letterIndexBits
		remain--
	}

	return string(b)
}

func GenerateRandomRuneString(length int) string {
	now := time.Now()
	defer infrastructure.EvaluateProcessTime("[GenerateRandomRuneString]", now)

	rand.Seed(time.Now().UnixNano())

	chars := []rune(bulkLetter)

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	// E.g. "ExcbsVQs"
	str := b.String()

	return str
}
