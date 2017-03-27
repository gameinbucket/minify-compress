package lib

func MiniJS(src []byte, begin int, end int, dest []byte, j * int) {
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
            if src[i] == ' ' {
                if i == 0 || i == end - 1 || src[i-1] == ' ' || !letter(src[i-1]) || !letter(src[i+1]) {
                    continue
                }
            } else if src[i] == '\r' || src[i] == '\n' {
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

func letter(b byte) (bool) {
    switch b {
    case 'a': return true
    case 'b': return true
    case 'c': return true
    case 'd': return true
    case 'e': return true
    case 'f': return true
    case 'g': return true
    case 'h': return true
    case 'i': return true
    case 'j': return true
    case 'k': return true
    case 'l': return true
    case 'm': return true
    case 'n': return true
    case 'o': return true
    case 'p': return true
    case 'q': return true
    case 'r': return true
    case 's': return true
    case 't': return true
    case 'u': return true
    case 'v': return true
    case 'w': return true
    case 'x': return true
    case 'y': return true
    case 'z': return true
    default: return false
    }
}