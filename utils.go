package main

import (
	"golang.org/x/crypto/blowfish"
	"crypto/cipher"
)



func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	var iv = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	plaintext = checksizeAndPad(plaintext)

	ecipher, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, blowfish.BlockSize+len(plaintext))

	ecbc := cipher.NewCBCEncrypter(ecipher, iv)
	ecbc.CryptBlocks(ciphertext[blowfish.BlockSize:], plaintext)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	var iv = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	dcipher, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err)
	}

	decrypted := ciphertext[blowfish.BlockSize:]
	if len(decrypted)%blowfish.BlockSize != 0 {
		panic("decrypted is not a multiple of blowfish.BlockSize")
	}

	dcbc := cipher.NewCBCDecrypter(dcipher, iv)
	dcbc.CryptBlocks(decrypted, decrypted)

	return decrypted, nil
}

// checksizeAndPad checks the size of the plaintext and pads it if necessary.
// Blowfish is a block cipher, thus the plaintext needs to be padded to
// a multiple of the algorithms blocksize (8 bytes).
// return the multiple-of-blowfish.BlockSize-sized plaintext
func checksizeAndPad(plaintext []byte) []byte {

	modulus := len(plaintext) % blowfish.BlockSize
	if modulus != 0 {
		padlen := blowfish.BlockSize - modulus

		// add required padding
		for i := 0; i < padlen; i++ {
			plaintext = append(plaintext, 0)
		}
	}

	return plaintext
}