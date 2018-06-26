package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getFilelist(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err

}

func replaceRem(content string, m float64) string {
	// -- ([1-9]\d*\.\d*|0\.\d*[1-9]\d*|\d+)rem
	// match like "3.5rem" or "0.35rem" or "35rem"
	// -- \b
	// matcher's left char can't be [0-9] or [a-z] or _
	remRe := regexp.MustCompile(`\b([1-9]\d*\.\d*|0\.\d*[1-9]\d*|\d+)rem`)
	newContent := remRe.ReplaceAllStringFunc(content, func(reString string) string {
		remValue := strings.Replace(reString, "rem", "", -1)
		if s, err := strconv.ParseFloat(remValue, 64); err == nil {
			new := strconv.FormatFloat(s*m, 'f', -1, 64) + "rem"
			return new
		}
		log.Fatalln("can't Parse" + remValue + "to float")
		return reString

	})
	return newContent
}

func main() {
	var m float64
	flag.Float64Var(&m, "m", 0.5, "设定倍数")
	flag.Parse()
	root := flag.Arg(0)
	paths, err := getFilelist(root)
	if err != nil {
		log.Fatalln("can't get files list")
		return
	}
	for _, path := range paths {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			log.Fatalln("cant open file " + path + " , will continue")
			continue
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln("cant read file " + path + " , will continue")
			continue
		}
		content := string(data)
		matched, err := regexp.MatchString(`\b([1-9]\d*\.\d*|0\.\d*[1-9]\d*|\d+)rem`, content)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		if matched {
			newContent := replaceRem(content, m)
			log.Println(path + " use rem")
			time.Sleep(1 * time.Second)
			fileWriter, err := os.OpenFile(
				path,
				os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
				0666,
			)
			if err != nil {
				log.Fatalln(err)
			}
			defer fileWriter.Close()
			byteSlice := []byte(newContent)
			_, err = fileWriter.Write(byteSlice)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("----update " + path + " success!\n")
		} else {
			log.Println(path + " never use rem\n")
			time.Sleep(100 * time.Millisecond)
		}
	}

}
