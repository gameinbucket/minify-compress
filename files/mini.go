package main

import (
    "fmt"
    "sync"
    "bytes"
    "os"
    "io"
    "io/ioutil"
    "compress/gzip"
    "strconv"
    "path"
    "image"
    "image/jpeg"
    "image/png"
    "image/gif"
)

import "./lib"

func assert(e error) {
    if e != nil {
        panic(e)
    }
}

func compress(src string, dest string) {
    input, err := os.Open(src)
    defer input.Close()
    assert(err)
    output, err := os.Create(dest)
    defer output.Close()
    assert(err)
    gz := gzip.NewWriter(output)
    defer gz.Close()
    val, err := io.Copy(gz, input)
    assert(err)
    fmt.Println("compressed " + src + " to " + dest + " (" + strconv.FormatInt(val, 10) + ")")
}

func decompress(src string, dest string) {
    input, err := os.Open(src)
    defer input.Close()
    assert(err)
    output, err := os.Create(dest)
    defer output.Close()
    assert(err)
    gz, err := gzip.NewReader(input)
    defer gz.Close()
    assert(err)
    val, err := io.Copy(output, gz)
    assert(err)
    fmt.Println("decompressed " + src + " to " + dest + " (" + strconv.FormatInt(val, 10) + ")")
}

func getPNG(src string) (image.Image) {
    input, err := os.Open(src)
    defer input.Close()
    assert(err)
    png, err := png.Decode(input)
    assert(err)
    return png
}

func makeJPEG(src string, dest string) {
    png := getPNG(src)
    output, err := os.Create(dest)
    defer output.Close()
    assert(err)
    jpeg.Encode(output, png, &jpeg.Options{80})
}

func makeGIF(src string, dest string) {
    png := getPNG(src)
    output, err := os.Create(dest)
    defer output.Close()
    assert(err)
    gif.Encode(output, png, nil)
}

func get(name string) ([]byte) {
    data, err := ioutil.ReadFile(name)
    assert(err)
    return data
}

func open(name string) (* os.File) {
    file, err := os.Open(name)
    assert(err)
    return file
}

func create(name string) (* os.File) {
    file, err := os.Create(name)
    assert(err)
    return file
}

func write(file * os.File, data []byte) {
    file.Write(data)
    file.Close()
}

func small(from string, name string, extension string, to string, group * sync.WaitGroup) {
    if extension == ".html" {
        original := get(from + "/" + name)
        mini := lib.MiniHTML(original)
        file := create(to + "/" + name)
        file.Write(mini)
        file.Close()
        compress(to + "/" + name, to + "/" + name + ".gz")
    } else if extension == ".css" {
        original := get(from + "/" + name)
        size := len(original)
        dest := make([]byte, size)
        j := 0
        lib.MiniCSS(original, 0, size, dest, &j)
        mini := bytes.TrimRight(dest, "\x00")
        file := create(to + "/" + name)
        file.Write(mini)
        file.Close()
        compress(to + "/" + name, to + "/" + name + ".gz")
    } else if extension == ".js" {
        original := get(from + "/" + name)
        size := len(original)
        dest := make([]byte, size)
        j := 0
        lib.MiniJS(original, 0, size, dest, &j)
        mini := bytes.TrimRight(dest, "\x00")
        file := create(to + "/" + name)
        file.Write(mini)
        file.Close()
        compress(to + "/" + name, to + "/" + name + ".gz")
    } else if extension == ".png" {
        makeJPEG(from + "/" + name, to + "/" + name[:len(name)-4] + ".jpg")
    }
    defer group.Done()
}

func reduce(from string, to string, group * sync.WaitGroup) {
    fmt.Println(from)
    if _ , err := os.Stat(to); os.IsNotExist(err) {
        os.Mkdir(to, os.ModePerm)
    }
    info, err := ioutil.ReadDir(from)
    assert(err)
    for _ , file := range info {
        name := file.Name()
        if file.IsDir() {
            reduce(from + "/" + name, to + "/" + name, group)
        } else {
            extension := path.Ext(name)
            group.Add(1)
            go small(from, name, extension, to, group)
        }
    }
}

func main() {
    var group sync.WaitGroup
    reduce("raw", "public", &group)
    group.Wait()
}