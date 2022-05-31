package repository

import (
	"log"
	"os"
	"sync/atomic"
	"time"
)

type KeyStore struct {
	Id          uint32
	Hash        string
	CreatedDate time.Time
}

type KeyStorage struct {
	logger     *log.Logger
	identifier uint32
	keyStore   map[uint32]*KeyStore
}

func NewKeyStorage() *KeyStorage {
	logger := log.New(os.Stdout, "KeyStorage Repository: ", log.LstdFlags)
	identifier := uint32(1)
	keyStore := make(map[uint32]*KeyStore)
	return &KeyStorage{logger, identifier, keyStore}
}

func (repository *KeyStorage) CreateKey() *KeyStore {

	newId := repository.identifier

	repository.keyStore[newId] = &KeyStore{
		Id:          newId,
		Hash:        "",
		CreatedDate: time.Now(),
	}

	atomic.AddUint32(&repository.identifier, 1)

	return repository.keyStore[newId]
}

func (repository *KeyStorage) GetKeyStore(identifier uint32) *KeyStore {
	return repository.keyStore[identifier]
}
