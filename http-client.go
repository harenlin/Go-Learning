package main
import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.nccu.edu.tw/app/home.php")
	// response, err := http.Get("http://trendmicro.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response status:", response.Status) // Response status: 200 OK
	// read out the webpage
    // body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))
	scanner := bufio.NewScanner(response.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

/* print out:
Response status: 200 OK
<!DOCTYPE html>
<html lang="zh-tw">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
*/