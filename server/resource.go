package server

import (
	"echo/Config"
	"echo/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var CONFIG = Config.LoadConfig()
var dbHandler = new(db.Handler)
var s = dbHandler.Init()

type Resource struct{}

type DBHandler struct {
	DBHandler db.Handler
}

//leehs 20220517 주어진 경로의 모든 데이터 파일들의 리스트를 반환
func getFiles(path string) []string {
	var dfile []string
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if strings.Contains(f.Name(), "(DATA)") {
			dfile = append(dfile, path+f.Name())
		}
	}
	return dfile
}

//leehs 20220519 주어진 경로의 지정된 파일만 반환
func getOnlyFiles(path string, name string) []string {
	var dfile []string
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if strings.ToLower(f.Name()) == "(data)"+name+".xlsx" {
			dfile = append(dfile, path+f.Name())
		}
	}
	if len(dfile) == 0 {
		log.Println("No such file")
	}
	return dfile
}

//leehs 20220519 지정된 파일만 db에 추가
func (this Resource) Upload(w http.ResponseWriter, req *http.Request) {
	root := "C:\\work\\client\\TS\\Server\\Data\\"

	tx := dbHandler.Begin()
	dbHandler.CheckConn()

	defer s.Close()
	var dfile []string
	name := strings.Split(req.FormValue("name"), ",")

	for _, i := range name {
		dfile = append(dfile, getOnlyFiles(root, i)...)
	}

	//leehs 20220526 잘못된 이름이 하나라도 있으면 전체 메소드 종료
	if len(dfile) != len(name) {
		msg := "Invalid file name"
		w.Write([]byte(msg))
		return
	}

	for _, i := range dfile {

		xlFile, err := xlsx.OpenFile(i)
		if err != nil {
			log.Println(err)
		}

		sheet := xlFile.Sheets[0]

		//leehs 20220517 테이블을 생성하는 query
		_, err = tx.Query("CREATE TABLE IF NOT EXISTS " + sheet.Name + " (" + sheet.Cell(1, 0).Value + " " + sheet.Cell(0, 0).Value + ")")
		if err != nil {
			log.Println(err)
		}

		//leehs 20220517 한국어 사용을 가능하게 하는 query
		_, err = tx.Query("ALTER TABLE " + sheet.Name + " convert to charset utf8")
		if err != nil {
			log.Println(err)
		}

		for i, _ := range sheet.Cols {
			if i == 0 {
				continue
			}
			tmp := sheet.Cell(0, i).Value
			if tmp == "string" {
				tmp = "varchar(256)"
			}
			if tmp == "#" {
				continue
			}
			//leehs 20220517 column을 추가하는 query
			_, err = tx.Query("ALTER TABLE " + sheet.Name + " ADD " + sheet.Cell(1, i).Value + " " + tmp)
			if err != nil {
				log.Println(err)
			}
		}

		for i := 2; i < len(sheet.Rows); i++ {
			var line []string
			for j, _ := range sheet.Cols {
				if sheet.Cell(0, j).Value != "#" {
					line = append(line, "'"+sheet.Cell(i, j).Value+"'")
				}
			}
			//leehs 20220517 데이터를 추가하는 query
			_, err = tx.Query("INSERT INTO " + sheet.Name + " VALUES (" + strings.Join(line, ", ") + ")")
			if err != nil {
				log.Println(err, sheet.Name)
			}
		}
	}

	err := dbHandler.Commit(tx)
	if err != nil {
		dbHandler.Rollback(tx)
		log.Println(err)
	}
	msg := "Upload complete"
	msg = strings.Join(name, ", ") + " " + msg
	w.Write([]byte(msg))
}

//leehs 20220517 데이터 파일들의 모든 데이터를 db에 추가
func (this Resource) UploadAll(w http.ResponseWriter, req *http.Request) {
	root := "C:\\work\\client\\TS\\Server\\Data\\"

	dbHandler.CheckConn()
	tx := dbHandler.Begin()

	dfile := getFiles(root)

	for _, i := range dfile {

		xlFile, err := xlsx.OpenFile(i)
		if err != nil {
			log.Println(err)
		}

		sheet := xlFile.Sheets[0]

		//leehs 20220517 테이블을 생성하는 query
		_, err = tx.Query("CREATE TABLE IF NOT EXISTS " + sheet.Name + " (" + sheet.Cell(1, 0).Value + " " + sheet.Cell(0, 0).Value + ")")
		if err != nil {
			log.Println(err)
		}

		//leehs 20220517 한국어 사용을 가능하게 하는 query
		_, err = tx.Query("ALTER TABLE " + sheet.Name + " convert to charset utf8")
		if err != nil {
			log.Println(err)
		}

		for i, _ := range sheet.Cols {
			if i == 0 {
				continue
			}
			tmp := sheet.Cell(0, i).Value
			if tmp == "string" {
				tmp = "varchar(256)"
			}
			if tmp == "#" {
				continue
			}
			//leehs 20220517 column을 추가하는 query
			_, err = tx.Query("ALTER TABLE " + sheet.Name + " ADD " + sheet.Cell(1, i).Value + " " + tmp)
			if err != nil {
				log.Println(err)
			}
		}

		for i := 2; i < len(sheet.Rows); i++ {
			var line []string
			for j, _ := range sheet.Cols {
				if sheet.Cell(0, j).Value != "#" {
					line = append(line, "'"+sheet.Cell(i, j).Value+"'")
				}
			}
			//leehs 20220517 데이터를 추가하는 query
			_, err = tx.Query("INSERT INTO " + sheet.Name + " VALUES (" + strings.Join(line, ", ") + ")")
			if err != nil {
				log.Println(err, sheet.Name)
			}
		}
	}

	err := dbHandler.Commit(tx)
	if err != nil {
		dbHandler.Rollback(tx)
		log.Println(err)
	}
	msg := "Upload complete"
	w.Write([]byte(msg))
}

//leehs 20220517 db의 모든 데이터를 삭제
func (this Resource) Clear(w http.ResponseWriter, req *http.Request) {
	dbHandler.CheckConn()
	tx := dbHandler.Begin()

	//leehs 20220517 db를 drop한 뒤 같은 이름의 db를 새로 생성
	_, err := tx.Query("DROP DATABASE test")
	if err != nil {
		log.Println(err)
	}

	_, err = tx.Query("CREATE DATABASE test")
	if err != nil {
		log.Println(err)
	}

	defer dbHandler.Rollback(tx)
	dbHandler.Commit(tx)

	msg := "Clear complete"
	w.Write([]byte(msg))
}
