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

	"github.com/wx13/version"
)

// GetVersion searches for a file named ".version" or "VERSION" in
// the current directory or any parent directory.  If found, it
// returns the first line of this file.  If not, it returns an
// empty string.
func GetVersion() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Start with the current dir, and step up through the parents.
	for {
		for _, name := range []string{".version", "VERSION"} {
			data, err := ioutil.ReadFile(path.Join(dir, name))
			if err == nil {
				text := string(data)
				v := strings.Split(text, "\n")[0]
				return v, nil
			}
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

	// Get the current commit hash.
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, _ := cmd.Output()
	commit := strings.Split(string(out), "\n")[0]

	// Get the git status.
	cmd = exec.Command("git", "status", "-s", "-uno")
	out, _ = cmd.Output()
	lines := strings.Split(string(out), "\n")
	if len(lines[0]) > 0 {
		commit += "*"
	}

	return commit, nil

}

// GetTIme returns a string representation of the current time.
func GetTime() (string, error) {
	return time.Now().Format(time.RFC3339), nil
}

func main() {

	version.Run()

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

	// Run the command.
	cmd := exec.Command("go", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))

}
