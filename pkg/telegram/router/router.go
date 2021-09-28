package router

import (
	"AdviserAPI/pkg/telegram/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var apiUrl = "https://www.boredapi.com/api/activity?"

func GetAdvice(person int) (string, error) {
	fmt.Println(person)
	if person < 1 || person > 5 {
		return "I can work with numbers range 1 - 5", nil
	}

	resp, err := http.Get(apiUrl + "participants=" + strconv.Itoa(person))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var bored models.BoredConfig
	json.Unmarshal(body, &bored)

	if bored.Type == "" {
		return "I can work only with numbers", nil
	}

	newMessage := strings.Title("Type: " + bored.Type + "\nActivity: " + bored.Activity +
		"\nFor " + strconv.Itoa(person) + " person..")

	return newMessage, nil
}
