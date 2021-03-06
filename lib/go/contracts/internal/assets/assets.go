// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleToken.cdc (7.43kB)
// ../../../contracts/FungibleToken.cdc (7.199kB)
// ../../../contracts/TokenForwarding.cdc (2.361kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _exampletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x4b\x6f\x1b\x37\x10\xbe\xfb\x57\x4c\x73\x68\x65\xc4\x96\x5d\xa0\xe8\xc1\x70\x1e\x76\x63\x17\x41\x5b\xb7\xc8\xa3\xbd\x9a\xbb\x3b\x92\x58\xef\x92\x02\xc9\x95\xac\x18\xfe\xef\x05\x87\xe4\x2e\x49\xed\xca\x76\xdc\xe4\x92\x48\xe2\x3c\x38\xfc\xe6\x9b\x8f\x0c\x6f\x96\x52\x19\xb8\x6c\xc5\x9c\x17\x35\x7e\x92\x37\x28\x60\xa6\x64\x03\xc7\xb7\x97\x9f\xaf\x7e\x7d\x7f\xfe\xfb\xc5\xa7\x3f\x7f\xbb\xb8\x3a\x7b\xf7\xee\xc3\xc5\xc7\x8f\x7b\x7b\xcb\xb6\x80\x52\x0a\xa3\x58\x69\xe0\xe2\x96\x35\x4b\x6f\x76\x92\x79\xb9\xdb\xdb\x03\x00\x38\x3a\x3a\x82\x4f\xd2\xb0\x1a\x74\xbb\x5c\xd6\x1b\x90\xb3\xc4\x4c\x03\x17\x80\xb7\x5c\x1b\x14\x25\x92\x89\x0d\xb1\x62\x0a\x8c\x35\xfb\x48\x56\x27\xf0\xf9\x92\xdf\xfe\xfc\x53\xec\xd3\x1a\xbf\x17\xdc\x70\x56\xf3\x2f\x58\x85\x5f\xfa\x15\x0b\x04\x5c\xa1\x30\x60\x16\xcc\x00\xd7\x80\x0d\x37\x06\x2b\x58\x2f\x50\x80\x59\x60\xbf\x11\xae\xa1\x54\xc8\x8c\x77\x63\x33\x70\xa6\x5b\x61\x26\xdc\xfd\x3b\xcd\x6b\x3f\x4f\xec\x1f\x6e\x16\x95\x62\x6b\xf1\xe4\xb4\x5c\x51\x98\x42\x58\x07\x1f\xee\x44\x18\xfc\xcd\xda\xda\x0c\x26\xd8\x85\x9b\xb0\x46\xb6\xc2\x84\xbc\x0e\xc8\xf4\x04\xce\xaa\x4a\xa1\xd6\x6f\xb6\xf2\x7c\x87\x4b\xa9\xb9\xf9\x8a\xf2\xf5\x79\x56\xc1\x07\x18\xb9\x33\xcb\x2e\xd8\x56\x96\x46\xee\xc8\xf1\x0f\x2e\xbe\x22\x41\x81\xeb\x38\xc9\xa6\x77\x92\xa7\xe5\xfc\x67\x39\x6d\x65\x71\xde\x2a\xf1\xcc\x32\x69\xa3\xe4\x66\x24\x09\xe7\x7e\x3c\x09\x4a\x52\xfd\x12\x81\xf4\x09\x59\x30\xaa\x06\x95\x40\x81\x42\x2d\x5b\x55\xe2\x38\xe8\x93\x58\x13\x56\xd7\x72\x8d\xd5\xd9\x58\x66\x94\xf9\xf3\x32\x2b\xc8\xc5\x23\x32\x4b\x62\x4d\xa2\x24\x7a\xd0\xc5\xc1\x2f\x58\xb9\x80\x56\xa3\x02\x6d\xa4\x42\x0d\x4c\x00\x17\xda\x30\x51\xa2\x25\x22\x29\xea\x0d\x11\x01\x99\x5b\x26\x32\x0b\xe4\x6e\x35\x9b\x63\xb2\x89\x59\x2b\x4a\xc3\xa5\x23\xac\xde\x86\x89\x0a\xe6\x72\x85\xf6\xf4\xa0\x70\xde\x96\x0a\xe9\xfb\xa5\xd4\xc6\x72\x4c\xc5\xc9\xb0\x73\xc7\x45\xc6\x95\x81\x90\x36\x04\x94\x92\xd5\x35\x56\xd3\x24\x7a\xb9\xc0\xf2\x46\xc3\x82\x2d\x97\xb6\x6a\x06\x54\x2b\x0c\x6f\x90\x4c\x71\x85\x0a\x58\x97\x21\x95\x2f\xf5\xd1\xf9\xfa\xe0\x4b\x6c\x57\x08\xb7\xff\x02\x43\xb1\xc3\xce\x2c\x2d\xe2\xad\xb1\x15\x4a\x58\x92\x4e\xd0\xa6\xd9\xb9\x73\xb8\x9e\x71\x41\xc6\x07\xa0\xa5\xfd\x5d\xd1\x09\x0a\x09\x6b\xb6\x81\x99\xb4\xb9\x35\xac\xe6\x25\x97\xad\x76\xc7\x61\xa4\x8f\xe9\xaa\xd8\x97\x46\xb6\x3e\x2c\x17\xc0\xb8\x9a\xc2\x19\xe8\x25\x96\x9c\xd5\x1e\x95\x3d\x48\x04\x62\xa5\xad\xa7\xa2\xcf\xc1\x48\x42\x79\xe7\xae\x27\x81\xb4\x14\x16\x51\x9d\x23\x4a\x21\x1b\x5f\xd3\xbf\x94\x5c\xf1\x0a\xd5\x41\xf6\xfd\x07\x2c\x91\xaf\xb6\xbf\x3f\x67\x35\xa1\xca\x8f\xbd\xf8\xec\x68\x8e\x41\xe1\x17\xb8\xdd\x69\x58\x75\x88\x8d\x67\x9e\x5f\x95\xce\x3b\xe7\x0c\x78\x37\x85\xe8\x58\x82\x43\x0b\x86\xb0\x15\x2a\xaa\x85\x80\xc5\x46\x67\x6b\x0d\x27\x99\xe7\x7d\xb8\xeb\x7e\xb7\x7f\x34\xd6\xb3\x69\x70\xf9\x2a\x38\xef\x96\xdc\xa7\xdb\x0a\xa3\x29\xfe\x32\x59\x70\x19\xb0\xe8\x30\xc3\x6e\x5c\xf3\x39\x7a\x03\xe6\x3e\xa8\x79\xdb\xa0\x30\x89\xa1\xed\x9b\xe0\x5d\x3b\x6b\x6f\x44\x43\xb0\x6b\xbc\xe9\x68\xe8\xf7\xc6\x63\x4b\x7b\x76\x31\x68\x55\x0e\x53\x1b\xdf\xb2\x81\x88\x5a\xed\x10\xb3\x90\x75\x95\x78\xb0\x41\x1a\x29\x70\xd3\x2d\x2d\x90\x8b\x39\x18\xc5\x84\x9e\xa1\x52\x58\x4d\x6d\x18\x85\xa6\x55\x42\xd3\x7a\x81\xeb\x7a\x93\x78\x09\x4d\xe5\x83\xca\xa4\xb5\xc8\xb1\x6b\x52\xdb\x34\xdc\x50\x3f\x16\xd1\x30\x4d\x7c\x61\xad\x71\x6d\x1b\x6b\x78\xdb\x16\x3d\xb3\x56\x74\x85\xcb\xc7\xc8\x09\xbc\x4d\xd1\xea\x72\xda\x89\x80\xe4\xe3\xa1\x3f\x84\xc4\xc0\x12\xf9\xa8\xfe\x70\x7f\x07\xfd\x41\xce\xe4\x5a\xa0\x7a\x33\x65\x6e\xce\xef\x27\xbe\x5c\x29\xe1\xf4\x30\xa6\x85\x1e\xb3\xce\xdb\xfe\x18\x1c\x7d\xd1\x9e\x86\x46\x7f\x30\xb2\xf8\x17\xcb\x1c\x92\x04\x43\x56\x55\x3a\x71\xc3\x8d\xee\xba\xce\x9f\x67\xd2\xd5\x08\xb4\x45\xfd\x08\x84\x72\x0d\x7e\xae\x5a\x4f\x5e\x1a\x90\x0b\x6d\xc3\xbb\xd4\x0a\x2c\x59\xab\xb1\x07\x7d\xda\x83\x36\xe5\x08\xdc\x16\xc6\xa8\x42\x26\x9e\xf5\x88\x80\xc8\xf6\x87\x3e\xf7\x05\x4b\xf7\x55\x20\x0a\x8b\x4c\xdd\x36\x58\xd1\xd6\x89\xc4\x67\x92\x86\x91\x87\xa5\x17\x2f\xbb\x01\xe8\x0f\x62\xe2\x4e\x7d\x08\x74\x39\xef\xd4\x68\x1c\x15\xc2\xe9\xa1\xd7\xb9\xfa\x3b\x78\x1b\x5f\x11\xa6\xe9\xde\x1f\xc2\xea\x4b\xe7\x6f\x9a\x53\x58\x06\xd9\x6d\x31\x9a\x98\x39\x4d\xfa\x20\x6e\x13\x1b\x78\x05\xc7\xd3\xe3\xe4\xf7\x70\xb2\x29\xdb\x47\xf0\xf5\x0b\x26\x79\x5d\x92\x02\x44\x37\x21\x78\x35\xfe\xd3\x61\x52\x88\x28\x5a\x14\xb3\x67\xa6\x8b\x66\x69\x36\xc3\xba\x29\xed\x97\x94\x4d\x1d\x34\x2d\xd3\x00\x8b\xe1\xff\x05\x95\xec\x75\x81\xa8\x3a\x76\xe4\x3d\xf9\xb1\xba\xb6\x3c\xea\x49\xd0\x0e\x77\x52\x03\x4d\xab\x1d\x19\xba\xc9\x18\x74\x4c\xe2\x8d\x04\x1c\x79\x71\x7e\x3b\x62\xcd\x45\x9b\xfd\x42\xaa\xca\x89\x0c\xea\x31\xf7\x7b\xef\xad\x2c\x69\x9e\x38\xe5\xc0\x8a\x9a\x9a\x59\xb9\xb9\x1e\x10\xac\xbb\x39\x4d\x8d\x04\x66\xb3\xc4\x6d\x09\x61\x21\x9f\x17\x73\x62\xd9\x36\xe7\xd7\x07\xe8\xed\x78\x7a\xbc\x1f\x1f\x52\x22\x4f\xce\xaa\x86\x0b\xae\x8d\x62\x46\xaa\x5c\x5f\x38\x7f\x57\xb8\x76\xea\xe8\x91\x0c\xd8\x9d\x68\x74\x4c\x83\x77\x84\x9d\xcd\x9e\xc5\x1e\xb9\x28\x9c\xc0\x5b\xaf\xdc\xee\xb6\x5b\x71\xe7\x4d\x23\xf9\xb8\x7b\x5c\x0c\x67\x30\xe2\xe0\x7e\xa4\x84\xee\x72\xf1\xec\x12\x66\x97\x99\xc7\x95\xd0\xc5\x26\xec\xb8\x7f\x0e\x55\x2b\xbf\xfd\xec\xaa\x48\x70\x38\xce\x02\x11\x62\x86\x6e\x08\x61\x30\xba\x91\x49\x4d\xc0\x2c\x12\x43\xff\xb8\x1b\x84\x1d\x3a\x41\x75\x3f\x4e\x6d\x77\x60\xd8\xd2\xc9\x5e\xe7\xd9\xc6\x73\xb7\xe6\x70\xdf\x08\xa8\x4c\x87\x66\x27\xf4\x21\xd2\xcf\x83\x18\x4c\x43\x59\x3b\x37\x02\x1e\x79\xd4\xd6\x40\x47\x9b\x3b\x20\x65\x60\x13\x6b\x02\xb3\x99\xe8\x85\xeb\x60\x4b\xcf\x46\x3a\xb1\x19\xe3\xc2\x9d\x30\xe9\x53\x1e\x50\x76\xdb\xa3\x32\xc3\x8e\xbd\x8e\xa6\xdf\xd8\x3f\xbe\xda\xaf\x2d\xf9\x9c\xc0\x0b\x57\x31\xff\x4c\xe2\x18\xb9\x40\x98\x13\x98\x94\xad\x83\x20\x86\x7f\x31\xe6\xe7\xd4\x4f\xe1\xec\x00\x46\xfc\xd6\xa8\xb5\x73\x6a\x6b\x11\x0e\xd5\xb9\x4a\x43\xdc\x3f\x7b\x2a\xbe\x1c\xd2\xae\xdb\xb9\xc2\xd0\x06\x1e\x14\xbe\xd9\xdb\x51\xae\x53\xe1\x59\xd2\x96\xee\x6d\xc3\xac\x3a\xa4\xdd\xf3\xed\x24\x9f\xc7\x79\x20\xa2\xbd\xe7\xf3\x80\x25\xbf\x87\x39\xa0\xa3\xb8\x54\x86\xb6\x4a\x3c\xa9\x31\xbd\x76\xea\xb5\x7c\x78\xdb\x39\x00\x9c\xcd\xb0\x34\x7c\x85\xf5\x86\xfc\xd2\xf5\xad\x97\xc5\xa3\x01\xae\xa4\xc1\x13\xa7\xec\x9d\xc8\x88\x1e\xf0\x58\x6b\x64\xc3\x0c\xb7\xad\xbb\x01\xdd\x16\xf4\x2a\x82\x55\x77\x33\x4d\x2f\x91\xf1\xa3\x77\xf2\x64\x44\x69\xb7\xa5\x91\x6a\x77\xd7\xf7\xf5\xf8\xe6\x7a\xda\x5a\xb1\x80\x9b\x71\xf9\x3c\xac\x66\xb3\x96\xc8\x5e\x32\xb7\xf1\x1d\xe1\x8f\x10\x1e\x6f\x81\x80\x9c\x36\xf6\x8f\xc7\xc7\x56\x55\xc7\x2f\x21\x6e\x04\x46\x25\x25\x39\x9a\x13\x71\x34\x49\x48\x41\xb2\x15\x5a\x31\xca\x45\xf2\xbe\xe7\x5c\xee\x0d\xd6\x6f\xb8\x63\xf3\x1c\xf7\xd3\xec\x7d\x5f\x4c\x6d\xbc\xc9\xe9\x21\x39\x73\xb7\x89\x23\x1f\xf7\x08\xa3\xd3\x70\xa7\x38\xb4\x3d\x66\x91\x50\xf3\x12\x4a\xb6\x64\x05\xaf\xb9\xd9\x84\xe9\x41\x6a\xb8\x8a\xdf\x33\xe8\x29\x0f\x6f\x97\x52\x63\xdc\x3e\xb4\xfa\xda\x8b\xda\x6b\x68\xd0\x2c\xa4\xbd\xde\x29\xd9\xce\x5d\xc5\xae\xc3\x8b\xd6\x35\xd0\x94\x9d\xb1\x72\xb0\x30\xc9\xde\x6a\x2e\x6e\x4e\xbf\xbf\x1b\x7e\x1a\xbb\x7f\x3d\x49\xb0\x71\xe4\xb6\x91\xec\xba\x7b\x46\x4b\x56\x1a\xa6\xe6\x68\x76\x15\xaa\x5b\xfe\x8d\x2b\xe6\x0f\xfb\x1a\x66\x1c\xeb\xac\x60\xe7\xe1\xb7\xa7\xd6\x6b\xbb\x09\xef\x06\x1f\x11\x1f\x53\x40\xbf\xf4\xff\xa8\x1f\x35\x3f\x51\x7a\x0f\xf9\xe4\xd2\x31\xd9\x8d\x70\xb2\xdd\x81\x70\xf2\x95\x9e\xd7\x85\xa5\x0b\x26\xe2\xff\x19\xd0\x0b\xb9\x8e\x24\x5f\xf7\xe8\xbc\x66\x3a\x7a\xf9\xac\x86\x4a\x1d\x91\xcf\x8e\xff\xa9\x1b\x6e\xdb\xfb\xbd\xfb\xbd\xff\x02\x00\x00\xff\xff\x71\x71\xa2\xf8\x06\x1d\x00\x00"

func exampletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_exampletokenCdc,
		"ExampleToken.cdc",
	)
}

func exampletokenCdc() (*asset, error) {
	bytes, err := exampletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1d, 0x5c, 0x4d, 0xa, 0x85, 0xbd, 0x97, 0x7b, 0x75, 0x21, 0xc6, 0x66, 0x3, 0x70, 0x7, 0xe6, 0x1a, 0x43, 0x4, 0xdc, 0x11, 0xb9, 0x6c, 0xe7, 0x93, 0xcb, 0x18, 0xb, 0xbc, 0xfe, 0x96, 0x20}}
	return a, nil
}

var _fungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x59\x4d\x93\xdb\xb8\x11\xbd\xf3\x57\x74\x79\x0f\x9e\x71\x64\xcd\x1e\x52\x39\x4c\xd5\x26\xd9\xad\x5d\x57\xf9\x92\x4a\x25\x4e\xf6\x2a\x88\x6c\x4a\xc8\x80\x00\x17\x00\xa5\xa1\x5d\xfe\xef\xa9\x6e\x7c\x91\x14\xad\x91\xed\xb9\x8c\x44\x02\x8d\xee\x87\xee\xd7\x0f\xd0\xc3\x9b\x37\x55\xf5\x03\x7c\x38\x22\xbc\x53\xe6\x0c\xef\x06\x7d\x90\x7b\x85\xf0\xc1\x3c\xa1\x06\xe7\x85\x6e\x84\x6d\xaa\xea\x87\x1f\x60\x97\x5e\xf2\xbb\x1d\xd4\x46\x7b\x2b\x6a\x0f\x52\x7b\xb4\xad\xa8\xb1\xaa\xc8\x50\xfe\x0a\xfe\x28\x3c\x08\xa5\xa0\x4d\x66\x3d\x9b\x4d\x33\x1d\x9c\xcd\xa0\x1a\x38\x8a\x13\xbd\xa2\xe7\xad\xb1\x1d\x78\xb3\xad\xde\xb7\x20\x60\x70\x68\x1d\x9c\x85\xf6\x8e\xde\x37\xd8\x2b\x33\x82\x00\x8d\xe7\x85\xa9\x0d\xf8\x23\x4a\x9b\xbf\x57\xc1\xb2\x46\x6c\x68\xa6\xec\x7a\x85\x1d\x6a\x4f\xc3\x60\x16\x48\xf1\x77\xcb\xfe\x4f\x8c\x2c\xdc\x6b\x8d\x22\x8c\x28\x20\xb2\x62\x07\x85\x0e\x84\x6e\x40\x8b\x4e\xea\x43\xc5\xe1\xfa\x19\x02\xae\xc7\x5a\xb6\x12\xdd\x36\x40\xf8\x5f\x31\x28\xbf\x03\x8b\xce\x0c\x96\x00\xfb\x4d\xd4\x47\x10\x75\x6d\x06\xf6\x4d\x78\x30\x67\xed\x42\x70\x09\x9e\x14\x04\xfb\x21\xc8\x61\xda\x97\x1a\x2b\xd3\xf2\x72\x6c\x34\xdb\x04\xe7\x8d\xc5\x06\xa4\x8e\x90\x24\xeb\xf4\x5c\x1c\x62\x94\xcb\x49\x47\xe1\xa0\x43\x7f\x34\x8d\x83\x1c\x87\x39\x6b\xb4\x1c\xa1\xf1\x47\xb4\x71\x3b\x6a\xa1\xa1\x16\x4a\xc5\x90\xfe\x69\xcd\x49\x36\x68\x77\x1b\xd8\xfd\x0b\x6b\x94\x27\xfe\x4c\xb3\x76\xbf\x08\x45\x8e\x96\x80\x0b\x34\x8e\xdd\x70\xd3\x27\xd0\x60\xad\x84\x45\xe8\x2d\xbe\xad\x8d\x6e\xa4\x97\x46\x07\x88\x7b\xe3\xfc\xf4\x19\xfb\x68\xd1\x79\x2b\x6b\x5f\x91\xb3\xf8\x8c\xf5\x40\x2f\x21\xc2\xd2\x0e\xba\x0e\x83\x03\x14\x21\xe4\x10\xfe\x08\xb4\x8e\xc3\x5e\x58\xe1\x11\xf6\x58\x8b\x81\x7c\xf1\x70\x90\x27\x74\x3c\x9c\xa2\xe5\x0f\x62\x2f\x95\xf4\x23\x6d\x81\x3b\x0a\x8b\x95\x00\x8b\x2d\x5a\xd4\x35\xe7\x45\x80\x39\x00\x1a\xb6\x50\xab\x11\xf0\xb9\x37\x2e\x9a\x6a\x25\xaa\xc6\x15\x8f\x2a\xa9\xc1\x68\x04\x63\xa1\x33\x16\x93\xc7\x05\x8a\x6d\x55\xbd\xa7\xd2\x71\x26\x3a\x14\xa0\x5f\x78\xd3\x89\x27\x84\x7a\x70\xde\x74\x19\xe1\x08\x4d\x4e\x78\xc2\x66\x8e\x32\x15\x92\x81\x93\xb0\xd2\x0c\x34\x5a\xea\x83\x83\xb3\xf4\x47\x36\x1f\x32\x6f\x5b\xbd\x33\x16\xf0\x59\x90\x99\x0d\x08\x68\xc5\x50\xa3\xe7\xbd\xdf\x63\xb1\x8e\x0d\xec\xc7\x54\xb7\x5c\x03\x0c\x07\xa4\xa4\x98\x15\xd7\x2f\x23\x0c\x4e\xea\xc3\xc4\x57\xda\xda\xe2\xda\x26\x86\x69\xda\x45\x89\x66\xc2\xa8\xc8\x01\x87\xba\xe1\x99\x36\xa4\x5b\xaa\x96\x1e\xd1\xbe\xf5\xe6\x2d\xfd\xdf\x70\x44\x66\xf0\x54\x35\xb4\x26\x91\x00\x2d\xc4\xdc\x40\xc1\x0a\xa8\x91\xac\x2a\x50\xd8\x1c\xd0\x82\xeb\x84\xf5\x79\xa9\x2d\x7c\x30\x61\xa5\x68\xdd\x1b\x10\xba\xd4\xc1\xa6\x0a\xf4\x14\x6b\xd4\x11\x24\x23\x2f\xda\x58\x71\x9e\x40\x09\xad\x35\xdd\x34\x47\x98\xaa\x42\x09\x71\xe2\x36\xd8\x1b\x27\x7d\xce\x0e\x30\x7a\xb6\xd2\x6b\x97\x72\x8b\x18\x92\x90\xf7\x18\xec\x5b\xa1\x5d\x8b\x76\x5b\x55\x6f\x1e\xaa\xea\xe1\xe1\x61\x0e\x1b\x3d\xe1\xa7\x2b\xac\xfc\x45\x46\xce\x5b\xbb\xe5\xe9\xfd\xb0\x5f\x21\xfa\xc5\xf6\x7c\xaa\x2a\x00\x80\xb4\x94\x37\x5e\x28\xd0\x43\xb7\x47\xcb\xa9\x1d\x70\x90\x1a\xf0\x59\x3a\x4f\x65\xb3\xcd\x13\xde\x7b\x90\x0e\x86\x3e\x16\xd2\x24\xb5\x2c\x3d\x42\xed\x06\x8b\x85\x92\x82\x6d\x37\xf4\xbd\x1a\xb3\x0d\xe7\xc5\xe8\x88\xe7\x06\xae\x66\x4a\x8d\x60\xb0\x11\x1e\xd3\x28\xfe\x4f\xe1\x9c\x84\x0d\x66\xfe\xcd\x56\x1e\xe1\x3f\xef\xe4\xf3\x5f\xfe\x3c\x89\x81\xfd\x7d\xaf\xa5\x97\x42\xc9\x8f\xd8\xcc\x4c\xa4\x28\xf1\x84\x89\xb2\xa5\x03\xec\xa4\xa7\x6a\x38\xd3\xd6\x92\xa3\x05\x34\x07\xb5\x45\xe1\x17\x66\xc8\x93\x60\xe2\x62\xb9\x3b\x19\x3e\xcf\xfd\xbb\x5f\x3a\xf8\x7b\xcc\x35\xfd\xd5\xee\x85\xfd\x20\x06\x4c\xf9\xaa\x43\x96\x8a\x90\x69\x57\x1d\xcd\xcb\xde\x89\x8e\xfa\x4a\xf2\x6f\xc3\x26\x1e\xe1\xe7\xa6\xb1\xe8\xdc\xdf\x2e\xfc\xfd\x35\xe4\xf9\x37\xc0\x59\xfc\x6d\x92\x0d\xca\x45\x73\x93\xbf\x79\xd9\x0b\x7f\xbd\x59\xf5\x36\x71\xd7\xaa\x9b\x8b\x32\x42\x22\xbe\x3a\xb2\xbc\xc5\x3f\x06\x69\x39\x79\x1d\xb4\xc6\x66\x74\x89\x18\x93\x91\x05\x29\x94\x7c\x67\x92\x1a\xfb\x52\x1a\xd3\x12\x69\x0c\x3a\xd0\x26\x2f\x38\x5f\xcb\x68\xd8\xed\x53\xab\x3d\xa2\xc5\x4d\x9e\x3b\xe9\x6c\x0a\x05\x75\x12\xd3\xc7\x0c\xed\x8d\x73\x32\x36\x13\xd3\x86\x24\x25\x27\x62\x43\xe9\x23\x0c\xae\xb8\x4e\x11\x37\x86\xfd\xd0\x58\xa3\x73\xc2\x4a\x35\x46\x7d\xc2\x04\x67\xce\x1a\xa2\x27\xdb\x8b\x5d\xb9\x14\x01\xa5\x4f\x44\x0a\x49\x4b\x65\x1e\x75\xc3\x3e\x12\xd3\x12\x38\x16\x27\x89\x1b\x67\x93\x43\x6b\xf0\x83\xa5\xa4\x89\xdc\x99\xfb\x9b\xc5\xce\x9c\xb0\xc9\x7d\x6e\x32\x71\x66\xe4\xc3\x44\x41\xbc\x66\x72\x41\xe7\x40\xe1\x09\x15\x25\x68\x3f\xec\x95\xac\x37\xb0\x1f\x28\x69\xa5\xa3\x67\x84\x8b\x20\xdc\xf6\x0a\xbb\x99\xb1\xb4\x0b\x2c\x0c\x8a\xb2\x22\x45\xc6\xdb\xce\x7e\x65\x70\xe6\xba\x6d\x66\xa8\x66\xf9\xc7\xec\xa0\x46\x6e\x21\x61\xf5\xe4\xe9\xf5\x78\xc2\xaa\x9d\x18\xe1\x60\x85\xf6\x51\xd5\xc5\x75\x72\x8c\xd4\xd0\x53\x2e\x50\x38\xf2\x94\x58\xb4\x78\xd1\x67\x15\x12\x25\xbe\x39\xbb\x24\x76\xeb\x99\x5a\xa4\x2a\x65\xbb\x33\x0b\x9c\x7f\x69\xef\x73\xe8\xfe\x68\xcd\x70\xa0\xd6\x9c\xf5\xd5\xad\x01\x05\xa9\xc4\x51\x11\x28\x2f\xc4\xc4\x9b\x77\x4b\x48\x64\x6b\x11\xc7\xcc\xf7\x99\x8d\xaf\x8f\x83\xaa\xa2\x1d\x74\x4e\xf7\x05\x45\xdd\x3f\xc2\xdf\x43\xfa\x7e\xca\x53\x78\x9a\x71\xcb\x47\xc1\x32\xec\x2c\xba\x78\xc2\x68\xa3\xd7\x21\xb9\xa8\x1a\xe0\x24\xd4\x80\x17\xd3\xc2\x94\x6d\x2c\x5b\xf8\xe9\x27\x88\x5e\x5c\x8c\xa4\xbf\x57\x89\xff\x85\x8a\xe3\xa0\x1b\x9c\x27\x55\x48\x2b\x39\xd1\x21\x88\x00\x52\xb2\x18\xd5\x6d\xe9\x35\x1c\xd3\xab\x99\xf9\xcf\xd5\xfc\xd3\xe7\xc2\xc7\xe9\x50\xf1\xfd\x7c\x1c\xbb\xc7\x0a\x1d\x73\x37\xb9\x91\x8e\x7f\xc7\x44\x82\x52\xd7\x6a\x68\x90\xa4\x64\x3a\x99\x04\x37\xea\x23\xd6\x4f\x73\x10\x22\x05\x64\x2b\x67\xe4\x73\x2d\xed\x10\x29\xfc\x5b\x04\x7e\x80\x21\x08\xfc\x6a\xca\x08\x8d\x49\x83\xd6\xd5\xfc\x06\x94\x7c\xa2\xc3\xa8\x92\xac\xa2\x3a\x92\x47\x42\x37\x45\x40\xb1\xce\xa5\x17\x24\x9a\x64\xcb\x49\xeb\xa1\x57\xe1\x2c\x02\x2f\x13\x79\xda\xa4\x25\x91\x27\x71\xeb\xc5\x13\x16\x36\x26\x86\x8e\x6f\x1c\xb5\xa6\x75\xf8\x4b\x3d\x8d\x3d\x5e\xad\x9f\x68\xeb\x2e\x28\x90\x50\x33\xf7\xcb\x3c\x8a\x87\xd1\x5b\xd2\x88\xc4\x9b\x90\x3a\xec\x47\x69\xad\x7c\x8c\x83\xe9\xa9\x3b\x1b\xa1\x88\x26\xc9\x27\x7c\x90\x2e\x1a\xcf\x61\x60\x90\x2f\x51\x08\x6e\xa6\x99\x91\x4d\x50\x13\x29\x22\x10\x6a\x63\x2d\xd6\x5e\x8d\x37\xe1\x1f\x83\x5b\xc2\x5f\xe4\xf8\xa4\x18\x05\x9c\x96\x3d\x73\x86\x28\x09\xe4\x38\x7c\x2e\x8e\xe9\x8f\x5c\xbc\x5b\xbc\xbd\xbf\x8d\x9f\x1c\xaa\x76\x4a\x33\xc9\xca\x3a\xcf\xa4\x88\x12\xbb\x4c\xb1\x49\xd9\x12\x1e\x25\x43\x37\x33\xca\xa5\x68\x4c\x58\x4d\x28\x7c\x99\x06\xe5\x3a\xc1\x9b\x2f\x1d\x41\xaf\x6c\x15\xaf\xf9\x98\x05\xcf\x26\x57\xcc\x66\x7d\xef\xd8\x9d\x70\x21\x22\xd2\xad\x06\xf3\x4c\x6d\xf9\xfc\x37\xf6\x2c\x15\xc4\xda\xe9\xac\x43\xa1\x27\x34\x11\x0d\xe2\x09\xed\xb8\xbc\xe1\xcb\xb3\xe7\xb7\x06\xee\xda\x3d\xd9\xd4\x28\xef\x4e\x83\xad\xd4\x38\x75\x6f\x79\xd1\x95\xf1\x6c\x8d\xed\x72\x5b\xfa\xc2\xdd\xd1\xd4\xfe\xfc\x1a\x69\x7a\x55\x10\x38\x84\x2f\x8c\x5c\x54\x4c\x91\xf0\x9b\x74\xdf\x42\x43\xca\x9d\xcb\xcb\x85\x41\x3e\x7d\x47\x69\x44\xb3\xe5\x36\x24\xec\x52\x84\x28\x5c\x6d\x15\xfd\x26\x3f\xce\xe4\xc3\x4c\x76\xf4\x56\x12\x30\x49\x1b\x2e\xf2\xfc\x92\x81\x82\x89\xeb\x35\xfa\xa2\xc0\xde\x85\x76\xbe\x2b\x12\x9b\x17\x78\xed\x66\x4c\x05\xab\x22\x3b\xf3\x5c\x69\x3d\xc9\x30\x36\x6b\xf3\xbf\x5b\x02\x59\x7c\x89\x61\xfe\xfa\x82\x90\xf9\x39\xa8\x97\x22\x4b\x12\xd3\xa8\xa0\xf2\x84\x06\x63\x01\xff\x18\x84\x0a\xdf\x56\x34\xcd\x55\x25\x03\x57\xa5\x1a\x9d\x07\x18\x27\x52\xcd\x42\x95\xeb\x9f\xdd\x1e\x5b\x63\x71\xc7\xd2\x00\x7d\xcc\x4a\x35\xe4\x45\x17\x0d\x69\xcd\x78\xbc\x2d\xd9\xe3\x41\x6a\x4d\x69\xb4\xb8\x13\x2d\xb7\xa5\x2b\xb3\x5f\x26\x6e\x76\xf0\x6e\xfa\xf8\x1e\xde\x5e\x47\xfb\x1f\x39\x43\xf6\x0b\x62\xe7\x3b\xb0\xa8\x39\x0a\xb2\xbd\xc5\x13\x5f\x50\xa6\xe1\x22\x48\x94\xdb\x65\xe4\x8d\x3a\x44\x34\x0d\x69\x90\xb2\x50\x24\xa7\xd9\x4e\xcb\x95\x73\xe6\x6d\x2a\xe4\x1b\x5b\xe3\x1a\xc2\x7f\x4a\x8f\x69\x81\xfc\xf8\x9b\xf0\x76\x43\xf7\x22\xd0\xe5\xae\xe5\xeb\xf4\x7a\x10\x39\xbf\x75\xbd\x1f\x23\xca\xf1\x2c\xa5\xc7\x78\xa7\x6e\xe2\x98\x19\x71\x30\xff\x1d\x05\x6d\xce\x47\xb4\x66\x29\x8f\xaa\x29\xd2\xcb\x25\xee\xd6\x68\x62\x05\xea\xcb\xa3\xce\x8f\xdb\x1f\x1f\xe1\x15\xd1\xb6\xc6\xb3\x1a\x93\x42\x8b\x3e\x31\x64\xfc\xb3\xcb\xd4\xa5\x57\x17\xb1\x7f\xae\x2a\x26\x7f\xe9\x4a\x43\x9d\xf6\x6e\xe9\xe0\x7f\x6c\x0a\x2d\xe7\x17\x23\x02\xc3\xf4\x47\xad\x8b\x7e\x4b\x16\x57\x7a\xbb\xd8\x9b\x13\x6e\xf2\x8d\xc2\xe5\x08\xfe\x81\x86\x4e\x27\x7b\x8c\xb6\xb1\x21\x5b\x46\x4f\xee\x68\x22\x3f\x74\x86\x7f\x22\x58\xde\xf3\xfe\x3a\x74\xdd\x08\x9f\x3e\x57\xff\x0f\x00\x00\xff\xff\x7b\x25\x6b\x5e\x1f\x1c\x00\x00"

func fungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_fungibletokenCdc,
		"FungibleToken.cdc",
	)
}

func fungibletokenCdc() (*asset, error) {
	bytes, err := fungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "FungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd7, 0x8, 0x18, 0xe5, 0xfc, 0xb3, 0x19, 0x7b, 0x71, 0x70, 0x7f, 0xf7, 0x58, 0x75, 0xbb, 0x39, 0x52, 0xad, 0x6f, 0x5a, 0xc3, 0x3b, 0x4c, 0xc, 0xe, 0x72, 0x27, 0x39, 0x37, 0x24, 0x77, 0xf}}
	return a, nil
}

var _tokenforwardingCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4d\x6f\xe3\x36\x14\xbc\xf3\x57\xcc\x6e\x81\xae\x1d\x64\xe5\x1e\x8a\x1e\x82\xa4\xdb\x6d\x3e\x8a\xa2\x45\x0a\x78\xb3\xed\xb1\xa0\xa9\x27\x8b\x1b\x99\x14\xc8\x27\x6b\x83\xc0\xff\xbd\x20\x25\xd1\x52\xa2\x14\xe9\xa5\xb9\xc4\x14\xf8\xe6\xcd\xbc\x19\x92\xab\x93\x13\x21\xbe\xc1\x4d\x63\xb6\x7a\x53\x11\xee\xec\x3d\x19\xdc\x58\xd7\x4a\x97\x6b\xb3\xc5\xa5\x35\xec\xa4\x62\x21\xee\x4a\xed\xa1\xfa\x25\x7c\x69\x5b\x8f\xd2\xb6\x90\x06\x52\x29\xdb\x18\x86\xb2\x4d\x95\xc3\x13\xa3\xa9\x21\xa1\x1a\xcf\x76\x97\xc0\x3b\xec\x35\x29\xd2\x7b\x72\x82\x2d\x64\x55\xd9\x16\x5c\xd2\x0e\x6c\x51\x74\x5d\xc1\x61\x9f\x0f\x5f\x24\x72\x5d\x14\xe4\xc8\x70\xea\xd1\x96\x64\x68\x4f\x2e\x94\x3d\xc0\x75\x68\x7d\x4d\x16\x58\xd2\x03\x94\x34\xa8\x9b\x4d\xa5\x7d\x09\x0e\xb4\x7b\x41\xe4\xe0\xc8\xdb\xc6\x29\x82\xf4\x90\x89\x0c\x94\xac\xe5\x46\x57\x9a\x1f\xf0\xa5\xf1\x8c\x4a\xdf\x13\x24\xfe\x94\x4d\xc5\xa7\x42\x9a\x3c\xb4\x83\x27\x13\x30\x72\x4b\xde\xbc\x63\xd0\x9e\x0c\x0c\x51\xa0\x8c\x7b\x63\x5b\x68\x86\xf6\x47\xd2\x99\x10\x7f\x95\x64\xc6\x23\x6a\xa5\xe1\xa8\x4d\x39\x92\x1c\x7a\x24\x6e\xa7\x9d\x24\x25\xab\x2a\x76\xeb\x76\xdc\x52\x9b\x76\x88\xa2\x31\x8a\xb5\x0d\x88\x39\x6a\x67\xf7\x3a\xa7\xd0\xb4\xd5\x5c\xc6\x9a\x24\xc8\x51\xa4\xa0\x08\x5c\x4a\xee\x90\x43\xef\xd1\xa0\x05\x97\xa4\xdd\x71\xdc\x99\x10\x27\x2b\x21\xf4\xae\xb6\x8e\x9f\xb8\x56\x38\xbb\xc3\x77\x5f\x6f\x3e\xdf\xfe\xf2\xeb\xcf\xbf\x5f\xdf\xfd\xf1\xdb\xf5\xed\xc7\xab\xab\xf5\xf5\xa7\x4f\x42\xd4\xcd\xe6\x18\x8c\xb8\x7f\x14\xa0\x47\x21\x00\x60\xb5\xc2\xf5\x3e\xf8\x18\xe9\x68\x0f\xda\x69\x66\xca\xa3\x9f\x03\x07\xe9\x08\x39\xd5\xd6\x6b\xee\x86\x1a\x24\xb1\x74\x5b\xe2\xc1\x69\x17\xd1\x42\x47\x8a\x70\xc3\x6c\xf2\xab\xae\x6e\x21\x77\x61\xce\x67\xf8\x7c\xa3\xbf\xfe\xf0\xfd\x69\x64\x7e\x86\x8f\x79\xee\xc8\xfb\x0f\x4b\x91\xea\x53\x12\xd2\x78\xcf\xa6\xa2\xb3\x34\xcc\x5e\x43\xaf\x23\x1e\x04\xed\x03\x73\x47\x91\xe2\x98\x73\x14\xd2\xea\xaa\xc2\x26\x06\x86\xb3\x69\x2d\x81\x1f\x6a\x82\x36\xb9\x56\x92\xc9\xf7\x03\x89\x33\x91\x63\xdb\x6c\x5c\x8e\x44\x77\x10\xe9\xa7\x54\x8a\xbc\x5f\x78\xaa\x8a\x25\xf6\x32\x58\xae\x74\xad\x29\x88\xbf\x4c\x71\x9e\x30\xef\x79\xce\xa1\xad\x56\x41\x7c\x17\xae\x2e\x31\xf2\x9e\xfc\x70\x04\x60\x37\x5f\x48\x71\x3c\x34\x06\xd2\x6d\x9b\x5d\x3c\x93\x26\x1f\xc2\xe4\xc7\x48\x9a\x07\xf3\x12\xa7\x77\xbe\x47\x6a\x7c\x48\x45\x3c\x4d\x6c\x1d\xe5\x47\xc9\x73\xb4\x82\x51\x45\x63\x06\xe6\x8b\xce\xcd\x9f\xa6\x3e\x45\xe0\x25\x1e\x53\x55\xf8\xab\x46\x99\x59\x53\x81\x0b\x84\x49\x65\x89\x50\xb6\xb1\xce\xd9\xf6\xfc\xdb\xc7\x79\xd3\x0f\x3f\x2e\x96\x6f\xc4\x33\xc8\x8d\xac\x64\xb0\xe7\x22\x06\x2b\xeb\x97\xa3\x6d\x93\x8a\x11\x81\x6c\x2a\xe1\xfc\x7d\xf8\xbf\x9c\x36\x08\x67\xe2\xe5\x44\xf7\xbd\x86\x48\x47\x39\xb6\x35\xe4\x3e\x64\xb2\x8b\xf7\x32\xa1\x1d\x26\xbe\xab\x52\x9a\x2d\xad\x07\xe9\xfd\xda\x4f\x1d\x82\x2d\xe2\x87\x22\xdd\x95\xbd\x87\xfd\x3d\x93\x1f\xb7\xfe\x9b\x53\x4f\x7a\x2d\xfe\x86\xa1\x76\x3d\x17\xcd\xa7\x8e\xd5\x8e\x9e\x7c\x09\x7f\xe3\xea\xd7\x78\x86\x37\x17\x30\xba\x3a\xc3\xdb\xcb\xf8\x1a\x19\xcb\xe8\xca\xe6\x2e\xc7\x78\xaf\x05\x91\x47\x5a\x6f\x27\x14\x0e\x93\xd5\x34\x42\xb8\x98\xb0\x9b\x1b\xbe\x36\x9a\x17\xb3\x07\xf3\x75\xea\xff\x53\x5c\xff\x5f\xe9\xcf\xd3\xd0\x15\x1c\xd2\x85\xff\xfc\x01\xeb\x3f\x85\x7b\xc5\x50\x3b\x79\x96\x07\x5a\xe9\x29\x7b\x21\x76\x7d\xe4\x52\xdc\x9e\xf5\x78\x61\xdc\xe1\xd6\x48\xed\x8e\x83\x76\xc4\x8d\x33\x38\x7f\xdf\xbf\xc7\xb3\x30\xe9\xe7\xb2\x57\x78\x10\xff\x04\x00\x00\xff\xff\xd1\xe8\x38\xa0\x39\x09\x00\x00"

func tokenforwardingCdcBytes() ([]byte, error) {
	return bindataRead(
		_tokenforwardingCdc,
		"TokenForwarding.cdc",
	)
}

func tokenforwardingCdc() (*asset, error) {
	bytes, err := tokenforwardingCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "TokenForwarding.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0xd4, 0x19, 0x5d, 0x89, 0x9e, 0x77, 0x14, 0xb5, 0x83, 0xdc, 0xa7, 0x54, 0x30, 0xe3, 0xff, 0x67, 0xa7, 0xd7, 0xd8, 0xb9, 0x2e, 0x4c, 0xf4, 0x3e, 0xb8, 0xd4, 0x52, 0xb7, 0xdd, 0x70, 0xd0}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"ExampleToken.cdc":    exampletokenCdc,
	"FungibleToken.cdc":   fungibletokenCdc,
	"TokenForwarding.cdc": tokenforwardingCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"ExampleToken.cdc":    &bintree{exampletokenCdc, map[string]*bintree{}},
	"FungibleToken.cdc":   &bintree{fungibletokenCdc, map[string]*bintree{}},
	"TokenForwarding.cdc": &bintree{tokenforwardingCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
