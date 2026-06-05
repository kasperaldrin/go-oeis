package oeis

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestSequenceIterator(t *testing.T) {
	root := t.TempDir()
	seqRoot := filepath.Join(root, "seq", "A000")
	if err := os.MkdirAll(seqRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(seqRoot, "A000001.seq"), []byte("%I A000001\n%N A000001 test\n%S A000001 1,2,3\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(seqRoot, "A000002.seq"), []byte("%I A000002\n%N A000002 other\n%S A000002 4,5\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	client := &OfflineClient{OfflinePath: root}
	it, err := client.NewSequenceIterator()
	if err != nil {
		t.Fatal(err)
	}

	seq1, err := it.Next()
	if err != nil {
		t.Fatal(err)
	}
	if seq1.ID != "A000001" || seq1.Name != "test" {
		t.Fatalf("first sequence: got ID=%q Name=%q", seq1.ID, seq1.Name)
	}

	seq2, err := it.Next()
	if err != nil {
		t.Fatal(err)
	}
	if seq2.ID != "A000002" {
		t.Fatalf("second sequence: got ID=%q", seq2.ID)
	}

	_, err = it.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}
}

func TestSequencesRange(t *testing.T) {
	root := t.TempDir()
	seqRoot := filepath.Join(root, "seq", "A001")
	if err := os.MkdirAll(seqRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(seqRoot, "A001000.seq"), []byte("%I A001000\n%N A001000 n\n%S A001000 9\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	client := &OfflineClient{OfflinePath: root}
	var ids []string
	for seq, err := range client.Sequences() {
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, seq.ID)
	}
	if len(ids) != 1 || ids[0] != "A001000" {
		t.Fatalf("got ids %v", ids)
	}
}
