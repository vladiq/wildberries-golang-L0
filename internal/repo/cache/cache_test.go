package cache

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"wb_l0/internal"
	"wb_l0/internal/services/data_generator"
)

const testDataPath = "./testdata/model.json"

func TestRepository_GetById(t *testing.T) {
	cacheRepo := NewCacheRepository()
	generatedData, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Errorf("TestRepository_GetById FAILED with error: %v", err)
	}

	var originalJson data_generator.ModelJSON
	err = json.Unmarshal(generatedData, &originalJson)
	if err != nil {
		t.Errorf("TestRepository_GetById FAILED with error: %v", err)
	}
	transformedJson := internal.MapGeneratedToStored(&originalJson)
	cacheRepo.Insert(*transformedJson)

	result, err := cacheRepo.GetById(transformedJson.OrderUid)
	if !cmp.Equal(result, transformedJson) {
		t.Error("TestRepository_GetById FAILED")
	} else {
		t.Log("TestRepository_GetById PASSED")
	}
}

func TestRepository_All(t *testing.T) {
	cacheRepo := NewCacheRepository()
	generatedData, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Errorf("TestRepository_All FAILED with error: %v", err)
	}

	var originalJson data_generator.ModelJSON
	err = json.Unmarshal(generatedData, &originalJson)
	if err != nil {
		t.Errorf("TestRepository_All FAILED with error: %v", err)
	}
	transformedJson := internal.MapGeneratedToStored(&originalJson)
	cacheRepo.Insert(*transformedJson)

	records, err := cacheRepo.All()
	result := records[0]
	if !cmp.Equal(result, *transformedJson) {
		t.Error("TestRepository_All FAILED")
	} else {
		t.Log("TestRepository_All PASSED")
	}
}

func TestRepository_Insert(t *testing.T) {
	cacheRepo := NewCacheRepository()
	generatedData, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Errorf("TestRepository_Insert FAILED with error: %v", err)
	}

	var originalJson data_generator.ModelJSON
	err = json.Unmarshal(generatedData, &originalJson)
	if err != nil {
		t.Errorf("TestRepository_Insert FAILED with error: %v", err)
	}
	transformedJson := internal.MapGeneratedToStored(&originalJson)
	cacheRepo.Insert(*transformedJson)

	result, err := cacheRepo.GetById(transformedJson.OrderUid)
	if !cmp.Equal(result, transformedJson) {
		t.Error("TestRepository_Insert FAILED")
	} else {
		t.Log("TestRepository_Insert PASSED")
	}
}
