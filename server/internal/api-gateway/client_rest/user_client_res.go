package client_rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"capsmhoo/internal/api-gateway/model"

	"github.com/spf13/viper"
)

type UserClient struct {
	client *http.Client
}

type UserClientRest interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id string) (model.User, error)
	CreateUser(user model.UserRequestBody) (model.User, error)
	UpdateUserByID(id string, user model.UserRequestBody) (model.User, error)
	DeleteUserByID(id string) (model.User, error)
	Login(user model.UserRequestBody) (model.LoginResponseBody, error)
}

func (s *UserClient) GetAllUsers() ([]model.User, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/user"

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserListResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != "200" {
		return nil, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *UserClient) GetUserByID(id string) (model.User, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/user/" + id

	// send request
	response, err := s.client.Get(path)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Code != "200" {
		return model.User{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *UserClient) CreateUser(user model.UserRequestBody) (model.User, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/user"

	// prepare request body
	byteData, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Code != "200" {
		return model.User{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *UserClient) UpdateUserByID(id string, user model.UserRequestBody) (model.User, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/user/" + id

	// prepare request body
	byteData, err := json.Marshal(user)
	if err != nil {
		return model.User{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	req, err := http.NewRequest(http.MethodPut, path, bodyReader)
	if err != nil {
		return model.User{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Code != "200" {
		return model.User{}, errors.New(resp.Error)
	}

	return resp.Data, nil
}

func (s *UserClient) DeleteUserByID(id string) (model.User, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/user/" + id

	// send request
	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return model.User{}, err
	}
	response, err := s.client.Do(req)
	if err != nil {
		return model.User{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.UserResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.User{}, err
	}
	if resp.Code != "200" {
		return model.User{}, errors.New(resp.Error)
	}
	return resp.Data, nil
}

func (s *UserClient) Login(user model.UserRequestBody) (model.LoginResponseBody, error) {
	path := viper.GetString("user-service.host") + ":" + viper.GetString("user-service.port") + "/login"

	// prepare request body
	byteData, err := json.Marshal(user)
	if err != nil {
		return model.LoginResponseBody{}, err
	}
	bodyReader := bytes.NewReader(byteData)

	// send request
	response, err := s.client.Post(path, "application/json", bodyReader)
	if err != nil {
		return model.LoginResponseBody{}, err
	}
	defer response.Body.Close()

	// read response
	var resp model.LoginResponseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return model.LoginResponseBody{}, err
	}
	if resp.Error != "" {
		return model.LoginResponseBody{}, errors.New(resp.Error)
	}

	return resp, nil
}

func ProvideUserClientRest(client *http.Client) *UserClient {
	return &UserClient{client: client}
}
