package goroutine

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func Ex2() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어:", word)
	PrintAllFiles(files, word)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string, word string) {

	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. err:", err)
			return
		}
		fmt.Println("찾으려는 파일 리스트")

		ch := make(chan FindInfo)
		cnt := len(filelist)
		recvCnt := 0

		for _, name := range filelist {
			go findWordInFile(name, word, ch)
		}

		for findInfo := range ch {
			recvCnt++

			for _, lineInfo := range findInfo.lines {
				fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
			}
			if recvCnt == cnt {
				break
			}
		}
	}
}

func findWordInFile(filename string, word string, ch chan FindInfo) {
	findInfos := FindInfo{filename, []LineInfo{}}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		ch <- findInfos
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfos.lines = append(findInfos.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfos
}
