
package main
import ("fmt"
        "math/rand"
        "time")

func main() {
    var url string
    var shortID string
    var newUrl string
    urlMap:= make( map[string]string)
    lookUp:= make( map[string]string)
    rand.Seed(time.Now().UnixNano())
    
  fmt.Println("URL SHORTENER")
  fmt.Println("\nPlease enter your url")
  fmt.Scanf("%s",&url)
  shortID= assign(url, urlMap, lookUp)
  fmt.Println(shortID)
  newUrl= "smallurl.com/"+shortID
  fmt.Println(newUrl)
   fmt.Println("\nPlease enter your url")
  fmt.Scanf("%s",&url)

  shortID= assign(url, urlMap,lookUp)
  fmt.Println(shortID)
  fmt.Println(urlMap)

  redirect:=retrieve(shortID, urlMap)
  fmt.Println(shortID ,"->",redirect)
  return
}

func assign(url string, urlMap map[string]string, lookUp map[string] string )string{
    var ans string
    value,exists:=lookUp[url]
    if exists{
        fmt.Println("URL shortid exists")
        return value
    }
    
    ans=randomString(urlMap)
    urlMap[ans]=url
    lookUp[url]=ans
    return ans
}
func retrieve(shortID string, urlMap map[string]string)string{
    return urlMap[shortID]
}
func randomString(urlMap map[string] string)string{
    var letters string
    letters="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result:=make([]byte,6)
    for i:=0;i<6;i++{
        index:=rand.Intn(len(letters))
        result[i]=letters[index]
    }
    value,exists := urlMap[string(result)]
    if exists{
        fmt.Println("found key with",value)
        return randomString(urlMap)
        
    }
    return string(result)
}
