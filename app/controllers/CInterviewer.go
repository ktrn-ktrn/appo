package controllers

import (
	"appo/app/helpers"
	"appo/app/models/entities"
	"appo/app/models/interviewer"
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

type CInterviewer struct {
	*revel.Controller
	provider *interviewer.InterviewerProvider
}

func LoginInterviewer(userName string, password string) (*entities.User, error) {

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
		return nil, fmt.Errorf("LoginInterviewer:%v", err)
	}

	user := &entities.User{
		ID:       c_id.Int64,
		UserName: c_user_name.String,
		Password: c_password.String,
	}

	return user, nil
}

func (c *CInterviewer) Init() {
	c.provider = new(interviewer.InterviewerProvider)

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

		user, err = LoginInterviewer(userName, password)
		if err != nil {
			fmt.Print("\n\n\nОШИБКА:\n", err)
		}
	}

	if user == nil {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}

	err = c.provider.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CInterviewer) GetInterviewerByID() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	interviewer, err := c.provider.GetInterviewerById(interviewerId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(interviewer))
}

func (c *CInterviewer) GetAllInterviewers() revel.Result {
	c.Init()
	interviewers, err := c.provider.GetAllInterviewers()
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(interviewers)
}

func (c *CInterviewer) GetInterviewers() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}
	interviewers, err := c.provider.GetInterviewers(assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(interviewers)
}

func (c *CInterviewer) PutInterviewer() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newInterviewer entities.Interviewer
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	createdInterviewer, err := c.provider.PutInterviewer(&newInterviewer, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdInterviewer))
}

func (c *CInterviewer) SetInterviewer() revel.Result {
	c.Init()

	var newInterviewer entities.Interviewer
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	createdInterviewer, err := c.provider.SetInterviewer(&newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdInterviewer))
}

func (c *CInterviewer) PostInterviewer() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newInterviewer entities.Interviewer
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedInterviewer, err := c.provider.PostInterviewer(&newInterviewer, interviewerId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedInterviewer))
}

func (c *CInterviewer) DeleteInterviewerByID() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	erro := c.provider.DeleteInterviewer(interviewerId, assessmentId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}

func (c *CInterviewer) DeleteInterviewer() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	erro := c.provider.DeleteInterviewerFromD(interviewerId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}
