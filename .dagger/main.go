package main

import (
	"context"
	"fmt"

	"dagger/simplelogin-cli/internal/dagger"
)

const (
	appName      = "simplelogin-cli"
	appSourceUrl = "github.com/juli3nk/simplelogin-cli"
)

type Git struct {
	Commit        string
	Tag           string
	Uncommitted   bool
	ModifiedFiles []string
}

type SimpleloginCli struct {
	Worktree *dagger.Directory
	Git      *Git
	Files    []*dagger.File
}

func fetchGitInfo(ctx context.Context, source *dagger.Directory) (*Git, error) {
	git := dag.Gitlocal(source)

	commit, err := git.GetLatestCommit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest commit: %w", err)
	}
	tag, err := git.GetLatestTag(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest tag: %w", err)
	}
	uncommitted, err := git.Uncommitted(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check uncommitted changes: %w", err)
	}
	modifiedFiles, err := git.GetModifiedFiles(ctx, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get modified files: %w", err)
	}

	return &Git{
		Commit:        commit,
		Tag:           tag,
		Uncommitted:   uncommitted,
		ModifiedFiles: modifiedFiles,
	}, nil
}

func New(
	ctx context.Context,
	source *dagger.Directory,
) (*SimpleloginCli, error) {
	git, err := fetchGitInfo(ctx, source)
	if err != nil {
		return nil, err
	}

	app := SimpleloginCli{Worktree: source, Git: git}

	return &app, nil
}
