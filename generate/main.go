package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Working directory is inside the v2 package
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// It's important that protoc uses an exact prefix path to the proto files
	// e.g. the -I argument must be a prefix of the individual proto files
	protoPath := filepath.Clean(
		filepath.Join(wd, "../src/main/proto"),
	)

	protoFilesPath := filepath.Join(protoPath, "feeder/v2")

	var protoFiles []string
	err = filepath.WalkDir(protoFilesPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".proto") {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// If dir does not exist, create it
	if _, err := os.Stat("gen"); os.IsNotExist(err) {
		if err = os.Mkdir("gen", 0755); err != nil {
			log.Fatal(err)
		}
	}

	args := []string{
		"--go_out=.",
		"--go_opt=module=github.com/spacecowboy/feeder-sync-api/v2",
		"--go-grpc_out=.",
		"--go-grpc_opt=module=github.com/spacecowboy/feeder-sync-api/v2",
		"-I", protoPath,
	}

	args = append(args, protoFiles...)

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
