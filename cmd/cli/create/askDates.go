package create

import (
	"fmt"
	"time"

	"github.com/everitosan/snimm-scrapper/internal/app/consult"
	"github.com/sirupsen/logrus"
)

func askDates(consultR *consult.Consult) {
	var startDate, endDate string

	startDate, _ = getDatePrompt(fmt.Sprintf("Fecha de inicio (dd/MM/AAAA | %s)", consult.Now))

	if startDate != consult.Now {
		endDate, _ = getDatePrompt("Fecha de fin (dd/MM/AAAA)")
		start, _ := time.Parse("02/01/2006", startDate)
		end, _ := time.Parse("02/01/2006", endDate)

		if start.Unix() > end.Unix() {
			logrus.Fatal("Fecha de fin no puede ser mayor a fecha de inicio")
		}
	} else {
		endDate = consult.Now
	}

	consultR.AddParameter("fechainicio", startDate)
	consultR.AddTextParameter("Desde", startDate)

	consultR.AddParameter("fechafinal", endDate)
	consultR.AddTextParameter("Hasta", endDate)
}
