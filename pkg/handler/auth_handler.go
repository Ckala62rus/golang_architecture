package handler

import (
	"net/http"

	"github.com/Ckala62rus/go/domain"
	"github.com/gin-gonic/gin"
)

type CreateAuthUser struct {
	Email    string `example:"agr.akyla@mail.ru"`
	Password string `example:"123123"`
}

// signUp
// @Summary      Authentication in system
// @Description  return id created user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body CreateAuthUser true "credentials"
// @Success      200  {object}  statusResponce
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input CreateAuthUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status:  true,
		Message: "User success created",
		Data:    id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required" example:"agr.akyla@mail.ru"`
	Password string `json:"password" binding:"required" example:"123123"`
}

// @Summary SignIn
// @Tags auth
// @Description login and return authorization bearer token
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success	200  {object}  statusResponce
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	if err != nil {
		// newErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusOK, statusResponce{
			Status:  false,
			Message: "authentication failed",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"token": token,
	// })

	c.JSON(http.StatusOK, statusResponce{
		Status:  true,
		Message: "authentication success!",
		Data: map[string]interface{}{
			"token": "Bearer " + token,
		},
	})
}
