package parser

import (
	"Nginx/internal"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func processFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败，%w", err)
	}
	defer file.Close()
	var reader io.Reader = file
	if strings.HasSuffix(strings.ToLower(filePath), ".gz") {
		gzReaader, err1 := gzip.NewReader(file)
		if err1 != nil {
			return fmt.Errorf("创建gzip失败！%w", err1)
		}
		defer gzReaader.Close()
		reader = gzReaader
	}
	scanner := bufio.NewScanner(reader)
	const maxCapacity = 10 * 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		if err2 := scanner.Err(); err2 != nil {
			return fmt.Errorf("逐行读取文件时出现错误：%w", err2)
		}
	}
	return nil
}

var nginxLogRegex = regexp.MustCompile(
	`^(\S+) - \S+ \[([^\]]+)\] "(\S+) (\S+) (\S+)" (\d{3}) (\d+) "([^"]*)" "([^"]*)"$`)

func ParseLogLine(line string) (*internal.Event, error) {
	matches := nginxLogRegex.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("日志格式不匹配%s", line)
	}
	time, err := time.Parse("02/Jan/2006:15:04:05-0700", matches[2])
}
