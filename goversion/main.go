package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func GetVersion() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		data, err := ioutil.ReadFile(path.Join(dir, ".version"))
		if err == nil {
			text := string(data)
			v := strings.Split(text, "\n")[0]
			return v, nil
		}
		if dir == "/" {
			return "", fmt.Errorf("cannot find .version file")
		}
		dir = filepath.Dir(dir)
	}

	return "", fmt.Errorf("no .version file")

}

func GetCommit() (string, error) {
	return "", nil
}

func GetTime() (string, error) {
	return time.Now().Format(time.RFC3339), nil
}

func main() {

	v, _ := GetVersion()
	ldflags := fmt.Sprintf(`-X github.com/wx13/version.Version=%s`, v)
	t, _ := GetTime()
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.Date=%s`, t)

	args := []string{}
	if len(os.Args) > 1 {
		args = append(args, os.Args[1])
	}
	args = append(args, "-ldflags")
	args = append(args, ldflags)
	args = append(args, os.Args[2:]...)

	cmd := exec.Command("go", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}