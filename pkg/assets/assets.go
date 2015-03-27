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

var _users_dweinand_projects_muxt_pkg_assets_data_config_example_toml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xce\x4d\x6a\xc4\x30\x0c\x86\xe1\xbd\x4e\x21\xdc\xf5\xd8\x27\xe8\x49\x86\x59\xa8\xb1\x3a\x18\x2c\x2b\xf8\x27\x49\x09\xc9\xd9\xeb\x04\x5a\x5a\x08\xcc\xf2\xb3\x9f\x17\xfb\x0d\x77\x67\xa5\x2d\xd5\xad\xab\xdd\x36\x5b\x55\x22\x40\x22\x61\x7c\x47\x73\x9e\x19\xc8\xaa\xf5\x98\xbb\x33\x00\xf7\xfb\x1c\x92\xd7\xf9\xf1\xf8\x65\xec\x43\xd5\x6c\x20\xd2\x97\xb6\x53\x0a\x85\x74\x9b\x38\xd7\x30\x50\xec\x11\xe2\x4f\x66\x47\x4a\xdc\x5b\xc4\x41\x45\x28\xf9\x83\x4f\x41\x5e\xa3\x67\xa3\xec\xaf\x3f\x50\x38\xf7\xc7\x0c\xfc\xd1\x1f\x2d\xf9\xc8\xc8\x0b\x0f\x98\x29\xc4\x82\xe5\xba\x8d\xfa\x2c\xff\xca\xda\x35\xde\x3e\xb1\x5f\x38\xcf\x13\x47\x1d\x85\x53\xb5\x7d\x1b\xf8\x0e\x00\x00\xff\xff\x0e\xa1\xb4\xd7\x30\x01\x00\x00")

func users_dweinand_projects_muxt_pkg_assets_data_config_example_toml_bytes() ([]byte, error) {
	return bindata_read(
		_users_dweinand_projects_muxt_pkg_assets_data_config_example_toml,
		"Users/dweinand/Projects/muxt/pkg/assets/data/config/example.toml",
	)
}

func users_dweinand_projects_muxt_pkg_assets_data_config_example_toml() (*asset, error) {
	bytes, err := users_dweinand_projects_muxt_pkg_assets_data_config_example_toml_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "Users/dweinand/Projects/muxt/pkg/assets/data/config/example.toml", size: 304, mode: os.FileMode(420), modTime: time.Unix(1427182408, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xd1\x6f\x9b\x3e\x10\x7e\xe7\xaf\xb8\x1f\x45\x6a\xa2\xdf\x1c\xd6\x3e\x36\x4a\xf7\xd0\xd7\x6d\xaa\xa6\x4d\x9b\xb4\x4d\x81\x82\x03\xd6\x88\x61\xd8\x84\x74\x94\xff\x7d\x67\x1b\x30\x24\x41\xab\x2a\x95\xe0\xbb\xef\xee\xfb\xbe\x33\x77\xf5\x9f\x5f\x89\xd2\x7f\x62\xdc\xa7\xfc\x00\x22\x75\x9c\x84\xca\xad\xdc\x57\xc7\x6d\x94\xf3\x1d\x4b\x16\x4b\x68\x1c\x00\x1a\xa5\x39\x04\xfa\xbf\x37\x8a\xc2\x0b\x24\x25\x2d\xc0\xbb\xc1\x5f\x51\x25\x81\xec\x6e\x81\xc4\x70\x8d\x7f\x2f\x20\x28\xfe\x12\xbe\xeb\xfb\xd7\x81\xd3\x3a\x0e\xe3\x6c\xa8\xbd\xdf\x87\x3c\x16\xe3\xea\xae\x90\x61\x29\x19\x4f\x40\xa5\x20\xb8\x3c\xd0\xd2\x1d\xa2\xc8\xcb\x06\x4d\x77\x15\x1c\x91\xd9\x04\x06\xa8\xca\x10\x03\xff\xb1\x46\x4d\x79\x4d\xf2\x42\xb2\x9c\x03\x49\x82\xb5\x83\xa0\xa6\xf1\x04\x15\x42\x1d\xdd\x6d\x60\xd5\xb6\x5d\xa1\x53\x6c\x1a\x0a\xd2\x27\x12\x89\xb0\xd5\xc7\x70\x4f\xdb\x16\x6e\xef\xfd\x98\x1e\x7c\x5e\x65\x99\xaa\xc7\x76\xf0\x1d\x5c\xef\x9d\x0b\x84\xfe\x86\x1b\xf8\xb9\x06\x99\x52\x8e\x11\x80\xa7\x50\xd0\x2d\xe3\x31\x3d\x6e\x82\x13\x6f\xc1\x55\x41\xa2\x83\x6e\xe0\xe8\xf4\xcf\x1f\xbe\x7c\xeb\x84\x70\x5a\xdb\xee\x31\x10\x61\x09\x34\x4d\xcd\x64\x0a\x0b\x0d\x85\xd5\x57\x7c\xe6\x35\xbc\x5d\x22\x35\xc2\xc7\x69\xc8\x6c\xf5\x29\xcf\xa5\x0a\x44\x2a\x60\x5e\x9a\x86\xf2\x78\x78\x04\xda\x92\x32\xe4\x09\x05\x8f\xc5\xc7\x37\xe0\xd5\xa6\xa2\x72\xc7\x14\xd7\x1e\x0d\x93\xea\x66\x51\x15\xd0\x25\x7a\x8b\x05\x9a\x8a\xd8\xb6\xfd\xdf\xb3\x92\x97\x4b\x57\xc3\x4c\xd6\x16\xdd\x45\x0b\x36\xd6\xfe\x8e\xe7\xdd\x2c\x5c\xa3\xb5\x8c\x45\x22\x35\x39\xad\x72\xf0\xa7\x6b\xaf\x45\x77\xa4\xc7\xda\xfb\xa3\x7f\x5a\x60\xfb\x0c\xd4\x1e\x4b\xda\x4b\xef\x2e\x24\x8f\xc9\x2f\xfa\x2c\xd4\x55\xf0\x26\x8a\xc0\x1d\x49\x1a\xe1\x5c\x78\x20\xfb\xf3\x0e\x1d\xa9\x07\xf3\x11\xbc\xb2\xfc\x29\xe8\x42\xed\x6e\x82\x85\x19\x61\x11\x72\xaa\x06\xd8\x23\x1f\xf1\x7d\x6e\x8a\x3a\xd7\x0c\xa1\x98\x1d\xe2\x78\x10\xc5\x64\x12\xa2\xc8\x98\xac\x67\x1d\x9f\x9e\x75\x9c\xcf\x65\x8e\xc5\xf4\x56\x29\x62\xaf\x37\x6a\x35\x2f\x41\x7b\x38\xad\x76\xc1\xc1\x99\x49\xbd\x0f\x9f\xf3\x4a\x0e\xfd\x33\x1a\x49\x92\xe9\xb3\x0b\xc3\xb2\x7a\x7b\x58\x5f\xd7\xb1\x2d\x74\x8f\x71\xb9\xfe\x22\x63\x39\x4b\x1b\xb3\x76\x6c\x58\x30\xe4\x0f\xee\x18\xb5\x1f\xdc\xc9\x82\xd1\x55\x48\x05\xa1\x94\x61\x94\x5e\xdc\x57\x6a\x89\x66\x82\x4e\xd2\x05\x6e\x10\x4c\x8f\x32\x46\xb9\x3c\xcd\xc6\xae\xc8\x31\x8a\xed\xe7\xe2\x98\x5d\x82\x97\x5b\xe9\xe9\x9f\x46\xcc\xf9\x5e\x87\xab\x7b\xb0\x1b\xf2\x6f\x00\x00\x00\xff\xff\x0a\x8e\x5a\x36\x66\x06\x00\x00")

func users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh_bytes() ([]byte, error) {
	return bindata_read(
		_users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh,
		"Users/dweinand/Projects/muxt/pkg/assets/data/tmux/command.sh",
	)
}

func users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh() (*asset, error) {
	bytes, err := users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "Users/dweinand/Projects/muxt/pkg/assets/data/tmux/command.sh", size: 1638, mode: os.FileMode(420), modTime: time.Unix(1427365559, 0)}
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
	"Users/dweinand/Projects/muxt/pkg/assets/data/config/example.toml": users_dweinand_projects_muxt_pkg_assets_data_config_example_toml,
	"Users/dweinand/Projects/muxt/pkg/assets/data/tmux/command.sh": users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh,
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
	"Users": &_bintree_t{nil, map[string]*_bintree_t{
		"dweinand": &_bintree_t{nil, map[string]*_bintree_t{
			"Projects": &_bintree_t{nil, map[string]*_bintree_t{
				"muxt": &_bintree_t{nil, map[string]*_bintree_t{
					"pkg": &_bintree_t{nil, map[string]*_bintree_t{
						"assets": &_bintree_t{nil, map[string]*_bintree_t{
							"data": &_bintree_t{nil, map[string]*_bintree_t{
								"config": &_bintree_t{nil, map[string]*_bintree_t{
									"example.toml": &_bintree_t{users_dweinand_projects_muxt_pkg_assets_data_config_example_toml, map[string]*_bintree_t{
									}},
								}},
								"tmux": &_bintree_t{nil, map[string]*_bintree_t{
									"command.sh": &_bintree_t{users_dweinand_projects_muxt_pkg_assets_data_tmux_command_sh, map[string]*_bintree_t{
									}},
								}},
							}},
						}},
					}},
				}},
			}},
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

