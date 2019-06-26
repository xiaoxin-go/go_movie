package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	resp, err :=http.Get("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1561522540859&di=eb1b2a5bb4089a8a87d439c5b2a3ec97&imgtype=jpg&src=http%3A%2F%2Fpic.962.net%2Fup%2F2016-7%2F14689002039500090.jpg")
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	file, err := os.Create("d:/project/go/src/goprj1/test.jpg")
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
