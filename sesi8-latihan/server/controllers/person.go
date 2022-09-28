package controllers

import (
	"fmt"
	"net/http"
	"sesi8-latihan/server/models"
	_ "sesi8-latihan/server/params"
	"sesi8-latihan/server/repositories"
	"sesi8-latihan/server/views"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	repo repositories.PersonRepo
}

func NewPersonController(repo repositories.PersonRepo) *PersonController {
	return &PersonController{
		repo: repo,
	}
}

// GetPeople godoc
// @Summary    Get all people
// @Decription Get list people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Router     /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {

	person, err := p.repo.GetPeople()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: person,
	})

}

// Create People godoc
// @Summary    Get people
// @Decription Get people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Param      id path int true "People ID"
// @Router     /people_detail/{id} [get]
func (p *PersonController) GetPeopleById(ctx *gin.Context) {
	getId := ctx.Params.ByName("Id")
	fmt.Println("----------------------")
	fmt.Println(getId)
	id, err := strconv.Atoi(getId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	person, err := p.repo.GetPeopleById(id)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: person,
	})

}

// Create People godoc
// @Summary    Create people
// @Decription Create people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Param People body params.Person true "Create People"
// @Router     /people [post]
func (p *PersonController) CreatePeople(ctx *gin.Context) {
	var req models.Person
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(req)
	errs := p.repo.CreatePeople(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "CREATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "CREATE_PEOPLE_SUCCESS",
		Payload: req,
	})
}

// Create People godoc
// @Summary    Delete people
// @Decription Delete people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Param      id path int true "People ID"
// @Router     /people/{id} [delete]
func (p *PersonController) DeletePeople(ctx *gin.Context) {
	getId := ctx.Params.ByName("Id")
	fmt.Println("----------------------")
	fmt.Println(getId)
	id, err := strconv.Atoi(getId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	er := p.repo.DeletePeople(id)
	if er != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_PEOPLE_SUCCESS",
		Payload: id,
	})
}

// Create People godoc
// @Summary    Update people
// @Decription Update people
// @Tags       person
// @Accept     json
// @Produce    json
// @Success    200 {object} views.GetAllPeopleSwagger
// @Param      id path int true "People ID"
// @Param People body params.Person true "Update People"
// @Router     /people/{id} [patch]
func (p *PersonController) UpdatePeople(ctx *gin.Context) {
	getId := ctx.Params.ByName("Id")
	var req models.Person
	id, err := strconv.Atoi(getId)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	er := ctx.ShouldBindJSON(&req)
	if er != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	errs := p.repo.UpdatePeople(id, &req)
	if errs != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PEOPLE_SUCCESS",
		Payload: req,
	})
}
