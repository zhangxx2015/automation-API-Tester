package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

func GenUUID32() string {
	uuidStr := uuid.NewString()
	return strings.ReplaceAll(uuidStr, "-", "")
}

func CurrDir(file string) string {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currDir, file)
}
func FileExist(fullPath string) bool {
	// 判断文件是否存在
	_, err := os.Stat(fullPath)
	if err != nil {
		// 如果文件不存在，err是nil，但是os.IsNotExist(err)会返回true
		if os.IsNotExist(err) {
			panic(fmt.Errorf("The file does not exist.[%v]", fullPath))
		}
		panic(fmt.Errorf("An error occurred:[%v]", err))
	}
	return true
}
func TextFromFile(fullPath string) string {
	FileExist(fullPath)
	data, err := os.ReadFile(fullPath) // Go 1.16及以后版本
	if err != nil {
		panic(err)
	}
	return string(data)
}
func JsonFromFile(fullPath string, v any) {
	text := TextFromFile(fullPath)
	err := json.Unmarshal([]byte(text), &v)
	if err != nil {
		panic(err)
	}
}
func DictFromString(jsonStr string) map[string]interface{} {
	dict := make(map[string]interface{}, 0)
	if len(jsonStr) > 0 {
		err := json.Unmarshal([]byte(jsonStr), &dict)
		if err != nil {
			panic(err)
		}
	}
	return dict
}

func FromJsonString(jsonData string, v any) bool {
	err := json.Unmarshal([]byte(jsonData), &v)
	if nil != err {
		panic(err)
	}
	return true
}

func ToJsonString(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if nil != err {
		panic(errors.New(fmt.Sprintf("Failed to marshal err:%v", err)))
	}
	return string(bytes)
}
func ToJsonStringIndent(obj interface{}, prefix string) string {
	bytes, err := json.MarshalIndent(obj, prefix, "  ")
	if nil != err {
		panic(errors.New(fmt.Sprintf("Failed to marshal err:%v", err)))
	}
	return string(bytes)
}
