package iv

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/darthguinea/golib/log"
)

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Save saves a representation of v to the file at path.
func Save(v interface{}, path *string) error {
	f, err := os.Create(*path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Load (path string) interface - load config file and return interface
func Load(v interface{}, path *string) error {
	data, err := ioutil.ReadFile(*path)
	json.Unmarshal(data, &v)

	return err
}

// ExpandPath - If the path is ~, this function will
// expand it to its full form (/home/<user>)
func ExpandPath(path *string) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if strings.Compare(*path, "~") == 0 {
		*path = dir
	} else if strings.HasPrefix(*path, "~/") {
		pathPtr := *path
		*path = filepath.Join(dir, pathPtr[1:])
	}
}

// CWD - Get the current binary location for this application
// i.e. /usr/local/bin
func CWD() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := path.Dir(ex)
	log.Info("My location [%v]", dir)
	return dir
}
