package lib

/*
Known bugs: what about spaces in special deliminators?
*/

func MiniCSS(src []byte, begin int, end int, dest []byte, j * int) {
    inString := false
    for i := begin; i < end; i++ {
        if inString {
            if src[i] == '"' {
                inString = false
            }
        } else {
            if src[i] == ' ' || src[i] == '\r' || src[i] == '\n' {
                continue
            } else if src[i] == '"' {
                inString = true
            } else if i + 1 < end && src[i] == '/' && src[i+1] == '/' {
                for {
                    i++
                    if i == end || src[i] == '\r' || src[i] == '\n' {
                        break
                    }
                }
                continue
            } else if i + 1 < end && src[i] == '/' && src[i+1] == '*' {
                for {
                    i++
                    if i == end || (src[i-1] == '*' && src[i] == '/') {
                        break
                    }
                }
                continue
            }
        }
        dest[*j] = src[i]
        *j++
    }
}