# About
Two programs that allow users to preview the formatting of a markdown file.

* CLI contains a command line program that parses and serves the markdown file in the brower.
* GUI contains a program that allows the user to write/ edit markdown files and view the format changes in real time. NOTE: the program displays the preview in plain text, since I haven't figured out how to webview working in Fyne.

# Warning
The cli program will try to gracefully shutdown when you press Ctrl+C. If you close the program without pressing Ctrl+C, the port will be occupied and you will need to either use a different port or kill the process.

# Why?
I am bad at remembering the syntax of markdown. This program allows me to preview the file and see how it looks before I commit it.

# Usage
```bash
    $ cd CLI/
    $ go run Preview.go -file=<file> -port=<port>
```
or 
```bash
    $ cd GUI/
    $ go run PreviewGUI.go
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
* [ ] Add support for copying and pasting in the GUI