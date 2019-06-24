package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goprj1/models"
	"strconv"
	"time"
)

func main(){
	db, err := sql.Open("mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie")
	if err != nil{
		fmt.Println("open mysql failed, err: ", err)
		return
	}
	err = db.Ping()
	if err != nil{
		fmt.Println("connection mysql failed err:", err)
	}
	m1 := models.Movie{
		Id:1,
		Title:"test",
		Url:"url",
		Info:"info",
		Logo:"logo",
		Genre:"genre",
		Series:"series",
		Director:"director",
		ReleaseTime:"2019-6-24",
		Length:"144分钟",
		Performer:"performer",
		Video:"video",
	}
	//result, err := db.Exec("insert Into movie(title, url, info, logo, genre, series, director, release_time, length, performer, video)",)
	stmt, err := db.Prepare(`insert Into movie(title, url, info, logo, genre, series, director, release_time,
							length, performer, video) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,)
	if err != nil{
		fmt.Println("预编译错误, ", err, stmt)
		return
	}
	star := time.Now()
	for i := 1045;i<2000;i++{
		//fmt.Println(m1.Title + strconv.Itoa(i))
		_, err := stmt.Exec(string(m1.Title + strconv.Itoa(i)), string(m1.Url + strconv.Itoa(i)), m1.Info + strconv.Itoa(i), m1.Logo + strconv.Itoa(i),
			m1.Genre + strconv.Itoa(i), m1.Series + strconv.Itoa(i), m1.Director + strconv.Itoa(i), m1.ReleaseTime, m1.Length, m1.Performer + strconv.Itoa(i), m1.Video + strconv.Itoa(i))
		if err != nil{
			fmt.Println("插入错误, ", err)
			return
		}
		//last_id, err := result.LastInsertId()
		//fmt.Println("last_id: ", last_id)
	}
	fmt.Println(star)
	fmt.Println(time.Now())

	if err != nil{
		fmt.Println("获取最后ID错误", err)
		return
	}
	defer db.Close()
}
