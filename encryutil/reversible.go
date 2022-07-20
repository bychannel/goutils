package encryutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// 可逆加密算法
// 对称算法 DES -> AES
// 不对称算法 RSA -> DSA

// DesEncrypt des加密逻辑
// salt为加解密的salt值，最大长度为8，超过长度会抛出相应异常
func DesEncrypt(text string, key ...string) (string, error) {
	k := []byte(defaultDesKey)
	if len(key) > 0 {
		k = []byte(key[0])
		if len(k) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
	}
	src := []byte(text)
	block, err := des.NewCipher(k)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("What is required is an integer multiple of the size")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

//DesDecrypt des解密逻辑
func DesDecrypt(decrypted string, key ...string) (string, error) {
	k := []byte(defaultDesKey)
	if len(key) > 0 {
		if len(key[0]) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
		k = []byte(key[0])
	}
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(k)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = zeroUnPadding(out)
	return string(out), nil
}

// AesEncrypt aes加密逻辑
// 如果salt存在，则使用传递的变量。如果不存在，则使用系统默认的值。
func AesEncrypt(original string, key ...string) (string, error) {
	var saltValue = defaultAesKey
	if len(key) > 0 {
		saltValue = key[0]
	}

	// Convert to byte array
	origData := []byte(original)
	k := []byte(saltValue)
	if len(k) != 16 && len(k) != 24 && len(k) != 32 {
		return "", errors.New("The length of the key should be 16 or 24 or 32")
	}
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// Completion code
	origData = pKCS7Padding(origData, blockSize)
	// encryption mode
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// create array
	crated := make([]byte, len(origData))
	// encryption
	blockMode.CryptBlocks(crated, origData)
	return base64.StdEncoding.EncodeToString(crated), nil
}

// AesDecrypt aes解密逻辑
func AesDecrypt(crated string, salt ...string) (string, error) {
	var saltValue = defaultAesKey
	if len(salt) > 0 {
		saltValue = salt[0]
	}
	// Convert to byte array
	cratedByte, _ := base64.StdEncoding.DecodeString(crated)
	k := []byte(saltValue)
	if len(k) != 16 && len(k) != 24 && len(k) != 32 {
		return "", errors.New("The length of the key should be 16 or 24 or 32")
	}
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// encryption mode
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// create array
	orig := make([]byte, len(cratedByte))
	// decrypt
	blockMode.CryptBlocks(orig, cratedByte)
	// to complete the code
	orig = pKCS7UnPadding(orig)
	return string(orig), nil
}

// fixme rsa加密逻辑

// fixme rsa解密逻辑

// fixme dsa加密逻辑

// fixme dsa解密逻辑
