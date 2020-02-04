package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mongodb-example/src/domain/dto"
	"mongodb-example/src/interfaces/api/utils"
	"net/http"
	"testing"
	"time"
)

func TestPostV1User(t *testing.T) {
	server := InitTestServer()
	defer server.Close()

	req := dto.User{
		Email:     "test@example.con",
		LastName:  "Smith",
		Country:   "US",
		City:      "New York",
		Gender:    "Male",
		BirthDate: dto.BirthDate(time.Now()),
	}

	raw, err := json.Marshal(req)

	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(
		server.URL+"/v1/user",
		utils.ApplicationJson,
		bytes.NewBuffer(raw),
	)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Status %v, body: %v", resp.StatusCode, string(body))
	}

	//TODO: try unmarshal to dto.Id and get user by id

}
