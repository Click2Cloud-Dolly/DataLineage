package datasets

const Parquet = "Parquet"
const AzureSQLDWTable = "AzureSqlDWTable"
const SqlPoolTable = "SqlPoolTable"
const AzureSQLTable = "AzureSqlTable"

type ParquetDetails struct {
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

type AzureSQLDWTableDetails struct {
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

type SqlPoolTableDetails struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type           string `json:"type"`
		TypeProperties struct {
			Schema string `json:"schema"`
			Table  string `json:"table"`
		} `json:"typeProperties"`
		SqlPool struct {
			Type          string `json:"type"`
			ReferenceName string `json:"referenceName"`
		} `json:"sqlPool"`
		LinkedServiceName struct {
			Type          string `json:"type"`
			ReferenceName string `json:"referenceName"`
		} `json:"linkedServiceName"`
		Schema []interface{} `json:"schema"`
	} `json:"properties"`
	Etag string `json:"etag"`
}

type GetDataset struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type           string `json:"type"`
		TypeProperties struct {
			FolderPath struct {
				Value string `json:"value"`
				Type  string `json:"type"`
			} `json:"folderPath"`
			FileName struct {
				Value string `json:"value"`
				Type  string `json:"type"`
			} `json:"fileName"`
			Format struct {
				Type string `json:"type"`
			} `json:"format"`
		} `json:"typeProperties"`
		Description       string `json:"description"`
		LinkedServiceName struct {
			ReferenceName string `json:"referenceName"`
			Type          string `json:"type"`
		} `json:"linkedServiceName"`
		Parameters struct {
			MyFolderPath struct {
				Type string `json:"type"`
			} `json:"MyFolderPath"`
			MyFileName struct {
				Type string `json:"type"`
			} `json:"MyFileName"`
		} `json:"parameters"`
	} `json:"properties"`
	Etag string `json:"etag"`
}
type GetDatasetDetails struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Type              string `json:"type"`
		Description       string `json:"description"`
		LinkedServiceName struct {
			ReferenceName string `json:"referenceName"`
			Type          string `json:"type"`
		} `json:"linkedServiceName"`
	} `json:"properties"`
	Etag string `json:"etag"`
}

type AzureSQLTableDetails struct {
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
	Parquet   ParquetDetails
	DWSQLPool AzureSQLDWTableDetails
	SQLPool   SqlPoolTableDetails
	SQLTable  AzureSQLTableDetails
}

type Mappings struct {
	Name         string   `json:"Name"`
	Type         string   `json:"Type"`
	Location     string   `json:"Location"`
	Source       string   `json:"Source"`
	TDestination []string `json:"TDestination"`
}
