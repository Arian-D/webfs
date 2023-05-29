package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

type urlNode struct {
	fs.Inode
	url     string
	content []byte
}

// var _ = (fs.NodeOpener)((*urlNode)(nil))
// func (u *urlNode) Open(ctx context.Context, open)

var _ = (fs.NodeLookuper)((*urlNode)(nil))

// TODO: Check if url returns a valid 200-level http response
func (n *urlNode) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	return nil, fs.OK
}

var _ = (fs.NodeOpener)((*urlNode)(nil))

// TODO: Implement Open
func (f *urlNode) Open(ctx context.Context, openFlags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
	if fuseFlags & (syscall.O_RDWR | syscall.O_WRONLY) != 0 {
		return nil, 0, syscall.EROFS
	}
	return nil, 0, syscall.EROFS
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
	err := os.Mkdir(*path, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
