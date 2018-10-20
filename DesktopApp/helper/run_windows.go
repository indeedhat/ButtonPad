package helper

import (
    "os"
    "os/exec"
)


// run an executable from its file path
func OpenExecutable(path string) error {
    if _, err := os.Stat(path); nil != err {
        return err
    }

    cmd := exec.Command("start", "/B", "Open Macro", path)
    return cmd.Run()
}