package service

import (
	"PCS-API/models"
	"PCS-API/repository"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// TODO: Mettre un champ "name" dans property notamment pour ce cas ci-dessous
func billGenerateContent(property models.Property, reservation models.Reservation) string {
	propertyName := repository.LessorGetById(property.LessorId)
	content := fmt.Sprintf(
		"%dx - %s de %s %s\n",
		int(reservation.EndDate.Sub(reservation.BeginDate).Hours()/24),
		property.Name,
		propertyName.FirstName,
		propertyName.LastName)
	return content
}

// TODO: Une bonne pratique de code serait de généraliser les pointeurs lors des appels de fonction pour éviter de faire
// bêtement une copie
func billGeneratePrice(property *models.Property, reservation *models.Reservation) float64 {
	var price float64
	price += property.Price * (reservation.EndDate.Sub(reservation.BeginDate).Hours() / 24)
	return price
}

func billCreate(property models.Property, reservation models.Reservation) (models.Bill, error) {
	var bill models.Bill
	bill.ID = uuid.New()
	bill.Date = time.Now()
	bill.Statut = "success"
	bill.Content = billGenerateContent(property, reservation)
	bill.Price = billGeneratePrice(&property, &reservation)

	bill, err := repository.BillCreate(bill, reservation.ID)

	return bill, err
}
