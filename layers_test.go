package geoserver

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetshpFiledsName(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	storename := gsCatalog.GetshpFiledsName("hisham.zip")
	assert.Equal(t, storename, "hisham")
}
func TestGetLayers(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	layers, err := gsCatalog.GetLayers("nurc")
	assert.NotNil(t, layers)
	assert.Nil(t, err)
}

func TestGetLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	layer, err := gsCatalog.GetLayer("topp", "tasmania_cities")
	assert.NotNil(t, layer)
	assert.Nil(t, err)
}
func TestUpdateLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	modified, err := gsCatalog.UpdateLayer("topp", "tasmania_cities", Layer{
		Attribution: &Attribution{
			Title: "Test Title",
		},
	})
	assert.True(t, modified)
	assert.Nil(t, err)
}
func TestUploadShapeFile(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	zippedShapefile := filepath.Join(gsCatalog.getGoGeoserverPackageDir(), "test_sample", "hurricane_tracks.zip")
	uploaded, err := gsCatalog.UploadShapeFile(zippedShapefile, "shapefileWorkspace", "")
	assert.True(t, uploaded)
	assert.Nil(t, err)
}
func TestDeleteLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	deleted, err := gsCatalog.DeleteLayer("sf", "bugsites", true)
	assert.True(t, deleted)
	assert.Nil(t, err)
}
func TestGeoserverImplemetLayerService(t *testing.T) {
	gsCatalog := reflect.TypeOf(&GeoServer{})
	LayerServiceType := reflect.TypeOf((*LayerService)(nil)).Elem()
	check := gsCatalog.Implements(LayerServiceType)
	assert.True(t, check)
}
