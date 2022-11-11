package handler

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	imageDir = "./" + "/images/"
)

type getAllUsers struct {
	Users []domain.User `json:"users"`
}

// Get all users
// @Summary      Get all users
// @Description  return all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  StatusResponce
// @Router       /users/ [get]
// @Security Authorization
func (h *Handler) GetAllUsers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	fmt.Println(userId)

	users := h.services.Users.GetAllUsers()
	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "all users",
		Data:    getAllUsers{Users: users},
	})
}

// @Summary      Get user by Name
// @Description  get user by Name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        name path string  true "User name"
// @Success      200  {object}  StatusResponce
// @Router       /users/user/{name} [get]
// @Security Authorization
func (h *Handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := h.services.Users.GetUserByName(name)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "one user",
		Data:    user,
	})
}

// @Summary      Get user by ID
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  StatusResponce
// @Router       /users/{id} [get]
// @Security Authorization
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

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "one user",
		Data:    user,
	})
}

type CreateUser struct {
	Name string
	Age  int
}

// @Summary CreateUser
// @Tags users
// @Description create new user
// @ID login
// @Accept  json
// @Produce  json
// @Param input body CreateUser true "credentials"
// @Success      200  {object}  StatusResponce
// @Router /users/ [post]
// @Security Authorization
func (h *Handler) CreateUser(c *gin.Context) {
	var user CreateUser

	err := c.BindJSON(&user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newUser, err := h.services.Users.CreateUser(domain.User{Name: user.Name, Age: user.Age})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "one user",
		Data:    newUser,
	})
}

// @Summary      Delete user by ID
// @Description  Delete user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  StatusResponce
// @Router       /users/{id} [delete]
// @Security Authorization
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

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: fmt.Sprintf("User was delete with id:%d", id),
	})
}

type UpdateUser struct {
	Name string
	Age  int
}

// @Summary 	 Update user
// @Tags         users
// @Description  update user
// @ID login
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param 	     input body UpdateUser true "credentials"
// @Success      200  {object}  StatusResponce
// @Router       /users/{id} [put]
// @Security Authorization
func (h *Handler) UpdateUser(c *gin.Context) {
	var user UpdateUser

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

	updateUser, err := h.services.Users.UpdateUser(domain.User{Id: id, Name: user.Name, Age: user.Age})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "updated user",
		Data:    updateUser,
	})
}

// @Summary 	 Upload file
// @Tags         upload
// @Description  upload other files
// @Accept       json
// @Produce      json
// @Param 	     file formData file true "this is a test file"
// @Success      200  {object}  StatusResponce
// @Router       /upload [post]
// @Security Authorization
func (h *Handler) UploadImage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	timeStamp := utils.TimeStamp()
	fileExtansion := filepath.Ext(file.Filename)
	fileNameHash := utils.GetMD5Hash(file.Filename + timeStamp)
	filepath := path.Join(imageDir + strconv.Itoa(userId) + "/" + fileNameHash + fileExtansion)

	saveToFile := "/images/" + strconv.Itoa(userId) + "/" + file.Filename
	_ = saveToFile

	h.services.UserImage.SaveImage(domain.Image{
		Filename: fileNameHash,
		Path:     filepath,
		UserId:   userId,
	})

	utils.CreateFolder(imageDir + strconv.Itoa(userId) + "/")

	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "images was updated",
		Data:    "http://" + c.Request.Host + "/images/" + strconv.Itoa(userId) + "/" + fileNameHash + fileExtansion,
	})
}

// @Summary 	SendEmail
// @Tags 		email
// @Description send email
// @Accept  	json
// @Produce  	json
// @Success     200  {object}  StatusResponce
// @Router /mail [get]
func (h *Handler) SendEmail(c *gin.Context) {
	// Sender data.
    from := os.Getenv("EMAIL")
    password := os.Getenv("PASSWORD")

    // Receiver email address.
    to := []string{
        "agr.akyla@mail.ru",
    }

    // smtp server configuration.
    smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

    // Message.
    message := []byte("This is a really unimaginative message, I know. Hello!")

    // Authentication.
    auth := smtp.PlainAuth("", from, password, smtpServer.host)

    // Sending email.
    err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
    if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

	c.JSON(http.StatusOK, StatusResponce{
		Status:  true,
		Message: "email send successful",
		Data: "email sended",
	})
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}