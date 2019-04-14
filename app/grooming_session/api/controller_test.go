package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dmmmd/scrumpoker/app/grooming_session"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/baloo.v3"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// TODO use testing container
const apiRoot = "http://127.0.0.1"

var api = baloo.New(apiRoot)

func TestPostReturnsCreatedModel(t *testing.T) {
	title := getRandomTitle()

	model := grooming_session.NewGroomingSession()
	model.Title = title

	stored := sendPostModel(t, model)
	storedId := stored.GetId().String()

	sendDeleteModel(t, storedId)
}

func TestGetReturnsCreatedModel(t *testing.T) {
	title := getRandomTitle()

	newModel := grooming_session.NewGroomingSession()
	newModel.Title = title

	stored := sendPostModel(t, newModel)
	storedId := stored.GetId().String()

	_ = api.Get("/grooming_sessions/" + storedId).
		Expect(t).
		Status(200).
		Type("application/json").
		JSON(stored).
		Done()

	sendDeleteModel(t, storedId)
}

func TestDeleteModel(t *testing.T) {
	title := getRandomTitle()

	newModel := grooming_session.NewGroomingSession()
	newModel.Title = title

	stored := sendPostModel(t, newModel)
	storedId := stored.GetId().String()

	sendDeleteModel(t, storedId)
	_ = api.Get("/grooming_sessions/" + storedId).
		Expect(t).
		Status(404).
		Type("application/json").
		BodyMatchString("Resource not found").
		Done()
}

func TestIndexContainsCreatedModel(t *testing.T) {
	title := getRandomTitle()

	newModel := grooming_session.NewGroomingSession()
	newModel.Title = title

	stored := sendPostModel(t, newModel)

	expectedSubstring, _ := json.Marshal(stored)

	_ = api.Get("/grooming_sessions").
		Expect(t).
		Status(200).
		Type("application/json").
		BodyMatchString(string(expectedSubstring)).
		Done()

	storedId := stored.GetId().String()
	sendDeleteModel(t, storedId)
}

func sendPostModel(t *testing.T, model *grooming_session.GroomingSession) *grooming_session.GroomingSession {
	body, _ := json.Marshal(model)

	resp, _ := http.Post(apiRoot+"/grooming_sessions", "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(responseBody))

	var stored grooming_session.GroomingSession
	err := json.Unmarshal(responseBody, &stored)
	if nil != err {
		assert.Fail(t, err.Error())
	}

	return &stored
}

func sendDeleteModel(t *testing.T, id string) {
	_ = api.Delete("/grooming_sessions/" + id).
		Expect(t).
		Status(204).
		BodyLength(0).
		Done()

	//req, _ := http.NewRequest(http.MethodDelete, apiRoot + "/grooming_sessions/" + id, nil)
	//
	//client := &http.Client{}
	//resp, _ := client.Do(req)
	//defer resp.Body.Close()
	//
	//responseBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(responseBody))
	//
	//assert.Equal(t, 204, resp.Status)
}

func getRandomTitle() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return fmt.Sprintf("Test title #%d-%d", time.Now().Unix(), r1.Intn(1000))
}
