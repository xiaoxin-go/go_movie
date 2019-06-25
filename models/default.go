package models

import (
	"time"
)

type User struct {
	Id int `orm:"column(id);auto"json:"id"`
	Username string `orm:"column(username)"json:"username"`
	Password string `orm:"column(password)"json:"password"`
	Nickname string `orm:"column(nickname)"json:"nickname"`
	Face string `orm:"column(face)"json:"face"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
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
	Director string `orm:"column(director)"json:"director"`
	ReleaseTime string `orm:"column(release_time)"json:"release_json"`
	Length string `orm:"column(length)"json:"length"`
	Performer string `orm:"column(performer)"json:"performer"`
	Video string `orm:"column(video)"json:"video"`
}

type Performer struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Age string `orm:"column(age)"json:"age"`
	Birthday string `orm:"column(birthday)"json:"birthday"`
	Height string `orm:"column(height)"json:"height"`
	Cup string `orm:"column(cup)"json:"cup"`
	Bust string `orm:"column(bust)"json:"bust"`
	Waist string `orm:"column(waist)"json:"waist"`
	Hips string `orm:"column(hips)"json:"hips"`
	Hometown string `orm:"column(hometown)"json:"hometown"`
	Hobby string `orm:"column(hobby)"json:"hobby"`
	Logo string `orm:"column(logo)"json:"logo"`
}

type Link struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Url string `orm:"column(url)"json:"url"`
	Title string `orm:"column(title)"json:"title"`
	ShareTime string `orm:"column(share_time)"json:"share_time"`
}

type Image struct{
	Id int `orm:"column(id);auto"json:"id"`
	Name string `orm:"column(name)"json:"name"`
	Url string `orm:"column(url)"json:"url"`
	Title string `orm:"column(title)"json:"title"`
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
	Title int `orm:"column(title)"json:"title"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type MovieCol struct{
	Id int `orm:"column(id);auto"json:"id"`
	Title string `orm:"column(title)"json:"title"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

type Follow struct{
	Id int `orm:"column(id);auto"json:"id"`
	Performer string `orm:"column(performer)"json:"performer"`
	UserId int `orm:"column(user_id)"json:"user_id"`
	Addtime time.Time `orm:"auto_now;column(addtime);type(datetime)"json:"addtime"`
}

