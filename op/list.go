package op

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

type Item struct {
	ID     string
	Title  string
	Vault  string
	Edited string
}

type Items []Item

func List() (Items, error) {
	var list Items

	cmd := exec.Command("op", "item", "list")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return list, err
	}

	scanner := bufio.NewScanner(&out)
	// Skip the header line
	if scanner.Scan() {
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				item := Item{
					ID:     fields[0],
					Title:  fields[1],
					Vault:  fields[2],
					Edited: strings.Join(fields[3:], " "),
				}
				list = append(list, item)
			}
		}
	}

	return list, nil
}
