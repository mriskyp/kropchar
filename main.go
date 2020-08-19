package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"

	inf "kropchar/infrastructure"
)

const (
	bulkLowerLetter  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bulkUpperletter  = "abcdefghijklmnopqrstuvwxyz"
	bulkNumberLetter = "0123456789"
	bulkLetter       = bulkLowerLetter + bulkUpperletter + bulkNumberLetter
	letterIndexBits  = 6
	letterIndexMask  = 1<<letterIndexBits - 1
	letterIndexMax   = 63 / letterIndexBits
)

func main() {
	// define size of random string
	randomSize := 32

	// example
	str := GenerateRandomRuneString(randomSize)
	fmt.Println(str)

	str = GenerateRandomByteString(randomSize)
	fmt.Println(str)

	str = GenerateRandomRemainderString(randomSize)
	fmt.Println(str)

	str = GenerateRandomStringBytesMask(randomSize)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrc(randomSize)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrcSB(randomSize)
	fmt.Println(str)

	str = GenerateRandStringBytesMaskImprSrcUnsafe(randomSize)
	fmt.Println(str)

	str = GenerateRandSeq(randomSize)
	fmt.Println(str)
}

// GenerateRandSeq is a func to generate random string by general use
func GenerateRandSeq(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandSeq]", now)

	rand.Seed(now.UnixNano())

	letters := []rune(bulkLetter)
	// Size rune bulk letter
	size := len(letters)

	byteData := make([]rune, randomSize)
	for i := range byteData {
		byteData[i] = letters[rand.Intn(size)]
	}

	// Return byte data as string
	return string(byteData)
}

// GenerateRandStringBytesMaskImprSrcUnsafe is a func to generate random string by byte mask, rand source unsafe type pointer
func GenerateRandStringBytesMaskImprSrcUnsafe(randomSize int) string {

	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrcUnsafe]", now)

	src := rand.NewSource(now.UnixNano())
	// Int63 returns a non-negative pseudo-random 63-bit integer as an int64 from the default Source.
	srcInt63 := src.Int63()

	byteData := make([]byte, randomSize)
	// Generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := (randomSize - 1), srcInt63, letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = srcInt63, letterIndexMax
		}

		idx := int(cache & letterIndexMask)
		if idx < len(bulkLetter) {
			byteData[i] = bulkLetter[idx]
			i--
		}

		cache >>= letterIndexBits
		remain--
	}

	//  pointers without the restrictions made for safe pointers
	unsf := unsafe.Pointer(&byteData)
	stringTemp := *(*string)(unsf)

	return stringTemp
}

// GenerateRandStringBytesMaskImprSrcSB is a func to generate random string by byte mask, rand source, string builder
func GenerateRandStringBytesMaskImprSrcSB(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrcSB]", now)

	src := rand.NewSource(now.UnixNano())
	// Int63 returns a non-negative pseudo-random 63-bit integer as an int64 from the default Source.
	srcInt63 := src.Int63()

	// A string builder
	sb := strings.Builder{}

	// Increase size of inner capacity to value capacity
	sb.Grow(randomSize)

	// Generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := randomSize-1, srcInt63, letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = srcInt63, letterIndexMax
		}
		if idx := int(cache & letterIndexMask); idx < len(bulkLetter) {
			sb.WriteByte(bulkLetter[idx])
			i--
		}
		cache >>= letterIndexBits
		remain--
	}

	// return string builder string
	return sb.String()
}

// GenerateRandomStringBytesMask is a func to generate random string by byte mask
func GenerateRandomStringBytesMask(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandomStringBytesMask]", now)

	byteData := make([]byte, randomSize)

	for i := 0; i < randomSize; {
		randInt63 := rand.Int63()
		if idx := int(randInt63 & letterIndexMask); idx < len(bulkLetter) {
			byteData[i] = bulkLetter[idx]
			i++
		}
	}

	// return byteData as string
	return string(byteData)
}

// GenerateRandomByteString is a func to generate random string by byte
func GenerateRandomByteString(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandomByteString]", now)

	charset := bulkLetter
	seededRand := rand.New(rand.NewSource(now.UnixNano()))
	byteData := make([]byte, randomSize)
	for i := range byteData {
		byteData[i] = charset[seededRand.Intn(len(charset))]
	}

	// return byteData as string
	return string(byteData)
}

// GenerateRandomRemainderString is a func to generate random string by remainder int64
func GenerateRandomRemainderString(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandomRemainderString]", now)

	byteData := make([]byte, randomSize)
	for i := range byteData {
		size := len(bulkLetter)
		randInt63 := rand.Int63()

		indexRemainder := randInt63 % int64(size)
		byteData[i] = bulkLetter[indexRemainder]
	}

	// return byteData as string
	return string(byteData)
}

// GenerateRandStringBytesMaskImprSrc is a func to generate random string by bytes mask random
func GenerateRandStringBytesMaskImprSrc(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandStringBytesMaskImprSrc]", now)

	src := rand.NewSource(now.UnixNano())
	// Int63 returns a non-negative pseudo-random 63-bit integer as an int64 from the default Source.
	srcInt63 := src.Int63()

	byteData := make([]byte, randomSize)

	// Generates 63 random bits, enough for letterIndexMax characters!
	for i, cache, remain := (randomSize - 1), srcInt63, letterIndexMax; i >= 0; {
		if remain == 0 {
			cache, remain = srcInt63, letterIndexMax
		}

		idx := int(cache & letterIndexMask)
		if idx < len(bulkLetter) {
			byteData[i] = bulkLetter[idx]
			i--
		}
		cache >>= letterIndexBits
		remain--
	}

	//Return byte as string
	return string(byteData)
}

// GenerateRandomRuneString is a func to generate random string by rune
func GenerateRandomRuneString(randomSize int) string {
	now := time.Now()
	defer inf.EvaluateProcessTime("[GenerateRandomRuneString]", now)

	rand.Seed(now.UnixNano())

	charset := []rune(bulkLetter)

	var builder strings.Builder
	for i := 0; i < randomSize; i++ {
		index := rand.Intn(len(charset))
		builder.WriteRune(charset[index])
	}

	str := builder.String()

	// return byteData as string
	return str
}
