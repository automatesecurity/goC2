// go-c2-crypto.go
// Supports encoding, encryption and obfuscation functions

package main

import (
    "io"
    "os"
    "encoding/base64"
    "crypto/sha512"
)

// shaHash generates a SHA512 hash and then hex encodes it
// Reference: https://stackoverflow.com/questions/10701874/generating-the-sha-hash-of-a-string-using-golang
func shaHash(text string) string {
	hasher := sha512.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

//
func computeSHA(filepath string) []byte {
	var result []byte
	file, err := os.Open(filepath)
	if err != nil {
		return nil // If there is an error do nothing
	}
	defer file.Close()
	hash := shaHash.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil // If there is an error do nothing
	}
	return hash.Sum(result)
}

// base64Encode encoded base64 strings
func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64Decode takes an encoded base64 string and decodes it
func base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

// Obfuscates the string data via basic rotation cipher
func obfuscate(Data string) string {
	var ObfuscateText string
	for i := 0, i < len(Data); i++ {
		ObfuscateText += string(int(Data[i]) + 1)
	}
	return ObfuscateText
}

// Deobfuscate returns the clear text value of the string
func deobfuscate(Data string) string {
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i]) - 1)
	}
	return ClearText
}