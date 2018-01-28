// Code generated by go-bindata.
// sources:
// repository/sql/migration/resources/001_initial_db.sql
// DO NOT EDIT!

package migration

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resources001_initial_dbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x5b\x6f\xe2\x48\x16\x7e\xe7\x57\x94\xd4\x0f\x80\x96\xb4\x5c\x77\xdb\xa3\x91\xd6\x01\xa7\x17\x2d\x31\x19\x30\xad\xe9\x27\x54\x60\x93\x58\xe1\x92\x05\xd3\x33\xf9\xf7\xab\xf2\xb5\x7c\xc5\x99\xee\x19\x4b\xa1\xed\x72\x9d\xdb\x77\x4e\x7d\xe7\xb8\xef\xee\xc0\xbf\x0e\xc1\xf3\x59\x84\x3e\x58\xbd\xf5\x7a\xe3\x85\x6d\xb9\x36\x70\xad\xfb\x99\x0d\xa6\x0f\xc0\x99\xbb\xc0\xfe\x7d\xba\x74\x97\x60\x2b\x42\xff\xf9\x74\x0e\xfc\x0b\x18\xf4\x40\xfa\xf8\xbe\x0e\x3c\x30\x75\xdc\x01\x84\x43\x00\x40\x24\xe0\xac\x66\x33\x00\xac\x95\x3b\x5f\x4f\x9d\xf1\xc2\x7e\xb4\x1d\x77\xd4\x03\xe0\x28\x0e\x3e\x48\xae\x1a\x11\xb9\xe5\x7b\x70\x09\x36\xfb\x78\x97\x3b\x75\xbe\x45\xdb\x86\x8a\xd6\x89\xfd\x60\xad\x66\x2e\xe8\xc3\xbe\xdc\x2f\xb6\x61\xf0\x3d\x51\xda\xbe\x5f\x8b\xf6\x07\x97\xf5\xf9\x74\x0a\x3b\xeb\x7f\x5a\x4c\x1f\xad\xc5\x37\xf0\x5f\xfb\x1b\x18\x28\x21\x0f\x7b\xc3\x1e\x00\xb6\xf3\x65\xea\xd8\xe0\x57\x30\x3d\x1e\x4f\x93\xfb\x5e\x2e\x3e\xfe\x8f\xb5\x58\xda\x2e\xf8\x15\x5c\xc3\x9d\x7e\xd8\x90\x5f\xba\x60\xfb\xbe\x7e\x13\x67\xff\x18\xae\x0f\xe2\x2d\x03\xf9\xf4\x1c\x83\x5c\x02\xae\x0d\xe8\x44\x4b\x5d\x8a\x54\xb0\x2b\xc1\x25\xa6\x46\x35\x0a\x7e\x72\xc0\xe1\x59\x1c\x2f\x32\x79\xa7\x63\x5c\x4e\x6f\xe2\xfd\x20\x4d\x16\xab\xa9\x58\x51\xea\x55\x0d\xfa\x7a\xf1\xcf\x09\x4e\x0d\x1a\x22\x68\xce\x27\xef\xba\xad\xb5\x93\x96\xa0\xd8\x5f\x93\x8a\x9a\xd8\xe3\xe9\xa3\x35\x1b\x40\x3a\x02\xa8\x88\x5d\x28\x9e\x53\xc1\xaf\xd6\x42\x46\x3f\x40\x94\x0e\x41\xd3\x95\xe2\x94\xca\x0b\xcf\xf3\xbd\xb5\x08\xa3\x3a\x7c\xb4\x97\xae\xf5\xf8\xd4\x16\x6f\x86\xf3\x6a\xb1\xb0\x1d\x77\x9d\x09\x45\xa1\xbf\x79\x22\x8c\xd5\x95\x95\x55\x14\xd5\x39\x53\x28\x84\x3c\x13\x43\xf9\x6e\xea\x4c\xec\xdf\xc1\x20\x41\xf7\x27\x97\x41\x92\x8d\xa4\x04\xf2\xd4\x34\xe7\xf0\x56\x15\xa8\x1c\x53\x49\x8d\x9a\xc0\x2f\xee\xd4\xc9\x37\x46\xbb\x20\x69\xce\x5f\x1d\x6c\x6f\xe7\x60\xeb\x2b\x6f\x1b\x8b\x45\x49\xb6\xdc\x68\xb9\xb6\x4c\x53\x6b\x64\x5d\xf3\x5d\xd5\xd7\xcd\x77\xcf\xdf\xfb\x3f\xa8\x22\xb8\xac\x15\xba\x56\xd8\xf4\x76\x54\x09\xbb\xfe\xef\x2a\x8e\x61\x10\xbe\xd7\x02\xd8\xcd\x89\x54\xc5\xfa\x7a\x0c\xc2\xbf\x76\x14\x8b\xd5\x9f\x15\xe1\xdf\x53\xea\x39\xa7\xa6\x1c\xaf\x96\x7d\x17\x72\xff\x10\xab\xe7\xca\x47\xe0\xef\x63\xf3\x4b\x78\xda\xbe\x46\xc1\x44\x77\xf1\x09\xce\x0f\x70\x6b\xab\xaa\xe3\x63\x35\x9c\x66\x5a\x6f\xad\x87\x5a\x07\x92\xb3\x98\x1e\xc6\x9c\x2b\xab\xfd\xbf\xf6\xdc\x15\xa0\x4d\x43\xfd\x61\x2c\x65\x84\x97\xde\xa0\x12\x6b\x09\xad\x9e\x5c\x56\x3c\x28\xb1\x5d\x5a\xfa\xd0\x80\x4d\xa5\xaf\xc2\xe0\x1f\x44\xb0\xef\x2e\x99\x48\x6d\xcf\x7e\x4a\x3c\x19\x65\x34\x22\x56\x4c\x63\x4e\x59\x1d\xc8\xa6\x86\xa5\x3e\x20\xb5\xd9\x9f\xb6\xaf\xb5\xed\xb0\x4d\x2a\xb8\xac\x85\x77\x08\x8e\x45\x2e\x53\xe6\xc7\x46\x2c\xc7\x73\x67\xe9\x2e\x2c\x99\xb2\x28\x97\x6b\x99\x96\xf5\x35\x38\x7a\xfe\x9f\x3d\x00\x56\xce\xf4\xb7\x95\x0d\x06\x72\xb5\xa9\x58\xea\x8a\xa4\xed\xc4\x45\x95\x22\xae\xe1\x4b\xdc\x3a\xb3\xc7\xca\x5c\xd3\x76\xf6\x94\x72\x03\xf5\x52\x72\xd7\xc1\x0f\x5f\x4e\xe9\xf8\x59\x20\xd8\xc2\xe0\x9e\x4f\x4d\x31\x7e\xf7\xb3\xf9\x7d\xe5\xcc\xa8\x7e\xfe\x95\x73\x63\xcd\x5c\x7b\x91\x00\xa2\x8e\x90\x3d\x00\xac\xc9\x44\x4d\xc4\xee\x75\xad\x6e\x58\x47\x89\x01\x0f\xf3\x85\x3d\xfd\xe2\x28\xce\x04\xde\x10\x2c\xec\x07\x7b\x61\x3b\x63\x3b\xc6\xf5\xa2\x4c\x3b\x00\xcc\x1d\x30\xb1\x67\xb6\x6b\x03\x67\x0e\xac\xb1\x3b\x9d\x3b\xf1\xea\xea\x49\x96\x64\xbe\x3a\xba\xed\x44\x36\xee\x14\xfc\x50\x5a\x8e\xea\x4a\x3e\x1b\xa9\x3d\xe9\x03\x0e\x95\xf0\xaa\x6b\x40\x55\x97\xeb\x76\xad\x83\xcd\xee\x75\x0d\x8b\x5e\xab\xdd\x44\x75\x5b\xfd\x4c\x2c\x74\x1c\xd5\xf3\xb1\xb5\x1c\x5b\x13\xbb\xe0\x77\xb2\x56\x03\x63\xb3\x4f\xe8\x27\x21\xd9\xe2\x4f\x09\xc5\x88\xf5\xab\x2e\x26\xcd\xa0\x06\xa7\x8f\x56\x99\x2c\x79\x79\xa8\x0a\xbe\xa4\x8b\x35\xe0\x28\x96\xff\x79\x34\x72\x1a\xaa\x2d\xfe\xfc\x75\x74\xfb\x63\xc0\xb4\x7b\xf5\x09\x4c\x9d\xa5\xbd\x70\x25\x8f\xcd\xd3\xaf\xc9\xec\x63\x66\x3a\x19\x45\xba\xe5\xbf\xc9\xd2\x57\x49\x58\xd9\x93\x2b\x9e\xf3\xfb\xe0\xe0\x5b\x72\x48\x28\xac\x8c\x5f\xc4\xf1\xd9\xf7\x86\xe0\xab\x35\x5b\xd9\xcb\xde\x27\x00\xc0\x00\x8e\x00\xd2\x46\xa0\x0f\x3f\x23\xdc\x1f\x81\xbe\xfc\x43\x1a\xe4\x77\x10\xdd\x41\x08\x20\x36\x29\x36\xb1\xde\x1f\x45\xd9\x1b\x8e\x62\x29\x9c\x48\x19\xea\xf5\xd9\x30\x1a\x34\x70\x53\xc3\x25\x0d\xe4\xc3\x1a\x48\x49\x03\xfd\x61\x0d\xec\xc3\x1a\x68\x49\x03\x4f\x34\x20\x88\x09\xf9\xac\x69\x8d\x92\x48\x2b\x49\xea\x1d\x24\xa9\x09\xb9\x09\x79\x49\xd2\xe8\x24\x89\xa0\xa9\xc1\x92\x24\xd4\xba\x89\x22\x93\x96\xa1\x82\xb0\x9b\x28\x36\x71\x39\x52\x88\x3a\x8b\x56\x1c\xc6\x9d\x45\xcb\xf5\x05\x49\x37\x51\x6a\xe2\x32\xc0\x90\x76\x16\x35\xca\xa2\xac\xab\x28\xad\x38\x9c\xd6\x12\x44\xb8\x59\x90\xd5\x64\x55\x57\x04\x31\x6a\x14\x44\xa8\x2c\x98\x56\xd2\x1d\xbc\x21\x4a\xcb\x39\x45\xda\x2d\x51\x0a\x20\x32\xa1\x61\x12\x96\x8a\xfe\xd2\xfb\x54\xe6\xb8\x98\xbd\x33\x1a\x8f\xb8\x2d\xbe\x75\xc4\xc1\xcf\x1e\xbe\xb8\x53\x27\x7b\x78\x3a\x07\xdb\xfc\x95\xca\x74\xf9\x4a\xc2\x74\x85\xb5\x49\x3c\x88\x67\x6b\x5f\xe3\x0f\xff\x11\x78\x8a\x9f\x7f\x4b\x3e\xbb\x2a\x0b\xab\x63\x10\x0e\x7b\x9f\x0a\xac\x89\x47\xa0\x3f\xde\x5f\x37\x77\x8f\x22\xf4\x81\x36\xa2\x7b\xf0\xb0\x17\x97\xed\x8b\x2f\x41\x20\x1a\x32\x38\x23\x9a\x06\x75\x4d\x16\x57\x1f\x7e\x8e\x10\x2c\xa0\x43\x4c\xa8\x99\x30\x25\xd6\xf4\x17\x8e\x40\x5f\x4b\x76\xef\x83\xd0\x3f\xf7\x15\xc2\xbc\x65\x12\xe1\xc4\x68\x8b\x4d\x68\x6a\xac\xb3\x4d\xda\xd1\x26\xc2\x51\x55\x36\xdb\xcc\x6a\xef\xb6\x4d\xd6\xc1\x66\x6a\xb4\xd5\x26\x21\x9d\x6d\xea\x5d\x6d\xb6\xc7\x89\x3e\x82\xad\xd1\xd9\x26\x6c\x8b\x13\x99\xb0\x7b\x9c\xb2\x01\xdc\x30\x9a\xd9\x6c\x2b\x5c\x6e\xe2\x36\xa3\x7d\xb5\x6b\x74\x08\xf2\x96\x3d\xa4\xe7\x0c\xd4\x12\x64\x95\x60\x56\x4b\x7b\x11\xcf\x62\xe9\xf8\x14\xf3\x8a\xbc\xb3\x0f\x22\xd8\xc7\xb7\x11\x65\xc4\xff\x2f\xa0\x2c\xa4\x1c\x92\x2e\x64\x04\x92\x2e\xdc\xc7\x1f\xe9\xc9\x5c\x76\xb1\xe4\xa7\x77\x99\x28\x64\x2f\xe9\x3f\x8a\xb3\x1f\x48\x47\x0f\xf2\xe6\xdf\x9b\xe0\x28\xce\xef\x77\xaf\x41\xb8\x7d\xf1\x8f\x9f\x3d\x3f\x8f\x18\xde\x21\x08\x20\x33\x35\xcd\x44\x46\x29\xe2\xf8\x57\x53\xfb\x45\xdf\x0d\x0e\x27\x29\x1d\x06\x87\x53\x47\xbd\x04\xdf\xd0\xab\xc7\x8e\x9e\xbe\xfb\x48\x8b\xfb\x50\xa4\x5d\xfc\xe1\xef\xe4\x5f\x59\x2f\x97\x3d\x82\x68\xa6\x56\xa6\xb2\x24\x4f\x6a\xcb\x48\xf4\x96\x9b\x85\xa6\xcb\x29\x09\x42\x13\xc2\x06\x1d\xd5\xd4\x4a\xcc\xad\x6b\xf8\x12\xa7\x57\xde\xa9\x13\x72\xba\xf6\x18\x7d\xf9\xe7\xcf\xd1\xc8\x5c\x1c\x80\x8d\x11\x88\x92\xf4\x26\x2e\x97\x3f\x4e\x67\xaf\x1f\xfb\x0b\xb4\x3f\x11\xc1\x88\x41\x44\x30\xc4\x1a\x22\x5c\xe7\x94\x6c\xb8\xc1\x29\x87\x4c\x27\x9c\x61\xe4\x73\x48\x18\xdb\x61\x4a\x18\x83\xc4\xe3\x88\x62\x8a\xa9\x40\x3e\xd5\x19\x8a\xde\x13\x4e\x19\x23\x94\x50\xca\x31\x61\x94\xf8\x8c\x53\x8d\x6d\xb1\x41\x04\x15\x64\x8b\x75\x6e\x60\x0d\x23\xaa\x31\x4a\x04\xe6\x84\x60\x4a\xb1\x7a\x5a\xa3\x2c\xdf\x72\x8d\x62\x8e\x89\x27\x1d\xc3\x94\x10\x4a\x88\x4e\x19\x33\x38\x27\x98\xed\xb0\xc1\x3c\x2e\xb0\x74\x7e\xc7\x0d\xe4\x4b\x37\x39\xc1\x3a\x45\x0c\x53\x88\xa1\xf4\x1d\x6b\x84\x92\x1d\xe3\x14\x51\x46\x74\xae\xb3\x0d\xdb\x11\xcc\x29\x26\x98\x32\xca\x04\x31\x38\xe2\x06\xdb\x62\xa6\x1e\xec\xa8\x50\x6e\xb9\x46\x74\xe6\x71\x4e\x21\xf1\x22\x34\x74\xb2\x65\x04\xed\x28\x22\x8c\xe8\xc8\xa7\x06\xf1\x08\xe1\x90\x31\xcc\x90\x4f\x74\x8e\x99\x8e\x21\xf1\xb0\xc1\x0c\x8c\x09\xe5\x84\x19\x84\x12\xce\x38\xe6\x64\x8b\x31\xe1\x64\xc3\x45\xf4\x4e\x67\x90\x6d\xb1\x4e\x25\xc2\x1e\xc6\x94\x57\x66\xc6\xbd\x27\xde\xb2\xd9\xe3\xff\x01\x00\x00\xff\xff\x5a\x98\x3a\xfd\x26\x1e\x00\x00")

func resources001_initial_dbSqlBytes() ([]byte, error) {
	return bindataRead(
		_resources001_initial_dbSql,
		"resources/001_initial_db.sql",
	)
}

func resources001_initial_dbSql() (*asset, error) {
	bytes, err := resources001_initial_dbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/001_initial_db.sql", size: 7718, mode: os.FileMode(436), modTime: time.Unix(1517160941, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"resources/001_initial_db.sql": resources001_initial_dbSql,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"resources": &bintree{nil, map[string]*bintree{
		"001_initial_db.sql": &bintree{resources001_initial_dbSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

