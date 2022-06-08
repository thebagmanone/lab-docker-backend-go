package routers

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"strconv"
)

func RandomNumber(rw http.ResponseWriter, r *http.Request) {

	type Response struct {
		Number int `json:"number"`
	}
	numbers, err := RandomNumbers(6)
	if err != nil {
		http.Error(rw, "Error obteniendo numero", 400)
		return
	}

	var res Response
	intNumb, _ := strconv.Atoi(numbers)
	res.Number = intNumb

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(res)

}

const otpChars = "1234567890"

func RandomNumbers(quantity int) (string, error) {

	buffer := make([]byte, quantity)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < quantity; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil

}
