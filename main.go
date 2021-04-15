package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main() {
	randFile := SecureRandomAlphaString(12)
	file, _ := os.Create("result" + randFile + ".txt")
	for i := 0; i < 100000000; i++ {
		tryBlock := SecureRandomAlphaString(64)
		h := sha256.New()
		h.Write([]byte(tryBlock))
		encodedString := fmt.Sprintf("%x", h.Sum(nil))
		zeros := findStartingZeros(encodedString)

		if i%100000 == 0 {
			fmt.Println(i)
		}
		if zeros > 10 {
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			fmt.Println(tryBlock)
			fmt.Println("------------------------------------------------------------------------")
			fmt.Println(encodedString)
			fmt.Println(zeros)
			file.WriteString(fmt.Sprintf("%s - %v \n", tryBlock, zeros))
		}
	}
	file.Close()

}

// //RandString get random string with high performance
// func RandString(n int) string {
// 	var src = rand.NewSource(time.Now().UnixNano())
// 	const letterBytes = "abcdef0123456789"
// 	const (
// 		letterIdxBits = 6                    // 6 bits to represent a letter index
// 		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
// 		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
// 	)
// 	b := make([]byte, n)
// 	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
// 	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
// 		if remain == 0 {
// 			cache, remain = src.Int63(), letterIdxMax
// 		}
// 		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
// 			b[i] = letterBytes[idx]
// 			i--
// 		}
// 		cache >>= letterIdxBits
// 		remain--
// 	}

// 	return *(*string)(unsafe.Pointer(&b))
// }

func SecureRandomAlphaString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func findStartingZeros(block string) int {
	runes := []rune(block)

	for i, c := range runes {
		if c != '0' {
			return i
		}
	}
	return len(block)
}
