package input

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadDay(day int) string {
	dayFile := filepath.Join(fmt.Sprintf("day%d", day), fmt.Sprintf("day%d.txt", day))
	content, err := os.ReadFile(dayFile)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", dayFile, err))
	}
	return strings.TrimSpace(string(content))
}

func ReadAndSplitDay(day int) []string {
	return strings.Split(ReadDay(day), "\n")
}
