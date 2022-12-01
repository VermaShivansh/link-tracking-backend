package controllers

import (
	"fmt"
	"net/http"

	"github.com/ShivanshVerma-coder/link-tracking/db"
	"github.com/ShivanshVerma-coder/link-tracking/helpers"
	"github.com/ShivanshVerma-coder/link-tracking/models"
	"github.com/ShivanshVerma-coder/link-tracking/repository"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type LinksController struct {
	Generate       gin.HandlerFunc
	GetTargetLink  gin.HandlerFunc
	GetInfo        gin.HandlerFunc
	UpdateSettings gin.HandlerFunc
}

func NewLinksController() *LinksController {
	linksController := &LinksController{}
	linksController.Generate = generate
	linksController.GetTargetLink = getTargetLink
	linksController.GetInfo = getInfo
	linksController.UpdateSettings = updateSettings

	return linksController
}

func generate(c *gin.Context) {

	type requestBody struct {
		Target_url string `json:"target_url" binding:"required,startswith=https://"`
		Tag_name   string `json:"tag_name" binding:"required"`
	}

	body := requestBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		helpers.SendResponse(c, http.StatusUnprocessableEntity, "Error in request body", err.Error())
		return
	}

	helpers.PrettyPrint(body)

	shortened_url_id, err := gonanoid.New(8)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Error in generating link", nil)
		return
	}

	generatedLink, err := repository.CreateLink(shortened_url_id, body.Target_url, body.Tag_name)

	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Unable to add link in database", nil)
		return
	}

	helpers.SendResponse(c, http.StatusCreated, "Link generated successfully", map[string]interface{}{"generatedLink": generatedLink})
}

func getTargetLink(c *gin.Context) {
	id, _ := c.Params.Get("id")
	ip := c.ClientIP()
	record, _ := db.IPClient.Get_all(ip)

	linkUnit, err := repository.GetTargetLink(id)
	if err != nil {
		fmt.Println("Error in finding target url in database")
		helpers.SendResponse(c, http.StatusBadRequest, "Error in finding target url in database", nil)
		return
	}

	//check all restrictions
	authorized, err := helpers.SBAC(linkUnit, ip, record)
	if err != nil {
		helpers.SendResponse(c, http.StatusInternalServerError, "Unable to check access to link", nil)
		return
	}
	if !authorized {
		helpers.SendResponse(c, http.StatusUnauthorized, "Sorry you can not access this link", nil)
		return
	}
	fmt.Println("Authorized. Sending Response...")

	// helpers.SendResponse(c, http.StatusOK, "Target link found", map[string]interface{}{"target_url": linkUnit.Target_url})
	c.Redirect(http.StatusTemporaryRedirect, linkUnit.Target_url)

	// update analytics in go routine
	go func() {
		repository.UpdateAnalytics(linkUnit, record.Country_short)
	}()
}

func getInfo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	linkUnit, err := repository.GetCompleteInfo(id)
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Error in fetch information related to link", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Successfully fetched data related to link", linkUnit)
}

func updateSettings(c *gin.Context) {
	id, _ := c.Params.Get("id")
	newLinkUnit := models.LinkUnit{}

	if err := c.ShouldBindJSON(&newLinkUnit); err != nil {
		helpers.SendResponse(c, http.StatusUnprocessableEntity, "Check request body", nil)
		return
	}

	fmt.Println(newLinkUnit)
	err := repository.UpdateSettings(id, newLinkUnit)
	if err != nil {
		helpers.SendResponse(c, http.StatusBadRequest, "Error in fetch information related to link", nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, "Settings updated for link", nil)
}
