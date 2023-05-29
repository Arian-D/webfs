package main

import (
	"context"
	"crypto/sha1"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type urlNode struct {
	fs.Inode
	url string
}

// var _ = (fs.NodeOpener)((*urlNode)(nil))
// func (u *urlNode) Open(ctx context.Context, open)

var _ = (fs.NodeLookuper)((*urlNode)(nil))

// TODO: Check if url returns a valid 200-level http response
func (f *urlNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	url := "https://" + name
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, fs.OK
	}
	h := sha1.New()
	h.Write([]byte(url))
	hash, _ := strconv.Atoi(string(h.Sum(nil)))

	fmt.Println(resp.Status)

	stable := fs.StableAttr{
		Mode: fuse.S_IFREG,
		Ino:  uint64(hash),
	}

	node := &urlNode{url: url}

	child := f.NewInode(ctx, node, stable)

	return child, fs.OK
}

var _ = (fs.NodeOpener)((*urlNode)(nil))

// TODO: Implement Open
func (f *urlNode) Open(ctx context.Context, openFlags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	// Disallow write
	if fuseFlags&(syscall.O_RDWR|syscall.O_WRONLY) != 0 {
		return nil, 0, syscall.EROFS
	}
	resp, _ := http.Get(f.url)
	output, _ := ioutil.ReadAll(resp.Body)
	fh = &bytesFileHandle{
		content: []byte(output),
	}
	return fh, fuse.FOPEN_DIRECT_IO, fs.OK
}

type bytesFileHandle struct {
	content []byte
}

var _ = (fs.FileReader)((*bytesFileHandle)(nil))

func (fh *bytesFileHandle) Read(ctx context.Context, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	end := off + int64(len(dest))
	if end > int64(len(fh.content)) {
		end = int64(len(fh.content))
	}

	// We could copy to the `dest` buffer, but since we have a
	// []byte already, return that.
	return fuse.ReadResultData(fh.content[off:end]), 0
}

func main() {
	path := flag.String("path", "/tmp/webfs", "The path where the directory will be mounted")
	flag.Parse()

	opts := &fs.Options{}
	opts.Debug = true
	server, _ := fs.Mount(*path, &urlNode{}, opts)
	server.Wait()
}
