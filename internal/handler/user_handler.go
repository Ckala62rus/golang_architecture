package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ckala62rus/go/domain"
	"github.com/gin-gonic/gin"
)

type getAllUsers struct {
	Users []domain.User `json:"users"`
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users := h.services.Users.GetAllUsers()
	c.JSON(http.StatusOK, getAllUsers{
		Users: users,
	})
}

func (h *Handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := h.services.Users.GetUserByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	user, err := h.services.Users.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newUser, err := h.services.Users.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func (h *Handler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	isDelete, err := h.services.Users.DeleteUserById(id)
	if !isDelete {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: true,
		Message: fmt.Sprintf("User was delete with id:%d", id),
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user domain.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	err = c.BindJSON(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Id = id

	updateUser, err := h.services.Users.UpdateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateUser)
}
