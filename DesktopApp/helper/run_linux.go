package helper


// run an executable from its file path
func OpenExecutable(path string) error {
    if _, err := os.Stat(path); nil != err {
        return err
    }

    cmd := exec.Command(path, "2>&1", "&")
    return cmd.Run()
}