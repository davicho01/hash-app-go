package services

import (
	"app/repository"
	"crypto/sha512"
	"encoding/base64"
	"log"
	"os"
	"time"
)

type Hash struct {
	logger               *log.Logger
	keyStorageRepository *repository.KeyStorage
}

func NewHash() *Hash {
	logger := log.New(os.Stdout, "Hash Handler: ", log.LstdFlags)
	keyStorage := repository.NewKeyStorage()
	return &Hash{logger, keyStorage}
}

func (service *Hash) CreateHash(password string) uint32 {

	keyStore := service.keyStorageRepository.CreateKey()

	go func() {
		time.Sleep(5 * time.Second)
		//Checking to see if there is a keyStore object and password
		if keyStore != nil && password != "" {
			newHash := service.HashPassword(password)
			keyStore.Hash = newHash
		}
	}()

	return keyStore.Id
}

func (service *Hash) GetHash(identifier uint32) *repository.KeyStore {

	return service.keyStorageRepository.GetKeyStore(identifier)
}

func (service *Hash) HashPassword(password string) string {

	sha512 := sha512.New()
	sha512.Write([]byte(password))

	service.logger.Printf("sha512 Base64:\t%s\t%s\n", password, base64.StdEncoding.EncodeToString(sha512.Sum(nil)))

	return base64.StdEncoding.EncodeToString(sha512.Sum(nil))
}
