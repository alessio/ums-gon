package createdmg

import (
	"context"
	"os/exec"
	"path/filepath"
)

// Cmd returns an *exec.Cmd that has the Path prepopulated to execute the
// create-dmg script. You MUST call Close on this command when you're done.
func Cmd(ctx context.Context) (*exec.Cmd, error) {
	fname, err := exec.LookPath("create_dmg")
	if err != nil {
		return nil, err
	}

	return exec.CommandContext(ctx, fname), nil
}

// Close cleans up the temporary resources associated with the command.
// This Cmd should've been returned by Cmd otherwise we may delete unrelated
// data.
func Close(cmd *exec.Cmd) error {
	// Protect against unset commands
	if cmd == nil || cmd.Path == "" || filepath.Base(cmd.Path) == cmd.Path {
		return nil
	}
	return nil
}
