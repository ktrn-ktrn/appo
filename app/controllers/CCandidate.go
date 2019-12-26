package controllers

import (
	"appo/app/helpers"
	"appo/app/models/candidate"
	"appo/app/models/entities"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"io/ioutil"

	"log"
	"strconv"

	"github.com/revel/revel"
)

type CCandidate struct {
	*revel.Controller
	provider *candidate.CandidateProvider
}

func LoginCandidate(userName string, password string) (*entities.User, error) {

	var erro error
	connStr := "user=postgres password=password port=5433 dbname=AssessmentManager sslmode=disable"
	db, erro := sql.Open("postgres", connStr)
	if erro != nil {
		return nil, erro
	}

	var (
		c_id        sql.NullInt64
		c_user_name sql.NullString
		c_password  sql.NullString
	)

	selectQuery := `SELECT c_id, c_user_name, c_password FROM t_user WHERE c_user_name = $1 AND
	c_password = $2`
	row := db.QueryRow(selectQuery, userName, password)

	err := row.Scan(&c_id, &c_user_name, &c_password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("LoginCandidate:%v", err)
	}

	user := &entities.User{
		ID:       c_id.Int64,
		UserName: c_user_name.String,
		Password: c_password.String,
	}

	return user, nil
}

func (c *CCandidate) Init() {
	c.provider = new(candidate.CandidateProvider)

	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}

	loginAndPassB64 := strings.TrimLeft(authorization, "Basic ")
	bLoginAndPass, err := base64.StdEncoding.DecodeString(loginAndPassB64)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR decode base64: %v", err))
		return
	}
	loginAndPass := string(bLoginAndPass)
	var user *entities.User = nil

	if len(loginAndPass) != 0 {
		loginAndPassSplited := strings.Split(loginAndPass, ":")

		userName := loginAndPassSplited[0]
		password := loginAndPassSplited[1]
		var err error

		user, err = LoginCandidate(userName, password)
		if err != nil {
			fmt.Print("\n\n\nОШИБКА:\n", err)
		}
	}

	// ЕСЛИ ЛОГИНА И ПАРОЛЯ НЕТ (ИЛИ ОНИ НЕПРАВИЛЬНЫЕ), ТО ВОЗВРАЩАЕМ ЗАПРОС
	// АУТЕНТИФИКАЦИИ (ЧЕРЕЗ ЗАГОЛОВОК WWW-Authentificate)
	if user == nil {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}

	err = c.provider.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CCandidate) GetCandidates() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}
	candidates, err := c.provider.GetCandidates(assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(candidates)
}

func (c *CCandidate) GetAllCandidates() revel.Result {
	c.Init()
	candidates, err := c.provider.GetAllCandidates()
	if err != nil {
		return c.RenderJSON(err)
	}

	return c.RenderJSON(candidates)
}

func (c *CCandidate) GetCandidateByID() revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	candidate, err := c.provider.GetCandidateById(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(helpers.Success(candidate))
}

func (c *CCandidate) GetCandidateStatus() revel.Result {
	c.Init()

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	sCandidatetId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidatetId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	candidate, err := c.provider.GetCandidateStatus(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(candidate))
}

func (c *CCandidate) SetStatus() revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sStatusId := c.Params.Get("statusID")
	statusId, err := strconv.ParseInt(sStatusId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	var newStatus entities.CandidateStatus
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newStatus)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedStatus, err := c.provider.SetCandidateStatus(&newStatus, statusId, candidateId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedStatus))
}

func (c *CCandidate) PutCandidate() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newCandidate entities.Candidate
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newCandidate)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	createdCandidate, err := c.provider.PutCandidate(&newCandidate, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdCandidate))
}

func (c *CCandidate) PostCandidateByID() revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newCandidate entities.Candidate
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newCandidate)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedCandidate, err := c.provider.PostCandidate(&newCandidate, candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedCandidate))
}

func (c *CCandidate) DeleteCandidateByID() revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Printf("ID Assessment from CCandidate:", assessmentId, ", ", candidateId)
	erro := c.provider.DeleteCandidate(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(erro)
	}
	return nil
}
