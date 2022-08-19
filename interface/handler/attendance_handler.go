package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/eccSchoolApp-api/config"
	"github.com/yumekiti/eccSchoolApp-api/domain"
	"github.com/yumekiti/eccSchoolApp-api/usecase"
)

type AttendanceHandler interface {
	Get() echo.HandlerFunc
}

type attendanceHandler struct {
	attendanceUsecase usecase.AttendanceUsecase
}

func NewAttendanceHandler(attendanceUsecase usecase.AttendanceUsecase) AttendanceHandler {
	return &attendanceHandler{attendanceUsecase: attendanceUsecase}
}

type requestAttendance struct{}

type responseAttendance struct {
	Title    string `json:"title"`
	Rate     string `json:"rate"`
	Absence  string `json:"absence"`
	Lateness string `json:"lateness"`
}

func (h *attendanceHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := config.GetUser(c)
		getAttendance, err := h.attendanceUsecase.Get(
			&domain.User{
				Id:       user.Id,
				Password: user.Password,
			},
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseAttendance{}
		for _, attendance := range getAttendance {
			res = append(res, responseAttendance{
				Title:    attendance.Title,
				Rate:     attendance.Rate,
				Absence:  attendance.Absence,
				Lateness: attendance.Lateness,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}
