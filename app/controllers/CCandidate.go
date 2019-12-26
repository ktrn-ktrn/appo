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

// проверка существованя пользователя с введенными именем и паролем
func LoginCandidate(userName string, password string) (*entities.User, error) {

	//подключение к БД
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

	//запрос к базе данных
	//выбрать пользователя с заданными именем и паролем
	selectQuery := `SELECT c_id, c_user_name, c_password FROM t_user WHERE c_user_name = $1 AND
	c_password = $2`
	row := db.QueryRow(selectQuery, userName, password)

	err := row.Scan(&c_id, &c_user_name, &c_password)

	//если нет такого пользователя, возвращаем nil, если есть ошибка - озвращаем ошибку
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

	//запрашиваем заголовок
	authorization := c.Request.Header.Get("Authorization")

	// если запрошенный заголовок не равен "Authorization", то возвращаем запрос
	// аутентификации через заголовок WWW-Authentificate:
	if authorization == "" {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}

	// получаем закодированные имя пользователя и пароль
	// убираем подстроку "Basic " и декодируем
	loginAndPassB64 := strings.TrimLeft(authorization, "Basic ")
	bLoginAndPass, err := base64.StdEncoding.DecodeString(loginAndPassB64)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR decode base64: %v", err))
		return
	}

	// конвертируем в string
	loginAndPass := string(bLoginAndPass)
	var user *entities.User = nil

	// если строка не пустая то разделяем её по символу ':' на имя пользователя и пароль...
	if len(loginAndPass) != 0 {
		loginAndPassSplited := strings.Split(loginAndPass, ":")

		userName := loginAndPassSplited[0]
		password := loginAndPassSplited[1]
		var err error

		// ... и вызываем функцию LoginAssessment с полученными данными в качестве параметров
		user, err = LoginCandidate(userName, password)
		if err != nil {
			fmt.Print("\n\n\nОШИБКА:\n", err)
		}
	}

	// если имя пользователя неправильные (то есть функция LoginAssessment вернула nil),
	//то возвращаем запрос аутентификации через заголовок WWW-Authentificate
	if user == nil {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}

	err = c.provider.Init()
	if err != nil {
		log.Fatal(err)
	}
}

//получить всех кандидатов, состоящих в выбранном ассессменте
func (c *CCandidate) GetCandidates() revel.Result {
	c.Init()

	//получаем ID выбранного ассессмента и конвертируем в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	//вызываем метод GetCandidates провайдера
	candidates, err := c.provider.GetCandidates(assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(candidates)
}

//получить список всех кандидатов, которые есть или будут
func (c *CCandidate) GetAllCandidates() revel.Result {
	c.Init()
	candidates, err := c.provider.GetAllCandidates()
	if err != nil {
		return c.RenderJSON(err)
	}

	return c.RenderJSON(candidates)
}

//получить выбранного кандидата
func (c *CCandidate) GetCandidateByID() revel.Result {
	c.Init()

	//получаем ID выбранного кандидата и конвертируем в int
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	//получаем ID выбранного ассессмента и конвертируем в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	//вызываем метод GetCandidateById провайдера
	candidate, err := c.provider.GetCandidateById(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(helpers.Success(candidate))
}

//получить возможные статусы кандидата
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

//задать статус выбранному кандидату
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

	//получаем значение статуса из фонта
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	//анмаршалим
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

//добавляем кандидата в выбранный ассессмент
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

//изменение кандидата
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

//удаление кандидата
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
