package datasets

const Parquet = "Parquet"
const AzureSQLTable = "AzureSqlTable"
const AzureSQLDWTable = "AzureSqlDWTable"
const SqlPoolTable = "SqlPoolTable"

type DFListDatasets struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type           string `json:"type"`
		TypeProperties struct {
			Schema string `json:"schema"`
			Table  string `json:"table"`
		} `json:"typeProperties"`
		LinkedServiceName struct {
			Type          string `json:"type"`
			ReferenceName string `json:"referenceName"`
		} `json:"linkedServiceName"`
	} `json:"properties"`
	Etag string `json:"etag"`
}

type DfParquetDetails struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type           string `json:"type"`
		TypeProperties struct {
			Location struct {
				Type       string `json:"type"`
				FileName   string `json:"fileName"`
				FileSystem string `json:"fileSystem"`
			} `json:"location"`
			CompressionCodec string `json:"compressionCodec"`
		} `json:"typeProperties"`
		LinkedServiceName struct {
			Type          string `json:"type"`
			ReferenceName string `json:"referenceName"`
		} `json:"linkedServiceName"`
		Schema []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"schema"`
	} `json:"properties"`
	Etag string `json:"etag"`
}
type DfAzureSQLTableDetails struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type           string `json:"type"`
		TypeProperties struct {
			Schema string `json:"schema"`
			Table  string `json:"table"`
		} `json:"typeProperties"`
		LinkedServiceName struct {
			Type          string `json:"type"`
			ReferenceName string `json:"referenceName"`
		} `json:"linkedServiceName"`
		Schema []interface{} `json:"schema"`
	} `json:"properties"`
	Etag string `json:"etag"`
}

type Dataset struct {
	DfParquet  DfParquetDetails
	DfSQLTable DfAzureSQLTableDetails
	////DWSQLPool AzureSQLDWTableDetails
	//SQLPool  SqlPoolTableDetails

}

type DFMappings struct {
	Name         string   `json:"Name"`
	Type         string   `json:"Type"`
	Location     string   `json:"Location"`
	Source       string   `json:"Source"`
	TDestination []string `json:"TDestination"`
}
