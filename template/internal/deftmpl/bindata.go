// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x7d\x73\xe2\x36\x1a\xff\xdf\x9f\xe2\xa9\x77\x6e\xba\x99\xe1\x2d\xbb\xed\x4e\x97\x40\x6e\x58\xe2\x6c\x3c\x47\x20\x63\x9c\xdd\xee\x74\x3a\x19\x61\x0b\xd0\xc6\x96\x5c\x49\x86\xd0\x94\xef\x7e\x23\xd9\x80\x0d\x26\x21\x99\x34\xe1\xae\x49\xa6\x5d\x2c\x4b\xcf\xcb\xef\x79\xb5\x25\x72\x7b\x0b\x3e\x1e\x12\x8a\xc1\xbc\xba\x42\x01\xe6\x32\x44\x14\x8d\x30\x37\x61\x3e\x6f\xa9\xeb\xf3\xe4\xfa\xf6\x16\x30\xf5\x61\x3e\x37\xb6\x2e\xb9\x74\x3a\x6a\xd5\xed\x2d\x54\xac\x1b\x89\x39\x45\xc1\xa5\xd3\x81\xf9\xbc\xfa\xa6\xaa\xe7\x89\x7f\x73\xec\x61\x32\xc1\xbc\xa9\x26\x39\xe9\x45\xb2\x26\xa5\x9e\x27\x2f\xe2\xc1\x77\xec\x49\x45\xf6\x37\xb5\xa4\x2f\x91\x8c\x05\xfc\x05\x92\x5d\x46\xd1\x62\x29\x19\x02\xfe\x63\x79\xd3\x1c\x12\x4e\xe8\x48\xad\xa9\xab\x35\x5a\x0b\x51\x39\xd5\xa3\xf0\x17\x04\x98\x66\x39\xfe\x0e\x6a\xd2\x67\xce\xe2\xa8\x83\x06\x38\x10\x95\x3e\xe3\x12\xfb\x17\x88\x70\x51\xf9\x82\x82\x18\x2b\x86\xdf\x19\xa1\x60\x82\xa2\x0a\x09\xcb\x91\x84\xb7\x8a\x56\xa5\xcd\xc2\x90\xd1\x64\xf1\x41\x3a\x96\xa1\x77\x00\xf3\xf9\xdb\xdb\x5b\x98\x12\x39\xce\x4f\xae\x38\x38\x64\x13\x9c\xe7\xde\x45\x21\x16\x29\x8c\x45\xdc\x97\x82\x1f\x2c\x3f\x6d\xb1\x8d\x8f\x85\xc7\x49\x24\x09\xa3\xe6\x1d\x18\x4b\x7c\x23\x13\x3b\x5e\x05\x44\xc8\x74\x2a\x47\x74\x84\xa1\x02\xf3\x79\x22\x57\xdd\x58\x0d\x6e\xe2\xa4\x50\x29\x6b\x20\x95\xf8\xea\xaa\x09\x4b\x05\x52\xc1\x12\xe6\x2d\x4a\x99\x44\x4a\xa6\x1c\xc9\xcc\xf0\xe3\xe8\xf6\x59\xcc\x3d\x5c\x4f\x8c\x89\x29\xe6\x48\x32\x9e\xb8\x9f\x51\x00\x54\x1e\x82\x90\x5f\xfb\x53\x7a\x17\x08\xc6\x9b\x37\x6f\xde\x40\xa2\xf8\x13\x21\x91\x90\xcc\x28\xfe\x94\x70\x24\xc4\x13\x50\x8c\x5d\x31\xc9\x82\x22\x02\xe4\x5d\x57\x7c\x3c\x44\x71\x20\x2b\x92\xc8\x00\xa7\xa0\x48\x1c\x46\x01\x92\xf9\xf8\xac\x6c\x73\xc3\x3c\x9d\x58\xa8\xb4\x10\x16\x91\xca\x27\x9f\x1d\xe9\x0d\x51\x10\x0c\x90\x77\xbd\x41\xaf\x50\x7c\x45\x14\xfe\x82\xfb\x26\x06\x84\x5e\xef\x2c\x41\xc4\xb1\x0a\x20\x73\xb7\xd9\x19\xfa\x77\x02\xa0\x53\xe9\x8e\x12\x10\x8f\x51\x1c\xb2\xef\x64\x47\x19\xd4\xfc\x98\x07\xbb\x4a\xbc\xbb\x72\x43\xc6\x64\x52\x38\x8a\xe3\xcc\x27\x74\x24\x51\x90\xa7\x7d\xe5\x31\x2a\x31\x95\xf7\x39\x97\x76\xe2\x24\x79\x6e\x09\x8e\xc2\x6c\x69\xac\xe5\xea\x5c\x31\x38\x80\x1a\x94\xe7\x73\x23\x19\x84\x64\x50\x67\xa5\xac\x18\x1b\xf9\x71\xad\xa2\x68\x26\xe5\x0c\x3a\x05\xfc\x1c\x2c\x58\x30\xc1\xfe\x1a\xc7\xc5\xf0\xee\x3c\x17\x2b\xd6\xb8\xb6\xb2\xbe\xc3\x83\x0d\x7a\x85\xbe\x95\x25\x70\xa7\x9d\xd2\xf4\xf8\xf8\x24\xb0\x95\x62\xea\x5c\x2a\x57\xbd\x94\xf1\x75\x12\xce\x3a\xc0\x3a\x76\x05\xc5\xe1\x09\x3c\x20\xc3\x76\x71\xeb\x01\x8c\xb7\xb8\xc1\x92\xe8\xca\x15\x8c\xe3\x75\x64\x1f\xe4\x0b\x63\x12\x79\x63\x24\x57\x31\xce\x59\xf8\xf8\xe4\xbd\x4e\x2d\xc4\x42\xa0\xd1\x03\x7c\x2a\x27\x5b\xa4\xb8\xf9\xb1\x9c\x2d\xe9\x6d\x76\x3c\x0f\xf3\xd3\x4d\x8a\x5e\x40\x8a\x93\xd3\x8e\x1a\x6f\xa3\xb8\xea\x95\x1f\x57\x02\x36\xe9\x12\x2a\x24\xa2\x1e\x16\x05\x74\x37\xd3\xc9\x76\x54\x59\x24\x46\x98\x12\xfc\x78\x23\xdd\x45\x6c\xd3\x42\xaf\x49\x7d\x5b\x34\xe7\x05\xd8\x0a\xa9\xd0\x4d\xde\xc3\xbd\x29\x67\xf5\x29\x7e\x4c\x60\xbe\x16\xe5\xbf\xb7\x28\xe7\x4c\x34\x21\x9e\x64\x9c\x45\x62\x65\x79\x89\x24\xbe\xca\xdb\xea\xd5\x1c\x0f\x0b\xa7\x4d\x54\x31\x95\x44\xce\xae\x7c\x22\xa2\x00\xcd\xae\xb6\x3c\xae\xdc\x9f\xfb\x36\x29\x87\x8c\x12\xc9\x14\x20\x57\x92\xb1\xe0\x81\x55\x25\x4b\x1b\x87\x88\x04\x2b\x3f\x58\xbd\x25\x79\xb0\x94\x79\x4a\x63\x19\x6a\xb1\x8c\xc6\x0f\x27\xbd\xb6\xfb\xed\xc2\x02\x35\x04\x17\x97\x9f\x3a\x76\x1b\xcc\x72\xb5\xfa\xf5\x7d\xbb\x5a\x3d\x71\x4f\xe0\xd7\x33\xf7\xbc\x03\x87\x95\x1a\xb8\x1c\x51\x41\x94\xb3\xa1\xa0\x5a\xb5\xba\x26\x98\x63\x29\xa3\x7a\xb5\x3a\x9d\x4e\x2b\xd3\xf7\x15\xc6\x47\x55\xd7\xa9\xde\x28\x5a\x87\x6a\x71\xfa\xb1\x2c\x33\x2b\x2b\xbe\xf4\xcd\x63\xa3\xf1\x43\xb9\x6c\xf4\xe5\x2c\xc0\x80\xa8\x0f\x9a\x89\x8f\x39\x51\x06\x55\xdd\x07\x28\xd2\xa2\x5e\xad\x8e\x88\x1c\xc7\x83\x8a\xc7\xc2\xaa\xd2\x61\x14\xd3\xaa\x26\x87\xbc\x84\x5e\x59\xab\x56\x5e\xc0\x21\x0c\xc3\x70\xc7\x18\xce\x6d\x17\x3a\xc4\xc3\x54\x60\x78\x7b\x6e\xbb\x07\x86\xd1\x66\xd1\x8c\x93\xd1\x58\xc2\x5b\xef\x00\xde\xd5\x0e\x7f\x82\xf3\x84\xa2\x61\x5c\x60\x1e\x12\x21\x08\xa3\x40\x04\x8c\x31\xc7\x83\x19\x8c\x38\xa2\x12\xfb\x25\x18\x72\x8c\x81\x0d\xc1\x1b\x23\x3e\xc2\x25\x90\x0c\x10\x9d\x41\x84\xb9\x60\x14\xd8\x40\x22\x42\x95\xff\x23\xf0\x58\x34\x33\xd8\x10\xe4\x98\x08\x10\x6c\x28\xa7\x88\x27\x1a\x22\x21\x98\x47\x90\xc4\x3e\xf8\xcc\x8b\x43\x4c\x93\xc0\x85\x21\x09\xb0\x80\xb7\x72\x8c\xc1\xec\xa7\x2b\xcc\x03\xcd\xc4\xc7\x28\x30\x08\x05\x75\x6f\x71\x4b\xbf\x60\x62\xb1\x04\x8e\x85\xe4\x44\xa3\x50\x02\x42\xbd\x20\x56\xfd\xf7\xf2\x76\x40\x42\x92\x72\x50\xcb\xb5\xe2\xc2\x90\x0c\x62\x81\x4b\x5a\xce\x12\x84\xcc\x27\x43\xf5\x2f\xd6\x6a\x45\xf1\x20\x20\x62\x5c\x02\x9f\x28\xd2\x83\x58\xe2\x12\x08\x35\xa8\x71\x2c\x29\x3d\xaa\x8c\x83\xc0\x41\x60\x78\x2c\x22\x58\x80\xd6\x75\x25\x9d\x9e\xa3\x44\x8f\x14\xa0\x32\x85\x48\xa8\x91\xe9\x98\x85\x79\x4d\x88\x30\x86\x31\xa7\x44\x8c\xb1\x5e\xe3\x33\x10\x4c\x73\x54\xde\xac\x46\xd4\xf4\x21\x0b\x02\x36\x55\xaa\x79\x8c\xfa\x24\x7d\xa7\xa4\x8d\x8c\x06\x6c\x82\xb5\x2e\x89\x5d\x29\x93\xc4\x4b\xe0\xd6\x06\x88\x56\x56\x4d\x6f\x89\x31\x0a\x02\x18\xe0\x14\x30\xec\x03\xa1\x80\x32\xea\x70\xc5\x5e\xb5\x58\x92\xa0\x00\x22\xc6\x35\xbf\x75\x35\x2b\x86\xe1\x9e\x59\xd0\xef\x9d\xba\x5f\x5b\x8e\x05\x76\x1f\x2e\x9c\xde\x17\xfb\xc4\x3a\x01\xb3\xd5\x07\xbb\x6f\x96\xe0\xab\xed\x9e\xf5\x2e\x5d\xf8\xda\x72\x9c\x56\xd7\xfd\x06\xbd\x53\x68\x75\xbf\xc1\x7f\xec\xee\x49\x09\xac\x5f\x2f\x1c\xab\xdf\x87\x9e\x63\xd8\xe7\x17\x1d\xdb\x3a\x29\x81\xdd\x6d\x77\x2e\x4f\xec\xee\x67\xf8\x74\xe9\x42\xb7\xe7\x42\xc7\x3e\xb7\x5d\xeb\x04\xdc\x1e\x28\x86\x29\x29\xdb\xea\x2b\x62\xe7\x96\xd3\x3e\x6b\x75\xdd\xd6\x27\xbb\x63\xbb\xdf\x4a\xc6\xa9\xed\x76\x15\xcd\xd3\x9e\x03\x2d\xb8\x68\x39\xae\xdd\xbe\xec\xb4\x1c\xb8\xb8\x74\x2e\x7a\x7d\x0b\x5a\xdd\x13\xe8\xf6\xba\x76\xf7\xd4\xb1\xbb\x9f\xad\x73\xab\xeb\x56\xc0\xee\x42\xb7\x07\xd6\x17\xab\xeb\x42\xff\xac\xd5\xe9\x28\x56\x46\xeb\xd2\x3d\xeb\x39\x4a\x3e\x68\xf7\x2e\xbe\x39\xf6\xe7\x33\x17\xce\x7a\x9d\x13\xcb\xe9\xc3\x27\x0b\x3a\x76\xeb\x53\xc7\x4a\x58\x75\xbf\x41\xbb\xd3\xb2\xcf\x4b\x70\xd2\x3a\x6f\x7d\xb6\xf4\xaa\x9e\x7b\x66\x39\x86\x9a\x96\x48\x07\x5f\xcf\x2c\x35\xa4\xf8\xb5\xba\xd0\x6a\xbb\x76\xaf\xab\xd4\x68\xf7\xba\xae\xd3\x6a\xbb\x25\x70\x7b\x8e\xbb\x5c\xfa\xd5\xee\x5b\x25\x68\x39\x76\x5f\x01\x72\xea\xf4\xce\x4b\x86\x82\xb3\x77\xaa\xa6\xd8\x5d\xb5\xae\x6b\x25\x54\x14\xd4\x90\xb3\x48\xcf\xd1\xd7\x97\x7d\x6b\x49\x10\x4e\xac\x56\xc7\xee\x7e\xee\xab\xc5\x4a\xc5\xc5\xe4\x8a\x51\x2e\x1f\x1b\x0d\x9d\x02\x6f\xc2\x80\x8a\x66\x41\x62\x3b\xfc\xf8\xf1\x63\x92\xcf\xcc\xdd\x26\x09\x95\xdc\x9a\xe6\x90\x51\x59\x1e\xa2\x90\x04\xb3\x3a\xfc\x78\x86\x83\x09\x96\xc4\x43\xd0\xc5\x31\xfe\xb1\x04\xcb\x81\x12\xb4\x38\x41\x41\x09\x04\xa2\xa2\x2c\x30\x27\xc3\x23\x18\xb0\x9b\xb2\x20\x7f\xaa\x5a\x0c\x03\xc6\x7d\xcc\xcb\x03\x76\x73\x04\x9a\xa8\x20\x7f\xe2\x3a\x1c\xfe\x14\xdd\x1c\x41\x88\xf8\x88\xd0\x3a\xd4\x8e\x54\x6e\x1d\x63\xe4\xbf\x24\xff\x10\x4b\x04\xaa\xa2\x36\xcd\x09\xc1\x53\x15\x45\x26\xa4\x6f\x80\x9a\xe6\x94\xf8\x72\xdc\xf4\xf1\x84\x78\xb8\xac\x2f\x5e\x0e\x2c\xa8\x2e\xc4\x55\xc6\x2c\xe3\x3f\x62\x32\x69\x9a\xed\x44\xd4\xb2\x3b\x8b\x70\x46\x70\xd5\x8a\x54\x95\x71\x8f\x74\x25\x10\x58\x36\x2f\xdd\xd3\xf2\x2f\x2f\x2c\xbe\x7e\x61\xf3\x72\xe6\xbe\xab\x17\x69\x54\xb5\x70\xc7\x86\xd1\xa8\x2a\xa7\x54\x1f\x06\xcc\x9f\x01\x91\x38\x14\x1e\x8b\x70\xd3\x34\xf5\x85\x9c\xa9\xcf\x69\x44\x09\x6f\x8c\x43\xa4\x23\xca\x52\xd5\xfd\x7c\xd1\xfb\x3e\xab\x92\xe5\x29\x1e\x5c\x13\x59\x4e\x6e\x84\x8c\xc9\xb1\x5e\x94\xd4\x06\x82\x04\xf6\x57\x93\x94\x6f\xe8\xd5\x65\xe4\x7f\x8f\x85\xac\x03\x65\x14\x1f\xc1\x18\xab\xca\x54\x87\xc3\x5a\xed\x5f\x47\x10\x10\x8a\xcb\xcb\xa1\xca\x07\x1c\x1e\x81\x8e\x80\x64\x02\xfc\x40\x42\x15\x2c\x88\xca\x23\x18\x20\xef\x7a\xc4\x59\x4c\xfd\xb2\xc7\x02\xc6\xeb\xf0\x66\xf8\x41\xfd\x66\xe1\x87\x08\xf9\xbe\x96\x4a\x79\xc3\x60\xa4\x67\x36\xcd\x74\xa6\xa9\xf0\x96\x68\xf0\xdc\xee\x91\x51\x69\x47\x3d\x0a\x65\x07\x68\x48\xfe\x82\x79\x0c\x40\x49\xf0\xcc\x99\x74\x82\xb9\x22\x12\x94\x51\x40\x46\xb4\x0e\x92\x45\x79\xa0\x26\xfa\x46\xd3\x94\x2c\x32\x8f\x1b\x55\xe9\xaf\x04\x4d\x32\xab\xf9\xa1\x56\x7b\xe6\x50\x29\x14\x3a\x7d\xb4\xaa\xc3\x20\x60\xde\x75\xce\xb7\x43\x74\x53\x4e\x9d\xe4\x43\xad\x16\xdd\xe4\x6e\x7a\x01\x46\x5c\x31\x94\xe3\xdc\xf8\xb6\x40\x59\x82\x03\x28\x96\x6c\x2d\x24\x72\x68\x69\xa0\x00\x1a\x3e\x99\x3c\xb7\x5b\xe5\xf5\x5d\x07\xe7\x6e\x25\x16\x72\x2b\x23\xeb\x60\x4e\xed\xac\x90\x30\xc1\xc3\x41\x90\xce\x6e\x9a\xb5\xe4\x5a\x44\xc8\x5b\x5c\x3f\xab\xa2\xe9\x4d\x8e\x7c\x12\x8b\x3a\xbc\xd7\x63\x05\x09\x60\x38\xcc\x65\xb1\x64\x59\x1d\x0e\xa3\x1b\x10\x2c\x20\x3e\xbc\xc1\x1f\xd5\x6f\x3e\x31\x0c\x87\x19\x2c\xf6\x21\x3b\xac\x24\x79\xbe\x2c\xf1\x61\x6b\xc0\xe5\xd0\xd5\x4b\xa6\x69\xa9\xf9\xb9\x56\x3b\x02\x5d\xa2\xd2\xf9\x1e\xa6\x12\xf3\x22\x7b\xe9\xff\x6a\xda\x28\x9b\x76\xb3\x3e\xfc\xfc\xee\x5d\xbb\xb8\x00\xbd\x53\x7e\x6d\x42\x1a\x6f\x09\x83\xac\xf5\x92\xb5\xc5\x11\xb9\xf8\x59\x1d\xe4\x58\x9e\xe0\x00\xfd\xb2\xa4\xf0\x5d\xd2\x01\x1c\xc2\x7c\x2e\x96\x2f\x3c\x60\xc8\x38\xac\x76\xd7\xb7\x1c\xf6\x80\xf9\x7c\x8d\x2b\x64\xf7\xda\x9b\xb9\x9d\xf6\x8d\x69\xe9\xab\x95\x9c\xf1\x97\x39\x78\x79\xcd\x5f\xdd\x74\x97\x62\xb6\x72\x9e\xc3\xc4\x79\xee\xf2\x8d\xbd\xcf\x7d\x5b\x61\xdf\x2f\x27\xd8\x77\x57\xa8\x41\x6d\x91\x4b\xee\x72\x87\x54\x0d\x04\x63\x8e\x87\x4d\x73\x97\x97\xee\xcf\xec\x0f\x8b\xa4\x79\x7a\x7a\x9a\x26\x5f\x1f\x7b\x8c\xeb\x77\x72\x8b\xc7\x83\xdc\x03\xc1\x3b\xf5\x38\x90\xcb\xdb\x03\x16\xf8\xc5\x89\xdb\x8b\xb9\x50\xd4\x23\x46\x92\x81\x65\x43\x41\xa8\x26\x9a\xf6\x15\x6b\x09\xfe\x67\x25\x98\xa6\xa7\x5f\xa2\x0e\x19\x0f\xeb\xe0\xa1\x88\x48\x14\x90\x3f\x71\x61\xd2\x7f\xff\xd3\x2f\xd8\x47\x05\xf5\x7a\x63\x46\x3a\xac\x51\xae\x27\x85\x7c\x39\xb8\xec\xde\xa2\x9b\xd4\xbc\xc7\x5f\x08\x9e\x02\xa1\x77\xed\x5c\x2f\x1e\x23\x51\xa1\x0f\xaf\x25\xde\xe2\xf4\x9b\xfc\xdc\xb7\xf9\x51\x50\x14\x5e\x43\xf6\xef\x09\x59\x21\x39\xa3\xa3\x97\x83\xf6\xb7\xed\xc7\x45\x7f\x4f\x77\xbe\x1a\xd5\x44\xc8\x27\xf0\xba\x82\x86\x21\xbd\xb3\x38\x04\xb8\xbe\x85\xf6\xea\x87\xff\x0c\x3f\x4c\x5a\xd3\xa5\xab\x35\x06\x2f\x67\x66\xa8\x16\x63\x74\xcf\x11\xd8\xed\x47\x54\x5f\x58\x99\xed\x71\x07\x05\xb5\x60\xb5\x89\x9e\x54\x82\x17\xf7\x8c\x8c\x44\xfb\xe2\x1e\xf7\x22\x7a\xef\x91\xe6\xff\x51\x67\xc9\x76\x98\xeb\xc7\xab\x5f\xa8\xa1\x5c\xb4\x5b\x1b\x3d\x65\x4c\x7d\xcc\x55\xf7\x97\x77\xa7\xe4\x7c\xb8\x6a\xa2\xf6\x2f\xc7\x3c\xae\x9a\xee\xd8\xde\x65\xcf\x9a\x14\x9a\xf7\xb5\x2b\xdc\x9b\x6a\xbc\x77\x9e\x09\xd0\x18\xef\xa1\x4c\x7b\x87\xd3\x43\x22\xf8\xae\x8e\xf8\x35\xb0\xfe\x3f\xdb\xdc\xec\xe3\xd6\xf2\xcc\xde\xea\x81\x6b\x31\xf4\x02\x8f\x5c\xd9\x13\x84\xaf\xde\xf8\xcf\xf0\xc6\xd7\x87\xae\xd7\x87\xae\xd7\x87\xae\x7d\x77\x96\xd7\x87\xae\xbd\x69\xd9\xb6\x19\xaa\x51\xd5\xfb\x71\xc7\x0f\xd8\x0a\x5d\x2e\x59\x8d\x3c\xfb\x49\x8c\xdc\xd1\xa4\xcc\x49\x93\x95\xa1\x3f\x7e\xfc\x78\xd7\x06\x77\x7e\x67\x77\x73\x4b\x72\x3f\x9a\x86\x7d\x6a\x5f\x9e\xb3\x75\x79\xb7\xb5\x75\x29\xdc\x44\xbb\xcf\xe4\x99\xde\x66\xed\x5c\x43\xfe\x14\x56\x36\x5d\xe5\xff\x28\xc6\xf3\x39\xc4\xbb\x6c\xb6\xd2\x1a\xed\x9c\xaa\x30\x95\x30\x98\xed\xb6\x0f\xb7\x99\x3b\x36\xce\x3b\xac\x67\x86\x46\xd5\x27\x93\xe3\xe4\xff\x46\x3e\x4d\xec\x5b\x5b\xbb\xe5\x78\x5d\xa2\xe2\x2a\x7f\x35\xaa\x03\xe6\xcf\xd4\xc8\x58\x86\xc1\xb1\x61\x14\x7f\x7f\x27\x8a\xc5\x98\x4d\x30\x7f\x82\x3f\xb0\xb0\x41\xea\xef\xff\x3e\xd8\xd3\x7c\x1d\x6c\xf7\x6f\x83\x3d\xdd\x97\xc1\x32\x3c\x77\x40\x72\xf5\x57\x12\x1e\xf0\xb5\xca\xff\x06\x00\x00\xff\xff\x80\x3a\xb2\x92\x52\x47\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 18258, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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

