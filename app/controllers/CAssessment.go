package controllers

import (
	"appo/app/helpers"
	"appo/app/models/assessment"
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

type CAssessment struct {
	*revel.Controller
	provider *assessment.AssessmentProvider
	db       *sql.DB
}

func LoginAssessment(userName string, password string) (*entities.User, error) {

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
		return nil, fmt.Errorf("LoginAssessment:%v", err)
	}

	user := &entities.User{
		ID:       c_id.Int64,
		UserName: c_user_name.String,
		Password: c_password.String,
	}

	return user, nil
}

func (c *CAssessment) Init() {
	c.provider = new(assessment.AssessmentProvider)
	/*
		c.Request.Header.Get("token")
		if !authentication.Get(c.Request.Header.Get("token")) {
			c.Redirect("/login")
		}
	*/
	// ПРОВЕРКА АУТЕНТИФИКАЦИИ
	authorization := c.Request.Header.Get("Authorization")

	if authorization == "" {
		fmt.Print("ЗДЕСЬ")
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

		user, err = LoginAssessment(userName, password)
		if err != nil {
			fmt.Print("\n\n\nОШИБКА:\n", err)
		}
	}

	// ЕСЛИ ЛОГИНА И ПАРОЛЯ НЕТ (ИЛИ ОНИ НЕПРАВИЛЬНЫЕ), ТО ВОЗВРАЩАЕМ ЗАПРОС
	// АУТЕНТИФИКАЦИИ (ЧЕРЕЗ ЗАГОЛОВОК WWW-Authentificate)
	if user == nil {
		fmt.Println("\n\n\nUSER = NIL")
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}
	err = c.provider.Init()
	if err != nil {
		log.Fatal(err)
	}
}
func (c *CAssessment) GetAssessmentByID() revel.Result {
	c.Init()

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	assessment, err := c.provider.GetAssessmentById(assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(assessment))
}

func (c *CAssessment) GetStatus() revel.Result {
	c.Init()

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	assessment, err := c.provider.GetAssessmentStatus(assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(assessment))
}

func (c *CAssessment) SetStatus() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sStatusId := c.Params.Get("statusID")
	statusId, err := strconv.ParseInt(sStatusId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	var newStatus entities.AssessmentStatus
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newStatus)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedStatus, err := c.provider.SetAssessmentStatus(&newStatus, statusId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedStatus))
}

func (c *CAssessment) GetAssessments() revel.Result {
	c.Init()
	assessments, err := c.provider.GetAssessments()
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(assessments))
}

func (c *CAssessment) PutAssessment() revel.Result {
	c.Init()
	var newAssessment entities.Assessment
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newAssessment)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	createdAssessment, err := c.provider.PutAssessment(&newAssessment)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdAssessment))
}

func (c *CAssessment) PostAssessmentByID() revel.Result {

	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	var newAssessment entities.Assessment
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newAssessment)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedAssessment, err := c.provider.PostAssessment(&newAssessment, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedAssessment))
}

func (c *CAssessment) DeleteAssessmentByID() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	//fmt.Printf("String Assessment from CAssessment:", sAssessmentId)
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	//fmt.Printf("ID Assessment from CAssessment:", assessmentId)
	erro := c.provider.DeleteAssessment(assessmentId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}
