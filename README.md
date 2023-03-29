# About
A simple program that previews the conetents of a markdown file in a browser.

# Warning
The program will try to gracefully shutdown when you press Ctrl+C. If you close the program without pressing Ctrl+C, the port will be occupied and you will need to either use a different port or kill the process.

# Why?
I am bad at remembering the syntax of markdown. This program allows me to preview the file in a browser and see how it looks before I commit it.

# Usage
```bash
    $ go run Preview.go -file=<file> -port=<port>
```

# Dependencies
* [Blackfriday](https://github.com/russross/blackfriday)
* [Gin](https://github.com/gin-gonic/gin)

# TODOs
* [x] Add a flag to specify the file
* [x] Add a flag to specify the port
* [ ] Create a markdown file that contains all the markdown syntax, to test the program
* [ ] Create my own markdown parser and web framework instead of using Blackfriday and Gin (Maybe I will do this in the future)
* [x] A GUI that allows you to edit the markdown file and preview it in real time