package lib

func MiniJS(src []byte, begin int, end int, dest []byte, j * int) {
    for i := begin; i < end; i++ {
        dest[*j] = src[i]
        *j++
    }
}