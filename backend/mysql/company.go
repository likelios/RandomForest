package mysql

import (
	"RandomForest/backend/models"
	"database/sql"
	"fmt"
	"log"
)

//CompanyQuery - SQL Запрос для получения данных для получения данных о компаниях
//Params:
//	%[1]v - Login
//	%[2]v - Password
const CompanyQuery = `select id, name, rating, description, email, phone, address from companies`

//GetCompany ...
func GetCompany(connect *sql.DB, id string) (rs []models.Compani) {
	query := CompanyQuery

	if len(id) > 0 {
		query = query + " where id = " + id
	}

	rows, err := connect.Query(fmt.Sprintf(query))
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.CompanyQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(query)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			r := models.Compani{}
			rows.Scan(
				&r.ID,
				&r.Name,
				&r.Rating,
				&r.Description,
				&r.Email,
				&r.Phone,
				&r.Address,
			)
			rs = append(rs, r)
		}
	}

	return
}
