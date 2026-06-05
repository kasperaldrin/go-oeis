package oeis

import (
	"io"
	"io/fs"
	"iter"
	"kasperaldrin/oeis/pkg/models"
	"kasperaldrin/oeis/pkg/services"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// SequenceIterator walks every .seq file under {OfflinePath}/seq/{prefix}/.
// Folders are visited in lexicographic order; files within each folder are too.
type SequenceIterator struct {
	offlinePath string
	root        string

	folders  []string
	folderIx int

	files  []string
	fileIx int
}

// NewSequenceIterator creates an iterator over all sequences in the offline dataset.
func (c *OfflineClient) NewSequenceIterator() (*SequenceIterator, error) {
	root := seqRoot(c.OfflinePath)
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var folders []string
	for _, e := range entries {
		if e.IsDir() {
			folders = append(folders, e.Name())
		}
	}
	slices.Sort(folders)

	it := &SequenceIterator{
		offlinePath: c.OfflinePath,
		root:        root,
		folders:     folders,
	}
	if err := it.loadNextFolder(); err != nil && err != io.EOF {
		return nil, err
	}
	return it, nil
}

// Next returns the next parsed sequence, or io.EOF when iteration is complete.
func (it *SequenceIterator) Next() (*models.OEISSequence, error) {
	for {
		if it.fileIx < len(it.files) {
			path := it.files[it.fileIx]
			it.fileIx++
			seq, err := parseSeqFile(path)
			if err != nil {
				return nil, err
			}
			return seq, nil
		}
		if err := it.loadNextFolder(); err != nil {
			if err == io.EOF {
				return nil, io.EOF
			}
			return nil, err
		}
	}
}

// Sequences returns a Go 1.23+ range iterator over all offline sequences.
//
//	for seq, err := range client.Sequences() {
//	    if err != nil { ... }
//	    ...
//	}
func (c *OfflineClient) Sequences() iter.Seq2[*models.OEISSequence, error] {
	return func(yield func(*models.OEISSequence, error) bool) {
		it, err := c.NewSequenceIterator()
		if err != nil {
			yield(nil, err)
			return
		}
		for {
			seq, err := it.Next()
			if err == io.EOF {
				return
			}
			if !yield(seq, err) {
				return
			}
		}
	}
}

func (it *SequenceIterator) loadNextFolder() error {
	it.files = nil
	it.fileIx = 0

	for it.folderIx < len(it.folders) {
		folder := it.folders[it.folderIx]
		it.folderIx++

		dir := filepath.Join(it.root, folder)
		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		var files []string
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".seq") {
				continue
			}
			id := strings.TrimSuffix(e.Name(), ".seq")
			if len(id) < 4 || id[:4] != folder {
				continue
			}
			files = append(files, filepath.Join(dir, e.Name()))
		}
		if len(files) == 0 {
			continue
		}
		slices.Sort(files)
		it.files = files
		return nil
	}
	return io.EOF
}

func parseSeqFile(path string) (*models.OEISSequence, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return services.ParseOEIS(f)
}

// WalkSeqFiles calls fn for every .seq file path in the offline dataset, in order.
// Useful when you only need paths or want to parse sequences yourself.
func WalkSeqFiles(offlinePath string, fn func(path string, id string) error) error {
	root := seqRoot(offlinePath)
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".seq") {
			return nil
		}
		id := strings.TrimSuffix(d.Name(), ".seq")
		dir := filepath.Base(filepath.Dir(path))
		if len(id) < 4 || id[:4] != dir {
			return nil
		}
		return fn(path, id)
	})
}

// NewSequenceIterator creates an iterator over all sequences in the online dataset
func (c *OnlineClient) NewSequenceIterator() (*SequenceIterator, error) {
	return nil, nil
}

// Sequences returns a Go 1.23+ range iterator over all online sequences.
//
//	for seq, err := range client.Sequences() {
//	    if err != nil { ... }
//	    ...
//	}
func (c *OnlineClient) Sequences() iter.Seq2[*models.OEISSequence, error] {
	return func(yield func(*models.OEISSequence, error) bool) {
		it, err := c.NewSequenceIterator()
		if err != nil {
			yield(nil, err)
			return
		}
		for {
			seq, err := it.Next()
			if err == io.EOF {
				return
			}
			if !yield(seq, err) {
				return
			}
		}
	}
}
