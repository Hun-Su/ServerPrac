package server

import (
	"database/sql"
	"echo/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"strings"
)

type Resource struct{}

//leehs 20220517 주어진 경로의 데이터 파일들의 리스트를 반환
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

//leehs 20220518 db connection
func openDB() *sql.DB {
	config := Config.LoadConfig()
	s, err := sql.Open("mysql", config.Db.User+":"+config.Db.Pwd+"@tcp("+config.Db.Port+")/"+config.Db.Name)

	if err != nil {
		log.Fatalln(err)
	}

	return s
}

//leehs 20220519 지정된 파일만 db에 추가
func (this Resource) Upload(name ...string) {
	root := "C:\\work\\client\\TS\\Server\\Data\\"
	s := openDB()
	defer s.Close()
	var dfile []string

	for _, i := range name {
		dfile = append(dfile, getOnlyFiles(root, i)...)
	}

	for _, i := range dfile {
		xlFile, err := xlsx.OpenFile(i)

		if err != nil {
			log.Fatalln(err)
		}

		sheet := xlFile.Sheets[0]

		//leehs 20220517 테이블을 생성하는 query
		NewTable := fmt.Sprint("CREATE TABLE IF NOT EXISTS ", sheet.Name, "(", sheet.Cell(1, 0), sheet.Cell(0, 0), ")")
		s.Exec(NewTable)

		//leehs 20220517 한국어 사용을 가능하게 하는 query
		alter := fmt.Sprint("ALTER TABLE ", sheet.Name, " convert to charset utf8")
		s.Exec(alter)

		for i, _ := range sheet.Cols {
			tmp := sheet.Cell(0, i).Value
			if tmp == "string" {
				tmp = "varchar(256)"
			}
			if tmp == "#" {
				continue
			}
			//leehs 20220517 column을 추가하는 query
			AddCol := fmt.Sprint("ALTER TABLE ", sheet.Name, " ADD ", sheet.Cell(1, i), " ", tmp)
			s.Exec(AddCol)
		}

		for i := 2; i < len(sheet.Rows); i++ {
			var line []string
			for j, _ := range sheet.Cols {
				if sheet.Cell(0, j).Value != "#" {
					line = append(line, "'"+sheet.Cell(i, j).Value+"'")
				}
			}
			//leehs 20220517 데이터를 추가하는 query
			ins := fmt.Sprint("INSERT INTO ", sheet.Name, " VALUES (", strings.Join(line, ", "), ")")
			s.Exec(ins)
		}
	}
}

//leehs 20220517 데이터 파일들의 모든 데이터를 db에 추가
func (this Resource) UploadAll() {
	root := "C:\\work\\client\\TS\\Server\\Data\\"
	s := openDB()
	defer s.Close()

	dfile := getFiles(root)

	for _, i := range dfile {
		xlFile, err := xlsx.OpenFile(i)

		if err != nil {
			log.Fatalln(err)
		}

		sheet := xlFile.Sheets[0]

		//leehs 20220517 테이블을 생성하는 query
		NewTable := fmt.Sprint("CREATE TABLE IF NOT EXISTS ", sheet.Name, "(", sheet.Cell(1, 0), sheet.Cell(0, 0), ")")
		s.Exec(NewTable)

		//leehs 20220517 한국어 사용을 가능하게 하는 query
		alter := fmt.Sprint("ALTER TABLE ", sheet.Name, " convert to charset utf8")
		s.Exec(alter)

		for i, _ := range sheet.Cols {
			tmp := sheet.Cell(0, i).Value
			if tmp == "string" {
				tmp = "varchar(256)"
			}
			if tmp == "#" {
				continue
			}
			//leehs 20220517 column을 추가하는 query
			AddCol := fmt.Sprint("ALTER TABLE ", sheet.Name, " ADD ", sheet.Cell(1, i), " ", tmp)
			s.Exec(AddCol)
		}

		for i := 2; i < len(sheet.Rows); i++ {
			var line []string
			for j, _ := range sheet.Cols {
				if sheet.Cell(0, j).Value != "#" {
					line = append(line, "'"+sheet.Cell(i, j).Value+"'")
				}
			}
			//leehs 20220517 데이터를 추가하는 query
			ins := fmt.Sprint("INSERT INTO ", sheet.Name, " VALUES (", strings.Join(line, ", "), ")")
			s.Exec(ins)
		}
	}
}

//leehs 20220517 db의 모든 데이터를 삭제
func (this Resource) Clear() {
	s := openDB()
	defer s.Close()

	//leehs 20220517 db를 drop한 뒤 같은 이름의 db를 새로 생성
	clear := fmt.Sprint("DROP DATABASE test")
	create := fmt.Sprint("CREATE DATABASE test")

	s.Exec(clear)
	s.Exec(create)
}
