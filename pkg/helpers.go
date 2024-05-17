package pkg

import (
	"fmt"
	"strings"
)

func GetProgressBar(currentValue, finalValue int) string {
	if currentValue < 0 || finalValue <= 0 {
		return "[                                        ] - 0%"
	}
	if currentValue > finalValue {
		currentValue = finalValue
	}
	progress := float64(currentValue) / float64(finalValue)
	percent := int(progress * 100)
	filled := percent / 2
	remaining := 50 - filled
	if remaining < 0 {
		remaining = 0
	}
	if filled < 1 {
		filled = 1
	}
	if percent == 100 {
		return fmt.Sprintf("[%s%s] - %d%%", strings.Repeat("=", filled), strings.Repeat(" ", remaining), percent)
	}
	return fmt.Sprintf("[%s>%s] - %d%%", strings.Repeat("=", filled-1), strings.Repeat(" ", remaining), percent)
}
