package main

import (
  "strings"
  "fmt"
)

func strip_string_left_and_right(my_line_with_hostname string){
	//fmt.Println("BEFORE: " + my_line_with_hostname)
	//fmt.Println(my_line_with_hostname)
	//str_trimright_bits := "': task 'netmiko_send_command' failed with traceback:"
	//str_trimright_bits := "2023-01-22 20:48:49,702 - nornir.core.task -    ERROR -      start() - Host '"
	//str_trimleft_bits := ""
	//res_out := strings.TrimLeft(strings.TrimRight(my_line_with_hostname))
	//res_out := strings.TrimLeft(strings.TrimRight(my_line_with_hostname, str_trimright_bits),
	//fmt.Println(res_out)  
}

func main(){
  my_line := "2023-01-22 20:48:49,702 - nornir.core.task -    ERROR -      start() - Host 'swi-lia-au177-11-012.au177-10.lia.corpintra.net': task 'netmiko_send_command' failed with traceback:"
  strip_string_left_and_right(my_line)  
}
