package files

import (
	"os"
	"fmt"
)

func WriteFile(fileName string, content string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully in", fileName)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}