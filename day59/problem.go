package day59

import "errors"

// File is a collection of bytes representing a file.
type File []byte

// Rsync maintains synchronization between a source and target server.
type Rsync struct {
	files   []File
	targets []*Rsync
}

// AddTarget adds a server as a target.
func (r *Rsync) AddTarget(t *Rsync) {
	r.targets = append(r.targets, t)
}

// AddFile inserts a new file into the server and returns
// a file id integer.
func (r *Rsync) AddFile(file File) int {
	r.files = append(r.files, file)
	for _, t := range r.targets {
		t.sendAddFile(file)
	}
	return len(r.files) - 1
}

// UpdateFile inserts a new file into the server and returns
// a file id integer.
func (r *Rsync) UpdateFile(id int, file File) error {
	if id < 0 || id >= len(r.files) {
		return errors.New("invalid file id")
	}
	// update and send diff
	for _, t := range r.targets {
		t.sendUpdateFile(id, file)
	}
	return nil
}

// DeleteFile removes the file from the source server
// and all the target servers.
func (r *Rsync) DeleteFile(id int) error {
	if id < 0 || id >= len(r.files) {
		return errors.New("invalid file id")
	}
	r.files[id] = nil
	for _, t := range r.targets {
		t.sendDeleteFile(id)
	}
	return nil
}

// GetFile returns the file associated with the id.
// Nil if the id is invalid.
func (r *Rsync) GetFile(id int) File {
	if id < 0 || id >= len(r.files) {
		return nil
	}
	return r.files[id]
}

// rpc calls below this point

func (r *Rsync) sendAddFile(f File) {
	r.AddFile(f)
}

func (r *Rsync) sendDeleteFile(id int) {
	_ = r.DeleteFile(id)
}

func (r *Rsync) sendUpdateFile(id int, file File) {
	// TODO: implement the incremental or changes update
	for i, b := range file {
		if i >= len(r.files[id]) {
			r.files[id] = append(r.files[id], b)
		}
		if r.files[id][i] != b {
			r.sendByte(id, i, b)
		}
	}
	r.files[id] = r.files[id][:len(file)]
}

func (r *Rsync) sendByte(id, index int, b byte) {
	r.files[id][index] = b
}
