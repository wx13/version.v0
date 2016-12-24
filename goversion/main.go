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

// GetVersion searches for a file named ".version" in
// the current directory or any parent directory.  If found,
// it returns the first line of this file.  If not, it
// returns an empty string.
func GetVersion() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Start with the current dir, and step up through the parents.
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

// GetCommit returns git commit and status info.
func GetCommit() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, _ := cmd.Output()
	return string(out), nil
}

// GetTIme returns a string representation of the current time.
func GetTime() (string, error) {
	return time.Now().Format(time.RFC3339), nil
}

func main() {

	// Construct the ldflags string.
	v, _ := GetVersion()
	ldflags := fmt.Sprintf(`-X github.com/wx13/version.Version=%s`, v)
	t, _ := GetTime()
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.Date=%s`, t)
	c, _ := GetCommit()
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.Commit=%s`, c)

	// Reconstruct the command line.
	args := []string{}
	if len(os.Args) > 1 {
		args = append(args, os.Args[1])
	}
	args = append(args, "-ldflags")
	args = append(args, ldflags)
	args = append(args, os.Args[2:]...)

	// Run the build command.
	cmd := exec.Command("go", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}
