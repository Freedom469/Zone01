package reloaded

import (
	"fmt"
	"os"
)

func WriteToOutputFile(file, s string) error {
	openFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer openFile.Close()

	_, err = openFile.WriteString(s)
	if err != nil {
		return fmt.Errorf("failed to write to file: ", err)
	}

	fmt.Println("Data has been written to the file:", file)
	return nil
}
