package main

import (
	"os/exec"
	"path/filepath"
	"time"

	"github.com/yutaro/slack-cmd-go"
)

func tpePath(name string) string {
	return filepath.Join(getCurPath(), "./template/"+name+".tpe")
}

func imgPath(name string) string {
	return filepath.Join(getCurPath(), "./image/"+name)
}

func print(c *scmd.Context, tpepath string) {
	options := c.GetOptions()
	flags := c.GetFlags()

	n := "1"
	if num, ok := options["n"]; ok {
		n = num
	}

	cmdMes := tpepath + "," + csvpath + "," + n
	if flags["t"] {
		cmdMes += ", /B -a " + imgpath
	}
	cmd := exec.Command(exepath, "-p", cmdMes)
	cmd.Run()

	if flags["t"] {
		go func() {
			time.Sleep(time.Second)
			c.SendFile(imgpath + "1.bmp")
		}()
	}

}
