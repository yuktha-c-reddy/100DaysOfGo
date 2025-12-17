package main
import ("fmt"
		"net/http"
		"math/rand"
		"time")

var(		
	urlMap=make(map[string]string)
	lookUp=make(map[string]string)
	BASE_URL="www.yurl.com/"
)
var shortID string
var newUrl string


func main(){
	
	fmt.Println("Helloooo")
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/",home)
	http.HandleFunc("/shorten",shorten)
	// http.HandleFunc("/retrieve",retrieve)
	http.ListenAndServe(":8080",nil)
}

func home(w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w,"Url shortener")
}

func shorten(w http.ResponseWriter, r *http.Request){

	
	shortID= calc_id()
	url:=r.URL.Query().Get("url")
	shortID=assign(shortID,url)
	newUrl=BASE_URL+shortID
	fmt.Fprintln(w,"You typed "+url)
	fmt.Fprintln(w,"Your link is "+ newUrl)
}

// func retrieve(w http.ResponseWriter, r *http.Request){
// 	http.Redirect(w,r,)
// }

func calc_id()string{
	result:=make([]byte,6)
	letters:="ABCDEFGHIJKLEMNOPQRSTUVWXYSabcdefghijklmnopqrstuvwxyz0123456789"
	for i:=0;i<6;i++{
		index:=rand.Intn(len(letters))
		result[i]=letters[index]
	}
	return string(result)
}

func assign(shortID string, url string)string{
	value,exists:=lookUp[url]
	fmt.Println("Assigning", value , url)
	if !exists{
		urlMap[shortID]=url
		lookUp[url]=shortID
	} else{
		fmt.Println(value+"already exists, using same shortID")
		return lookUp[url]
	}
	return shortID
}
