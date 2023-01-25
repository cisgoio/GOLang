//ref: https://stackoverflow.com/questions/73017827/golang-convert-string-in-to-an-array

package main

import(
 "fmt" 
 "strings"
)

func trim_and_split__eg01(my_string string)[]string{
  
  trimmedStr := strings.Trim("[\"firsName\",\"lastName\", \"email\"]", "[]")

  //res := strings.Split(trimmedStr, ",")
  //fmt.Println()
  return strings.Split(trimmedStr, ",")
  
}

func trim_and_split__eg02(my_string string)[]string{
  
  //trimmedStr := strings.Trim(my_string)

  //res := strings.Split(trimmedStr, ",")
  //fmt.Println()
  return strings.Split(my_string, "'")
  
}


func main(){
  
  fmt.Println(trim_and_split__eg01("mystring'device"))

  res_eg01 :=trim_and_split__eg02("mystring'device")
  fmt.Println(trim_and_split__eg02("mystring'device"))
  fmt.Println(res_eg01[1])
}
