package main
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{0,1,2,3,4,5,6}
	b := a[0:4]

	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Printf("a=%v, len=%d, cap=%d, type=%d \n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v, len=%d, cap=%d, type=%d \n", b, len(b), cap(b), bs.Data)

	b = append(b, 10, 11, 12, 13)
	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("a=%v, len=%d, cap=%d, type=%d \n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v, len=%d, cap=%d, type=%d \n", b, len(b), cap(b), bs.Data)


	b = append(b, 10, 11, 12)
	as = (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	fmt.Printf("a=%v, len=%d, cap=%d, type=%d \n", a, len(a), cap(a), as.Data)
	fmt.Printf("b=%v, len=%d, cap=%d, type=%d \n", b, len(b), cap(b), bs.Data)


}