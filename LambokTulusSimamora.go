package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	fileExt := flag.String("t", "text", "File extension (default .txt). (Optional)")
	fileDest := flag.String("o", "error", "File destination (default current directory). (Optional)")
	flag.Parse()

	var dest string
	src := os.Getenv("ERROR_LOG_PATH")
	if src == "" {
		src = "/var/log/nginx/error.log"
	}

	if *fileExt == "text" {
		*fileExt = "txt"
	}

	dest = *fileDest + "." + *fileExt

	ext := strings.Split(*fileDest, ".")
	extLen := len(ext)

	if len(ext) > 1 && (ext[extLen-1] == "txt" || ext[extLen-1] == "json") {
		dest = *fileDest
	}

	moveFile(src, dest)
}

func moveFile(sourceFile string, destFile string) {
	e := os.Rename(sourceFile, destFile)

	if e != nil {
		log.Fatal(e)
		return
	}
	log.Print("Success")
}
