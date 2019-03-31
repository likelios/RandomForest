package mysql

import (
	"database/sql"
	"fmt"
	"RandomForest/backend/models"
	"log"
)

//LoginQuery - SQL Запрос для авторизации
//Params:
//	%[1]v - Login
//	%[2]v - Password
const LoginQuery = `select id, -1
						from users
						where login = '%[1]v'
						and password_hash = '%[2]v'`

//GetLogin ...
func GetLogin(connect *sql.DB, login, pass string) (models.LoginGoodResponse, models.LoginBadResponse) {

	rows, err := connect.Query(fmt.Sprintf(LoginQuery, login, pass))
	if err != nil {
		log.Panicln(fmt.Sprintf("Ошибка выполнения запроса query.LoginQuery:\n"+
			"%[1]v\n\n", fmt.Sprintf(LoginQuery, login, pass)), err)
	} else {
		defer rows.Close()
		for rows.Next() {
			r := models.LoginGoodResponse{}
			rows.Scan(&r.ID, &r.Status)
			return r, models.LoginBadResponse{}
		}
	}

	return models.LoginGoodResponse{ID: -1}, models.LoginBadResponse{Error: "Пользователь не найден", ID: -1}
}
