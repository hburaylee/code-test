// Code generated by go-bindata.
// sources:
// resources/conf.xml
// resources/init.xml
// DO NOT EDIT!

package main

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

var _resourcesConfXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xb1\xaf\xc8\xcd\x51\x28\x4b\x2d\x2a\xce\xcc\xcf\xb3\x55\x32\xd4\x33\x50\x52\x48\xcd\x4b\xce\x4f\xc9\xcc\x4b\xb7\x55\x0a\x0d\x71\xd3\xb5\x50\xb2\xb7\xe3\xb2\x49\xce\xcf\x4b\xb3\xe3\x52\x50\xb0\x49\xca\x4f\xa9\xb4\xf3\x48\xcd\xc9\xc9\x57\x28\xcf\x2f\xca\x49\x51\xb4\xd1\x07\x0b\x71\xd9\xe8\x43\xd4\x00\x02\x00\x00\xff\xff\xbc\x07\xc1\x14\x52\x00\x00\x00")

func resourcesConfXmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesConfXml,
		"resources/conf.xml",
	)
}

func resourcesConfXml() (*asset, error) {
	bytes, err := resourcesConfXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/conf.xml", size: 82, mode: os.FileMode(420), modTime: time.Unix(1586005532, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesInitXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xb1\xaf\xc8\xcd\x51\x28\x4b\x2d\x2a\xce\xcc\xcf\xb3\x55\x32\xd4\x33\x50\x52\x48\xcd\x4b\xce\x4f\xc9\xcc\x4b\xb7\x55\x0a\x0d\x71\xd3\xb5\x50\xb2\xb7\xe3\xb2\xc9\xcc\xcb\x2c\xb1\xe3\x52\x50\xb0\x49\xca\x4f\xa9\xb4\xf3\xc8\x2f\x57\x48\x2c\x4a\x55\xa8\xcc\x2f\x55\xb4\xd1\x07\x0b\x71\xd9\xe8\x43\xd4\x00\x02\x00\x00\xff\xff\x99\x0f\xed\xbc\x52\x00\x00\x00")

func resourcesInitXmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesInitXml,
		"resources/init.xml",
	)
}

func resourcesInitXml() (*asset, error) {
	bytes, err := resourcesInitXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/init.xml", size: 82, mode: os.FileMode(420), modTime: time.Unix(1586005523, 0)}
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
	"resources/conf.xml": resourcesConfXml,
	"resources/init.xml": resourcesInitXml,
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
		"conf.xml": &bintree{resourcesConfXml, map[string]*bintree{}},
		"init.xml": &bintree{resourcesInitXml, map[string]*bintree{}},
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

