package functions

func Split(a, b int) (res, rem int){
    res = a/b
    rem = a%b
    return res, rem
}
