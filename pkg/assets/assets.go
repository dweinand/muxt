package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _config_example_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xce\x4d\x6a\xc4\x30\x0c\x86\xe1\xbd\x4e\x21\xdc\xf5\xd8\x27\xe8\x49\x86\x59\xa8\xb1\x3a\x18\x2c\x2b\xf8\x27\x49\x09\xc9\xd9\xeb\x04\x5a\x5a\x08\xcc\xf2\xb3\x9f\x17\xfb\x0d\x77\x67\xa5\x2d\xd5\xad\xab\xdd\x36\x5b\x55\x22\x40\x22\x61\x7c\x47\x73\x9e\x19\xc8\xaa\xf5\x98\xbb\x33\x00\xf7\xfb\x1c\x92\xd7\xf9\xf1\xf8\x65\xec\x43\xd5\x6c\x20\xd2\x97\xb6\x53\x0a\x85\x74\x9b\x38\xd7\x30\x50\xec\x11\xe2\x4f\x66\x47\x4a\xdc\x5b\xc4\x41\x45\x28\xf9\x83\x4f\x41\x5e\xa3\x67\xa3\xec\xaf\x3f\x50\x38\xf7\xc7\x0c\xfc\xd1\x1f\x2d\xf9\xc8\xc8\x0b\x0f\x98\x29\xc4\x82\xe5\xba\x8d\xfa\x2c\xff\xca\xda\x35\xde\x3e\xb1\x5f\x38\xcf\x13\x47\x1d\x85\x53\xb5\x7d\x1b\xf8\x0e\x00\x00\xff\xff\x0e\xa1\xb4\xd7\x30\x01\x00\x00")

func config_example_toml_bytes() ([]byte, error) {
	return bindata_read(
		_config_example_toml,
		"config/example.toml",
	)
}

func config_example_toml() (*asset, error) {
	bytes, err := config_example_toml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "config/example.toml", size: 304, mode: os.FileMode(420), modTime: time.Unix(1427182408, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _config_test_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4b\xcc\x4d\x55\xb0\x55\x50\x2a\x49\x2d\x2e\xc9\xcc\x4b\x57\xe2\x2a\xca\xcf\x2f\x01\x09\xd4\xe9\x2b\x71\x71\x45\x47\x97\x67\xe6\xa5\xe4\x97\xc7\xc6\x72\xc1\x14\xa6\xa6\x64\x96\xe4\x17\x29\x71\x25\xe7\xe7\xe6\x26\xe6\xa5\x80\x84\xca\x32\x73\x95\xb8\x00\x01\x00\x00\xff\xff\x46\x13\xdb\xf8\x49\x00\x00\x00")

func config_test_toml_bytes() ([]byte, error) {
	return bindata_read(
		_config_test_toml,
		"config/test.toml",
	)
}

func config_test_toml() (*asset, error) {
	bytes, err := config_test_toml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "config/test.toml", size: 73, mode: os.FileMode(420), modTime: time.Unix(1427065897, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _tmux_command_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\xb1\x4e\x03\x31\x10\x44\xfb\xfd\x8a\x21\x50\x40\xe1\x18\x68\x91\xe0\x0f\x28\x68\x28\x10\x85\x39\x6f\x64\x4b\xc9\x9e\xb8\x5d\xdf\x45\xb2\xfc\xef\x18\x88\x38\x8a\x54\xa3\xb1\xbc\xef\x69\x2e\x2f\x7c\xd1\xc9\x7f\x64\xf1\x2c\x33\x34\x11\xd9\xa1\x1c\xa1\x16\x26\x73\xca\xd3\xcc\xd3\xef\x4b\x0a\xda\xbb\x6a\x1e\x05\xce\x50\xeb\xf6\x39\x1c\xb8\x35\xdc\x3f\xfa\xc8\xb3\x97\xb2\xdf\x13\xe5\x1d\xde\xb0\xb9\x7a\xda\xc0\xf1\x27\xee\xf0\xfe\x00\x4b\x2c\x04\x0c\xf1\xfb\xe6\x65\x1c\xad\x35\xea\xfd\x07\x2a\xbc\xac\xd0\x08\xa7\x2b\xb7\xd6\x25\x5b\xc2\x75\x96\xc8\x47\x6c\x5f\x7b\x8e\x0b\x6e\x6f\xba\xd1\xc9\xff\x6f\x5d\x79\xc2\xc2\x0d\xab\xa3\x56\x96\xf8\x17\xb4\xcb\xa7\x65\xae\x20\x98\x85\x21\x9d\x5d\x43\x5f\x01\x00\x00\xff\xff\x2d\xd7\x8a\xa5\x11\x01\x00\x00")

func tmux_command_sh_bytes() ([]byte, error) {
	return bindata_read(
		_tmux_command_sh,
		"tmux/command.sh",
	)
}

func tmux_command_sh() (*asset, error) {
	bytes, err := tmux_command_sh_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "tmux/command.sh", size: 273, mode: os.FileMode(420), modTime: time.Unix(1427182363, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/example.toml": config_example_toml,
	"config/test.toml": config_test_toml,
	"tmux/command.sh": tmux_command_sh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"config": &_bintree_t{nil, map[string]*_bintree_t{
		"example.toml": &_bintree_t{config_example_toml, map[string]*_bintree_t{
		}},
		"test.toml": &_bintree_t{config_test_toml, map[string]*_bintree_t{
		}},
	}},
	"tmux": &_bintree_t{nil, map[string]*_bintree_t{
		"command.sh": &_bintree_t{tmux_command_sh, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

