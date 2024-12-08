package util

import (
	"fmt"
	"os/exec"
)

func CopyToClipboard(text string) {
	cmd := exec.Command("wl-copy", text)
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Errorf("failed to copy to clipboard: %w", err))
	}
}
