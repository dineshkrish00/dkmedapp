package download

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func LogFileDownload() {
	path := "./example_dir" // Replace this with the user-provided path

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		fileContent, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("Error reading file: %s\n", file.Name())
			continue
		}

		newFilePath := filepath.Join(path, "new_", file.Name())
		err = ioutil.WriteFile(newFilePath, fileContent, 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", file.Name())
			continue
		}

		fmt.Printf("File %s copied to %s\n", file.Name(), newFilePath)
	}
}
