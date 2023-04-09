package ports_utils

import (
	"errors"
	"os"

	"github.com/go-git/go-git/v5"
)

type GitRepo struct {
	url  string
	path string
}

func NewGitRepo(url, path string) *GitRepo {
	return &GitRepo{
		url:  url,
		path: path,
	}
}

func (r *GitRepo) Clone() error {
	_, err := git.PlainClone(r.path, false, &git.CloneOptions{
		URL:      r.url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *GitRepo) Exists() bool {
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return false
	}
	_, err := git.PlainOpen(r.path)
	return err == nil
}

func (r *GitRepo) Update() error {
	repo, err := git.PlainOpen(r.path)
	if err != nil {
		return err
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

func (r *GitRepo) Delete() error {
	if _, err := os.Stat(r.path); os.IsNotExist(err) {
		return errors.New("repository doesn't exist")
	}
	return os.RemoveAll(r.path)
}

// Implement the GitRepository interface
func (r *GitRepo) GitClone() error {
	return r.Clone()
}

func (r *GitRepo) GitExists() bool {
	return r.Exists()
}

func (r *GitRepo) GitUpdate() error {
	return r.Update()
}

func (r *GitRepo) GitDelete() error {
	return r.Delete()
}
