package runner

import (
	"fmt"
	"path/filepath"
	"github.com/brunohgv/adbatch/utils"
	"github.com/brunohgv/adbatch/files"
)

type Process struct {
	Command  string
	DeviceId string
	FilePath string
}

func (p *Process) Run() {
	out, _ := utils.ExecuteCommand(fmt.Sprintf("adb -s %s %s", p.DeviceId, p.Command))
	fmt.Printf("Output for %v:\n\n", p.DeviceId)
	fmt.Println(out)
	filePath := filepath.Join(p.FilePath, fmt.Sprintf("%v-out.txt", p.DeviceId))
	files.WriteFile(filePath, out)
	fmt.Println("================================================")
}