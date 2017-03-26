package main

import (
    "fmt"
    "sync"
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

func compressor(directory string) {
    fmt.Println("directory " + directory)
    info, err := ioutil.ReadDir(directory)
    assert(err)
    for _ , file := range info {
        if file.IsDir() {
            compressor(directory + "/" + file.Name())
        } else {
            extension := path.Ext(file.Name())
            if extension == ".html" || extension == ".css" || extension == ".js" {
                compress(directory + "/" + file.Name(), directory + "/" + file.Name() + ".gz")
            }
        }
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

func small(path string, name string, extension string, group * sync.WaitGroup) {
    original := get(path + "/" + name + "." + extension)
    fmt.Print("original: " + string(original))
    mini := lib.MiniHTML(original)
    fmt.Print("mini: " + string(mini))
    file := create(path + "/" + name + ".min." + extension)
    file.Write(mini)
    file.Close()
    defer group.Done()
}

func main() {
    // compressor("public")
    // makeJPEG("public/solaire-of-astora.png", "public/solaire-of-astora.jpg")
    // makeGIF("public/solaire-of-astora.png", "public/solaire-of-astora.gif")
 
    var group sync.WaitGroup
    group.Add(1)
    go small("../example", "app", "html", &group)
    group.Wait()
}