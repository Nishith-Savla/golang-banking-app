package domain

import (
	"encoding/json"
	"fmt"
	"github.com/Nishith-Savla/golang-banking-app/logger"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	response, err := http.Get(u)
	if err != nil {
		logger.Error("Error while sending verification request: " + err.Error())
		return false
	}

	m := map[string]bool{}
	if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
		logger.Error("Error while decoding response from auth server: " + err.Error())
		return false
	}
	return m["isAuthorized"]
}

/*
  This will generate an url for token verification in the below format

  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}

  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVER_ADDRESS"), os.Getenv("AUTH_SERVER_PORT")),
		Path:   "/auth/verify",
	}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
