package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init(){
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8", 30)
}

type User struct {
	Id int `orm:"column(id);auto"`
	Name string `orm:"column(name)"`
	Pwd string `orm:"column(pwd)"`
	Email string `orm:"column(email)"`
	Phone string `orm:"column(phone)"`
	Info string `orm:"column(info)"`
	Face string `orm:"column(face)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
	Uuid string `orm:"column(uuid)"`
}

type Userlog struct{
	Id int `orm:"column(id);auto"`
	UserId int `orm:"column(user_id)"`
	Ip string `orm:"column(ip)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

type Tag struct{
	Id int `orm:"column(id);auto"`
	Name string `orm:"column(name)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

type Movie struct{
	Id int `orm:"column(id);auto"`
	Title string `orm:"column(title)"`
	Url string `orm:"column(url)"`
	Info string `orm:"column(info)"`
	Logo string `orm:"column(logo)"`
	Genre string `orm:"column(genre)"`
	Series string `orm:"column(series)"`
	Vender string `orm:"column(vender)"`
	Director string `orm:"column(director)"`
	Studio string `orm:"column(studio)"`
	Star int `orm:"column(str)"`
	PlayNum int `orm:"column(playnum)"`
	CommentNum int `orm:"column(commentnum)"`
	TagId int `orm:"column(tag_id)"`
	Area string `orm:"column(area)"`
	ReleaseTime time.Time `orm:"column(release_time);type(datetime)"`
	Length string `orm:"column(length)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
	Performer string `orm:"column(performer)"`
	State int `orm:"column(state)"`
	Video string `orm:"column(video)"`
	IdDelete int `orm:"column(is_delete)"`
}

type Performer struct{
	Id int `orm:"column(id);auto"`
	Name string `orm:"column(name)"`
	Age int `orm:"column(age)"`
	Birthday time.Time `orm:"column(birthday);type(datetime)"`
	Height string `orm:"column(height)"`
	Cup string `orm:"column(cup)"`
	Bust string `orm:"column(bust)"`
	Waist string `orm:"column(waist)"`
	Hips string `orm:"column(hips)"`
	Hometown string `orm:"column(hometown)"`
	Hobby string `orm:"column(hobby)"`
	IdDelete int `orm:"column(is_delete)"`
	Image string `orm:"column(image)"`
}

type Link struct{
	Id int `orm:"column(id);auto"`
	Name string `orm:"column(name)"`
	Url string `orm:"column(url)"`
	MovieId int `orm:"column(movie_id)"`
}

type Image struct{
	Id int `orm:"column(id);auto"`
	Name string `orm:"column(name)"`
	Url string `orm:"column(url)"`
	MovieId int `orm:"column(movie_id)"`
}

type Preview struct{
	Id int `orm:"column(id);auto"`
	Title string `orm:"column(title)"`
	Logo string `orm:"column(logo)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

type Comment struct{
	Id int `orm:"column(id);auto"`
	Content string `orm:"column(content)"`
	MovieId int `orm:"column(movie_id)"`
	UserId int `orm:"column(user_id)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

type MovieCol struct{
	Id int `orm:"column(id);auto"`
	MovieId int `orm:"column(movie_id)"`
	UserId int `orm:"column(user_id)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

type Follow struct{
	Id int `orm:"column(id);auto"`
	PerformerId int `orm:"column(performer_id)"`
	UserId int `orm:"column(user_id)"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"`
}

