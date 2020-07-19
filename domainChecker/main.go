package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//TODO add your Full Path
const fullPath = ""

var cmds = []*exec.Cmd{
	exec.Command(fullPath + "synonyms"),
	exec.Command(fullPath + "sprinkle"),
	exec.Command(fullPath + "coolify"),
	exec.Command(fullPath + "domainify"),
	exec.Command(fullPath + "available"),
}

func main() {
	cmds[0].Stdin = os.Stdin
	cmds[len(cmds)-1].Stdout = os.Stdout
	for i := 0; i < len(cmds)-1; i++ {
		currCmd := cmds[i]
		nextCmd := cmds[i+1]
		stdout, err := currCmd.StdoutPipe()
		if err != nil {
			log.Fatalln("Line_25" + err.Error())
		}
		nextCmd.Stdin = stdout
		for _, cmd := range cmds {
			fmt.Println("Running: " + cmd.String())
			if err := cmd.Start(); err != nil {
				log.Fatalln("Line_30: " + err.Error())
			} else {
				defer cmd.Process.Kill()
			}
		}
		for _, cmd := range cmds {
			if err := cmd.Wait(); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
