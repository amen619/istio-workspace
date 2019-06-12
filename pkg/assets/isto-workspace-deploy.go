// Package assets Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml
// deploy/istio-workspace/operator.yaml
// deploy/istio-workspace/role.yaml
// deploy/istio-workspace/role_binding.yaml
// deploy/istio-workspace/service_account.yaml
package assets

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3d\x4e\x04\x31\x0c\x85\xfb\x9c\xc2\x27\x18\x34\x1d\x4a\x0b\x1d\x88\x02\x24\x7a\x6f\xc6\xbb\x6b\x6d\x26\xb6\x62\x67\x84\x84\xb8\x3b\xca\x64\xf9\x29\xb6\xcc\x97\x4f\xef\xbd\x04\x95\xdf\xa9\x1a\x4b\x89\x80\xca\xf4\xe1\x54\xfa\xc9\xa6\xcb\xbd\x4d\x2c\x77\xdb\x7c\x20\xc7\x39\x5c\xb8\x2c\x11\x1e\x9a\xb9\xac\xaf\x64\xd2\x6a\xa2\x47\x3a\x72\x61\x67\x29\x61\x25\xc7\x05\x1d\x63\x00\x28\xb8\x52\x04\x23\x1b\x41\x6c\xce\x32\x89\x52\xb1\x33\x1f\x7d\x4a\xb2\x06\x53\x4a\x5d\x3d\x55\x69\x1a\xe1\x96\x32\x72\xac\x5b\x00\xa3\xfd\x6d\x44\xee\x24\xb3\xf9\xd3\x7f\xfa\xcc\xe6\xfb\x8d\xe6\x56\x31\xff\x0d\xd8\xa1\x71\x39\xb5\x8c\xf5\x17\x07\x00\x4b\xa2\x14\xe1\xa5\xd7\x28\x26\x5a\x02\xc0\xf6\xf3\x19\xdb\x8c\x59\xcf\x38\x77\xaf\x1d\xea\xf5\xc5\xd7\x39\xe6\xe8\xcd\x22\x7c\x7e\x85\xef\x00\x00\x00\xff\xff\x66\x1b\x53\x70\x41\x01\x00\x00")

func deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml,
		"deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml",
	)
}

func deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml", size: 321, mode: os.FileMode(436), modTime: time.Unix(1554454542, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceOperatorYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xc1\x6e\xe2\x30\x10\xbd\xe7\x2b\x46\xa8\xd7\x52\xb8\xfa\x16\x81\x97\x45\x6d\x49\x14\xa2\xae\xf6\x84\xa6\xce\x50\xbc\x71\xec\xac\xed\x50\xa1\x8a\x7f\x5f\x05\x42\x61\x43\x20\x73\xcb\xf3\xcc\x9b\x37\x33\x2f\xb9\xd4\x19\x83\x94\x8a\x52\xa1\xa7\x00\x4b\xf9\x46\xd6\x49\xa3\x19\xf8\x06\x1c\x9a\x92\xb4\xdb\xc8\xb5\x1f\x4a\xf3\xb4\x1d\x07\x25\x5a\x2c\xc8\x93\x75\x2c\x00\x78\x04\x8d\x05\x31\x98\x3f\xf3\xd5\x34\x9a\x3c\xf3\x64\x95\xf0\xd9\x7c\x99\x26\xbf\x03\x00\x00\x4b\x7f\x2b\x69\x29\x63\xe0\x6d\x45\x07\x68\x8b\xaa\x22\x06\x99\x11\x39\xd9\xa1\x34\xb7\x58\xe2\x68\x39\x4f\xa3\x5e\x1e\x74\x0a\xf3\x5c\x57\xde\x91\x6e\x51\xcd\x5f\xc3\x19\x5f\x2d\xc2\x57\xde\xc3\x21\x9d\x97\xe6\xf1\xd3\xd8\xdc\x95\x28\xa8\x93\x26\x0d\x67\x3d\x2c\xf5\xba\x9c\xbf\x28\x4e\xf9\x0b\x8f\x13\xbe\xe4\x8b\x09\x5f\xbd\xf1\x64\x39\x8f\x16\x3d\x1c\x83\xd1\x70\x3c\x1a\x0d\x02\xf3\xfe\x87\x84\x6f\x56\x7c\x3c\xd3\x94\x4a\x65\x76\x05\x69\x7f\x28\xb8\x3c\x16\x96\xa5\xab\x6f\x53\xe3\x05\x79\xcc\xd0\x23\x3b\x7c\x41\x23\xe5\x7a\x42\x00\x57\x92\x38\x65\x59\x2a\x95\x14\xe8\x18\x8c\x1b\xc4\x91\x22\xe1\x8d\x3d\x65\x00\x14\xe8\xc5\xe6\x05\xdf\x49\xb9\x33\x78\xaf\x01\x7c\xbb\xe8\x82\xa4\x25\xaf\x0e\x75\xc5\x79\x9f\xf5\x7f\xe9\x47\xb1\x76\x2b\x05\x85\x42\x98\x4a\xfb\xc5\xdd\x5a\x00\x61\xb4\x47\xa9\x1b\x0b\x9f\xe3\xb1\xa7\xeb\x31\x64\x81\x1f\xc4\xe0\xe1\xab\xc3\xf3\xfb\xa7\x16\x7c\x32\xf1\xe9\xe1\x6c\xc9\x3d\xbb\x44\xd2\x70\xb6\x6f\xf5\x11\xa6\x28\x50\x67\xac\x05\xd7\x32\x65\xde\x16\x85\xf6\xc3\x75\x65\xd6\x8b\xe9\x1c\x20\xae\x94\x8a\x8d\x92\x62\xc7\x20\x54\x9f\xb8\x73\xad\x2c\xd2\xdb\x2e\xc2\xe3\x86\x7e\x85\xe9\xe4\xe7\x61\x8c\x65\x1c\x4e\xf8\x55\xde\xd9\xcf\x83\x9b\x1c\x71\x34\x3d\xff\x9b\x1d\xc5\x3f\xac\x29\xae\x15\xd4\xb1\x96\xa4\xb2\x84\xd6\xdd\xaf\xcd\x7b\x8c\x7e\xc3\xbe\xed\x36\xac\x7b\xde\x94\x12\xc5\x3c\x09\xd3\x28\xb9\xab\x87\xc1\xa0\x65\x8c\xdb\xb3\xdd\xfc\xf3\xbb\x79\x1f\xbe\xba\x0a\xf6\x83\xe0\x5f\x00\x00\x00\xff\xff\x36\x7b\xea\xee\xa1\x05\x00\x00")

func deployIstioWorkspaceOperatorYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceOperatorYaml,
		"deploy/istio-workspace/operator.yaml",
	)
}

func deployIstioWorkspaceOperatorYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceOperatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/operator.yaml", size: 1441, mode: os.FileMode(436), modTime: time.Unix(1560358000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4f\x6b\xfb\x30\x0c\xbd\xe7\x53\x98\x5e\x0a\x3f\xa8\x7f\xec\x36\x72\xdd\x61\xf7\x31\x76\x57\x1d\xb5\x15\xb5\x2d\x23\xc9\x19\xdb\xa7\x1f\x49\xc3\xea\x42\xdb\xd1\x9c\xe2\x17\x3f\xbd\x3f\x11\x14\xfa\x40\x51\xe2\xdc\x3b\xd9\x42\xf0\x50\xed\xc0\x42\xdf\x60\xc4\xd9\x1f\x9f\xd5\x13\xff\x1f\x9f\xba\x23\xe5\xa1\x77\x2f\xb1\xaa\xa1\xbc\x71\xc4\x2e\xa1\xc1\x00\x06\x7d\xe7\x5c\x10\x9c\x09\xef\x94\x50\x0d\x52\xe9\x5d\xae\x31\x76\xce\x65\x48\xd8\x3b\x52\x23\xde\x7c\xb2\x1c\xb5\x40\xc0\x4e\x6a\x44\x9d\x88\x1b\x07\x85\x5e\x85\x6b\x99\x8f\xd3\xb3\x71\xab\xd5\xfc\x2a\xa8\x5c\x25\x60\xf3\xa5\xf0\xa0\xbf\x07\x45\x19\x29\xe0\x19\xc0\x3c\x14\xa6\x6c\x67\xa4\x4c\xd9\xd4\x30\xdb\xc8\xb1\x26\x0c\x11\x28\x35\x84\x11\xdb\xdb\x81\xf3\x8e\xf6\x09\x4a\xab\x11\x04\x97\x2b\x23\xca\xb6\xf1\xb2\xfe\xb7\x7e\x3c\xc0\x54\xc7\x5c\xc1\xd5\x91\x7b\xb4\x5b\x23\xa1\x2c\xae\xae\x0c\x1d\xb0\x44\xfe\x4a\x17\x59\x06\xc0\xc4\x59\xb1\x81\x04\x4b\xa4\x00\x17\x98\x1a\x18\xee\x6a\xd4\xc7\x43\x4e\x8e\x3c\x17\xcc\x7a\xa0\x9d\x79\xe2\xbf\xed\x9d\x0a\x7e\x54\x28\x71\x26\x63\xa1\xbc\xf7\x81\x05\x59\x7d\xe0\x74\x4b\x6c\x59\x8a\x85\x73\xa7\xe5\xe5\x97\x4f\x8b\x8b\xb7\x94\xe7\xb5\x6d\x32\xde\xd1\x3d\xf9\xbf\x1a\xeb\x27\x00\x00\xff\xff\xde\xf7\xfc\x19\x64\x03\x00\x00")

func deployIstioWorkspaceRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceRoleYaml,
		"deploy/istio-workspace/role.yaml",
	)
}

func deployIstioWorkspaceRoleYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/role.yaml", size: 868, mode: os.FileMode(436), modTime: time.Unix(1560259028, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceRole_bindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x4e\x33\x31\x0c\x84\xef\x79\x0a\x1f\xfe\x6b\xf7\x17\x37\x94\x5b\xa9\x2a\x4e\x20\xd4\x22\xee\x6e\xd6\xa5\x66\x77\xe3\xe0\x38\x8b\x04\xe2\xdd\x51\xd3\x5d\x40\x20\x95\xdc\x32\x63\x8d\x3f\x4f\xc7\xb1\xf5\x70\x4f\x43\xea\xd1\xc8\x61\xe2\x07\xd2\xcc\x12\x3d\xd8\x24\x36\x92\x28\xe6\x03\xef\xad\x61\xf9\x3f\x5e\xb8\x84\x8a\x03\x19\x69\xf6\x0e\x60\x01\x11\x07\xf2\x70\xbb\xbc\x59\x6f\xef\x96\xab\xb5\x03\x00\x50\x7a\x2e\xac\xd4\x7a\x30\x2d\x54\xa5\x11\xfb\x42\x1e\x38\x1b\xcb\xe2\x45\xb4\xcb\x09\x03\x2d\x24\x91\xa2\x89\x3a\xd9\x3d\x51\xb0\x29\xf3\xc4\xb5\xea\x4b\x36\xd2\x8d\xf4\x74\xc5\xb1\xe5\xf8\x58\x93\xbe\x53\xea\x0e\x43\x83\xc5\x0e\xa2\xfc\x8a\xc6\x12\x9b\xee\x32\x4f\xa4\xc7\xe1\x81\x0c\x5b\x34\xf4\xf5\x07\x13\xed\x0f\x8a\xea\xe5\xf2\x45\x70\x7c\x33\xc5\x96\x74\xe4\x40\xcb\x10\xa4\x44\x9b\xcc\x73\x41\xb3\x5b\x15\x0f\xff\xde\x3e\xbb\x79\x3f\x95\x23\x3d\x6d\x68\x3f\xef\xf9\x75\xeb\x9f\xa4\xb5\x82\x6b\x95\x92\xce\x14\xe0\x3e\x02\x00\x00\xff\xff\x9e\x78\x82\x8b\xdc\x01\x00\x00")

func deployIstioWorkspaceRole_bindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceRole_bindingYaml,
		"deploy/istio-workspace/role_binding.yaml",
	)
}

func deployIstioWorkspaceRole_bindingYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceRole_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/role_binding.yaml", size: 476, mode: os.FileMode(436), modTime: time.Unix(1560358000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _deployIstioWorkspaceService_accountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0e\x80\x30\x08\x00\xc0\x9d\x57\xf0\x01\x07\x57\x36\xdf\x60\xe2\x4e\x28\x03\x69\x0a\x4d\xc1\xfa\x7d\x8f\xa7\x3d\xba\xd2\xc2\x09\xf7\x09\xdd\xbc\x11\xde\xba\xb6\x89\x5e\x22\xf1\x7a\xc1\xd0\xe2\xc6\xc5\x04\x88\xce\x43\x09\x2d\xcb\xe2\xf8\x62\xf5\x9c\x2c\x0a\x7f\x00\x00\x00\xff\xff\x94\xa0\xb7\x3f\x46\x00\x00\x00")

func deployIstioWorkspaceService_accountYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployIstioWorkspaceService_accountYaml,
		"deploy/istio-workspace/service_account.yaml",
	)
}

func deployIstioWorkspaceService_accountYaml() (*asset, error) {
	bytes, err := deployIstioWorkspaceService_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/istio-workspace/service_account.yaml", size: 70, mode: os.FileMode(436), modTime: time.Unix(1554454542, 0)}
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
	"deploy/istio-workspace/crds/istio_v1alpha1_session_crd.yaml": deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml,
	"deploy/istio-workspace/operator.yaml":                        deployIstioWorkspaceOperatorYaml,
	"deploy/istio-workspace/role.yaml":                            deployIstioWorkspaceRoleYaml,
	"deploy/istio-workspace/role_binding.yaml":                    deployIstioWorkspaceRole_bindingYaml,
	"deploy/istio-workspace/service_account.yaml":                 deployIstioWorkspaceService_accountYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"istio-workspace": &bintree{nil, map[string]*bintree{
			"crds": &bintree{nil, map[string]*bintree{
				"istio_v1alpha1_session_crd.yaml": &bintree{deployIstioWorkspaceCrdsIstio_v1alpha1_session_crdYaml, map[string]*bintree{}},
			}},
			"operator.yaml":        &bintree{deployIstioWorkspaceOperatorYaml, map[string]*bintree{}},
			"role.yaml":            &bintree{deployIstioWorkspaceRoleYaml, map[string]*bintree{}},
			"role_binding.yaml":    &bintree{deployIstioWorkspaceRole_bindingYaml, map[string]*bintree{}},
			"service_account.yaml": &bintree{deployIstioWorkspaceService_accountYaml, map[string]*bintree{}},
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