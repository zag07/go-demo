package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	num     int
	logFile string
	f       *os.File
)

var events = []string{"visit", "start", "register", "login", "pay"}

var msgs = []string{"哈哈哈哈", "嘿嘿嘿嘿", "呵呵呵呵"}

func main() {
	flag.IntVar(&num, "num", 100, "number of logs")
	flag.StringVar(&logFile, "logFile", "./log/messages/zag13.log", "file of log")
	flag.Parse()

	_, err := os.Stat(logFile)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(path.Dir(logFile), 0750)
		f, err = os.Create(logFile)
	} else {
		f, err = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0666)
	}

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	write := bufio.NewWriter(f)
	for i := 0; i < num; i++ {
		data := map[string]string{
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"event":     events[rand.Intn(len(events))],
			"msg":       msgs[rand.Intn(len(msgs))],
		}
		jsonString, _ := json.Marshal(data)
		write.WriteString(string(jsonString) + "\n")
	}

	write.Flush()

	fmt.Println("已生成日志" + strconv.Itoa(num) + "条")
}
