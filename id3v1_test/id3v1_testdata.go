// Code generated by go-bindata.
// sources:
// id3v1_test/sample_ms932_v1.1.mp3
// id3v1_test/sample_ms932_v1.mp3
// id3v1_test/sample_usascii_v1.1.mp3
// id3v1_test/sample_usascii_v1.mp3
// id3v1_test/sample_utf8_v1.1.mp3
// id3v1_test/sample_utf8_v1.mp3
// DO NOT EDIT!

package id3v1_test

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

var _id3v1_testSample_ms932_v11Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x6f\x0e\x69\x9e\xdc\x5c\xd6\xdc\xdd\x74\xa6\x39\xae\xd9\xb9\x39\xbd\xb9\x9b\x01\x09\x20\xc9\x3a\x36\x46\x37\xa7\x36\x3b\x35\x47\x34\xa7\x63\x93\x6d\xee\x6e\xce\x6f\x6e\x40\xd6\x6b\x68\x69\x69\x10\x9c\x98\x5b\x90\x93\xaa\xe0\x92\x58\x92\xa8\x50\x9e\x59\x92\xa1\xe0\x1b\x6c\x69\x6c\xa4\x60\x6c\x62\x6a\x66\xce\xc0\x28\x03\x08\x00\x00\xff\xff\x9d\x69\xa1\x40\x58\x01\x00\x00")

func id3v1_testSample_ms932_v11Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_ms932_v11Mp3,
		"id3v1_test/sample_ms932_v1.1.mp3",
	)
}

func id3v1_testSample_ms932_v11Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_ms932_v11Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_ms932_v1.1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _id3v1_testSample_ms932_v1Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x6f\x0e\x69\x9e\xdc\x5c\xd6\xdc\xdd\x74\xa6\x39\xae\xd9\xb9\x39\xbd\xb9\x5b\x01\x09\x20\xc9\x3a\x36\x46\x37\xa7\x36\x3b\x35\x47\x34\xa7\x63\x93\x6d\xee\x6e\xce\x6f\x6e\x40\xd6\x6b\x68\x69\x69\x10\x9c\x98\x5b\x90\x93\xaa\xe0\x92\x58\x92\xa8\x50\x9e\x59\x92\xa1\xe0\x1b\x6c\x69\x6c\xa4\x60\x6c\x62\x6a\x66\x6e\x61\x29\x03\x08\x00\x00\xff\xff\xfa\x22\x50\xe2\x58\x01\x00\x00")

func id3v1_testSample_ms932_v1Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_ms932_v1Mp3,
		"id3v1_test/sample_ms932_v1.mp3",
	)
}

func id3v1_testSample_ms932_v1Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_ms932_v1Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_ms932_v1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _id3v1_testSample_usascii_v11Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x0f\x4e\xcc\x2d\xc8\x49\x55\x08\xc9\x2c\xc9\x49\x65\xc0\x00\x50\x59\xc7\xa2\x92\xcc\xe2\x12\x9c\xb2\x39\x49\xa5\xb9\x98\x7a\x0d\x2d\x2d\x0d\xa0\x2a\x5c\x12\x4b\x12\x15\xca\x33\x4b\x32\x14\x42\x83\x75\x1d\x83\x9d\x3d\x3d\x15\xcc\xcc\x19\x18\x65\x00\x01\x00\x00\xff\xff\x0b\x7f\x35\x28\x58\x01\x00\x00")

func id3v1_testSample_usascii_v11Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_usascii_v11Mp3,
		"id3v1_test/sample_usascii_v1.1.mp3",
	)
}

func id3v1_testSample_usascii_v11Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_usascii_v11Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_usascii_v1.1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _id3v1_testSample_usascii_v1Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x0f\x4e\xcc\x2d\xc8\x49\x55\x08\xc9\x2c\x01\x92\x18\x00\x2a\xeb\x58\x54\x92\x59\x5c\x82\x53\x36\x27\xa9\x34\x17\x53\xaf\xa1\xa5\xa5\x01\x54\x85\x4b\x62\x49\xa2\x42\x79\x66\x49\x86\x42\x68\xb0\xae\x63\xb0\xb3\xa7\xa7\x82\x99\xb9\x85\xa5\x0c\x20\x00\x00\xff\xff\x2f\x76\xcd\x7e\x58\x01\x00\x00")

func id3v1_testSample_usascii_v1Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_usascii_v1Mp3,
		"id3v1_test/sample_usascii_v1.mp3",
	)
}

func id3v1_testSample_usascii_v1Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_usascii_v1Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_usascii_v1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _id3v1_testSample_utf8_v11Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x7f\xdc\xb4\xf5\x71\xf3\xe6\xc7\xcd\xd3\x1f\x37\xaf\x7e\xdc\xb8\xee\x71\xd3\xfe\xc7\x4d\x4b\x1e\x37\x77\x00\xb9\x40\x33\x31\x64\x17\x3d\x6e\xde\xf3\xb8\xb9\xed\x71\xd3\xe2\xc7\x4d\x3b\xb1\xc9\xae\x7e\xdc\x3c\xe1\x71\xf3\x02\xa0\x5e\x43\x4b\x4b\x83\xe0\xc4\xdc\x82\x9c\x54\x05\x97\xc4\x92\x44\x85\xf2\xcc\x92\x0c\x85\xd0\x10\x37\x0b\x05\x23\x63\x13\x53\x33\x73\x06\x46\x19\x40\x00\x00\x00\xff\xff\x0c\x8a\xd5\x94\x58\x01\x00\x00")

func id3v1_testSample_utf8_v11Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_utf8_v11Mp3,
		"id3v1_test/sample_utf8_v1.1.mp3",
	)
}

func id3v1_testSample_utf8_v11Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_utf8_v11Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_utf8_v1.1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _id3v1_testSample_utf8_v1Mp3 = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xfa\xff\x58\x22\x85\x81\x81\x81\x71\x09\x03\x18\x30\x7b\x80\x48\x1f\x47\x5f\x57\x63\x3d\x4b\x4b\x3d\xd3\x50\x62\x01\xc8\x1c\x63\x74\x73\x88\xd6\x8d\x6a\x4e\x1a\x15\xcc\x09\x71\x74\x7f\xdc\xb4\xf5\x71\xf3\xe6\xc7\xcd\xd3\x1f\x37\xaf\x7e\xdc\xb8\xee\x71\xd3\xfe\xc7\x4d\x4b\x1e\x37\x77\x00\xb9\x0a\x0a\x0a\x18\xb2\x8b\x1e\x37\xef\x79\xdc\xdc\xf6\xb8\x69\xf1\xe3\xa6\x9d\xd8\x64\x57\x3f\x6e\x9e\xf0\xb8\x79\x01\x50\xaf\xa1\xa5\xa5\x41\x70\x62\x6e\x41\x4e\xaa\x82\x4b\x62\x49\xa2\x42\x79\x66\x49\x86\x42\x68\x88\x9b\x85\x82\x91\xb1\x89\xa9\x99\xb9\x85\xa5\x0c\x20\x00\x00\xff\xff\x90\xde\x3d\x7d\x58\x01\x00\x00")

func id3v1_testSample_utf8_v1Mp3Bytes() ([]byte, error) {
	return bindataRead(
		_id3v1_testSample_utf8_v1Mp3,
		"id3v1_test/sample_utf8_v1.mp3",
	)
}

func id3v1_testSample_utf8_v1Mp3() (*asset, error) {
	bytes, err := id3v1_testSample_utf8_v1Mp3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "id3v1_test/sample_utf8_v1.mp3", size: 344, mode: os.FileMode(438), modTime: time.Unix(1468246151, 0)}
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
	"id3v1_test/sample_ms932_v1.1.mp3": id3v1_testSample_ms932_v11Mp3,
	"id3v1_test/sample_ms932_v1.mp3": id3v1_testSample_ms932_v1Mp3,
	"id3v1_test/sample_usascii_v1.1.mp3": id3v1_testSample_usascii_v11Mp3,
	"id3v1_test/sample_usascii_v1.mp3": id3v1_testSample_usascii_v1Mp3,
	"id3v1_test/sample_utf8_v1.1.mp3": id3v1_testSample_utf8_v11Mp3,
	"id3v1_test/sample_utf8_v1.mp3": id3v1_testSample_utf8_v1Mp3,
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
	"id3v1_test": &bintree{nil, map[string]*bintree{
		"sample_ms932_v1.1.mp3": &bintree{id3v1_testSample_ms932_v11Mp3, map[string]*bintree{}},
		"sample_ms932_v1.mp3": &bintree{id3v1_testSample_ms932_v1Mp3, map[string]*bintree{}},
		"sample_usascii_v1.1.mp3": &bintree{id3v1_testSample_usascii_v11Mp3, map[string]*bintree{}},
		"sample_usascii_v1.mp3": &bintree{id3v1_testSample_usascii_v1Mp3, map[string]*bintree{}},
		"sample_utf8_v1.1.mp3": &bintree{id3v1_testSample_utf8_v11Mp3, map[string]*bintree{}},
		"sample_utf8_v1.mp3": &bintree{id3v1_testSample_utf8_v1Mp3, map[string]*bintree{}},
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

