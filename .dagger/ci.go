package main

import (
	"context"

	"dagger/simplelogin-cli/internal/dagger"
)

// Lint commit messages
func (m *SimpleloginCli) LintCommitMsg(
	ctx context.Context,
	args []string,
) (string, error) {
	return dag.Commitlint().
		Lint(m.Worktree, dagger.CommitlintLintOpts{Args: args}).
		Stdout(ctx)
}
