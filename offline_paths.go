package oeis

import (
	"fmt"
	"path/filepath"
)

// seqFilePath returns the path to a sequence file in the offline dataset.
// Files live at {offlinePath}/seq/{first4OfID}/{ID}.seq (e.g. seq/A000/A000001.seq).
func seqFilePath(offlinePath, id string) (string, error) {
	if len(id) < 4 {
		return "", fmt.Errorf("oeis: sequence ID %q is too short", id)
	}
	folder := id[:4]
	return filepath.Join(offlinePath, "seq", folder, id+".seq"), nil
}

// seqRoot returns the directory containing per-prefix sequence folders.
func seqRoot(offlinePath string) string {
	return filepath.Join(offlinePath, "seq")
}
