package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init(){
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8", 30)
}

type User struct {
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Pwd string `orm:"column(pwd)"json:"pwd"`
	Email string `orm:"column(email)"json:"email"`
	Phone string `orm:"column(phone)"json:"phone"`
	Info string `orm:"column(info)"json:"info"`
	Face string `orm:"column(face)"json:"face"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
	Uuid string `orm:"column(uuid)"json:"uuid"`
}

type Userlog struct{
	Id int `orm:"column(id);auto"json:"id"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Ip string `orm:"column(ip)"json:"ip"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type Tag struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type Movie struct{
	Id int `orm:"column(id);auto"json:"id"`
	Title string `orm:"column(title)"json:"title"`
	Url string `orm:"column(url)"json:"url"`
	Info string `orm:"column(info)"json:"info"`
	Logo string `orm:"column(logo)"json:"logo"`
	Genre string `orm:"column(genre)"json:"genre"`
	Series string `orm:"column(series)"json:"series"`
	Vender string `orm:"column(vender)"json:"vender"`
	Director string `orm:"column(director)"json:"director"`
	Studio string `orm:"column(studio)"json:"studio"`
	Star int `orm:"column(str)"json:"star"`
	PlayNum int `orm:"column(playnum)"json:"playnum"`
	CommentNum int `orm:"column(commentnum)"json:"commentnum"`
	TagId int `orm:"column(tag_id)"json:"tag_id"`
	Area string `orm:"column(area)"json:"area"`
	ReleaseTime time.Time `orm:"column(release_time);type(datetime)"json:"release_json"`
	Length string `orm:"column(length)"json:"length"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
	Performer string `orm:"column(performer)"json:"performer"`
	State int `orm:"column(state)"json:"state"`
	Video string `orm:"column(video)"json:"video"`
	IdDelete int `orm:"column(is_delete)"json:"is_delete"`
}

type Performer struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Age int `orm:"column(age)"json:"age"`
	Birthday time.Time `orm:"column(birthday);type(datetime)"json:"birthday"`
	Height string `orm:"column(height)"json:"height"`
	Cup string `orm:"column(cup)"json:"cup"`
	Bust string `orm:"column(bust)"json:"bust"`
	Waist string `orm:"column(waist)"json:"waist"`
	Hips string `orm:"column(hips)"json:"hips"`
	Hometown string `orm:"column(hometown)"json:"hometown"`
	Hobby string `orm:"column(hobby)"json:"hobby"`
	IdDelete int `orm:"column(is_delete)"json:"is_delete"`
	Image string `orm:"column(image)"json:"image"`
}

type Link struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Url string `orm:"column(url)"json:"url"`
	MovieId int `orm:"column(movie_id)"json:"movie_id"`
}

type Image struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Url string `orm:"column(url)"json:"url"`
	MovieId int `orm:"column(movie_id)"json:"movie_id"`
}

type Preview struct{
	Id int `orm:"column(id);auto"json:"id"`
	Title string `orm:"column(title)"json:"title"`
	Logo string `orm:"column(logo)"json:"logo"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type Comment struct{
	Id int `orm:"column(id);auto"json:"id"`
	Content string `orm:"column(content)"json:"content"`
	MovieId int `orm:"column(movie_id)"json:"movie_id"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type MovieCol struct{
	Id int `orm:"column(id);auto"json:"id"`
	MovieId int `orm:"column(movie_id)"json:"movie_id"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type Follow struct{
	Id int `orm:"column(id);auto"json:"id"`
	PerformerId int `orm:"column(performer_id)"json:"performer_id"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

