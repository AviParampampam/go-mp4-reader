package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	content, err := ioutil.ReadFile("./example.mp4")

	if err != nil {
		log.Fatal(err)
	}

	readVideoFile(content)
}

func readVideoFile(data []byte) {
	lastPosition := 0
	for {
		lastPosition = readAtom(data, lastPosition)
		fmt.Printf("\n\n")
		time.Sleep(time.Second)
		if lastPosition == -13 {
			break
		}
	}
}

func readAtom(data []byte, start int) int {
	if start < len(data) {
		atomSize := bytesInt(data[start : start+4]) // 0..3
		atomType := string(data[start+4 : start+8]) // 4..7
		atomContent := string(data[start+7 : start+atomSize])

		formatAtomContent := formatAtom(atomContent)

		fmt.Printf("Atom:\n Size: %d\n Type: %s\n Content: %s\n Last position: %d",
			atomSize, atomType, formatAtomContent, start)

		if atomType == "moov" {
			fmt.Println(atomContent)
		}
		return start + atomSize
	} else {
		return -13
	}
}

func formatAtom(content string) string {
	if len(content) > 32 {
		return content[:32] + "..."
	}
	return content
}

func trim(b []byte) []byte {
	return bytes.Trim(b, "\x00")
}

func bytesInt(arr []byte) (ret int) {
	for i, b := range arr {
		ret |= int(b) << (8 * int(len(arr)-i-1))
	}
	return ret
}
