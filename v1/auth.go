package clam

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

type loginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthService struct {
	client       *gocloak.GoCloak
	realm        string
	clientId     string
	clientSecret string
}

func NewAuthService(realm string, clientId string, clientSecret string) *AuthService {
	return &AuthService{
		client:       gocloak.NewClient("https://sso.haitu.io"),
		realm:        realm,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (s *AuthService) Login(w http.ResponseWriter, r *http.Request) error {
	params := loginParams{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return errors.New("bad request")
	}

	_, err := s.client.Login(
		context.Background(),
		s.clientId,
		s.clientSecret,
		s.realm,
		params.Username,
		params.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return err
	}

	return nil
}
