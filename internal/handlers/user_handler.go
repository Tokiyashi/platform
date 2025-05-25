package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"platform/internal/db"
	jwt_auth "platform/internal/jwt"

	"github.com/gorilla/mux"
)

type User_handler struct {
  db *db.DB
}

func NewUserHandler(db *db.DB) *User_handler {
  h := &User_handler{
    db: db,
  }
  return h
}

// GetUser godoc
// @Summary Получить данные пользователя
// @Description Получить данные пользователя по ID
// @Security BearerAuth
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Error"
// @Router /users/{id} [get]
func (h *User_handler) GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  userID := vars["id"]

  res, err := h.db.Users.GetUser(userID)
  if err != nil {
    w.Write([]byte("Ошибка"))
    return
  }

  json.NewEncoder(w).Encode(res)
}

type AuthBody struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

// Auth godoc
// @Summary Авторизовать пользователя
// @Description Логин с емаилом и паролем
// @Tags Авторизация
// @Accept json
// @Produce json
// @Param request body AuthBody true "Креды пользователя"
// @Success 200 {string} string "JWT Token"
// @Failure 404 {string} string "Пользователь не найден"
// @Failure 500 {string} string "Ошибка генерации токена"
// @Router /auth [post]
func (h *User_handler) Auth(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)

  if err != nil {
    return
  }

  var ab AuthBody
  err = json.Unmarshal([]byte(body), &ab)

  if err != nil {
    fmt.Print(err)
    return
  }

  user, err := h.db.Users.Auth(ab.Email, ab.Password)

  if err != nil {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Пользователь не найден"))
    return
  }

  jwt := &jwt_auth.JWTAuth{}

  fmt.Printf("User after Auth: %+v\n", user)

  token, err := jwt.GenerateToken(user)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Ошибка генерации токена"))
    return
  }

  w.Write([]byte(token))
}

type RegisterBody struct {
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Email     string `json:"email"`
  Password  string `json:"password"`
}

// AddOne godoc
// @Summary Зарегать нового пользователя
// @Description Создать нового пользователя
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param request body RegisterBody true "User registration data"
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Неправильно составлено тело запроса"
// @Failure 500 {string} string "Не удалось создать пользователя"
// @Router /users [post]
func (h *User_handler) AddOne(w http.ResponseWriter, r *http.Request) {
  body, err := io.ReadAll(r.Body)

  if err != nil {
    return
  }

  var rb RegisterBody
  err = json.Unmarshal([]byte(body), &rb)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Неправильно составлено тело запроса"))
  }

  err = h.db.Users.AddUser(rb.FirstName, rb.LastName, rb.Email, rb.Password)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Не удалось создать пользователя"))
  }
}

// GetUserCourses godoc
// @Summary Получить курсы пользователя
// @Description Получить список курсов, в которых состоит пользователь
// @Tags Пользователи
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} models.Course
// @Failure 400 {string} string "Ошибка получения курсов"
// @Failure 404 {string} string "Пользователь не найден"
// @Router /users/{id}/courses [get]
func (h *User_handler) GetUserCourses(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  userID := vars["id"]

  courses, err := h.db.Courses.GetUserCourses(userID)
  if err != nil {
    fmt.Println(err)
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Ошибка получения курсов"))
    return
  }
    if courses == nil {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Курсы не найдены"))
    return
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(courses)
}