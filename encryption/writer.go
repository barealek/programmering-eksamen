package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Kilde: https://gist.github.com/ayubmalik/2c973c2a7ae7e0d22ece7f5c4dfbd726
// EncryptedWriter wraps w with an OFB cipher stream.
func EncryptedWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {

	// generate random initial value
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// write clear IV to allow for decryption
	n, err := w.Write(iv)
	if err != nil || n != len(iv) {
		return nil, errors.New("could not write initial value")
	}

	block, err := newBlock(key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	return &cipher.StreamWriter{S: stream, W: w}, nil
}
