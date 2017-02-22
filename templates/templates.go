// Code generated by go-bindata.
// sources:
// templates/swift/enum.tpl
// templates/swift/service.tpl
// templates/swift/struct.tpl
// DO NOT EDIT!

package templates

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

var _templatesSwiftEnumTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\xbd\x4a\xc4\x40\x14\x46\xfb\xfb\x14\x5f\x91\x72\x49\xfa\x85\x54\xb2\x0b\x36\x96\xf6\xe3\xe6\x46\x06\x92\x49\xd8\xcc\x20\xcb\x30\x85\x95\xff\x60\x11\x6c\x6d\x44\x6c\x54\x6c\xfc\x41\xc1\x97\x89\xb3\xdb\xf9\x0a\x32\x63\x04\xcb\x7b\xee\x39\x7c\x59\x46\x59\x06\x6b\x91\xee\x88\x9a\xe1\x5c\xda\x1d\xc8\x52\xd3\x2f\xf7\xf7\x37\xfe\xea\x68\x78\x7f\x5e\xf7\x4f\xd8\x6a\x0a\x9e\x57\x2b\xac\xfb\x6b\x7f\x7c\xf9\xfd\x71\xbe\x79\x7c\x19\x5e\x2f\x36\xb7\x87\xfe\xe4\xec\xeb\xf4\x6e\xf8\x7c\xf0\xfd\x5b\x28\x49\xd6\x6d\xb3\xd4\x98\x37\x46\x15\x42\xcb\x46\x11\xb5\x66\xaf\x92\x0b\xb0\x32\xf5\xff\xb9\x29\xb6\x95\x9e\x60\x16\x31\x11\x00\x8c\xa6\x5e\xb5\x2c\x2a\x29\x3a\xcc\x90\x07\x2b\x3e\xad\xc5\x52\xa8\x7d\x46\x22\x27\x48\x4a\x4c\x73\xa4\x21\x4e\x77\x45\x65\xb8\x83\x73\x51\x5b\x88\x8e\x83\x9b\x94\x7f\x43\xc8\xc7\x3b\x8a\x01\x58\x0b\x56\x45\x28\x1c\xfd\x04\x00\x00\xff\xff\x83\x20\x0d\x86\x07\x01\x00\x00")

func templatesSwiftEnumTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSwiftEnumTpl,
		"templates/swift/enum.tpl",
	)
}

func templatesSwiftEnumTpl() (*asset, error) {
	bytes, err := templatesSwiftEnumTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/swift/enum.tpl", size: 263, mode: os.FileMode(420), modTime: time.Unix(1487071673, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSwiftServiceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x3d\x6f\xd4\x40\x10\xed\xfd\x2b\x46\xd6\x15\x77\xd2\xe1\xa3\xb6\x74\x28\x47\x94\x48\x14\xa0\x28\x41\x50\x9c\xae\xd8\xb3\xc7\x97\x45\xf6\xda\xec\xae\x83\x22\xb3\x05\x15\xdf\x12\x45\x44\x4b\x83\x10\x0d\x20\x1a\x3e\x04\x12\x7f\x26\x5c\xd2\xf1\x17\xd0\xae\xd7\x67\xfb\xcc\x25\xb8\xf2\xae\xdf\xcc\xbc\x79\x6f\xc6\xa3\x11\x14\x05\xf4\x84\x00\x7f\x0c\x1e\x28\xe5\xac\x6e\xbc\x5b\x24\x41\x50\xca\x13\x0f\x68\x24\x9d\xd1\x48\x7f\x5a\x7e\x78\xbb\x7c\xfd\xf8\xf4\xc7\x97\xb3\x93\xcf\xb0\x9d\x86\xb8\x1b\x1f\xc3\xd9\xc9\x9b\xe5\x93\x57\x7f\x7e\xbe\x38\xff\xf4\xf5\xf4\xdb\xcb\xf3\x77\x8f\x96\x4f\x9f\xff\x7e\xf6\xfe\xf4\xd7\xc7\xe5\xc9\x77\x1d\xe9\xd0\x24\x4b\xb9\x84\xdd\x34\x67\x21\x91\x34\x65\x8e\x93\xe5\xf3\x98\x06\x20\x24\xcf\x03\xb9\x56\x13\x0a\x07\x00\xf4\x25\x27\x6c\x81\xd0\xa3\x43\xe8\x25\x9a\xa3\x06\x1d\x20\x3f\xa2\x01\x7a\x37\x51\x1e\xa6\xa1\x00\xa5\x74\x34\x47\x99\x73\x76\xfb\x38\xc3\x15\x6e\x7b\xe2\xe9\xf3\x81\xe4\x94\x2d\xa0\x97\x78\xfb\x35\x46\x29\x53\x62\x2b\xa4\x22\x20\x3c\x24\xf3\x18\xf7\x51\xe4\xb1\x34\xd7\x2b\x72\x44\xd2\x00\xa2\x9c\x05\x15\xc3\xb2\xa8\xe1\xd9\x4b\x40\xa9\x7e\x9b\x64\x64\x8a\x27\xde\x84\x2f\xf2\x04\x99\xac\xd8\x45\x55\x6b\x7e\x95\x68\x9d\x5d\xe4\x59\x5e\x30\x2c\x0a\x64\xa1\x52\x41\x9a\x64\x31\x6a\xb5\x7c\xd8\x42\x11\x90\x4c\x03\xfb\x6b\xcd\x2a\x35\x80\x2b\xd7\xe0\x4e\x4a\xc3\x21\x44\x84\xc6\x39\xc7\x16\x7e\x87\xf3\x94\xaf\x30\xe6\xe5\x7a\x9a\xc6\x50\x38\xa6\x55\xfd\x2c\x72\xc2\x43\x88\x51\x02\xb2\x05\x65\x08\x63\xb8\x8b\xf3\xc9\xde\x8d\x1d\x73\xf4\xec\x2d\xc6\x02\xa1\x80\xb2\x38\x44\x44\x1f\x55\x9d\x45\xc7\x67\x44\x1e\xc2\x18\xdc\xaa\x49\x6b\x95\x6d\x7e\xa4\xaf\x93\xea\x64\x40\xc4\x0a\xb5\x9d\xe6\x4c\x76\xc4\x7b\x08\x31\xb2\xca\x2b\x3b\x12\x34\x02\x86\xeb\x81\x57\x9b\xa0\x23\xc2\x21\x23\x9c\x24\x02\xc6\x30\x2d\x05\xf6\x61\xc2\x8e\x67\xfd\xc1\xe5\x76\xad\xd2\x94\x29\xa6\x6e\xcb\x3f\x77\x06\x63\x68\xdd\x78\xf7\x44\xca\x8a\x02\x90\x85\xa5\xd9\xe5\x4b\x93\xf2\x26\xce\x4d\x58\x88\xf3\x7c\xb1\xc7\x29\x93\x7d\x2d\xe2\x10\xdc\x7d\xbc\x9f\xa3\x90\x3e\xb8\x43\xcb\x45\xb3\x37\x26\x28\xb5\x19\xee\x0e\x61\xea\xcf\x06\x5d\x22\xab\x97\xd2\x4e\x2f\x4b\x45\x19\xec\x43\x99\xa2\x2c\xe2\x5f\x28\x72\x89\xa9\x79\x4c\xfd\xd9\xaa\xd2\x10\x9a\x23\xab\x07\x45\x64\x29\x13\x08\x94\xd5\x53\xb2\xa9\xd7\x12\xaa\xd9\x57\x61\x83\x56\x4c\x4d\xaa\x31\xfc\xae\x1e\x69\xb7\xd9\xa4\x7e\x68\x64\x66\x91\x9b\x7d\xb6\x76\xb5\x36\xa6\xaf\x2d\xf3\xeb\x42\xf6\x7f\xd3\x7c\xea\x4e\xfa\x65\x9e\x36\x1b\x65\x97\xa1\x13\x67\x96\x48\x6f\xdc\xda\x0e\x99\x2d\xf4\x28\x3b\x22\x31\x0d\xab\x6e\x3b\xd1\x5d\x61\x4c\x9c\x56\xc5\x24\x1d\x74\x22\xec\xc6\xf7\xff\xf1\x59\xad\xeb\x67\x2d\xfb\x4f\x2b\xec\x1c\x39\x1b\x54\xe9\xb8\xb3\x36\x6d\xaa\xf1\x37\x2a\xac\x24\x94\x5d\x52\xfb\xa2\x6e\x37\x74\xaa\x06\xf5\x6c\xd9\x3f\x93\xe4\x79\x29\xad\x72\xda\xdc\x94\xf3\x37\x00\x00\xff\xff\x97\xbb\x8c\xa0\xf1\x06\x00\x00")

func templatesSwiftServiceTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSwiftServiceTpl,
		"templates/swift/service.tpl",
	)
}

func templatesSwiftServiceTpl() (*asset, error) {
	bytes, err := templatesSwiftServiceTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/swift/service.tpl", size: 1777, mode: os.FileMode(420), modTime: time.Unix(1487680507, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSwiftStructTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x3f\x4b\xc3\x40\x18\xc6\xf7\x7c\x8a\x87\xd2\xa1\x05\x4d\xf6\x62\x08\xb5\xd0\xd1\xa5\x6e\xa5\x43\x6c\xae\xe5\x24\xbd\x94\xbb\x8b\x52\xc2\x0d\x4e\xfe\x07\x87\xe2\xea\x22\xe2\xa2\xe2\xe2\x1f\x14\xfc\x32\x35\xed\xe6\x57\x90\xb7\x0d\xa9\x0d\x82\x78\x43\x48\xf2\x3e\xef\xf3\x3e\xbf\xf7\x1c\x07\x49\x82\xb2\x52\xa8\xb9\xb0\x61\x8c\x95\xff\xb1\xb7\xfc\x01\x83\x31\xb6\xda\xe7\x3d\x6d\x39\x0e\x95\xd2\xbb\xeb\xf4\xf2\x70\xf2\xf6\x34\x1d\x3f\xa2\x11\x05\xac\x19\x8e\x30\x1d\x5f\xa5\x47\x17\x5f\xef\x67\xb3\x87\xe7\xc9\xcb\xf9\xec\xe6\x20\x3d\x3e\xfd\x3c\xb9\x9d\x7c\xdc\xa7\xe3\x57\xea\xb4\xf8\x60\x18\x49\x8d\x66\x14\x8b\xc0\xd7\x3c\x12\x96\x35\x8c\x77\x42\xde\x85\xd2\x32\xee\xea\xc2\xcc\x1a\x36\x7d\xc5\x90\x58\x00\xa8\x24\x7d\xd1\x67\x28\xf3\x35\x94\x7b\xf3\xa4\x4d\xce\xc2\x40\x51\x5e\x52\x64\x56\x7b\xbe\x9c\xfb\xf4\x96\x36\x99\x6d\xab\x51\xb7\xb7\x47\x43\xd6\xd2\x92\x8b\x3e\x29\xe8\x0b\xc6\x78\x49\x02\x26\x02\x32\xfa\xe9\xc4\x05\xd7\x5e\x65\x57\x45\xa2\x86\xba\x18\x79\xd5\x2c\x0a\x9d\xfc\xa5\x1f\xfb\x32\x40\xc8\x34\x02\xde\xd5\x70\x41\x7a\xf8\xca\x43\x7b\x31\x67\xde\xdb\x01\x0b\x09\x05\x92\xe9\x58\x0a\x08\x1e\xc2\xe4\x1e\x7f\xb3\x65\xaa\x25\x14\xdc\xc5\xbc\x8d\x75\x94\x56\x0a\xa5\x25\x0b\x35\x2d\x9e\xc5\xfd\xe4\x4c\xbf\x11\x91\x20\x63\x59\x41\xa8\x54\xff\x19\x98\x4c\xda\x85\x74\x1d\xb8\xab\x20\x9e\x4d\x61\x0a\x17\x40\x27\xdb\x14\x99\x64\x24\xc6\xfa\x0e\x00\x00\xff\xff\x4e\x3d\x65\x52\xa9\x02\x00\x00")

func templatesSwiftStructTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSwiftStructTpl,
		"templates/swift/struct.tpl",
	)
}

func templatesSwiftStructTpl() (*asset, error) {
	bytes, err := templatesSwiftStructTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/swift/struct.tpl", size: 681, mode: os.FileMode(420), modTime: time.Unix(1487680507, 0)}
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
	"templates/swift/enum.tpl": templatesSwiftEnumTpl,
	"templates/swift/service.tpl": templatesSwiftServiceTpl,
	"templates/swift/struct.tpl": templatesSwiftStructTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"swift": &bintree{nil, map[string]*bintree{
			"enum.tpl": &bintree{templatesSwiftEnumTpl, map[string]*bintree{}},
			"service.tpl": &bintree{templatesSwiftServiceTpl, map[string]*bintree{}},
			"struct.tpl": &bintree{templatesSwiftStructTpl, map[string]*bintree{}},
		}},
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
