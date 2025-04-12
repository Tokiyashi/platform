package handlers

import (
	"encoding/json"
	"net/http"
	"platform/internal/db"
	"platform/internal/models"
	"strconv"

	"github.com/gorilla/mux"
)

type section_handler_interface interface {
	GetSections(w http.ResponseWriter, r *http.Request)
	AddOneSection(w http.ResponseWriter, r *http.Request)
	DeleteOneSection(w http.ResponseWriter, r *http.Request)
	UpdateSection(w http.ResponseWriter, r *http.Request)
}

type Section_handler struct {
	db *db.DB
}

func NewSectionHandler(db *db.DB) *Section_handler {
	h := &Section_handler{
		db: db,
	}
	return h
}

type SectionBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseId    int    `json:"course_id"`
}

// GetSections godoc
// @Summary Получить секции курса
// @Description Получить все секции определенного курса
// @Tags Секции
// @Accept json
// @Produce json
// @Param course_id query int true "ID курса"
// @Success 200 {array} models.Section
// @Failure 400 {string} string "Ошибка при получении секций"
// @Security BearerAuth
// @Router /sections [get]
func (h *Section_handler) GetSections(w http.ResponseWriter, r *http.Request) {
	courseId := r.URL.Query().Get("course_id")

	sections, err := h.db.Sections.GetSections(courseId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при получении секций"))
		return
	}
	json.NewEncoder(w).Encode(sections)
}

// AddOneSection godoc
// @Summary Добавить секцию
// @Description Создать новую секцию в курсе
// @Tags Секции
// @Accept json
// @Produce json
// @Param request body SectionBody true "Данные секции"
// @Success 201 {string} string "Секция создана"
// @Failure 400 {string} string "Ошибка при создании секции"
// @Security BearerAuth
// @Router /sections [post]
func (h *Section_handler) AddOneSection(w http.ResponseWriter, r *http.Request) {
	var sb SectionBody
	if err := json.NewDecoder(r.Body).Decode(&sb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный формат данных"))
		return
	}

	err := h.db.Sections.AddSection(&models.Section{
		Title:       sb.Title,
		Description: sb.Description,
		CourseId:    sb.CourseId,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при создании секции"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Секция успешно создана"))
}

// DeleteOneSection godoc
// @Summary Удалить секцию
// @Description Удалить секцию по ID
// @Tags Секции
// @Accept json
// @Produce json
// @Param id path int true "ID секции"
// @Success 200 {string} string "Секция удалена"
// @Failure 400 {string} string "Ошибка при удалении секции"
// @Security BearerAuth
// @Router /sections/{id} [delete]
func (h *Section_handler) DeleteOneSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный ID секции"))
		return
	}

	err = h.db.Sections.DeleteOneSection(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при удалении секции"))
		return
	}
	w.Write([]byte("Секция успешно удалена"))
}

// UpdateSection godoc
// @Summary Обновить секцию
// @Description Обновить данные секции
// @Tags Секции
// @Accept json
// @Produce json
// @Param id path int true "ID секции"
// @Param request body SectionBody true "Новые данные секции"
// @Success 200 {string} string "Секция обновлена"
// @Failure 400 {string} string "Ошибка при обновлении секции"
// @Security BearerAuth
// @Router /sections/{id} [put]
func (h *Section_handler) UpdateSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный ID секции"))
		return
	}

	var sb SectionBody
	if err := json.NewDecoder(r.Body).Decode(&sb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверный формат данных"))
		return
	}

	err = h.db.Sections.UpdateSection(&models.Section{
		Id:          id,
		Title:       sb.Title,
		Description: sb.Description,
		CourseId:    sb.CourseId,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при обновлении секции"))
		return
	}

	w.Write([]byte("Секция успешно обновлена"))
}
