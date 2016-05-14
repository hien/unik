// Code generated by go-bindata.
// sources:
// instance-listener/main.go
// DO NOT EDIT!

package bindata

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

var _instanceListenerMainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\xdd\x73\xdb\x36\x12\x7f\xb6\xfe\x0a\x94\x33\xcd\x50\xb1\x42\xb9\x73\x99\xde\x9d\x52\x3f\xb8\x75\x92\x6a\x26\x8e\x75\x72\x32\x7d\x70\x72\x2e\x4c\x42\x36\x2f\x14\xc1\x10\xa0\x1c\x4f\xce\xff\xfb\xfd\x76\x01\x7e\x89\x92\xdb\xb9\xaf\xa9\x1e\x44\x12\xfb\x81\xdd\xc5\xee\xe2\x07\x14\x32\xfe\x24\x6f\x94\x58\xcb\x34\x1f\x8d\xa6\xd3\x77\xe7\xa7\xe7\x33\x7c\x7d\x52\xc2\x54\xa5\x12\x56\x0b\x99\xdd\xc9\x7b\x23\xae\xd3\x9c\x98\x85\xbd\x4d\x8d\x58\xa5\x19\xd3\x30\x98\x48\x2b\xa7\x69\x6e\xac\xcc\x63\x75\x95\xa5\xc6\xaa\x5c\x95\x57\x34\x2c\x74\x2e\x4a\x15\xeb\x75\x01\x76\x28\x97\x49\x22\xa4\x53\x6e\x65\x79\xa3\xec\x68\x94\xae\x0b\x5d\x5a\x11\x8e\x0e\x02\x95\xc7\x3a\x49\xf3\x9b\xe9\x3f\x8c\xce\x03\x1a\x28\x4b\x5d\x1a\x7a\x5b\xad\x2d\x3d\x52\x3d\x4d\x75\x65\xd3\x8c\x3e\x72\x65\xfd\x63\x7a\x6b\x6d\x41\xef\xc6\x96\x90\x67\x09\x73\x9f\xc7\xf4\xcc\xf4\x0d\x3d\x6c\xba\x56\xac\x28\x93\xfc\xad\xc1\x34\x1e\x8d\x62\x0d\xbb\x05\x4c\xb7\x8a\x3d\x3a\x16\xc1\x94\xfd\x69\x86\x22\x67\xcc\xc8\xde\x17\xca\x31\xe2\xbf\xac\x62\x2b\xbe\x8e\x0e\xce\x64\x3c\x2f\xce\x64\x21\xe0\x53\x71\xe9\x66\xff\xe8\x1e\xa2\xf3\xfb\x95\x74\xcc\x82\x79\x61\x82\x5f\x59\xe8\x65\xbe\x21\xa9\x8e\xd0\x50\xde\x0b\x81\x95\xa4\x1e\x46\xa3\x55\x95\xc7\xbc\x4c\xe1\x98\xe6\x46\xfc\x8c\x98\x1d\x0b\x6d\xa2\x13\xbc\x8e\x0e\x56\xba\x14\xe9\x44\x60\x9c\x86\x4b\x99\x63\x55\x99\x09\xcc\x07\x08\x43\xb4\x80\x5e\xbb\x0a\x03\xe2\xf8\x76\x33\x13\xdf\x9a\x60\xe2\x25\xc6\xa3\x83\x87\xd1\x01\x79\xbe\x28\xe1\xf6\x17\xd2\x40\xa1\x8a\x2e\xd8\x98\x30\x28\x78\x18\xfc\x41\x95\xa7\x9f\xae\xe8\xc5\x0d\x09\x9a\x96\xd7\xda\xa8\xdc\x8a\x4d\x2a\x45\x95\x14\x22\x64\x6b\x12\x0c\xa5\xab\x34\x96\x36\x45\x26\x14\x55\x59\x68\xa3\x4c\x80\xd9\x58\xf9\x42\x96\x46\x85\xe3\xdd\xa6\x33\x07\x79\xc6\xee\x6e\xb9\x40\xc4\xbd\x6e\xa4\x2b\xf1\xb4\xe3\xca\xf1\x71\x6d\xf4\x20\x12\x2f\x97\xcb\xf3\x25\xb2\xbd\x42\x12\x14\xa5\xde\xc0\x60\xf1\xcc\xbb\x0a\x5d\x07\xa5\xb2\x55\x99\xef\x53\xba\x57\x9f\x57\x21\x62\x99\xe7\xda\x8a\x6b\x25\x3e\x04\x1f\x82\x81\x46\xca\x9c\x37\x1a\x15\x05\x9f\x29\x5d\xa3\xe5\x2f\x67\x95\x55\x5f\xbe\x82\xa8\x38\x43\xf6\x51\x8d\xdc\xa8\x1e\xad\xa1\x6c\x64\x29\x8c\xcb\x53\xb0\x45\x4d\x82\x1e\x73\xd5\x85\x83\x34\x1b\x7b\x2e\x9f\x91\x43\xb6\x1d\x12\x2e\x51\x26\x02\xe5\x49\x06\xb8\x8a\x8c\x96\x4a\x26\xaf\x50\x30\x61\x53\x3a\x63\x8e\x1a\x71\x7d\x73\x2c\xf2\x34\x1b\x84\x2b\xd6\x55\x96\x08\x0a\x51\x09\xe1\xb6\x0c\x27\xb0\xe2\x1e\x51\xe3\x4e\xc3\xcd\xa6\xc4\x02\x5d\x6b\x6d\x67\x22\x38\x84\xc2\xe8\x25\x75\x86\x70\x4c\xeb\x2d\x54\x66\x14\xab\xf6\x93\xc1\x24\x2a\x9d\xe8\x7d\xbe\x46\x7e\xdd\xca\x2c\x74\xe6\x3e\x31\xe3\x17\xdb\xd6\xf4\x93\x4a\x62\xee\x84\xba\x5a\x41\x89\xe9\xab\x9d\xcb\x70\x30\x2d\x2d\x20\x6a\xf2\xa0\x6e\x77\xf3\xa2\x89\x07\xda\x1a\xd6\x46\x66\xf3\x22\xfc\xcd\x08\xf8\x84\x69\x67\x86\xb0\xc8\x48\x5a\xa4\x05\x52\x7b\x13\xb0\xda\x7e\xe6\x8c\x7a\x2a\x2e\xd0\x49\x2d\x35\x0c\xca\x71\x91\xa4\x26\xd6\x1b\x55\xde\x8b\x90\xca\xf0\x56\x81\x78\xad\x24\xa2\x57\x6a\x99\xc4\xd2\xd8\xb1\xb8\x4b\xed\x2d\xd4\xbb\xc2\x69\x1d\xa8\x6b\x9d\xdc\x4b\xf3\x95\x26\x5f\x2e\x3f\x5e\xdf\x5b\x15\x76\x33\xff\x30\x98\x05\x87\x7b\xc4\xda\xd1\x33\x69\x38\x3b\x3b\x7c\xa7\x6a\x25\xab\xcc\x12\x85\x02\xf3\xe3\xf2\xfc\xe4\xf4\xa7\x93\x8b\x77\x57\xf3\xc5\xe6\x39\xd7\xbd\x82\xe1\x46\x31\x43\x37\xae\x7d\xad\x2e\xa6\x5b\x33\x1d\x3f\x1a\xde\x9a\xf9\x99\x8b\xe9\x8b\xee\xc0\xb3\x35\xe4\xdd\x68\x9b\x8e\x88\x7f\x5c\x65\xb4\xfa\x4d\xd8\x04\xf6\xae\x52\x99\x7e\xc8\x76\xd8\xd6\xad\x70\x83\x12\x55\xb6\x49\x0c\xec\x54\xd1\x69\x2a\xb3\xf7\xa7\x8b\x30\xc0\xea\x3c\x87\x2e\x58\x8d\xcc\x24\x0a\x46\x4f\x30\x05\xb9\x30\x5f\xcc\xb0\x6b\xf4\x03\x34\xc1\xf8\x02\x3b\xe5\x4c\xfc\xf5\x2f\x7f\xfe\x1e\x5f\x0f\xbf\x95\x5c\xd8\x33\xa3\x8b\xa2\x1f\x89\xc6\x9d\x26\x14\x6d\xea\x25\x30\x8d\x7b\x77\xeb\x33\xb6\xc7\x5c\xc5\xd4\xbc\x61\x6a\xdf\x9e\xf1\x96\xaf\x37\x5a\xd0\x06\xe5\xf6\x26\x6e\xe8\x5c\x5f\x57\xce\x7b\xb4\x29\x8e\x45\xf4\x4b\x99\x22\x9f\x28\xbd\x48\x7e\x87\x03\x3b\x17\x70\xaf\xd9\x77\x50\x47\xb9\x4f\x48\xa4\xb1\x9a\x7c\x70\xd3\x6d\xd5\xed\xc0\x07\x9e\xaf\x76\x82\x6b\xba\x3f\x7d\xa3\x13\x73\x44\x51\xc4\x1d\xfc\x80\xa0\x44\x74\x91\x29\x55\x84\xdf\x1d\x1d\x1d\x89\xa7\x82\x47\xce\xd2\x0c\xc9\x00\xb0\x93\x27\x75\x7f\xa0\x34\x5f\xd3\xca\x13\x3e\x89\xde\xaa\xbb\x0b\x55\x6e\xd4\x59\xf5\x85\x09\xd1\xcf\x32\x4f\x32\xf5\x8a\xa2\x16\x4c\x4b\x75\x43\xa9\x54\x22\xd0\x1c\x47\xe4\x9a\x13\x5b\x2a\x53\x00\xa5\x28\x8e\x5c\x39\x41\x95\x7c\x16\x4f\x3d\xe5\x73\xa5\xa8\xa0\x7d\xf3\x03\x25\x3a\x53\xf6\x56\x27\x14\xd0\x60\x71\x7e\xf1\xce\x6d\x51\x70\xd1\xb8\xc8\xff\x8c\x3e\xab\xca\x90\xc5\xd1\x38\x6c\x65\xde\x6a\xfb\x4a\x57\xce\xe6\x36\x14\x0f\x02\x7f\xa6\xc8\x52\x4b\x39\xc9\xdb\x8c\x83\x56\xc8\x28\x0c\x86\x34\xd5\x52\xad\xb5\x55\x44\x07\x18\x98\x71\x6c\xa8\x30\x55\x1e\x36\x82\x63\xf1\x83\xf8\x6e\xd8\x68\xfb\xd2\xb4\xa0\x93\x4e\xf1\xb9\xe6\x5b\x32\x9d\xab\x4e\x40\x4a\xa3\x61\x4d\x19\x29\x02\x4d\x02\x74\x4a\x9f\x94\x7d\x55\x5b\x4e\x90\x45\x1e\x97\xce\x0b\x76\xa2\x36\xec\xf2\xe8\x23\x88\x6b\x19\x9f\xb8\xaa\x76\xed\xe7\x73\xf4\x7e\xf9\x26\xfa\x5b\x85\xf6\x19\x8e\xa3\xd7\xca\x86\x01\x58\xae\xea\xca\x1f\x6f\x75\x97\xb9\xd7\x2d\xea\xb5\x53\xc9\x80\xa7\xe9\xe2\xad\x21\x7b\x59\x5a\x73\x88\x65\x3a\xc5\x77\xdd\x75\x50\x40\xf6\x56\x35\x4a\x80\xac\xf0\x2e\x36\xa6\xb8\xc5\xac\xd3\xcd\xb5\xfe\x02\x89\x7e\x09\xb6\xf8\x22\xa2\xbf\x90\x63\x93\xa8\x95\x42\x3c\x1b\xca\xfb\x3c\x6b\x68\x2d\x5a\xb8\x6c\x2d\xf9\x88\x99\x5b\xd3\x89\x0d\xb3\x10\xfc\x08\xcd\x44\xd4\x30\x84\xf3\x9d\x95\xb4\xb0\x25\x5a\x36\xb3\xba\x49\xbb\xa4\xce\xb4\x18\x9e\x08\xed\x90\x4c\x0b\x44\xba\x16\xb8\xbc\xfa\x06\x3c\xec\x16\x04\x1e\x43\x33\xfd\xd8\x42\xcd\xce\xe0\xf6\x98\xa0\xd1\x33\x75\x2c\x18\x70\xe5\x9a\x5c\x00\xca\xb5\x0c\x79\xeb\xa0\x50\x0a\x16\xd9\x3d\x35\x21\xde\x57\xd5\xba\xb0\xf7\x84\xec\x03\xdf\x05\xfa\x60\x89\x91\xc9\x99\xc7\x25\xd0\x37\x1e\xed\x6c\x82\xbb\x51\x92\xc7\x33\x6c\x07\x2a\x62\x0f\x32\xe9\x17\x40\xbf\xee\xa8\x95\x24\x8d\xad\x64\x99\xc7\xcf\xf4\x4a\xc2\xb4\x69\xbc\xf2\x9b\x06\xd8\x51\xd9\x5d\xf2\xc3\xb0\x6f\x21\x1e\x57\xcd\xe1\x0f\x86\xfd\x41\xfa\xd7\xbf\x57\xdd\xfb\x70\xed\x49\x96\x71\xcb\xfb\x51\x27\xf7\x7b\x56\xec\x11\x13\xe7\x39\x1c\xcf\x65\xc6\xdd\xbf\xe4\x95\x1a\xf7\x44\x42\x0f\xb3\xba\x0b\x39\xf4\xc7\x95\x51\x6d\x46\xf4\x53\xa6\xdd\xe1\x89\x31\x3f\xa5\xc4\xa0\x1c\x7e\x07\x26\xa6\x14\x1c\xa2\xe2\xff\xbd\x33\xbd\x4d\x7e\x47\x59\x0d\x1a\xe9\xbe\x4a\xde\x5d\xc8\xbe\xb2\x3a\x3d\x67\x7f\x37\xea\x34\xa3\x3d\x1d\x08\xfd\x06\xfc\xa3\xbd\xbd\xef\x91\x70\x9d\xc4\xb1\x2a\xac\x4a\x76\x17\x4f\xed\xad\xf9\xef\x55\xcd\xeb\x97\xff\x51\xd1\xb4\x5b\xc3\x76\xff\xee\x50\x3a\x11\xdb\xdf\xdc\xda\xed\xe4\xff\x58\x31\x0f\xdd\xd5\x08\x3b\x5d\xab\x9b\x26\x0e\xb3\x53\x17\xa4\x4b\x09\xc2\x14\x7f\x02\x8e\xa3\x8c\x63\x13\xde\x30\xf9\x24\x4f\x78\xf6\x30\x98\x31\x15\x79\x37\x6e\x6e\x61\xba\x47\x3c\x11\x12\x7a\x9f\x2f\x38\x0a\xb0\x93\x9c\x4b\x57\x32\xa6\xe6\xd9\x01\xfe\xec\x12\x0f\xef\x3e\x15\xba\x55\x70\xac\x8b\xaf\x0f\x5e\x9b\x21\xd4\x48\x8d\x1b\xf5\xac\x36\x64\x32\x18\xee\x74\xf9\x89\x30\x91\xd7\x17\x88\x43\xb1\x75\x26\x76\x37\x2a\xc0\xde\x6c\x48\x7b\xa7\xe2\xec\x1a\x1c\x93\x56\x94\x0b\x42\xe6\xad\x52\x2a\xa3\x0f\x84\xaf\x58\x82\x22\x4b\x6d\xb2\xf5\x88\x87\x23\xaa\x0f\x76\x67\xd7\xf2\x02\x08\x03\x36\x57\xca\xaf\x8a\x37\x48\x7a\x38\xe9\xaf\xa7\x48\xc1\x70\xc7\x43\x59\x14\x74\xec\x80\xbb\xbe\x3b\xfb\xb2\x96\x35\xc6\x33\xd8\xbf\xe2\x5b\xb1\x21\x55\x34\x18\x85\x74\x3f\xe7\x81\x0f\x00\xbb\x12\x4f\x5d\x24\xdf\xe2\x10\xc0\x28\x9f\x20\xc4\x06\x03\xd1\xdc\xbc\xd1\xba\xb8\x96\x94\xc0\xe2\xc9\x13\xe1\x07\x5f\x67\xfa\x1a\x47\xb3\x3c\x25\xb8\xdf\xa1\xbc\xd3\xcf\xf1\xd5\x3b\xa0\xd4\x6b\xd5\xd2\xf9\x1c\xc7\xb4\x87\xfa\x1c\xc1\xf7\x03\x8f\xaf\x6a\x7b\xf4\x5a\xa5\x88\x3f\x4e\xe4\xc8\x47\xb7\x44\xd8\xd6\xb1\xaa\xbd\x13\x9c\x43\x91\x4c\x1d\x73\x26\x4e\xa7\x62\xd9\x9e\x98\x85\x9b\xca\x30\x4e\x44\xc8\x70\xd2\x16\x7a\x25\xe8\x74\xcb\x07\x24\x8c\xce\x17\x0d\x9c\xc4\x5c\x8c\x00\x88\x1c\xb9\x9c\xee\x9e\xbe\x41\xae\x93\x9a\x38\xfc\x07\x9f\x70\xfd\x3b\x85\x22\xe7\x03\x3e\xf0\x7e\x5a\xb8\x94\xce\x29\x4e\x34\xb0\x66\xce\x6e\x5a\x53\x78\x10\x0f\x6c\xa9\x24\xc4\xe8\xad\x9e\x21\xaf\x2f\x00\x89\x72\xf4\x02\xcf\x1f\x44\x8e\xc7\xe1\x21\x6b\x80\xc8\x65\xca\x30\xb4\xa0\xe7\x3f\x45\xf8\x77\x52\x8f\xf7\x71\x37\xc4\x60\x6b\xca\xd3\xf5\x68\x77\x7f\x83\x23\x7a\xe7\x8a\xcc\x57\x67\x9d\xc6\x1e\x2e\xf3\xa2\xb8\xaa\xd8\xde\x27\xb2\xe8\xf7\x35\xbb\xbd\x3d\x8e\x8d\xc3\x70\xdd\x5f\x57\x5b\xf0\x82\x3b\x55\xff\xde\xcc\xc1\xad\x89\x38\xfa\xfe\xf9\xf3\x9d\x9b\xf3\x96\xce\x5e\x8c\xc3\xa1\xc4\xee\x3b\x2e\x0a\x92\xbf\xe2\xe2\xab\x6f\xc6\x79\xdd\xbb\xbb\x07\xc4\xf3\x5f\x01\x00\x00\xff\xff\xd4\x73\x74\x64\x1e\x18\x00\x00")

func instanceListenerMainGoBytes() ([]byte, error) {
	return bindataRead(
		_instanceListenerMainGo,
		"instance-listener/main.go",
	)
}

func instanceListenerMainGo() (*asset, error) {
	bytes, err := instanceListenerMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "instance-listener/main.go", size: 6174, mode: os.FileMode(420), modTime: time.Unix(1463326342, 0)}
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
	"instance-listener/main.go": instanceListenerMainGo,
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
	"instance-listener": &bintree{nil, map[string]*bintree{
		"main.go": &bintree{instanceListenerMainGo, map[string]*bintree{}},
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
