package main 

/*
#cgo CFLAGS:-I./include
#cgo LDFLAGS:-L./lib 
#include "test_c.h"
*/                                                                                                                                                                                
import "C"
import "fmt" 

func main(){
	fmt.Println(C.test_c())
}
