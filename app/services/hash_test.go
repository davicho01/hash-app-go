package services

import (
	"testing"
	"time"
)

func TestHash_HashPassword(t *testing.T) {

	hashService := NewHash()

	hash := hashService.HashPassword("angryMonkey")
	if hash != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Error("Incorrect hash")
	}

	hash = hashService.HashPassword("password")
	if hash != "sQnzu7wkTrgkQZF+0G1hi5AI3Qmzvv0bXgc5THBqi7mAsdd4Xll27ASbRt9fEyavWi6m0QP9B8lThf+rDKy8hg==" {
		t.Error("Incorrect hash")
	}
}

func TestHash_CreateHash(t *testing.T) {

	hashService := NewHash()
	identifier := hashService.CreateHash("angryMonkey")

	if identifier != 1 {
		t.Error("Identifier should be value 1 got:", identifier)
		return
	}

}

func TestHash_GetHash(t *testing.T) {

	hashService := NewHash()

	keyStore := hashService.GetHash(1)

	if keyStore != nil {
		t.Error("KeyStore object must be nil")
		return
	}

	identifier := hashService.CreateHash("angryMonkey")

	if identifier != 1 {
		t.Error("Identifier should be value 1 got:", identifier)
		return
	}

	keyStore = hashService.GetHash(identifier)

	if keyStore.Hash != "" {
		t.Error("KeyStore.Hash must be empty")
		return
	}

	time.Sleep(6 * time.Second)

	keyStore = hashService.GetHash(identifier)

	if keyStore.Hash == "" {
		t.Error("KeyStore.Hash must NOT be empty")
		return
	}

	if keyStore.Hash != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Error("Incorrect hash got:", keyStore.Hash)
		return
	}

}
