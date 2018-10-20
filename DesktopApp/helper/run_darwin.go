package helper


// run an executable from its App Name
func OpenExecutable(appName string) error {
    if _, err := os.Stat(path); nil != err {
        return err
    }

    cmd := exec.Command("open", "-a", appName)
    return cmd.Run()
}