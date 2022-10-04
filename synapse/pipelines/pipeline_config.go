package pipelines

import "time"

type PipelinesList struct {
	Value []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Etag       string `json:"etag"`
		Properties struct {
			Activities []struct {
				Type           string `json:"type"`
				TypeProperties struct {
					Source struct {
						Type            string `json:"type"`
						PartitionOption string `json:"partitionOption"`
					} `json:"source"`
					Sink struct {
						StoreSettings struct {
							Type string `json:"type"`
						} `json:"storeSettings"`
						FormatSettings struct {
							Type string `json:"type"`
						} `json:"formatSettings"`
						Type string `json:"type"`
					} `json:"sink"`
					EnableStaging           bool `json:"enableStaging"`
					ValidateDataConsistency bool `json:"validateDataConsistency"`
				} `json:"typeProperties"`
				Policy struct {
					Retry                  int    `json:"retry"`
					Timeout                string `json:"timeout"`
					RetryIntervalInSeconds int    `json:"retryIntervalInSeconds"`
					SecureOutput           bool   `json:"secureOutput"`
					SecureInput            bool   `json:"secureInput"`
				} `json:"policy"`
				Name   string `json:"name"`
				Inputs []struct {
					Type          string `json:"type"`
					ReferenceName string `json:"referenceName"`
				} `json:"inputs"`
				Outputs []struct {
					Type          string `json:"type"`
					ReferenceName string `json:"referenceName"`
				} `json:"outputs"`
				UserProperties []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"userProperties"`
			} `json:"activities"`
			LastPublishTime time.Time `json:"lastPublishTime"`
		} `json:"properties"`
	} `json:"value"`
}

//type Pipeline struct {
//	ID         string `json:"id"`
//	Name       string `json:"name"`
//	Type       string `json:"type"`
//	Properties struct {
//		Description string `json:"description"`
//		Activities  []struct {
//			Type           string `json:"type"`
//			TypeProperties struct {
//				IsSequential bool `json:"isSequential"`
//				Items        struct {
//					Value string `json:"value"`
//					Type  string `json:"type"`
//				} `json:"items"`
//				Activities []struct {
//					Type           string `json:"type"`
//					TypeProperties struct {
//						Source struct {
//							Type string `json:"type"`
//						} `json:"source"`
//						Sink struct {
//							Type string `json:"type"`
//						} `json:"sink"`
//						DataIntegrationUnits int `json:"dataIntegrationUnits"`
//					} `json:"typeProperties"`
//					Inputs []struct {
//						ReferenceName string `json:"referenceName"`
//						Parameters    struct {
//							MyFolderPath string `json:"MyFolderPath"`
//							MyFileName   string `json:"MyFileName"`
//						} `json:"parameters"`
//						Type string `json:"type"`
//					} `json:"inputs"`
//					Outputs []struct {
//						ReferenceName string `json:"referenceName"`
//						Parameters    struct {
//							MyFolderPath string `json:"MyFolderPath"`
//							MyFileName   struct {
//								Value string `json:"value"`
//								Type  string `json:"type"`
//							} `json:"MyFileName"`
//						} `json:"parameters"`
//						Type string `json:"type"`
//					} `json:"outputs"`
//					Name string `json:"name"`
//				} `json:"activities"`
//			} `json:"typeProperties"`
//			Name string `json:"name"`
//		} `json:"activities"`
//		Parameters struct {
//			OutputBlobNameList struct {
//				Type string `json:"type"`
//			} `json:"OutputBlobNameList"`
//		} `json:"parameters"`
//		Variables struct {
//			TestVariableArray struct {
//				Type string `json:"type"`
//			} `json:"TestVariableArray"`
//		} `json:"variables"`
//	} `json:"properties"`
//	Etag string `json:"etag"`
//}

type GetPipelineDet struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		Parameters struct {
			CwItems struct {
				Type         string `json:"type"`
				DefaultValue []struct {
					Source struct {
						Table string `json:"table"`
					} `json:"source"`
					Destination struct {
						FileName string `json:"fileName"`
					} `json:"destination"`
				} `json:"defaultValue"`
			} `json:"cw_items"`
		} `json:"parameters"`
		Activities []struct {
			Type           string `json:"type"`
			TypeProperties struct {
				Activities []struct {
					Type           string `json:"type"`
					TypeProperties struct {
						Source struct {
							Type            string `json:"type"`
							PartitionOption string `json:"partitionOption"`
						} `json:"source"`
						Sink struct {
							StoreSettings struct {
								Type string `json:"type"`
							} `json:"storeSettings"`
							FormatSettings struct {
								Type          string `json:"type"`
								QuoteAllText  bool   `json:"quoteAllText"`
								FileExtension string `json:"fileExtension"`
							} `json:"formatSettings"`
							Type string `json:"type"`
						} `json:"sink"`
						Translator              interface{} `json:"translator"`
						EnableStaging           bool        `json:"enableStaging"`
						ValidateDataConsistency bool        `json:"validateDataConsistency"`
					} `json:"typeProperties"`
					Policy struct {
						Retry                  int    `json:"retry"`
						Timeout                string `json:"timeout"`
						RetryIntervalInSeconds int    `json:"retryIntervalInSeconds"`
						SecureOutput           bool   `json:"secureOutput"`
						SecureInput            bool   `json:"secureInput"`
					} `json:"policy"`
					Name   string `json:"name"`
					Inputs []struct {
						Type          string `json:"type"`
						ReferenceName string `json:"referenceName"`
						Parameters    struct {
							CwTable string `json:"cw_table"`
						} `json:"parameters"`
					} `json:"inputs"`
					Outputs []struct {
						Type          string `json:"type"`
						ReferenceName string `json:"referenceName"`
						Parameters    struct {
							CwFileName string `json:"cw_fileName"`
						} `json:"parameters"`
					} `json:"outputs"`
					UserProperties []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"userProperties"`
				} `json:"activities"`
				Source struct {
					Type            string `json:"type"`
					PartitionOption string `json:"partitionOption"`
				} `json:"source"`
				Sink struct {
					StoreSettings struct {
						Type string `json:"type"`
					} `json:"storeSettings"`
					FormatSettings struct {
						Type string `json:"type"`
					} `json:"formatSettings"`
					Type string `json:"type"`
				} `json:"sink"`
				EnableStaging           bool `json:"enableStaging"`
				ValidateDataConsistency bool `json:"validateDataConsistency"`
			} `json:"typeProperties"`
			Policy struct {
				Retry                  int    `json:"retry"`
				Timeout                string `json:"timeout"`
				RetryIntervalInSeconds int    `json:"retryIntervalInSeconds"`
				SecureOutput           bool   `json:"secureOutput"`
				SecureInput            bool   `json:"secureInput"`
			} `json:"policy"`
			Name   string `json:"name"`
			Inputs []struct {
				Type          string `json:"type"`
				ReferenceName string `json:"referenceName"`
			} `json:"inputs"`
			Outputs []struct {
				Type          string `json:"type"`
				ReferenceName string `json:"referenceName"`
			} `json:"outputs"`
			UserProperties []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"userProperties"`
		} `json:"activities"`
		LastPublishTime time.Time `json:"lastPublishTime"`
	} `json:"properties"`
	Etag string `json:"etag"`
}
