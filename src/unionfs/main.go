package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/hanwen/go-fuse/v2/fs"
)

func newNode(rootData *fs.LoopbackRoot, parent *fs.Inode, name string, st *syscall.Stat_t) fs.InodeEmbedder {
	return &fs.LoopbackNode{
		RootData: rootData,
	}
}

// ExampleLoopbackReuse shows how to build a file system on top of the
// loopback file system.
func main() {
	mntDir := "/home/alex/Desktop/mount"
	origDir := "/home/"

	rootData := &fs.LoopbackRoot{
		NewNode: newNode,
		Path:    origDir,
	}

	sec := time.Second
	opts := &fs.Options{
		AttrTimeout:  &sec,
		EntryTimeout: &sec,
	}

	server, err := fs.Mount(mntDir, newNode(rootData, nil, "", nil), opts)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	fmt.Printf("files under %s cannot be deleted if they are opened", mntDir)
	server.Wait()
	server.Unmount()
}
