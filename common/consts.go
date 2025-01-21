package common

type DataSourceType int

const (
	CLICKHOUSE DataSourceType = iota
	DORIS
	MYSQL
	ELASTICSEARCH
	HIVE
	STARROCKS
	ODPS
	POSTGRES
)
