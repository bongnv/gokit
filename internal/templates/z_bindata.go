package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _tmpl_crud_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xcd\x6e\xdb\x30\x10\x84\xcf\xe1\x53\x2c\x78\x4a\x82\x54\xba\x14\x3d\xe8\x6a\x19\x46\x00\xa3\x31\x92\xf8\x54\x14\x30\x21\xaf\x15\xd6\xb6\xa8\x90\xeb\x36\x85\xaa\x77\x2f\x68\xca\x92\x7f\x48\xc1\x69\x9d\x93\x7e\x66\xb4\xfb\x61\xc4\xdd\x52\x64\x4b\x91\x23\x54\x15\x44\x13\x77\xff\x55\xac\x11\xea\x9a\x31\xb9\x2e\x95\x26\xe0\x99\x2a\x08\xdf\x88\x33\x46\xbf\x4b\x67\x6d\x3c\x60\x48\x6f\x32\x82\x8a\xd5\x8c\xc5\xf1\xbe\xf4\x84\xfa\xa7\xcc\x10\xa4\x01\x01\x83\xc7\x69\x0a\xb2\x20\xd4\x0b\x91\x21\x2c\x94\x06\x71\x58\xc7\xb9\x23\x16\xc7\xb9\x4a\x72\x2c\x50\x0b\x42\xc8\xd5\x52\x52\x8e\xc5\x4e\x87\x4f\x5d\x91\xd3\x5e\x27\x78\x2d\x43\xf7\x11\x03\x00\xb0\x4d\x96\x92\x12\x58\x23\xbd\xa8\x79\xc2\x47\xc3\x67\x0e\xa5\xa0\x97\x84\xc7\xed\xf7\x7f\x80\xd4\x58\xfd\x42\xbd\xbd\x9b\xac\x36\x5a\xac\xa0\xae\xe3\xea\x3e\xad\x39\xbb\x1a\x21\x5d\x67\xf4\x06\x4d\x3a\xd1\xc0\x5d\xef\x40\xe3\x2b\xdc\x8e\x90\xf6\x40\x1e\xf1\x75\x83\x86\x6e\xe0\xfa\x44\x30\xa5\x2a\x0c\xde\x01\x6a\xad\xf4\x0d\xfb\x7f\x3e\xce\xae\xc6\xd2\xf4\xb1\x59\xd9\x0f\x77\xa2\x9c\x45\x37\x79\x78\x7a\x0f\x9e\xad\x31\xd0\x28\x08\x7b\x18\x9d\xc1\x4f\xe9\xd1\xce\xe3\x9c\xbe\xfb\x2f\xdb\x3a\xd3\x72\xde\xcf\xea\x0c\x7e\x56\x8f\x76\x16\x6b\x3a\x1c\x0f\x9f\x87\xff\x82\x9b\xe2\x0a\x7b\x71\x9d\xc1\x87\xbb\x25\x6a\x46\xd9\x7b\x80\x21\x8a\x22\x37\x65\x7e\xb9\x5d\x07\x16\xe4\x3e\xb5\xcf\xb2\xc8\x61\xf6\xc3\xa8\x22\xe1\x72\xce\x67\x81\xea\x2e\x93\x70\xf9\x46\x3f\xa8\xbf\xbf\x3f\x6e\xf7\x1f\x9a\x76\x55\xd5\x66\xb5\x93\xda\xfe\xfe\x11\xe8\xfa\x07\xf4\x83\xfe\x0f\x8b\x85\x41\xb2\xcb\xe5\xcb\xe7\x5d\x4f\xb5\x7d\xc7\x67\x5b\xc3\x58\xae\x25\xc1\xa1\x61\x65\xdf\x85\x31\x8e\x73\x08\x19\x8e\x83\x68\x8f\x42\x1b\xc2\xb7\xef\x81\x4c\xfa\xa7\x73\x47\x16\x9a\xbf\x0e\x2d\xe8\xb8\xd8\x4f\x0a\xce\x79\x2f\xc4\xa5\x8f\x4a\x68\xba\x3b\x88\xa0\xe3\x23\x19\x8e\x93\x08\x5b\x2e\x46\x11\x5a\x1c\x1d\x44\xd0\x71\xde\x62\xf8\x1b\x00\x00\xff\xff\x00\xd3\x9a\x32\x8d\x08\x00\x00")

func tmpl_crud_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_crud_tmpl,
		"tmpl/crud.tmpl",
	)
}

var _tmpl_crud_handler_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x51\x6f\x9b\x3a\x14\x7e\xc6\xbf\xe2\x5c\x74\x7b\x85\x2b\x4a\x7a\x5f\x2b\xe5\x4a\x57\xed\xd6\x4d\xaa\xda\x6e\xda\x9e\xa2\x3c\xd0\xe4\x40\xad\x82\x4d\x6c\xb3\x35\x42\xfc\xf7\xc9\x36\x24\x34\x81\x16\xa4\x55\x7b\x6a\x39\x9c\xef\xf8\x3b\xdf\xf7\xc9\xa4\xaa\xce\xe0\x6f\x96\x17\xd9\x6d\x9c\x23\x5c\xcc\xa1\x90\x8c\xeb\x04\x7c\x53\x3b\x51\x9f\x62\xbe\xce\x50\xfa\x10\xd9\xf7\x67\x75\x4d\x2c\x62\x1d\x8b\x43\xc0\x89\xba\xfa\xff\xae\xdb\x48\x8a\x78\xf5\x14\xa7\x08\x8f\x6e\x88\x22\x84\xe5\x85\x90\x1a\x02\xe2\xf9\x2b\xc1\x35\x3e\x6b\x9f\x78\x3e\x4a\x29\xa4\xf2\x09\xf1\xfc\xaa\x82\xe8\xbe\x81\xd5\xb5\x7f\x54\x99\x31\xae\x51\xf2\x38\x9b\x29\x2d\x64\x9c\xa2\x69\x49\x99\x7e\x2c\x1f\xa2\x95\xc8\x67\x0f\x82\xa7\xfc\xc7\x2c\x15\x4f\x4c\xcf\x4a\xcd\xb2\x19\x72\xcd\xf4\x76\xb6\x29\x51\x6e\xdf\x6c\xce\x44\xea\x13\x4a\x08\xd1\xdb\x02\xa1\xaa\x3a\xd2\xd4\x35\x28\x2d\xcb\x95\x86\x8a\x78\xe6\x4d\x2b\x81\x7d\x61\xb9\x44\x2f\xcb\xa4\x26\x24\x29\xf9\x0a\x02\x06\xa7\x07\xb3\x28\x5c\xa3\x0e\x56\xfa\x19\x1a\x1d\xa2\x4b\xf7\x37\x04\x89\x1b\xdb\xde\x6e\xdd\x20\xa2\x6b\xd4\xa6\xda\x3c\x7e\xc5\x4d\x89\x4a\x53\x08\xc6\xf4\xaa\x42\x70\x85\x21\x58\xa5\xa9\xd9\xc0\xca\xc2\x50\xd9\x9a\xb1\x91\x1d\xb0\x8f\xbe\x18\xc5\x0c\xc5\x90\x78\x9e\x95\x2f\xfa\xb0\x29\xe3\x2c\xf0\xd9\xda\xb7\x34\xa3\xcf\x57\x74\xff\xf2\x86\xe5\x4c\x07\xff\x9a\x0a\x25\xc4\x63\x89\x9d\xfc\xd7\x1c\x38\xcb\xcc\x89\x9e\x44\x5d\x4a\x6e\x1e\xed\xa1\xc4\xab\x5d\x5b\x86\x3c\x68\xe9\x50\x98\xcf\xe1\xbc\xaf\x5d\x48\x15\xdd\xe2\xcf\xc0\x6f\xd3\x74\x01\x12\x95\x28\xe5\x0a\x81\x0b\x0d\x89\x28\xf9\xda\xa7\x6e\x6a\x83\xfd\x07\x9f\x0b\xe4\x0a\xd5\x80\x20\xe6\x98\x4e\xf9\x02\x5c\x58\xbe\x89\x4e\x71\x47\x6d\x71\xbe\x34\xbb\xd5\xa1\xa1\xf4\xba\xb7\x37\x4c\x4d\x34\xd7\x20\x46\xbb\x7b\xd4\x7c\x6c\x6f\xc2\x32\x8d\x52\x19\x63\x17\x4b\xe7\x8f\xf5\xb3\x6a\x34\x37\xee\xdd\x25\x89\x42\x0d\xff\x35\x7a\xb7\x90\x39\xc4\x45\x81\x7c\x1d\x34\x85\x10\x1c\xde\xb5\x07\x7b\x24\xa5\x3b\x0b\x4d\xd1\xfa\x3f\x7a\x9a\x4b\xcb\x0e\xd7\xce\x9a\x12\x4b\x68\x46\x46\x51\x34\x25\x70\x1a\x73\xab\x4b\x1e\x3f\x61\xb0\x58\xf6\x09\xdc\x11\x37\x7c\x19\x4f\x4a\xbc\x44\x48\x60\x61\x13\x15\x33\x48\xc6\x3c\x45\x68\x5b\xec\xc9\xf6\x90\x05\x5b\xc2\x7c\x38\x52\xdb\x83\xac\x8e\x37\xba\x89\xad\x16\xf7\x59\x29\xe3\xac\x93\x5f\x7b\xee\xc8\x90\x5e\x4a\x8c\x35\x4e\x8b\xa9\xc3\x8c\x0e\x6a\x4f\xfb\xc0\x4d\x64\x85\x74\xff\x7d\x94\x22\xef\x8a\x65\x32\xd2\x79\xa6\xc4\x1b\x8a\xc6\x7e\xa3\xd6\x1e\x3a\x3a\x16\xaf\xb8\x30\xb8\xc5\xf8\xeb\x63\x3b\xf6\xe6\xf8\x5e\xac\x27\x9b\xe2\x30\xa3\x4d\xe9\x69\x7f\x4f\x53\xf6\x1b\xfd\x56\x53\x06\xb7\x78\x07\x53\xae\x30\xc3\xa9\xa6\x38\x4c\x9f\x29\x56\xe4\x3f\xfb\x09\x9e\xf4\xf5\x7d\x61\xc5\x11\xc5\xbd\x36\x21\x74\xbf\x94\x3b\x45\x87\xe5\x87\xd3\xce\xef\xa6\x5d\x92\x7a\xd5\xec\x34\x18\x76\x66\x31\x37\x62\xde\x97\xa0\x37\xc3\xd3\x19\x57\xd5\x07\x4c\x0f\x93\x6e\xae\xd4\xb7\x38\xd1\xde\x55\x1a\xa6\x76\xc0\x08\x9e\x3d\x13\x2c\xb9\x5f\x01\x00\x00\xff\xff\xd6\xbd\x96\x14\xa2\x0b\x00\x00")

func tmpl_crud_handler_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_crud_handler_tmpl,
		"tmpl/crud_handler.tmpl",
	)
}

var _tmpl_dao_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xc1\x6e\xdb\x3a\x10\x3c\x8b\x5f\xb1\x10\x5e\x02\x29\x70\xc8\x7b\x80\x77\x48\xa2\x16\xc8\xc5\x69\x91\xf6\x54\x14\x01\x25\xad\x65\x36\x12\xa9\x90\x2b\xb7\x8e\xa0\x7f\x2f\x28\x31\xae\x6a\x27\xa8\x01\x9f\x6c\x68\x67\x67\x67\x66\x97\x7d\x7f\x09\xff\x95\xd2\x2c\x65\x83\x70\xf5\x3f\xb4\x56\x69\x5a\x41\x7c\xe6\xb2\xeb\xfb\x18\xf8\xf8\xfd\x72\x18\xd8\x2b\xf0\xae\x69\xeb\x39\x50\x35\x6d\x7d\xe6\xe2\x3f\x24\x1e\xcc\x5a\x59\x3c\xc9\x0a\xa1\xef\x81\x7f\x9a\xfe\x8f\x45\x5f\x53\x4d\x6b\x2c\x41\xc2\x00\x00\xe2\xc2\x68\xc2\x5f\x14\x33\x16\xc5\x95\xa2\x75\x97\xf3\xc2\x34\x22\x37\xba\xd2\x1b\x51\x99\x27\x45\xa2\x23\x55\x0b\xd4\xa4\x68\x2b\x4c\x4b\xca\xe8\x78\x6a\x3e\xa6\xe1\xb9\x43\xbb\x3d\xc4\xff\x50\xfa\x65\xdd\x89\xca\xd8\x26\x66\xd1\xe3\x7b\x35\x51\x2a\x59\x63\x41\x4e\x34\x5b\xf7\x5c\xc7\x20\x04\x28\xad\x48\xc9\x5a\xbd\x20\x8c\x1f\x59\xca\x18\x6d\xdb\xd1\xed\x2e\x86\x61\x00\xa5\x09\xed\x4a\x16\x08\x3d\x8b\x6e\x2d\x4a\xc2\x24\xd8\xe5\xb7\xd3\xef\x02\x2e\x7c\x44\xa1\x23\x05\xb4\xd6\x58\x16\x7d\xf6\x9a\x0f\xb1\x9c\xf3\xd1\x0d\x1f\xeb\x29\x24\xdf\xbe\xcf\xdb\x17\x53\x7b\x3a\x7a\xfd\xda\x96\x47\xcf\xf3\xf8\x0c\x6b\x3c\x16\x3f\x30\x26\x04\x2c\xf1\xe7\x9e\x61\xd3\xa2\x76\x50\x18\xad\xb1\xf0\x5b\x72\x40\x06\xb2\x1b\xd8\x28\x09\xb4\x46\xa8\xd4\x06\x35\x64\x0f\x4b\x0e\x77\x04\x16\xa9\xb3\xda\x81\xd4\x90\x5d\xdf\x7b\xa8\x2c\x0a\x74\x6e\x84\x96\x92\x64\x2e\x1d\x72\xb6\xea\x74\x71\x38\x2b\x29\x73\xb8\xf0\x0b\xe2\xd9\x4d\xba\x1f\x7c\xcf\xa2\x89\x1c\xce\x43\x65\xbc\xda\x61\xe8\x59\x14\x95\xf9\x15\x94\xf9\x82\x45\x83\xf7\x31\xdf\x5b\x00\x81\x23\xdb\x15\xe4\x59\x66\x43\x3c\x78\x94\x92\xa8\x31\x93\x59\x43\x0a\x61\xb9\x8f\x70\x10\xdf\x74\x84\x6f\xa5\x38\xa9\x74\x5d\x4d\xfe\x39\x29\x5e\xe6\x3c\xd0\x4c\x3d\xe9\xce\xc4\x84\xe2\x1f\x5e\xb3\x7f\x57\xc6\x74\x36\x6f\xa8\xf0\x57\xa3\xd0\x1d\x7b\x40\x5e\xda\x46\xda\x30\xd8\xc1\xdf\xa8\xb9\xec\x89\xed\x8b\x95\xda\xad\x8c\x6d\x12\x6f\x63\x37\x8e\x73\x9e\xf2\x8f\x4a\x97\xc9\x79\x60\xda\xf7\xe4\x16\xc7\x9b\x0b\x07\x7d\x6a\xc6\x0f\x72\x73\x42\xc2\xe1\x95\x9c\x2a\x22\xd0\xfc\x43\xc6\xef\x00\x00\x00\xff\xff\xa9\x49\xe3\x60\x9c\x05\x00\x00")

func tmpl_dao_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_dao_tmpl,
		"tmpl/dao.tmpl",
	)
}

var _tmpl_endpoints_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xdf\x6a\xe3\x3c\x10\xc5\xaf\xa5\xa7\x18\x4c\x0b\x36\x34\x0a\xbd\x0d\xe4\xea\xa3\x17\x1f\x2c\x4b\xd9\xdd\x17\xd0\xda\x93\x44\x24\x96\x9c\x91\x9c\x66\x11\x7a\xf7\x45\xf2\x9f\x28\x4d\xb6\xed\x45\x08\x96\xce\x9c\x39\xfa\x8d\xe4\xfd\x02\x1e\x3a\x59\xef\xe5\x16\xbf\xcb\x16\x61\xb5\x06\xf1\x9a\x7d\x2f\x42\xe0\x49\x64\x91\x4e\xaa\xc6\xff\xb5\x43\xda\xc8\x3a\x29\x3b\x52\xda\x6d\xa0\x78\xb4\xe2\xd1\x16\xd7\x85\xe2\x6e\xf9\xdc\x63\xde\xe5\x63\x77\xf0\x1e\x9c\xf9\x66\xde\x90\xae\xd5\x51\xa3\xda\xce\x90\x83\x92\xb3\xa2\x36\xda\xe1\xd9\x15\x9c\xb3\x62\xab\xdc\xae\xff\x2d\x6a\xd3\x2e\xb7\x66\xb1\x57\x6e\x19\x7f\xa8\x9b\xce\x28\xed\x0a\xce\x0a\xef\xe7\x54\x10\x42\xc1\x2b\xce\x97\x4b\x98\x14\x16\x84\x10\xdc\xfd\xe9\x30\x5b\xb2\x8e\xfa\xda\x81\xe7\xde\x03\x49\xbd\x45\x78\x98\x36\x53\xf2\x97\x59\x19\x02\x67\xe9\x70\xd3\xbe\x18\x03\x4f\x92\xd9\x75\x2e\x8a\xa6\xa8\x9b\x74\xf2\x90\xb2\xb4\x72\x8f\x17\x4b\x42\xd7\x93\xb6\x20\x35\xbc\xbc\x4f\xf4\xb6\x43\x42\x40\x59\xef\x66\x5f\x50\xfa\x64\xf6\x68\xa3\x91\xdb\x21\xd4\x86\x08\x6d\x67\x74\xa3\xf4\x16\x5a\x74\x3b\xd3\x80\xd1\x69\xaf\x23\x73\x52\x0d\x36\x30\xc2\x15\x7c\xd3\xeb\xfa\xba\x7f\x69\xe3\x18\x6e\x67\x1d\x42\x95\x11\xf2\x9c\x0d\x39\x2f\x6b\x3e\x92\xf8\x02\xae\x8f\x79\xad\x52\x9a\x98\xe0\x5f\x8a\xd2\x56\x4f\xa9\xd5\x04\x91\x25\x8e\x5f\xe9\xfd\x8e\x75\x8e\x7a\x2e\x3a\x29\x39\xb0\x92\xd6\xde\x25\xf5\x71\xb6\xcf\xe1\xcd\x91\x32\x88\xd1\xbc\xac\xdd\x19\xc6\xab\x2d\xfe\x1b\xfe\x9f\x80\xf0\xd8\xa3\x8d\x53\x1e\xbd\x7c\xa8\xa0\xcc\xbe\x9e\x00\x89\x0c\x55\xd1\x8d\x9d\x24\xc5\x17\xc2\x58\xbc\x02\x79\x4d\x5c\x43\xa2\x41\xcb\x19\xab\xd2\x1c\xe0\xe1\x80\x1b\xf7\x53\x35\xe9\x41\x16\xb1\x28\xd9\x15\x03\xd7\x34\x29\xb5\x01\x3c\x42\x79\x40\x9d\x1d\xfb\x07\xda\xfe\xe0\x6c\x05\xcf\xc3\x48\x87\x99\xce\x66\x6b\x28\xae\x5c\xd2\xa8\xe2\x23\x66\xde\xdf\xf5\x7b\x95\x24\xdb\x68\x37\x15\x5c\xbc\x42\x80\x35\x58\x71\x8f\x7a\x24\x56\x0d\x57\xe1\x60\x31\xbb\x5c\x23\xb4\x5f\xf1\x55\xaf\xd6\x11\x57\x83\xe7\x9b\x76\xf0\x5c\x89\x24\x49\x85\x84\xc7\xa8\x1d\x4b\x45\x19\x1b\xe6\x3e\x21\xdc\x30\xfb\x24\x5a\x1a\x5e\xc5\xb3\xf3\xb3\x69\xde\x33\xe9\xe1\xf2\x7a\x8f\xba\x09\xe1\x6f\x00\x00\x00\xff\xff\xd0\xc2\x71\xf6\x8a\x05\x00\x00")

func tmpl_endpoints_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_endpoints_tmpl,
		"tmpl/endpoints.tmpl",
	)
}

var _tmpl_entity_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcb\x31\x0b\xc2\x30\x14\x04\xe0\xb9\xf9\x15\x8f\xee\x26\x8b\x38\x74\x13\x5d\x5c\x9c\x74\xb6\xc1\x3e\x42\x68\xf3\x52\x9e\xd7\x21\x94\xfe\x77\xa9\x28\xe8\xe6\x4d\xc7\xf1\x9d\x71\x8e\xe6\x99\xec\xd9\x27\xa6\x65\x21\x6b\xad\x71\x2e\xe4\x26\xb0\xb0\x7a\x30\x85\xdc\x47\x04\x16\x62\x41\x44\xa1\x8d\xac\xf4\xeb\x63\x50\xc6\x9f\x81\x1e\xd0\xe9\x0e\x9a\x4d\x75\x3a\xd2\x3b\x51\xb0\xdb\xbe\x5a\x1b\xb2\xa6\xa6\x1e\x35\x26\xaf\xe5\xd6\x73\xa9\x5b\x53\x1d\x94\x3d\xb8\xdb\x83\x10\x13\xdb\x4b\x4c\xfc\x91\x92\x41\x32\x0d\xc3\xca\xae\x63\xf7\x07\x5b\xcc\x33\x00\x00\xff\xff\xf3\x14\x39\xda\xd8\x00\x00\x00")

func tmpl_entity_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_entity_tmpl,
		"tmpl/entity.tmpl",
	)
}

var _tmpl_handlers_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\x41\x4e\xc3\x30\x10\x45\xd7\x9d\x53\x7c\x65\x81\x92\x2a\x4a\xef\x50\x36\xb0\x89\x10\x9c\xc0\x98\xa1\xb1\x70\xec\x74\x3c\x21\x8d\xa2\xdc\x1d\x95\x46\x90\x45\x41\xb0\xb1\x17\xfe\xdf\xef\xfd\xce\xd8\x37\x73\x60\x34\x26\xbc\x78\x96\x44\xe4\xda\x2e\x8a\x22\xa7\x4d\x66\x63\x50\x3e\x69\x46\xb4\xc9\xa6\x09\xd5\xc3\x12\x9e\xe7\x8c\x0a\x22\x1d\xbb\xaf\xe2\x7d\xdb\x79\x24\x95\xde\x2a\x26\x9a\x89\x5e\xfb\x60\x91\x37\xd8\xae\x02\x05\xee\xd8\xfb\x98\x5b\x3d\x61\xf9\xbb\xba\xbd\xdc\x25\x84\x8f\xd8\xae\x28\xb5\x69\xcf\xa4\xea\x91\x8f\x3d\x27\x2d\x90\x5f\x7f\x4d\x5d\x0c\x89\x4b\xb0\x48\x94\x02\x13\x6d\x84\xb5\x97\x80\xe0\x7c\x79\x3e\x7e\xb6\xd9\x8f\xfc\x3f\x97\xfd\xc8\xbf\xeb\x7c\x06\xfe\x64\xb4\xdb\xa1\xe6\x01\x56\xd8\x28\x27\x18\x04\x1e\x90\x58\xde\x9d\x65\x0c\x4e\x1b\x3c\xf7\xc9\x05\x4e\x09\x3e\x1e\x9c\xad\x2e\x1b\x6a\x1e\xf2\x02\x57\xc8\x4f\x4b\xf5\x9b\x76\xb3\xda\x3a\xcd\x34\xd3\x47\x00\x00\x00\xff\xff\x91\xab\xfd\x00\xec\x01\x00\x00")

func tmpl_handlers_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_handlers_tmpl,
		"tmpl/handlers.tmpl",
	)
}

var _tmpl_main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x8a\xc3\x30\x0c\x86\x67\xeb\x29\x84\x27\x1b\x82\xb2\x17\x3a\xdf\x76\x57\xe8\x78\xdc\xe0\x26\x3a\xc7\x34\xb1\x8d\xe2\xa4\x43\xc8\xbb\x1f\x4d\x72\xd0\xa9\x93\x10\xff\xaf\x4f\x5f\x76\xcd\xdd\x79\xc6\xc1\x85\x08\x10\x86\x9c\xa4\xa0\x01\xa5\xfb\xe4\x35\x80\xd2\xcb\x82\x74\x39\x4a\xeb\x5a\x87\x58\x58\xa2\xeb\xeb\xce\xc5\xb6\x67\x19\xf5\x9b\xce\xc8\x32\x87\x86\x35\x28\x9f\xee\xa1\x5c\x59\x66\x16\xd4\x3e\x94\x6e\xba\x51\x93\x86\xfa\x96\xa2\x8f\x73\xbd\xc5\xf5\x54\xc2\x7e\xc3\xa2\xc1\x02\xfc\x4e\xb1\xd9\xc4\x8c\xc5\x05\x54\xca\x65\xc4\xd3\x19\xbf\x7f\x5e\x68\xf4\x95\x4b\x48\x71\x59\x8f\xfc\x8c\x2e\x67\x8e\xad\x79\x6e\x15\x1e\x06\xf4\xc1\x65\x2f\x8e\xe6\x5f\x9c\x3e\xf9\x61\xac\x25\x22\x0b\xa0\x58\xe4\xc9\x7e\x25\x6f\xc3\x80\xda\xc0\x44\x54\x81\xb2\x00\xa0\xfa\xe4\xe9\x22\x21\x96\x3e\x1a\x7d\xdd\x1f\xe0\x58\x52\xce\xdc\xe2\x23\x94\xee\xa4\x2b\x64\x11\x0b\x2b\xfc\x05\x00\x00\xff\xff\xb8\x92\x81\x77\x5f\x01\x00\x00")

func tmpl_main_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_main_tmpl,
		"tmpl/main.tmpl",
	)
}

var _tmpl_server_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xdf\x6f\xd3\x3e\x10\x7f\xb6\xff\x8a\xfb\x46\x9b\x94\x7c\xd5\x39\x1a\x8f\x43\x7b\x02\x04\x48\xfc\xa8\xd0\x04\x0f\x08\x21\x2f\xb9\xa6\x66\xad\x9d\x3a\xee\x18\xb2\xfc\xbf\xa3\xb3\xe3\x74\x6d\x47\x11\x95\x9a\x38\x77\x9f\xbb\xfb\xdc\x0f\x9f\xf7\x17\x70\xd6\xcb\xe6\x4e\x76\xf8\x41\xae\x11\xae\xae\x41\xcc\x1f\x7d\x87\xc0\x23\x66\x40\x7b\xaf\x1a\x7c\xab\x1d\xda\x85\x6c\x22\xb0\xb7\x4a\xbb\x05\x14\xe7\x83\x38\x1f\x8a\x7d\x3b\x11\x9f\x17\x07\xe6\x53\x88\xec\x9b\x8f\xb1\xc1\x7b\x70\xe6\x9d\xf9\x89\x76\x1f\x4c\x18\xb5\xee\x8d\x75\x50\x72\x56\x34\x46\x3b\x7c\x70\x05\x67\x05\xea\xc6\xb4\x4a\x77\xf5\x8f\xc1\x68\x12\x28\x43\x4f\x8d\xae\x5e\x3a\xd7\x17\x9c\xb3\xa2\x53\x6e\xb9\xbd\x15\x8d\x59\xd7\xb7\x46\x77\xfa\xbe\xee\xcc\x9d\x72\xf5\xd6\xa9\x55\x4d\x61\xd0\x92\x8d\xf7\x13\x79\x08\xe1\x58\x52\x2b\x4a\x5b\xcb\x55\x8d\xba\xed\x8d\xd2\x91\xc0\x49\xe7\x44\x81\x0e\x05\xaf\x38\xaf\x6b\x78\x8d\xee\x63\xef\x94\xd1\x03\x08\x21\xf8\x62\xab\x9b\x47\xb2\x72\xa0\x02\x1c\x17\x19\x42\xa8\xe0\xeb\xb7\xc4\x54\x24\x30\x78\xce\x92\xe0\xd5\x48\x66\xa0\x92\xae\xe5\x1d\x4e\x82\x72\xa8\x38\x67\xa6\x4f\xaa\x03\x07\x3e\x64\xdd\x35\xc8\xbe\x47\xdd\x96\xf4\x35\xeb\xd0\xbd\xb9\xb9\x99\x4f\x94\xf6\x63\x54\x42\x88\x19\x79\xb5\xe8\xb6\x56\x03\x99\xf0\xc0\x53\x26\xa7\x4d\x01\x27\x27\x4f\x25\xf3\x8f\x34\x39\xf3\xde\x4a\xdd\x21\x9c\x65\xbf\x71\xa4\x76\xe1\x68\xea\xd8\x58\x23\xf1\x45\xb9\x25\x71\xcb\xea\x91\xdc\x04\xf7\x9c\x31\xf6\x1e\xdd\xd2\xb4\x57\x90\x7f\xd4\xff\xc9\xbb\x48\x5a\x9a\x8c\x19\x81\xe7\xd2\x2d\x77\xd0\x43\x30\xc5\x22\xc4\x04\xcf\x81\xb2\xc9\x41\x71\xc4\x9e\xf5\x38\xf3\x59\x4b\x0e\xe8\xfe\xa8\x05\xe0\x06\xca\x15\xea\x47\xd8\xb9\xb4\x72\x3d\x54\x70\x19\xd3\x65\x9f\x70\xb3\xc5\xc1\xbd\xc4\xc6\xb4\x68\xaf\x40\xab\x55\x32\x07\x5c\x0d\xf1\x22\x3d\x01\x6a\xe3\xe1\x88\x42\x08\x23\x30\x7b\xd0\xed\x58\xd5\x50\xc5\x06\xec\x24\x4f\x8c\x84\xf7\x70\xba\x41\x71\x2b\x50\x52\x1d\xfe\x29\xa9\x67\x21\xa4\xd1\xfa\x1b\xc3\xf2\x3b\x8c\x4b\x41\xbc\x48\xef\x19\x58\xf8\x9f\xee\x9f\x18\x21\x15\x94\x36\x9d\x40\xe5\xab\xe5\xc3\x0c\xd0\x5a\xfa\x1b\x5b\xd1\x18\xc6\x3d\x35\xe2\x6e\x7e\xf5\x71\x4f\x95\x4a\xb7\xf8\x70\x44\x0f\x2e\x2b\x11\x21\xb1\x02\x16\x37\x84\xa5\x8c\xb4\x72\x9f\xe5\x6a\x8b\xfb\x8e\x68\x87\x31\x6a\xa1\xb5\x04\xcc\xab\x41\xa4\x2e\xe4\x3c\xec\x0c\x2c\x6e\xaa\xe7\x11\xf6\xdf\x35\xf5\x8f\x68\xe5\xe2\x52\x3b\x49\xc5\x59\xd8\x55\xdc\xe2\x26\x09\x63\x41\xa9\x27\xb1\xb4\xa8\xdb\x10\xf8\xef\x00\x00\x00\xff\xff\xe5\x79\xba\x6d\xdc\x05\x00\x00")

func tmpl_server_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_server_tmpl,
		"tmpl/server.tmpl",
	)
}

var _tmpl_service_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x6f\xc2\x30\x0c\x85\xcf\xe4\x57\x58\x3d\x01\x62\xe9\x9d\x23\xbb\xec\x34\x4d\x6c\x3f\x80\x28\x7a\x44\x19\x34\x09\x8e\x99\x40\x88\xff\x3e\xd1\x94\xb6\x42\x9b\xc4\xa9\xed\xf3\xf3\xf3\x67\x37\x19\xbb\x33\x0e\x74\xb9\x90\xfe\x28\xef\xef\xa6\x01\x5d\xaf\x4a\xf9\x26\x45\x16\xaa\x6c\x0c\x82\x93\x54\x4a\xd5\x35\x7d\x82\x7f\xbc\x05\xf9\x4c\x86\xb2\x6f\xd2\x1e\xe4\x83\x80\xb7\xc6\x82\xb6\x91\x6f\x72\xf1\x68\x55\xd7\x2e\x2e\x1d\x02\xd8\x08\xc8\xc5\x9d\x17\x87\x70\xaf\xd3\xcb\xd0\xd8\xc5\x2a\x39\x27\x0c\x33\xfa\xf2\x45\x4d\xde\xb0\xdf\xc7\xa9\x95\x13\x75\x3c\xfa\xb5\x3c\x17\x94\x68\xbe\xc6\xe1\x88\x2c\x33\x9a\xce\xd7\xc8\x29\x86\x8c\x05\x81\x39\xf2\x4c\x4d\x56\x67\xfc\xdd\xc8\x38\xd0\x7c\x75\xc6\xa8\xbb\xfd\x7a\x08\xb8\xb6\x9b\x77\x26\x4a\x8c\x8c\x20\xb7\xfd\xb9\x48\xba\x60\xdf\x0d\x59\xf8\x68\xe5\xc6\xdc\x9e\x32\x0b\xfb\xe0\x68\xf3\x9d\x63\x58\x56\xc1\x34\xa8\x36\x7d\x64\x99\x44\x5a\xf7\x19\x9d\x32\x84\x7c\xe1\x24\x0f\x21\xed\xef\xb8\x87\x0c\xfc\xff\xa3\x8d\x3c\xcf\xd3\x8d\x4e\x31\x00\x8e\xc5\xe7\x18\x7f\x03\x00\x00\xff\xff\x09\xae\x5b\x8e\x63\x02\x00\x00")

func tmpl_service_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_service_tmpl,
		"tmpl/service.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"tmpl/crud.tmpl": tmpl_crud_tmpl,
	"tmpl/crud_handler.tmpl": tmpl_crud_handler_tmpl,
	"tmpl/dao.tmpl": tmpl_dao_tmpl,
	"tmpl/endpoints.tmpl": tmpl_endpoints_tmpl,
	"tmpl/entity.tmpl": tmpl_entity_tmpl,
	"tmpl/handlers.tmpl": tmpl_handlers_tmpl,
	"tmpl/main.tmpl": tmpl_main_tmpl,
	"tmpl/server.tmpl": tmpl_server_tmpl,
	"tmpl/service.tmpl": tmpl_service_tmpl,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
		"crud.tmpl": &_bintree_t{tmpl_crud_tmpl, map[string]*_bintree_t{
		}},
		"crud_handler.tmpl": &_bintree_t{tmpl_crud_handler_tmpl, map[string]*_bintree_t{
		}},
		"dao.tmpl": &_bintree_t{tmpl_dao_tmpl, map[string]*_bintree_t{
		}},
		"endpoints.tmpl": &_bintree_t{tmpl_endpoints_tmpl, map[string]*_bintree_t{
		}},
		"entity.tmpl": &_bintree_t{tmpl_entity_tmpl, map[string]*_bintree_t{
		}},
		"handlers.tmpl": &_bintree_t{tmpl_handlers_tmpl, map[string]*_bintree_t{
		}},
		"main.tmpl": &_bintree_t{tmpl_main_tmpl, map[string]*_bintree_t{
		}},
		"server.tmpl": &_bintree_t{tmpl_server_tmpl, map[string]*_bintree_t{
		}},
		"service.tmpl": &_bintree_t{tmpl_service_tmpl, map[string]*_bintree_t{
		}},
	}},
}}
