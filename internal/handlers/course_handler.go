package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"platform/internal/db"
	"platform/internal/models"
	"strconv"

	"github.com/gorilla/mux"
)

type Course_handler_interface interface {
	GetCurrentCourses(w http.ResponseWriter, r *http.Request)
	AddOneCourse(w http.ResponseWriter, r *http.Request)
	DeleteOneCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	JoinCourse(w http.ResponseWriter, r *http.Request)
}

type Course_handler struct {
	db *db.DB
}

func NewCourseHandler(db *db.DB) *Course_handler {
	h := &Course_handler{
		db: db,
	}
	return h
}

// GetUserCourses godoc
// @Summary Получить курсы
// @Description Получить список курсов
// @Tags Курсы
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Security BearerAuth
// @Failure 400 {string} string "Ошибка при получении курсов"
// @Router /courses [get]
func (h *Course_handler) GetCurrentCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.db.Courses.GetAllCourses(context.Background())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при получении курсов"))
		return
	}
	json.NewEncoder(w).Encode(courses)
}

type CourseBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatorId   string `json:"creatorId"`
}

// AddOneCourse godoc
// @Summary Добавить новый курс
// @Description Создать новый курс
// @Tags Курсы
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CourseBody true "Данные курса"
// @Success 201 {string} string "Курс создан"
// @Failure 400 {string} string "Ошибка при создании курса"
// @Router /courses [post]
func (h *Course_handler) AddOneCourse(w http.ResponseWriter, r *http.Request) {
	var cb CourseBody
	if err := json.NewDecoder(r.Body).Decode(&cb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный формат данных"))
		return
	}

	course := models.Course{
		Title:       cb.Title,
		Description: cb.Description,
		Creator_id: cb.CreatorId,
	}

	_, err := h.db.Courses.AddCourse(r.Context(), course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при создании курса"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Курс успешно создан"))
}

// DeleteOneCourse godoc
// @Summary Удалить курс
// @Description Удалить курс по ID
// @Tags Курсы
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID курса"
// @Success 200 {string} string "Курс удален"
// @Failure 400 {string} string "Ошибка при удалении курса"
// @Router /courses/{id} [delete]
func (h *Course_handler) DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный ID курса"))
		return
	}

	err = h.db.Courses.DeleteCourse(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при удалении курса"))
		return
	}

	w.Write([]byte("Курс успешно удален"))
}

type JoinCourseBody struct {
	UserId int `json:"userId"`
	CourseId int `json:"courseId"`
}
// JoinCourse godoc
// @Summary Присоединиться к курсу
// @Description Присоединиться к курсу
// @Tags Курсы
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body JoinCourseBody true "ID пользователя и ID курса"
// @Success 200 {string} string "Пользователь успешно добавлен в курс"
// @Failure 400 {string} string "Ошибка при добавлении пользователя в курс"
// @Router /courses/join [post]
func (c *Course_handler) JoinCourse(w http.ResponseWriter, r *http.Request) {
	var cb JoinCourseBody
	if err := json.NewDecoder(r.Body).Decode(&cb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный формат данных"))
		return
	}

	log.Println(cb.UserId, cb.CourseId)
	err := c.db.Courses.JoinCourse(r.Context(), cb.UserId, cb.CourseId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при добавлении пользователя в курс"))
		return
	}

}

// UpdateCourse godoc
// @Summary Обновить курс
// @Description Обновить данные курса
// @Tags Курсы
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID курса"
// @Param request body CourseBody true "Новые данные курса"
// @Success 200 {string} string "Курс обновлен"
// @Failure 400 {string} string "Ошибка при обновлении курса"
// @Router /courses/{id} [put]
func (h *Course_handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cb CourseBody
	if err := json.NewDecoder(r.Body).Decode(&cb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный формат данных"))
		return
	}

	course := models.Course{
		Id:          id,
		Title:       cb.Title,
		Description: cb.Description,
	}

	err := h.db.Courses.UpdateCourse(r.Context(), course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при обновлении курса"))
		return
	}

	w.Write([]byte("Курс успешно обновлен"))
}
