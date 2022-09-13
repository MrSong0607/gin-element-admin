package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/pquerna/otp/totp"
)

func HashEncryptToHex(data []byte) string {
	hashed := md5.Sum(data)
	return Hex(hashed[:])
}

func HmacSHA256(key string, data []byte) ([]byte, error) {
	h := hmac.New(sha256.New, []byte(key))
	if _, err := h.Write(data); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func HmacMd5(key, data []byte) ([]byte, error) {
	h := hmac.New(md5.New, key)
	if _, err := h.Write(data); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func GenOtpSecret(user, org string) (secret, uri string, e error) {
	key, e := totp.Generate(totp.GenerateOpts{
		Issuer:      org,
		AccountName: user,
	})

	if e != nil {
		return
	}

	return key.Secret(), key.URL(), e
}

func ValidOtpCode(secret, passcode string) bool {
	return totp.Validate(passcode, secret)
}

func AesGcmEncrypt(plaintext, key []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText = aesGcm.Seal(nil, nonce, plaintext, nil)
	return
}

func AesGcmDecrypt(cipherText, key []byte) (plaintext []byte, err error) {
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = aesGcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return
}

func EncryptMap(fileds map[string]string, secret string) (signStr string, sign []byte) {
	key := []string{}
	for k := range fileds {
		key = append(key, k)
	}

	sort.Strings(key)
	pair := []string{}
	for _, k := range key {
		if len(fileds[k]) == 0 {
			continue
		}
		pair = append(pair, k+"="+fileds[k])
	}

	signStr = strings.Join(pair, "&")
	sign, _ = HmacSHA256(secret, []byte(signStr))

	return
}

func BytesEqual(b1, b2 []byte) bool {
	return bytes.Equal(b1, b2)
}
