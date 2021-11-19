package tool

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	pwd, err := LoadFile("private_config/db_pwd")
	checkErr(err)

	dataSourceName := fmt.Sprintf("root:%s@tcp(localhost:3306)/seekerDB?charset=utf8", pwd)
	db, err = sql.Open("mysql", dataSourceName)
	checkErr(err)
}

// StoreResultInMap 将query的结果存进map结构
func StoreResultInMap(rows *sql.Rows) (map[string][]string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve columns in rows")
	}
	elements := make([]interface{}, len(columns))
	for i, _ := range elements {
		var a interface{}
		elements[i] = &a
	}

	res := make(map[string][]string)
	for rows.Next() {
		err = rows.Scan(elements...)
		checkErr(err)

		values := make([]string, len(columns))
		for i, e := range elements {
			values[i] = fmt.Sprintf("%s", *e.(*interface{}))
		}
		if len(values) < 1 {
			return nil, errors.New("too few values in row")
		}
		res[values[0]] = values[1:]
		fmt.Println(values)
	}

	return res, nil
}

// Query 查询语句
func Query(key string, table string, extra string) (map[string][]string, error) {
	sentence := "SELECT " + key + " FROM " + table + extra
	rows, err := db.Query(sentence)
	checkErr(err)

	m, err := StoreResultInMap(rows)
	if err != nil {
		return nil, errors.Wrap(err, "fail to store result in map")
	}

	fmt.Println(m)
	return m, nil
}

// Insert 插入数据
func Insert(keys []string, values []interface{}, table string) error {
	for i, _ := range keys {
		keys[i] = keys[i] + "=?"
	}
	keysCombined := strings.Join(keys, ",")
	sentence := "Insert " + table + " SET " + keysCombined
	stmt, err := db.Prepare(sentence)
	if err != nil {
		return errors.Wrap(err, "fail to prepare sentence")
	}

	res, err := stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "fail to execute sentence")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "fail to retrieve last insert id")
	}

	fmt.Println(id)
	return nil
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func createUnfilledRow(keyNum int) string {
	unfilled := "("
	for i := 0; i < keyNum; i++ {
		unfilled += "?"
		if i != keyNum-1 {
			unfilled += ","
		}
	}
	unfilled += ")"
	return unfilled
}

func InsertOrUpdate(keys []string, valuesMap map[string][]string, table string) error {

	unfilled := createUnfilledRow(len(keys))
	for i, _ := range keys {
		keys[i] = keys[i] + "=VALUES(" + keys[i] + ")"
	}

	counter := 1
	percent := 0.1

	for token, strList := range valuesMap {

		counter += 1
		if float64(counter) >= float64(len(valuesMap))*percent {
			percent += 0.1
			fmt.Printf("progress: %d percent\n", int(percent*100))
		}

		filename := fmt.Sprintf("index_files/c_inverted_list_%s", token)
		sentence := "Insert INTO " + table + " VALUES" + unfilled + " ON DUPLICATE KEY UPDATE " + strings.Join(keys[1:], ",")
		//fmt.Printf("trying to execute statement: %s", sentence)
		stmt, err := db.Prepare(sentence)
		if err != nil {
			return errors.Wrap(err, "fail to prepare sentence")
		}

		values := []interface{}{token, strList[0], filename}
		//fmt.Printf("  With values: %v\n", values)

		_, err = stmt.Exec(values...)
		if err != nil {
			return errors.Wrap(err, "fail to execute sentence")
		}

		err = stmt.Close()
		if err != nil {
			return errors.Wrap(err, "fail to close statement")
		}

		err = WriteFile(filename, strList[1])
		if err != nil {
			return errors.Wrap(err, "fail to write index file")
		}

		//fmt.Printf("execution succeeded.\n")
	}
	return nil
}
