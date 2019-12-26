package auth

import (
	"appo/app/models/entities"
	"appo/app/models/mappers"
	"database/sql"

	"fmt"

	//"bytes"

	//"io/ioutil"
	//"log"
	//"strings"

	_ "github.com/lib/pq"
)

type AuthProvider struct {
	db   *sql.DB
	auth *mappers.AuthMapper
}

func (p *AuthProvider) Init() error {

	p.auth = new(mappers.AuthMapper)
	p.auth.Init(p.db)
	return nil
}

/*
func digestPost(host string, uri string, postBody []byte) bool {
	url := host + uri
	method := "POST"
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		log.Printf("Recieved status code '%v' auth skipped", resp.StatusCode)
		return true
	}
	digestParts := digestParts(resp)
	digestParts["uri"] = uri
	digestParts["method"] = method
	digestParts["username"] = "username"
	digestParts["password"] = "password"
	req, err = http.NewRequest(method, url, bytes.NewBuffer(postBody))
	req.Header.Set("WWW-Authenticate", "Basic realm=\"My Realm\"")
	req.Header.Set("Content-Type", "application/json")


	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println("response body: ", string(body))
		return false
	}
	return true
}
*/

func (p *AuthProvider) Login(userName string) (*entities.User, error) {

	user, err := p.auth.Login(userName)
	if err != nil {
		fmt.Printf("AuthProvider::Login:%v", err)
		return nil, err
	}

	return user, nil
}

func (p *AuthProvider) Logout() error {
	defer p.db.Close()

	return nil
}
