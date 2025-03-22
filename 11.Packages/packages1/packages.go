package main  
  
import "encoding/json"  
  
type jSON struct {  
    Name string `json:"name"`  
}  
  
func main() {  
    var data []byte  
    json.Unmarshal(data, &jSON{})  
}

/*
چون پکیج 
json 
نیاز داره تا فیلد 
Name
 رو تغییر بده پس در نتیجه در کد بالا حتما باید این فیلد رو با حرف بزرگ می‌نوشتیم.
*/