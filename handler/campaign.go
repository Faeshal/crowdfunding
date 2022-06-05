package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// * api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {

	// karena c.query otomatis convert ke string , jadi mesti di kurung dgn method integer converted dulu
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("error get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("list campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// * api/v1/campaigns/:ID
func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed get detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	detailCampaign, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("failed get detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(detailCampaign)
	response := helper.APIResponse("detail campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// * api/v1/campaigns
func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("failed create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// * get current user id
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
