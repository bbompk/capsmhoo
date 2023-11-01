package client_rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"capsmhoo/internal/api-gateway/model"

	"github.com/spf13/viper"
)

type StudentClient struct {
	client *http.Client
}

type StudentClientRest interface {
	GetAllStudents() ([]model.Student, error)
	GetStudentByID(id string) (model.Student, error)
	CreateStudent(student model.StudentRequestBody) (model.Student, error)
	UpdateStudentByID(id string, student model.StudentRequestBody) (model.Student, error)
	DeleteStudentByID(id string) (model.Student, error)
}

func (s *StudentClient) GetAllStudents() ([]model.Student, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/student"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp model.StudentListResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != "200" {
		return nil, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *StudentClient) GetStudentByID(id string) (model.Student, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/student/" + id

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.Student{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.StudentResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Student{}, err
	}
	if resp.Code != "200" {
		return model.Student{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *StudentClient) CreateStudent(student model.StudentRequestBody) (model.Student, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/student"

	// prepare request body
	byteData, err := json.Marshal(student)
	if err != nil {
		return model.Student{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.Student{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.StudentResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Student{}, err
	}
	if resp.Code != "200" {
		return model.Student{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *StudentClient) UpdateStudentByID(id string, student model.StudentRequestBody) (model.Student, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/student/" + id

	// prepare request body
	byteData, err := json.Marshal(student)
	if err != nil {
		return model.Student{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	req, err := http.NewRequest(http.MethodPut, path, bodyReader)
	if err != nil {
		return model.Student{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.Student{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.StudentResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Student{}, err
	}
	if resp.Code != "200" {
		return model.Student{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *StudentClient) DeleteStudentByID(id string) (model.Student, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/student/" + id

	// send request
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return model.Student{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.Student{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.StudentResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.Student{}, err
	}
	if resp.Code != "200" {
		return model.Student{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func ProvideStudentClientRest(client *http.Client) *StudentClient {
	return &StudentClient{client: client}
}
