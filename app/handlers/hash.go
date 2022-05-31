package handlers

import (
	"app/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Hash struct {
	logger       *log.Logger
	hashService  *services.Hash
	statsService *services.Stats
}

func NewHash() *Hash {
	logger := log.New(os.Stdout, "Hash Handler: ", log.LstdFlags)
	hashService := services.NewHash()
	statsService := services.NewStats()
	return &Hash{logger, hashService, statsService}
}

func (handler *Hash) GenerateHash(writer http.ResponseWriter, request *http.Request) {

	start := time.Now()

	err := request.ParseForm()
	if err != nil {
		handler.logger.Println(err)
		http.Error(writer, "Form Error Occurred!", http.StatusBadRequest)
		return
	}

	password := request.FormValue("password")
	if password == "" {
		http.Error(writer, "Password Required!", http.StatusBadRequest)
		return
	}

	identifier := handler.hashService.CreateHash(password)

	diff := time.Now().Sub(start)

	println(diff.Microseconds())

	handler.statsService.AddStats(diff.Microseconds())

	fmt.Fprint(writer, identifier)
}

func (handler *Hash) GetHashByIdentifier(writer http.ResponseWriter, request *http.Request) {

	regex := regexp.MustCompile("[0-9]+")
	identifier := regex.Find([]byte(request.URL.Path))

	if len(identifier) == 0 {
		http.Error(writer, "Identifier Required!", http.StatusBadRequest)
		return
	}

	id32, _ := strconv.ParseUint(string(identifier), 10, 32)

	keyStore := handler.hashService.GetHash(uint32(uint(id32)))

	if keyStore == nil {
		http.Error(writer, "Password For Key Not Found!", http.StatusNotFound)
	} else if keyStore.Hash == "" {
		http.Error(writer, "Password Hash Not Yet Been Processed!", http.StatusTooEarly)
	} else {
		fmt.Fprint(writer, keyStore.Hash)
	}

	return

}

func (handler *Hash) GetHashStats(writer http.ResponseWriter, request *http.Request) {

	statResults := handler.statsService.GetStats()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(statResults)
}
