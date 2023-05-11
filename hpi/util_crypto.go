package hpi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

// Constant Value
const chipherKey64Bit = "9514365ed2bfd049a3877465ebd547210a492db4edf03e9506a3f94213532158"
const minChar = 16

// Read Me https://gist.github.com/kkirsche/e28da6754c39d5e7ea10

// EncryptWithAes256Gcm is function to encrypt with AES256GCM
func (c *Transform) EncryptWithAes256Gcm(text string, passcode string) (string, error) {

	//Binary of cipher key
	dStr, err := hex.DecodeString(chipherKey64Bit)
	if err != nil {
		return "", err
	}
	plaintext := []byte(text)

	// Prepare
	//===============================
	hasher := sha512.New()
	hasher.Write(dStr)
	out := hex.EncodeToString(hasher.Sum(nil))

	newKey, err := hex.DecodeString(out[:64])
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(out[64:(64 + 24)])
	if err != nil {
		return "", err
	}

	// check passcoce
	aData := c.checkPassCode(passcode, minChar)

	//===============================

	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesgcm.Seal(nil, nonce, plaintext, aData)

	return hex.EncodeToString(cipherText), nil
}

// DecryptWithAes256Gcm is function to decrypt with AES256GCM
func (c *Transform) DecryptWithAes256Gcm(encryptedText string, passcode string) (string, error) {

	//Binary of cipher key
	dStr, err := hex.DecodeString(chipherKey64Bit)
	if err != nil {
		return "", err
	}

	// Prepare
	//===============================

	hasher := sha512.New()
	hasher.Write(dStr)
	out := hex.EncodeToString(hasher.Sum(nil))

	newKey, err := hex.DecodeString(out[:64])
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(out[64:(64 + 24)])
	if err != nil {
		return "", err
	}

	//===============================

	aData := c.checkPassCode(passcode, minChar)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	output, err := aesgcm.Open(nil, nonce, cipherText, aData)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// checkPassCode is function to check passcode 64 bit
func (c *Transform) checkPassCode(passcode string, minChar int) []byte {

	// validate bit
	if len(passcode) > minChar {
		minChar = len(passcode)
	}

	var sb strings.Builder
	pLen := len(passcode)
	for i := 0; i < minChar; i++ {
		b := byte('0')
		if i < pLen {
			b = passcode[i]
		}
		sb.WriteByte(b)
	}

	return []byte(sb.String())
}
