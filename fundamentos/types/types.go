package types

import "fmt"

//bool
//
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 utin64 uintptr
//
// byte = uint8
//
// rune = int32
//
//  float32 float64
//
// complex64 complex128

func TakeByte(b byte) {
    fmt.Println(b)
}

func TakeString(s string){
    fmt.Println(s)
}

func TakeInt32(n int32){
    fmt.Println(n)
}

func TakeInt64(n int64){
    fmt.Println(n)
}

func TakeFloat64(f float64){
    fmt.Println(f)
}
