package common

import (
	"fmt"
	"testing"
)

func TestGetDataSourceType(t *testing.T) {
	sourceType := GetDataSourceType("clickhouse")

	fmt.Println(sourceType)
}
