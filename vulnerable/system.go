package vulnerable

import (
	"context"
	"os/exec"
)

func System(ctx context.Context, cmd string) ([]byte, error) {
	return exec.CommandContext(ctx, "sh", "-c", cmd).CombinedOutput()
}
