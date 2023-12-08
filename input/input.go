package input

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ForDay(day int) string {
	dayFile := filepath.Join(fmt.Sprintf("day%d", day), fmt.Sprintf("day%d.txt", day))
	content, err := os.ReadFile(dayFile)
	if err != nil {
		panic(fmt.Sprintf("%s: error %s", dayFile, err))
	}
	return strings.TrimSpace(string(content))
}

func LinesForDay(day int) []string {
	return strings.Split(ForDay(day), "\n")
}
