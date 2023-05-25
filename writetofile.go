package inkwell

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFile(filename string, content string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	fmt.Fprint(writer, content)

	writer.Flush()

	return nil
}
