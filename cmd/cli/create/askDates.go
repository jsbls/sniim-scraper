package create

import (
	"time"

	"github.com/everitosan/snimm-scrapper/internal/app/consult"
	"github.com/sirupsen/logrus"
)

func askDates(consult *consult.Consult) {
	var startDate, endDate string

	startDate, _ = getDatePrompt("Fecha de inicio")
	endDate, _ = getDatePrompt("Fecha de fin")

	start, _ := time.Parse("02/01/2006", startDate)
	end, _ := time.Parse("02/01/2006", endDate)

	if start.Unix() > end.Unix() {
		logrus.Fatal("Fecha de fin no puede ser mayor a fecha de inicio")
	}

	consult.AddParameter("fechainicio", startDate)
	consult.AddParameter("fechafinal", endDate)
}
