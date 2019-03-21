package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"unsafe"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filename := "response.txt"
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		str := fmt.Sprintln(<-ch)
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			file, err = os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
		}
		defer file.Close()
		str_byte := sbytes(str)
		file.Write(str_byte)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func sbytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
