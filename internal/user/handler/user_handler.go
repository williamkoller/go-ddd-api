package handler

import (
	"github.com/williamkoller/go-ddd-api/internal/auth"
	"github.com/williamkoller/go-ddd-api/internal/user/domain"
	"github.com/williamkoller/go-ddd-api/internal/user/usecase"
	"github.com/williamkoller/go-ddd-api/pkg/response"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc}
}

// Register godoc
// @Summary Criar usuário
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.User true "Usuário a ser registrado"
// @Success 201 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.uc.Register(&user); err != nil {
		log.Printf("[Register] Error %v", err.Error())
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") &&
			strings.Contains(err.Error(), "users_email_key") {
			response.Error(c, http.StatusConflict, "email already in use")
			return
		}
		response.Error(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	response.Created(c, user)
}

// Login godoc
// @Summary Login do usuário
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credenciais de login"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.uc.Login(loginData.Email, loginData.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "could not generate token")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// List godoc
// @Summary Listar todos os usuários
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) List(c *gin.Context) {
	users, err := h.uc.List()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, users)
}

// Update godoc
// @Summary Atualizar usuário
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param user body domain.User true "Dados atualizados"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.uc.Update(uint(id), &user); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "updated successfully"})
}

// Delete godoc
// @Summary Deletar usuário
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.uc.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"message": "deleted successfully"})
}
