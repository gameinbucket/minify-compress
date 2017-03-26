package lib

/*
Known bugs: what about spaces in special deliminators?
*/

func MiniCSS(src []byte, begin int, end int, dest []byte, j * int) {
    for i := begin; i < end; i++ {
        if src[i] == ' ' || src[i] == '\r' || src[i] == '\n' {
            continue
        }
        dest[*j] = src[i]
        *j++
    }
}