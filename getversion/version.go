package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
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

func main() {

	v, _ := GetVersion()
	vs := fmt.Sprintf(`-X github.com/wx13/version.Version=%s `, v)
	ldflags := vs
	args := []string{"build", "-ldflags", ldflags}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("go", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}
