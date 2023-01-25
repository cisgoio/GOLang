//ref: https://stackoverflow.com/questions/73017827/golang-convert-string-in-to-an-array

import(
 "fmt" 
)

fun trim_and_split(my_string)string{
  
  trimmedStr := strings.Trim("[\"firsName\",\"lastName\", \"email\"]", "[]")
  fmt.Println(strings.Split(trimmedStr, ","))
  
}


func main(){
  
  fmt.Println(trim_and_split("mystring'device"))

}
