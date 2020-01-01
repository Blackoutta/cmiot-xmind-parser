package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	maps            []interface{}
	totalCount      int
	contentFileName = "content.json"
)

func main() {
	// check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("使用方法：./xmindparser [pathToFile]...")
		os.Exit(1)
	}

	xmindFilePaths := os.Args[1:]

	for _, v := range xmindFilePaths {
		// extract the content of content.json from .xmind file
		xmindContent := extractContent(v, contentFileName)
		if err := json.Unmarshal(xmindContent, &maps); err != nil {
			log.Fatalf("error unmarshaling: %v\n", err)
		}

		// parse content.json and calculate the total testing checkpoints
		m := maps[0]
		mm := m.(map[string]interface{})
		t := mm["rootTopic"]
		tt := t.(map[string]interface{})
		c := tt["children"]

		// dealing with maps with 0 nodes
		if c == nil {
			log.Printf("The map: %v doesn't have any children node, please populate it first. Skipping this file.\n\n", v)
			continue
		}

		cc := c.(map[string]interface{})
		a := cc["attached"]
		aa := a.([]interface{})
		nextAttached(aa)
	}
	fmt.Println(fmt.Sprintf("总测试点数量: %v", totalCount))
}

func nextAttached(a []interface{}) {
	for _, v := range a {
		aa := v.(map[string]interface{})
		if b := aa["children"]; b != nil {
			bb := b.(map[string]interface{})
			na := bb["attached"]
			naa := na.([]interface{})
			nextAttached(naa)
		} else {
			totalCount++
		}
	}
}

func extractContent(xmindFileName string, contentFileName string) []byte {
	var result []byte
	zipReader, err := zip.OpenReader(xmindFileName)
	defer zipReader.Close()
	HandleError("error opening zip reader", err)

	for _, f := range zipReader.File {
		if f.Name != contentFileName {
			continue
		}

		file, err := f.Open()
		HandleError("error opening file", err)
		defer file.Close()
		buf := bufio.NewReader(file)
		scanner := bufio.NewScanner(buf)
		scanner.Scan()
		result = scanner.Bytes()
		break
	}
	return result
}

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalln(fmt.Sprintf(msg+": %v", err))
	}
}
