package lib

import (
    "bytes"
)

func MiniHTML(src []byte) ([]byte) {
    size := len(src)
    dest := make([]byte, size)
    inString := false
    spaceInString := false
    inTag := false
    inComment := false
    inScript := false
    inStyle := false
    j := 0
    for i := 0; i < size; i++ {
        if inString {
            if src[i] == '"' {
                inString = false
                if inTag && !spaceInString {
                    k := j-1
                    for {
                        if dest[k] == '"' {
                            break
                        } else {
                            k--
                        }
                    }
                    for {
                        if k == j {
                            break
                        } else {
                            dest[k] = dest[k+1]
                            k++
                        }
                    }
                    j--
                    continue
                }
            } else if src[i] == ' ' {
                spaceInString = true
            }
        } else if inComment {
            if src[i] == '>' && src[i-1] == '-' && src[i-2] == '-' {
                inComment = false
            }
            continue
        } else if inScript && !inTag {
            k := i
            for k < size {
                if k + 3 < size && src[k] == '<' && src[k+1] == '/' && src[k+2] == 's' && src[k+3] == 'c' {
                    break
                }
                k++
            }
            MiniJS(src, i, k, dest, &j)
            inScript = false
            i = k-1
            continue
        } else if inStyle && !inTag {
            k := i
            for k < size {
                if k + 3 < size && src[k] == '<' && src[k+1] == '/' && src[k+2] == 's' && src[k+3] == 't' {
                    break
                }
                k++
            }
            MiniCSS(src, i, k, dest, &j)
            inStyle = false
            i = k-1
            continue
        } else {
            switch src[i] {
                case '<':
                if i + 3 < size && src[i+1] == '!' && src[i+2] == '-' && src[i+3] == '-' {
                    inComment = true
                    continue
                } else {
                    inTag = true
                    if i + 6 < size && src[i+1] == 's' && src[i+2] == 'c' && src[i+3] == 'r' && src[i+4] == 'i' && src[i+5] == 'p' && src[i+6] == 't' {
                        inScript = true
                    } else if i + 5 < size && src[i+1] == 's' && src[i+2] == 't' && src[i+3] == 'y' && src[i+4] == 'l' && src[i+5] == 'e' {
                        inStyle = true
                    }
                }
                case '>':
                inTag = false
                case '\n':
                //continue
                case '\r':
                //continue
                case '\t':
                continue
                case ' ':
                if j > 0 && (dest[j-1] == ' ' || dest[j-1] == '>' || dest[j-1] == '\n' || dest[j-1] == '\r') {
                    continue
                }
                case '"':
                inString = true
                spaceInString = false
            }
        }
        dest[j] = src[i]
        j++
    }
    return bytes.TrimRight(dest, "\x00")
}