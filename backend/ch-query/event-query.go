package query

import (
	"database/sql"
	"fmt"
	"RandomForest/backend/models"
	"log"
	"time"
)

//EventQuery - SQL Запрос для получения данных для отчета Event
//Params:
//	%[1]v - Type
//	%[2]v - CustomerID
const EventQuery = `select groupArray(Date),
						   groupArray(%[1]v)
						from hack.event
						where CustomerID = %[2]v`

const EventTableQuery = `select DeviceID, concat(CAST(CoordX AS String), ',', CAST(CoordY AS String)), Temp, Humidity, WindSpeed, Date
					 from hack.event
					 where CustomerID = %[1]v`

//GetEventRows ...
func GetEventRows(connect *sql.DB, typeName, customer string) (rs models.EventDataNice) {

	rows, err := connect.Query(fmt.Sprintf(EventQuery, typeName, customer))
	r := models.EventData{}
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.EventQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(EventQuery, typeName, customer)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&r.Date, &r.Value)
			log.Println("test ", r)
		}
	}

	rs.Value = r.Value
	rs.Date = parseDate(r.Date)

	return
}

//GetEventRows ...
func GetEventTable(connect *sql.DB, customer string) (rs []models.EventNiceGrid) {

	rows, err := connect.Query(fmt.Sprintf(EventTableQuery, customer))
	rr := []models.EventGrid{}
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.EventTableQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(EventTableQuery, customer)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			r := models.EventGrid{}

			rows.Scan(
				&r.DeviceID,
				&r.Coord,
				&r.Temp,
				&r.Humidity,
				&r.WindSpeed,
				&r.Date,
			)

			rr = append(rr, r)
		}
	}

	rs = parseGridDate(rr)

	return
}

func parseDate(d []time.Time) (s []string) {
	for _, e := range d {
		s = append(s, e.Format("02.01.06 15:04:05"))
	}
	return
}

func parseGridDate(ob []models.EventGrid) (rs []models.EventNiceGrid) {
	for _, e := range ob {
		rs = append(rs, models.EventNiceGrid{
			DeviceID:  e.DeviceID,
			Coord:     e.Coord,
			Temp:      e.Temp,
			Humidity:  e.Humidity,
			WindSpeed: e.WindSpeed,
			Date:      e.Date.Format("02.01.06 15:04:05"),
		})
	}

	return
}
