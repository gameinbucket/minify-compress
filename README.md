# minify-compress
golang program to minify and compress common files for use on the web

## how to compile
the main function is located in mini.go. Other go files are located in the lib folder. To build an executable run: go build -o "minify"

## how to use
simply run the executable and all html, css, js, and png files found in a given "raw" folder will be minified and compressed into an adjacent "public" folder.
