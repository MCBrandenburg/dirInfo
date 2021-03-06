package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli"
)

type FileData struct {
	Directory          string      `json:"directory"`
	FilePath           string      `json:"filepath"`
	Name               string      `json:"name"`
	Size               int64       `json:"size"`
	Mode               os.FileMode `json:"mode"`
	MachineName        string      `json:"machineName"`
	MD5Hash            string      `json:"md5,omitempty"`
	SHA1Hash           string      `json:"sha1,omitempty"`
	SHA256Hash         string      `json:"sha256,omitempty"`
	Extension          string      `json:"extension,omitempty"`
	LastModified       time.Time   `json:"lastModified"`
	SystemInfo         interface{} `json:"systemInfo,omitempty"`
	Note               string      `json:"note,omitempty"`
	DuplicateFilePaths []string    `json:"duplicates,omitempty"`
}

var dupEnabled bool
var noArray bool
var systemInfo bool
var md5Enabled bool
var note string
var root string
var output string
var sha1Enabled bool
var sha256Enabled bool
var verbose bool

var machineName string

func main() {

	app := cli.NewApp()
	app.Name = "Directory Info"
	app.Usage = "Looks in specified directory or directory executed from to get all file info"
	app.Version = "0.0.7"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "duplicate, d", Usage: "Looks for duplicates. Uses SHA1 hashing if no other are selected.", Destination: &dupEnabled},
		cli.BoolFlag{Name: "info,i", Usage: "Get system info data", Destination: &systemInfo},
		cli.BoolFlag{Name: "md5, m", Usage: "Enable MD5 hashing of files.", Destination: &md5Enabled},
		cli.StringFlag{Name: "note, n", Usage: "Note that will be attached to the data. Example:  '-n working'", Destination: &note},
		cli.BoolFlag{Name: "noArray, na", Usage: "Outputs data in a JSON Array", Destination: &noArray},
		cli.StringFlag{Name: "path, p", Usage: "Directory to start searching", Destination: &root},
		cli.StringFlag{Name: "output, o", Usage: "Filename of the output. Example: '-o something'", Destination: &output},
		cli.BoolFlag{Name: "sha1, s", Usage: "Enable SHA1 hashing of files.", Destination: &sha1Enabled},
		cli.BoolFlag{Name: "sha256, s2", Usage: "Enable SHA256 hashing of files.", Destination: &sha256Enabled},
		cli.BoolFlag{Name: "talkative, t", Usage: "Verbose output for each file", Destination: &verbose},
	}
	app.Action = func(c *cli.Context) error {
		startTime := time.Now()
		if root == "" {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return err
			}
			root = wd
		}

		mn, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			return err
		}
		machineName = mn

		if output == "" {
			output = fmt.Sprintf("%s-%s", machineName, "results")
		}

		if !(md5Enabled || sha256Enabled) {
			sha1Enabled = true
		}

		fileInfo, err := getFileInfo()
		if err != nil {
			fmt.Println(err)
			return err
		}

		if dupEnabled {
			files := getListOfDuplicates(fileInfo)

			for idx := 0; idx < len(fileInfo); idx++ {
				f := &fileInfo[idx]
				dups, ok := files[f.getIdentifier()]
				if ok {
					for _, dup := range dups {
						if dup != f.FilePath {
							f.DuplicateFilePaths = append(f.DuplicateFilePaths, dup)
						}
					}
				}
			}

			fmt.Println(fmt.Sprintf("Found %d Duplicate Items", len(files)))
		}

		if noArray {
			err = writeNoArray(fileInfo)
		} else {
			err = writeArray(fileInfo)
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(fmt.Sprintf("File Read Completed in %s", time.Since(startTime).String()))
		fmt.Println(fmt.Sprintf("Found %d items", len(fileInfo)))
		fmt.Println(fmt.Sprintf("Data written to: %s.json", output))
		return nil

	}

	app.Run(os.Args)
}

func getFileInfo() ([]FileData, error) {

	directoryFileInfo := []FileData{}
	fullPath, err := filepath.Abs(root)

	if err != nil {
		return directoryFileInfo, err
	}
	funkyWalk := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || !f.Mode().IsRegular() {
			return nil
		}

		fd := FileData{
			FilePath:     path,
			Name:         f.Name(),
			Directory:    strings.TrimRight(path, f.Name()),
			Mode:         f.Mode(),
			Size:         f.Size(),
			MachineName:  machineName,
			Extension:    filepath.Ext(f.Name()),
			LastModified: f.ModTime()}

		if systemInfo {
			info := f.Sys()
			if info != nil {
				fd.SystemInfo = info
			}
		}

		if note != "" {
			fd.Note = note
		}

		if sha1Enabled {
			sha1Result, err := getHash(path, sha1.New())
			if err != nil {
				return err
			}
			fd.SHA1Hash = hex.EncodeToString(sha1Result)
		}

		if sha256Enabled {
			sha256Result, err := getHash(path, sha256.New())
			if err != nil {
				return err
			}
			fd.SHA256Hash = hex.EncodeToString(sha256Result)
		}

		if md5Enabled {
			md5Result, err := getHash(path, md5.New())
			if err != nil {
				return err
			}
			fd.MD5Hash = hex.EncodeToString(md5Result)
		}

		directoryFileInfo = append(directoryFileInfo, fd)

		if verbose {
			fmt.Println(fd.Name, fd.getIdentifier())
		}

		return nil
	}
	err = filepath.Walk(fullPath, funkyWalk)
	return directoryFileInfo, err
}

func getHash(filePath string, hash hash.Hash) ([]byte, error) {
	var results []byte

	f, err := os.Open(filePath)
	if err != nil {
		return results, err
	}
	defer f.Close()

	if _, err := io.Copy(hash, f); err != nil {
		return results, err
	}

	bytes := hash.Sum(results)
	return bytes, nil
}

func (fd FileData) getIdentifier() string {
	if md5Enabled {
		return fd.MD5Hash
	}

	if sha256Enabled {
		return fd.SHA256Hash
	}

	return fd.SHA1Hash
}

func getListOfDuplicates(fileInfo []FileData) map[string][]string {
	duplicates := make(map[string][]string)

	for _, f := range fileInfo {
		hash := f.getIdentifier()
		d, ok := duplicates[hash]
		if ok {
			duplicates[hash] = append(d, f.FilePath)
		} else {
			duplicates[hash] = []string{f.FilePath}
		}
	}

	for index, d := range duplicates {
		if len(d) < 2 {
			delete(duplicates, index)
		}
	}

	return duplicates
}

func writeArray(directoryFileInfo []FileData) error {
	data, err := json.Marshal(directoryFileInfo)
	err = ioutil.WriteFile(fmt.Sprintf("%s.json", output), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func writeNoArray(directoryFileInfo []FileData) error {
	newFile, err := os.Create(fmt.Sprintf("%s.json", output))
	if err != nil {
		return err
	}
	defer newFile.Close()

	for idx, fd := range directoryFileInfo {
		if idx != 0 {
			_, err = newFile.WriteString("\n")
			if err != nil {
				return err
			}
		}
		data, err := json.Marshal(fd)
		_, err = newFile.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
