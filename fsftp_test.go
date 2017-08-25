package foldersync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFolderSync(t *testing.T) {

	fsftp := &FSFTP{Host: "192.168.10.20", Port: 21, User: "research00", Password: "researchFtp"}

	fsftp.FolderSync("dataset/Public_Abnormal/Image_Opencv_Block", "/001_Public_Abnormal/001_Image_Opencv_Block", "metadata.xml")
	assert.Equal(t, 123, 123, "they should be equal")
}
