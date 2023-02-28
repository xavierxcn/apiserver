package image

import (
	"github.com/xavierxcn/apiserver/internal/serve/pkg/storage"
	"gorm.io/gorm"
)

// TImage 镜像版本
type TImage struct {
	gorm.Model
	ImageName string `gorm:"column:image_name;type:varchar(255);not_null;default:'';uniqueIndex:t_image_version_index"` // 镜像名称
	Version   string `gorm:"column:version;type:varchar(31);not_null;uniqueIndex:t_image_version_index"`                // 版本
}

func (m *TImage) TableName() string {
	return "t_image"
}

// Create 创建
func (m *TImage) Create() error {
	return storage.Client().Create(m).Error
}

// Update 更新
func (m *TImage) Update() error {
	return storage.Client().Save(m).Error
}

// Delete 删除
func (m *TImage) Delete() error {
	return storage.Client().Delete(m).Error
}

// FindByImageName 通过镜像ID查找
func (m *TImage) FindByImageName(imageName string) ([]*TImage, error) {
	var imageVersions []*TImage
	err := storage.Client().Where("image_name = ?", imageName).Find(&imageVersions).Error
	return imageVersions, err
}

// FindLatestVersion 查找最新版本
func (m *TImage) FindLatestVersion() error {
	return storage.Client().Where("image_name = ?", m.ImageName).Order("created_at desc").First(m).Error
}
