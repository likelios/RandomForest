package mysql

import (
	"database/sql"
	"fmt"
	"RandomForest/backend/models"
	"log"
)

//ClientQuery - SQL Запрос для получения данных для получения данных о клиенте
//Params:
//	%[1]v - id
const ClientQuery = `select id,
							type,
							date_start,
							date_end,
							sum,
							is_accedents,
							company_id,
							name,
							contacts,
							areas
						from clients`

//GetClient ...
func GetClient(connect *sql.DB, id string) (rs []models.NiceClient) {
	query := ClientQuery
	rr := []models.Client{}

	if len(id) > 0 {
		query = query + " where id = " + id
	}

	rows, err := connect.Query(fmt.Sprintf(query))
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.ClientQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(query)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			r := models.Client{}
			rows.Scan(
				&r.ID,
				&r.Type,
				&r.DateStart,
				&r.DateEnd,
				&r.Sum,
				&r.IsAccedents,
				&r.CompanyID,
				&r.Name,
				&r.Contacts,
				&r.Areas,
			)
			rr = append(rr, r)
		}
	}
	rs = prepareNiceClient(rr)

	return
}

//GetClientByCompany ...
func GetClientByCompany(connect *sql.DB, id, condition string) (rs []models.NiceClient) {
	query := ClientQuery
	rr := []models.Client{}

	if len(id) > 0 {
		query = query + " where company_id " + condition + id
	}

	rows, err := connect.Query(fmt.Sprintf(query))
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.ClientQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(query)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			r := models.Client{}
			rows.Scan(
				&r.ID,
				&r.Type,
				&r.DateStart,
				&r.DateEnd,
				&r.Sum,
				&r.IsAccedents,
				&r.CompanyID,
				&r.Name,
				&r.Contacts,
				&r.Areas,
			)
			rr = append(rr, r)
		}
	}
	rs = prepareNiceClient(rr)

	return
}

func prepareNiceClient(ob []models.Client) (rs []models.NiceClient) {
	for _, e := range ob {
		rs = append(rs, models.NiceClient{
			e.ID,
			e.Type,
			e.DateStart.Format("02.01.06 15:04:05"),
			e.DateEnd.Format("02.01.06 15:04:05"),
			e.Sum,
			e.IsAccedents,
			e.CompanyID,
			e.Name,
			e.Contacts,
			e.Areas,
		})
	}

	return
}
