package loadrepo

import (
	"os"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Clones or pulls a Git repository.
// dir: workspace folder
// url: Git repository URL
// user: Git user login
// pass: Git user password
// return value:
// 0: pull ok, no changes
// 1: pull ok, with changes
// 2: clone ok
// <0: error
func LoadRepo(dir string, url string, user string, pass string) (int, error) {
	auth := http.BasicAuth{
		Username: user,
		Password: pass,
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {  // folder not exists?
		_, err := git.PlainClone(dir, false, &git.CloneOptions{URL: url, Auth: &auth})
		if err != nil {
			return -2, err
		} else {
			return 2, nil // clone ok
		}
	} else {
		r,_ := git.PlainOpen(dir)
		w,_ := r.Worktree()
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Auth: &auth})
		if err != nil {
			if err.Error() == "already up-to-date" {
				return 0, nil // pull ok, no changes
			} else {
				return -1, err
			}
		} else {
			return 1, nil // pull ok, with changes
		}
	}
}

// Clones or pulls a Git repository.
// (no user and password arguments)
// dir: workspace folder
// url: Git repository URL
// return value:
// 0: pull ok, no changes
// 1: pull ok, with changes
// 2: clone ok
// <0: error
func LoadPublicRepo(dir string, url string) (int, error) {
	return LoadRepo(dir, url, "", "")
}
