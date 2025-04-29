package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"errors"
	"io"
)

// Kilde: https://gist.github.com/ayubmalik/2c973c2a7ae7e0d22ece7f5c4dfbd726
// EncryptedReader wraps r with an OFB cipher stream.
func EncryptedReader(key string, r io.Reader) (*cipher.StreamReader, error) {

	// read initial value
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if err != nil || n != len(iv) {
		return nil, errors.New("could not read initial value")
	}

	block, err := newBlock(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func newBlock(key string) (cipher.Block, error) {
	hash := md5.Sum([]byte(key))
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, err
	}
	return block, nil
}
