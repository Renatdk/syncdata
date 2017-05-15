package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	l "scsyncs/synccommoditymarkets/services/logging"

	_ "github.com/denisenkom/go-mssqldb"
)

var connation *sql.DB

func initConnection() *sql.DB {

	l.InitLogging()

	file, e := ioutil.ReadFile("./configs/database.json")
	checkErr(e)

	var jsontype Connection
	json.Unmarshal(file, &jsontype)

	connString := fmt.Sprintf("server=%s;user id=%s; database=%s; password=%s;port=%d", jsontype.Server, jsontype.User, jsontype.Database, jsontype.Password, jsontype.Port)
	if jsontype.Debag {
		fmt.Println(connString)
	}

	conn, e := sql.Open("mssql", connString)
	checkErr(e)

	//defer conn.Close()
	return conn
}

func GetConnection() *sql.DB {

	if connation == nil {
		connation = initConnection()
	}
	return connation
}

func IsHaveData(code, date string) bool {

	conn := GetConnection()

	query := "SELECT COUNT(*) as count FROM teoview.oper_rep.commodity_markets where Code='" + code + "' and Date = '" + date + "'"

	rows, err := conn.Query(query)
	checkErr(err)

	result := false

	for rows.Next() {
		var count string
		err = rows.Scan(&count)

		checkErr(err)

		fmt.Println(count)

		intCount, err := strconv.Atoi(count)
		checkErr(err)

		if intCount > 0 {
			result = true
		}
	}

	return result
}

func AddData(code string, date, value json.Number) {
	conn := GetConnection()
	// insert
	//stmt, err := conn.Prepare("INSERT INTO [teoview].[oper_rep].[commodity_markets]([Code],[Value],[Date])VALUES(?, ?,CAST('?' AS DATETIME2))")
	stmt, err := conn.Prepare(`INSERT INTO [teoview].[oper_rep].[commodity_markets]
           ([Code]
           ,[Value]
           ,[Date])
     VALUES
           (?
           ,?
           ,?)`)
	checkErr(err)
	res, err := stmt.Exec(code, string(value), string(date))
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Добавлено")
	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		l.Error.Println(err)
	}
}
