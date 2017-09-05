package buildah

import (
	"io"

	dockertypes "github.com/docker/engine-api/types"
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/docker"
	s2itar "github.com/openshift/source-to-image/pkg/tar"
	"github.com/openshift/source-to-image/pkg/util/fs"
)

type Buildah struct {
}

func (b *Buildah) IsImageInLocalRegistry(name string) (bool, error) {
	return false, nil
}

func (b *Buildah) IsImageOnBuild(string) bool {
	return false
}
func (b *Buildah) GetOnBuild(string) ([]string, error) {
	return []string{}, nil
}
func (b *Buildah) RemoveContainer(id string) error {
	return nil
}
func (b *Buildah) GetScriptsURL(name string) (string, error) {
	return "", nil
}
func (b *Buildah) GetAssembleInputFiles(string) (string, error) {
	return "", nil
}

func (b *Buildah) RunContainer(opts docker.RunContainerOptions) error {
	return nil
}
func (b *Buildah) GetImageID(name string) (string, error) {
	return "", nil
}
func (b *Buildah) GetImageWorkdir(name string) (string, error) {
	return "", nil
}
func (b *Buildah) CommitContainer(opts docker.CommitContainerOptions) (string, error) {
	return "", nil
}
func (b *Buildah) RemoveImage(name string) error {
	return nil
}
func (b *Buildah) CheckImage(name string) (*api.Image, error) {
	return nil, nil
}
func (b *Buildah) PullImage(name string) (*api.Image, error) {
	return nil, nil
}
func (b *Buildah) CheckAndPullImage(name string) (*api.Image, error) {
	return nil, nil
}
func (b *Buildah) BuildImage(opts docker.BuildImageOptions) error {
	return nil
}
func (b *Buildah) GetImageUser(name string) (string, error) {
	return "", nil
}
func (b *Buildah) GetImageEntrypoint(name string) ([]string, error) {
	return []string{}, nil
}

func (b *Buildah) GetLabels(name string) (map[string]string, error) {
	labels := make(map[string]string)
	return labels, nil
}
func (b *Buildah) UploadToContainer(fs fs.FileSystem, srcPath, destPath, container string) error {
	return nil
}
func (b *Buildah) UploadToContainerWithTarWriter(fs fs.FileSystem, srcPath, destPath, container string, makeTarWriter func(io.Writer) s2itar.Writer) error {
	return nil
}
func (b *Buildah) DownloadFromContainer(containerPath string, w io.Writer, container string) error {
	return nil
}
func (b *Buildah) Version() (dockertypes.Version, error) {
	return dockertypes.Version{}, nil
}
func (b *Buildah) CheckReachable() error {
	return nil
}
