package db

import (
	"database/sql"
	"echo/Config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

var CONFIG = Config.LoadConfig()

type Handler struct {
	Tx *Tx
	db *DB
}

type Tx struct {
	tx      *sql.Tx
	handler *Handler
}

type DB struct {
	DB      *sql.DB
	handler *Handler
}

//leehs 20220526 DB연결 확인 후 없으면 재접속
func (this *Handler) CheckConn() {
	err := this.db.DB.Ping()

	if err != nil {
		log.Println(err)
		this.Init()
	} else {
		return
	}
}

//leehs 20220526 DB 접속 및 Timeout 설정
func (this *Handler) Init() *sql.DB {
	db, _ := sql.Open("mysql", CONFIG.Db.User+":"+CONFIG.Db.Pwd+"@tcp("+CONFIG.Db.Port+")/"+CONFIG.Db.Name)
	t, _ := time.ParseDuration(CONFIG.Db.Timeout)
	db.SetConnMaxLifetime(t)

	this.db = &DB{db, this}
	this.Tx = &Tx{}
	return db
}

//leehs 20220526 쿼리 전달
func (this *DB) Query(sql string, arg ...string) (*sql.Rows, error) {
	if len(arg) == 0 {
		return this.DB.Query(sql)
	} else {
		return this.DB.Query(sql, strings.Join(arg, ", "))
	}
}

func (this *Tx) Query(sql string, arg ...string) (*sql.Rows, error) {
	if len(arg) == 0 {
		return this.tx.Query(sql)
	} else {
		return this.tx.Query(sql, strings.Join(arg, ", "))
	}
}

//leehs 20220526 Tx 시작
func (this *Handler) Begin() *Tx {
	tx, _ := this.db.DB.Begin()
	return &Tx{tx, this}
}

//leehs 20220526 Tx 실제 DB반영
func (this *Handler) Commit(tx *Tx) error {
	return tx.tx.Commit()
}

func (this *Handler) Rollback(tx *Tx) error {
	return tx.tx.Rollback()
}
