package lib

/*
Known bugs: what about spaces in special deliminators?
*/

func MiniCSS(src []byte, begin int, end int, dest []byte, j * int) {
    inString := 0
    for i := begin; i < end; i++ {
        if inString == 1 {
            if src[i] == '"' {
                inString = 0
            }
        } else if inString == 2 {
            if src[i] == '\'' {
                inString = 0
            }
        } else {
            if src[i] == ' ' || src[i] == '\r' || src[i] == '\n' {
                continue
            } else if src[i] == '"' {
                inString = 1
            } else if src[i] == '\'' {
                inString = 2
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