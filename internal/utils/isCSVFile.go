package utils

import (
	"path/filepath"
)

func isCSVFile(fileName string) bool {
	return filepath.Ext(fileName) == ".csv"
   }