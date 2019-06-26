package pachong1

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var image_path = "d:/project/go/src/goprj1/image"
var htmltext string = ""


// 1. 通过演员获取演员信息，以及演员电影url
// 2. 循环电影URL， 通过获取电影信息
// 3. 获取电影图片信息
// 4. 获取磁力链接信息

type Movie struct{
	Title string `json:"title"`
	Logo string	`json:"logo"`
	Info string `json:"info"`
	ReleaseTime string `json:"release_time"`
	Length string `json:"length"`
	Director string `json:"director"`
	Producers string `json:"producers"`
	Issuer string `json:"issuer"`
	Series string `json:"series"`
	Genre string `json:"genre"`
	Performer string `json:"performer"`
	Url string `json:"url"`
}

type Performer struct{
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Age string `json:"age"`
	Height string `json:"height"`
	Cup string `json:"cup"`
	Bust string `json:"bust"`
	Waist string `json:"waist"`
	Hips string `json:"hips"`
	Hometown string `json:"hometown"`
	Hobby string `json:"hobby"`
	Logo string `json:"logo"`
	Url string `json:"url"`
}

type Link struct{
	Title string `json:"title"`
	Url string `json:"url"`
	Name string `json:"name"`
	Size string `json:"size"`
	Date string `json:"date"`
}

type Image struct{
	Title string `json:"title"`
	Url string `json:"url"`
	Name string `json:"name"`
}

type GetMovie struct{
	Title string `json:"title"`
}

func getHtml(url string)(string, error){
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("获取网页出错, ", err)
		return "", err
	}
	html_text := ""
	buf := make([]byte, 4096)
	for{
		n, err := resp.Body.Read(buf)
		if err != nil && err.Error() != "EOF"{
			fmt.Println("读取异常, " ,err)
			return "", err
		}
		if n == 0{
			break
		}
		html_text += string(buf[:n])
	}
	defer resp.Body.Close()
	html_text = strings.ReplaceAll(strings.ReplaceAll(html_text,"\n", ""),"  ","")
	return html_text, err
}
func (this *GetMovie) getPerformerUrls(html_text string)([]string, error){
	var urls []string
	urls_ret := regexp.MustCompile(`href="(https://avmoo.asia/cn/star/.*?)"`)
	urls_list := urls_ret.FindAllStringSubmatch(html_text, -1)
	for _, url_item := range urls_list{
		urls = append(urls, url_item[1])
	}
	return urls, nil
}

func (this *GetMovie) getPerformerInfo(url string)(Performer,int, error){
	performer := Performer{}
	html_text, err := getHtml(url)
	if err != nil{
		fmt.Println(err)
		return performer,0 , err
	}
	info_ret := regexp.MustCompile(`<span class="pb-10">.*?更多无码影片</a>`)
	info_texts := info_ret.FindStringSubmatch(html_text)
	info_text := info_texts[0]

	name_ret := regexp.MustCompile(`<span class="pb-10">(.*?)</span>`)
	name_text := name_ret.FindStringSubmatch(info_text)
	performer.Name = name_text[1]

	birthday_ret := regexp.MustCompile(`<p>生日: (.*?)</p>`)
	birthday_texts := birthday_ret.FindStringSubmatch(info_text)
	if len(birthday_texts) > 1{
		performer.Birthday = birthday_texts[1]
	}

	age_ret := regexp.MustCompile(`<p>年龄: (.*?)</p>`)
	age_texts := age_ret.FindStringSubmatch(info_text)
	if len(age_texts) > 1{
		performer.Age = age_texts[1]
	}

	height_ret := regexp.MustCompile(`<p>身高: (.*?)</p>`)
	height_texts := height_ret.FindStringSubmatch(info_text)
	if len(height_texts) > 1{
		performer.Height = height_texts[1]
	}

	cup_ret := regexp.MustCompile(`<p>罩杯: (.*?)</p>`)
	cup_texts := cup_ret.FindStringSubmatch(info_text)
	if len(cup_texts) > 1{
		performer.Cup = cup_texts[1]
	}

	bust_ret := regexp.MustCompile(`<p>胸围: (.*?)</p>`)
	bust_texts := bust_ret.FindStringSubmatch(info_text)
	if len(bust_texts) > 1{
		performer.Bust = bust_texts[1]
	}

	waist_ret := regexp.MustCompile(`<p>腰围: (.*?)</p>`)
	waist_texts := waist_ret.FindStringSubmatch(info_text)
	if len(waist_texts) > 1{
		performer.Waist = waist_texts[1]
	}

	hips_ret := regexp.MustCompile(`<p>臀围: (.*?)</p>`)
	hips_texts := hips_ret.FindStringSubmatch(info_text)
	if len(hips_texts) > 1{
		performer.Bust = hips_texts[1]
	}

	page_ret := regexp.MustCompile(`<a name="numbar"href="/cn/star/.+?/page/\d+?">(\d+?)</a>`)
	page_texts := page_ret.FindAllStringSubmatch(html_text, -1)
	page, err := strconv.Atoi(page_texts[len(page_texts)-1][1])
	if err != nil{
		fmt.Println(err)
		page = 1
	}
	return performer, page, nil
}

func (this *GetMovie) getMovieUrls(url string)([]string, error){
	var movie_urls []string
	html_text, err := getHtml(url)
	if err != nil{
		fmt.Println(err)
		return movie_urls, err
	}
	urls_ret := regexp.MustCompile(`<a class="movie-box" href="(https://avmoo.asia/cn/movie/.*?)">`)
	urls_list := urls_ret.FindAllStringSubmatch(html_text, -1)
	for _, url_item := range urls_list{
		movie_urls = append(movie_urls, url_item[1])
	}
	return movie_urls, nil
}

func (this *GetMovie) getMovieInfo(url string)(Movie, error){
	html_text, err := getHtml(url)
	if err != nil{
		fmt.Println(err)
		return Movie{}, err
	}
	// 将回车替换掉
	movie := Movie{}
	// 获取container中少量元素
	container_ret := regexp.MustCompile(`<h3>.*下载`)
	container_texts := container_ret.FindStringSubmatch(html_text)
	container_text := container_texts[0]

	info_text_ret := regexp.MustCompile(`<h3>.*?<h4>`)
	info_texts := info_text_ret.FindStringSubmatch(container_text)
	info_text := info_texts[0]

	// 获取info
	ret := regexp.MustCompile(`<h3>(.*?)</h3`)
	infos := ret.FindStringSubmatch(info_text)

	//fmt.Println(infos[1])

	// 获取title
	title := strings.Fields(infos[1])[0]
	movie.Title = title
	this.Title = title

	// 获取logo封面链接
	logo_ret := regexp.MustCompile(`<a class="bigImage" href="(.*?)"`)
	logos := logo_ret.FindStringSubmatch(info_text)
	movie.Logo = title + ".jpg"

	// 保存封面
	this.SaveImage(logos[1], title)

	// 获取发行时间
	release_time_ret := regexp.MustCompile(`发行时间:</span> (.*?)</p>`)
	release_times := release_time_ret.FindStringSubmatch(info_text)
	if len(release_times) > 1{
		movie.ReleaseTime = release_times[1]
	}

	// 获取长度
	length_ret := regexp.MustCompile(`长度:</span> (.*?)</p>`)
	lengths := length_ret.FindStringSubmatch(info_text)
	if len(lengths) > 1{
		movie.Length = lengths[1]
	}
	//fmt.Println(length)

	// 获取导演
	director_ret := regexp.MustCompile(`导演:</span> <a href="https://avmoo.asia/cn/director/.*?">(.*?)</a></p>`)
	directors := director_ret.FindStringSubmatch(info_text)
	if len(directors) > 1{
		movie.Director = directors[1]
	}

	// 获取制作商
	producers_ret := regexp.MustCompile(`制作商: </p><p><a href="https://avmoo.asia/cn/studio/.*?">(.*?)</a></p>`)
	producers := producers_ret.FindStringSubmatch(info_text)
	if len(producers) > 1{
		movie.Producers = producers[1]
	}

	// 获取发行商
	issuer_ret := regexp.MustCompile(`发行商: </p><p><a href="https://avmoo.asia/cn/label/.*?">(.*?)</a></p>`)
	issuer := issuer_ret.FindStringSubmatch(info_text)
	if len(issuer) > 1{
		movie.Issuer = issuer[1]
	}

	// 获取系列
	series_ret := regexp.MustCompile(`系列: </p><p><a href="https://avmoo.asia/cn/series/.*?">(.*?)</a></p>`)
	seriess := series_ret.FindStringSubmatch(info_text)
	if len(seriess) > 1{
		movie.Series = seriess[1]
	}

	// 获取类别
	genre_ret := regexp.MustCompile(`<a href="https://avmoo.asia/cn/genre/.*?">(.*?)</a>`)
	genres := genre_ret.FindAllStringSubmatch(info_text, -1)
	genre := ""
	for _, genre_item := range genres{
		genre += "," + genre_item[1]
	}
	movie.Genre = genre[1:]

	// 获取演员HTML
	performer_html_ret := regexp.MustCompile(`<h4>演员.*?推荐</h4>`)
	performer_htmls := performer_html_ret.FindStringSubmatch(container_text)
	performer_html := ""
	if len(performer_htmls) != 0{
		performer_html = performer_htmls[0]
	}

	// 获取演员
	performer_ret := regexp.MustCompile(`<span>(.*?)</span>`)
	performer_list := performer_ret.FindAllStringSubmatch(performer_html, -1)
	performer := ""
	for _, performer_item := range performer_list{
		performer += "," + performer_item[1]
	}
	movie.Performer = performer[1:]
	return movie, nil
}

func (this *GetMovie) getImage(url string)([]Image, error){
	html_text, err := getHtml(url)
	images := []Image{}
	if err != nil{
		return images, err
	}
	images_ret := regexp.MustCompile(`<h4>样品图像</h4>.*?<h4>推荐</h4>`)
	images_texts := images_ret.FindStringSubmatch(html_text)
	if len(images_texts) == 0{
		return images, nil
	}
	images_text := images_texts[0]

	image_ret := regexp.MustCompile(`href="(.*?)" title="(.*?)"`)
	image_list := image_ret.FindAllStringSubmatch(images_text, -1)
	for _, image_item := range image_list{
		image := Image{}
		image.Title = this.Title
		image.Url = image_item[1]
		image.Name = image_item[2]
		images = append(images, image)
	}
	return images, nil
}

func (this *GetMovie) SaveImage(url,name string){
	resp, err :=http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	file, err := os.Create(image_path + "/" + name + ".jpg")
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		n, err := resp.Body.Read(buf)
		if err != nil && err.Error() != "EOF"{
			fmt.Println("读取异常, " ,err)
			return
		}
		if n == 0{
			break
		}
		_, err = file.Write(buf[:n])
		if err != nil{
			fmt.Println(err)
			return
		}
	}
}

type Mysql struct{
	Db sql.DB
}

func (this *Mysql)Init()error{
	db, err := sql.Open("mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8")
	if err != nil{
		return err
	}
	err = db.Ping()
	if err != nil{
		return err
	}
	this.Db = *db
	return nil
}
func(this *Mysql) SaveMovie(m Movie) error{
	stmt, err := this.Db.Prepare(`insert into movie(title, url, info, logo, genre, series, 
		director, release_time, length, performer) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil{
		return err
	}
	_, err = stmt.Exec(m.Title, m.Url, m.Info, m.Logo,m.Genre, m.Series, m.Director, m.ReleaseTime, m.Length, m.Performer)
	return err
}
func (this *Mysql) SavePerformer(p Performer)error{
	stmt, err := this.Db.Prepare(`insert into performer(name, age, birthday, height, cup, bust, 
		waist, hips, hometown, hobby, logo, url) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil{
		return err
	}
	_, err = stmt.Exec(p.Name, p.Age, p.Birthday, p.Height,p.Cup, p.Bust, p.Waist, p.Hips, p.Hometown, p.Hobby, p.Logo, p.Url)
	return err
}
func (this *Mysql) SaveLink(l Link)error{
	stmt, err := this.Db.Prepare(`insert into link(name, url, title, share_time, size) values(?, ?, ?, ?, ?)`)
	if err != nil{
		return err
	}
	_, err = stmt.Exec(l.Name, l.Url, l.Title, l.Date, l.Size)
	return err
}
func (this *Mysql) SaveImage(i Image)error{
	stmt, err := this.Db.Prepare(`insert into image(title, name, url) values(?, ?, ?)`)
	if err != nil{
		return err
	}
	_, err = stmt.Exec(i.Title, i.Name, i.Url)
	return err
}
func (this *Mysql)updatePerformer(p Performer)error{
	var logo string
	stmt, err := this.Db.Prepare(`select logo from movie where performer=?`)
	if err != nil{
		return err
	}
	err = stmt.QueryRow(p.Name).Scan(&logo)
	if err != nil{
		return err
	}
	stmt, err = this.Db.Prepare(`update performer set logo=? where name=?`)
	_, err = stmt.Exec(logo, p.Name)
	return err
}

func ErrorExit(err error, message string){
	if err != nil{
		fmt.Println(message, err)
		os.Exit(-1)
	}
}

func main(){
	get_movie := GetMovie{}
	mysql := Mysql{}
	err := mysql.Init()
	ErrorExit(err, "数据库连接异常")

	// 获取演员HTML
	performer_text, err := getHtml("https://avmoo.asia/cn/actresses")
	ErrorExit(err, "获取演员页面异常")

	// 获取演员URL列表
	performer_urls, err := get_movie.getPerformerUrls(performer_text)
	ErrorExit(err, "获取演员URL列表异常")
	fmt.Println("performer_urls: ", performer_urls)
	//
	for index, per_url := range performer_urls{
		performer, page, err := get_movie.getPerformerInfo(per_url)
		if err != nil{
			fmt.Println(err)
			continue
		}
		performer.Url = per_url
		if err = mysql.SavePerformer(performer); err != nil{
			fmt.Println(err)
		}
		fmt.Println("perforemr: ", performer)
		for i := 1;i<page+1;i++{
			performer_url := performer.Url
			if i != 1{
				performer_url = performer.Url + "/page/" + strconv.Itoa(i)
			}
			fmt.Println("performer_url: ", performer_url)
			movie_url, err := get_movie.getMovieUrls(performer_url)
			fmt.Println("movie_urls:", movie_url)
			if err != nil{
				fmt.Println("获取电影url异常", err)
				continue
			}
			for _, url := range movie_url{
				fmt.Println("movie_url:", url)
				m, err := get_movie.getMovieInfo(url)
				if err != nil{
					fmt.Println("获取电影信息异常", err, url)
					continue
				}
				err = mysql.SaveMovie(m)
				if err != nil{
					fmt.Println(err, "保存电影信息异常")
				}
				images, err := get_movie.getImage(url)
				fmt.Println("images: ", images)
				if err != nil{
					fmt.Println(err, "获取图片异常")
				}
				for _, image := range images{
					err = mysql.SaveImage(image)
					if err != nil{
						fmt.Println("保存图片异常", err)
					}
					go get_movie.SaveImage(image.Url, image.Name)
				}
			}
		}
		err = mysql.updatePerformer(performer)
		if err != nil{
			fmt.Println(err, "更新logo异常")
		}
		fmt.Println(performer, page)
		if index > 5{
			break
		}
	}
}
