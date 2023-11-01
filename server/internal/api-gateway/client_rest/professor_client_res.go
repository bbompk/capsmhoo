package client_rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"capsmhoo/internal/api-gateway/model"

	"github.com/spf13/viper"
)

type ProfessorClient struct {
	client *http.Client
}

type ProfessorClientRest interface {
	GetAllProfessors() ([]model.Professor, error)
	GetProfessorByID(id string) (model.Professor, error)
	CreateProfessor(professor model.ProfessorRequestBody) (model.Professor, error)
	UpdateProfessorByID(id string, professor model.ProfessorRequestBody) (model.Professor, error)
	DeleteProfessorByID(id string) (model.Professor, error)
}

func (s *ProfessorClient) GetAllProfessors() ([]model.Professor, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/professor"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp model.ProfessorListResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != "200" {
		return nil, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *ProfessorClient) GetProfessorByID(id string) (model.Professor, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/professor/" + id

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.Professor{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.ProfessorResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Professor{}, err
	}
	if resp.Code != "200" {
		return model.Professor{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *ProfessorClient) CreateProfessor(professor model.ProfessorRequestBody) (model.Professor, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/professor"

	// prepare request body
	byteData, err := json.Marshal(professor)
	if err != nil {
		return model.Professor{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.Professor{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.ProfessorResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Professor{}, err
	}
	if resp.Code != "200" {
		return model.Professor{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *ProfessorClient) UpdateProfessorByID(id string, professor model.ProfessorRequestBody) (model.Professor, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/professor/" + id

	// prepare request body
	byteData, err := json.Marshal(professor)
	if err != nil {
		return model.Professor{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	req, err := http.NewRequest(http.MethodPut, path, bodyReader)
	if err != nil {
		return model.Professor{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.Professor{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.ProfessorResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Professor{}, err
	}
	if resp.Code != "200" {
		return model.Professor{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *ProfessorClient) DeleteProfessorByID(id string) (model.Professor, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/professor/" + id

	// send request
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return model.Professor{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.Professor{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.ProfessorResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Professor{}, err
	}
	if resp.Code != "200" {
		return model.Professor{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func ProvideProfessorClientRest(client *http.Client) *ProfessorClient {
	return &ProfessorClient{client: client}
}
