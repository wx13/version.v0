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

// GetFileVersion searches for a file named ".version" or "VERSION" in
// the current directory or any parent directory. If found, it
// returns the first line of this file. If not, it returns an
// empty string. The only input parameter is the max search depth.
// If match search depth is 0, it only looks in the current directory.
func GetFileVersion(depth int) (string, error) {

	// Get the current working directory.
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Start with the current dir, and step up through the parents.
	for i := 0; i <= depth; i++ {
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
func GetCommit() (string, string, error) {

	// Get the current commit hash.
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, _ := cmd.Output()
	commit := strings.Split(string(out), "\n")[0]

	// Get the git status.
	cmd = exec.Command("git", "status", "-s", "-uno")
	out, _ = cmd.Output()
	lines := strings.Split(string(out), "\n")
	status := ""
	if len(lines[0]) > 0 {
		status = "*"
	}

	return commit, status, nil

}

// GetTag gets the most recent version tag in the history.
func GetTagVersion() (string, int, error) {
	// Find all the version tags.
	cmd := exec.Command("git", "tag", "-l", "v[0-9]*")
	out, err := cmd.Output()
	if err != nil {
		return "", 0, err
	}
	tags := strings.Split(string(out), "\n")

	// Find the ancestor tags
	ancestors := []string{}
	for _, tag := range tags {
		cmd := exec.Command("git", "merge-base", "--is-ancestor", tag, "HEAD")
		_, err := cmd.Output()
		if err == nil {
			ancestors = append(ancestors, tag)
		}
	}

	// Find the most recent ancestor
	numCommits := -1
	bestTag := ""
	for _, tag := range ancestors {
		cmd := exec.Command("git", "log", "--oneline", tag+"..HEAD")
		out, err := cmd.Output()
		if err != nil {
			continue
		}
		n := len(strings.Split(string(out), "\n")) - 1
		if numCommits < 0 || n < numCommits {
			numCommits = n
			bestTag = tag
		}
	}
	if numCommits < 0 || bestTag == "" {
		return "", 0, fmt.Errorf("No version tags")
	}
	return bestTag, numCommits, nil
}

// GetVersion returns two version strings (regular and full).
func GetVersion(hash string) (string, string) {

	var version string
	var fullVersion string

	// Get the version from git tags.
	tag_version, distance, err := GetTagVersion()

	// If we got a version from git, then only look for a version file
	// in the current directory (let's the user override git). If we didn't
	// get a version number from git, then look deeper.
	depth := 0
	if err != nil || tag_version == "" {
		depth = 100
	}

	// Look for a version file.
	file_version, err := GetFileVersion(depth)
	if err == nil {
		version = file_version
		fullVersion = version
	} else {
		version = tag_version
		// If we are not on the current tag, append a commit hash.
		if distance > 0 && len(hash) > 8 {
			fullVersion = version + "-" + hash[:8]
		}
	}

	return version, fullVersion
}

// GetTIme returns a string representation of the current time.
func GetTime() (string, error) {
	return time.Now().Format(time.RFC3339), nil
}

func main() {

	// It's self-referential, man.
	version.Print()

	// Get the version information.
	buildTime, _ := GetTime()
	gitHash, gitStatus, _ := GetCommit()
	version, fullVersion := GetVersion(gitHash)
	fullVersion += gitStatus

	// Construct the ldflags string.
	ldflags := fmt.Sprintf(` -X github.com/wx13/version.Date=%s`, buildTime)
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.Commit=%s`, gitHash+gitStatus)
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.Version=%s`, version)
	ldflags += fmt.Sprintf(` -X github.com/wx13/version.FullVersion=%s`, fullVersion)

	// Reconstruct the command line.
	args := []string{}
	if len(os.Args) > 1 {
		args = append(args, os.Args[1])
	}
	args = append(args, "-ldflags")
	args = append(args, ldflags)
	if len(os.Args) > 2 {
		args = append(args, os.Args[2:]...)
	}

	// Run the command.
	cmd := exec.Command("go", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Printf(string(out))

}
