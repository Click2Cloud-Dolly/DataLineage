package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "192.168.1.35"
	port     = 5432
	user     = "postgres"
	password = "dbroot@123"
	dbname   = "datalineage"
)

type NodeDetails struct {
	gorm.Model
	ClientId       string
	SubscriptionId string
	NodeData       string
	DirectionData  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	//err = db.AutoMigrate(&NodeDetails{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	cid := "c75694ad-2861-42fc-8e24-ec6e04d93eab"
	//sid := "33d5c9e6-809b-4dc3-adc0-102402363283"
	//nodes := "[\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_cwb\",\n  \"label\": \"SourceDataset_cwb\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagesql.database.windows.net,datalineagedb;Campaign_Analytics\",\n  \"label\": \"Campaign_Analytics\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagesql.database.windows.net,datalineagedb \\n Type: AzureSqlTable\"\n },\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_cwb\",\n  \"label\": \"DestinationDataset_cwb\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;Campaign_Analytics\",\n  \"label\": \"Campaign_Analytics\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagews1.sql.azuresynapse.net,ddsqlpool \\n Type: AzureSqlDWTable\"\n },\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_s4v\",\n  \"label\": \"SourceDataset_s4v\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagesql.database.windows.net,datalineagedb;SalesDataBefore\",\n  \"label\": \"SalesDataBefore\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagesql.database.windows.net,datalineagedb \\n Type: AzureSqlTable\"\n },\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_s4v\",\n  \"label\": \"DestinationDataset_s4v\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataBefore\",\n  \"label\": \"SalesDataBefore\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagews1.sql.azuresynapse.net,ddsqlpool \\n Type: AzureSqlDWTable\"\n },\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_a1v\",\n  \"label\": \"SourceDataset_a1v\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagesql.database.windows.net,datalineagedb;SalesDataAfter\",\n  \"label\": \"SalesDataAfter\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagesql.database.windows.net,datalineagedb \\n Type: AzureSqlTable\"\n },\n {\n  \"id\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_a1v\",\n  \"label\": \"DestinationDataset_a1v\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://datalineagews1.dev.azuresynapse.net \\n Type: DatasetReference\"\n },\n {\n  \"id\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataAfter\",\n  \"label\": \"SalesDataAfter\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: datalineagews1.sql.azuresynapse.net,ddsqlpool \\n Type: AzureSqlDWTable\"\n },\n {\n  \"id\": \"https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53;Salesdata\",\n  \"label\": \"Salesdata\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53 \\n Type: PowerBIReport\"\n },\n {\n  \"id\": \"https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53;Salesdata\",\n  \"label\": \"Salesdata\",\n  \"color\": \"#ffc3bf\",\n  \"title\": \"Location: https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53 \\n Type: PowerBIReport\"\n }\n]"
	//json.Unmarshal([]byte(nodes), &node1)
	//nodeDir := "[\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_cwb\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_cwb\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagesql.database.windows.net,datalineagedb;Campaign_Analytics\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_cwb\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_cwb\",\n  \"to\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;Campaign_Analytics\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;Campaign_Analytics\",\n  \"to\": \"https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53;Salesdata\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_s4v\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_s4v\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagesql.database.windows.net,datalineagedb;SalesDataBefore\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_s4v\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_s4v\",\n  \"to\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataBefore\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataBefore\",\n  \"to\": \"https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53;Salesdata\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_a1v\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_a1v\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagesql.database.windows.net,datalineagedb;SalesDataAfter\",\n  \"to\": \"https://datalineagews1.dev.azuresynapse.net;SourceDataset_a1v\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"https://datalineagews1.dev.azuresynapse.net;DestinationDataset_a1v\",\n  \"to\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataAfter\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n },\n {\n  \"from\": \"datalineagews1.sql.azuresynapse.net,ddsqlpool;SalesDataAfter\",\n  \"to\": \"https://app.powerbi.com/groups/9224e2c1-c244-425e-b9d9-355a88498df4/reports/7a35e98d-5dfc-4bb4-94d1-9a10abf0bc53;Salesdata\",\n  \"label\": \"Sends to\",\n  \"color\": \"#ffc3bf\"\n }\n]"
	//json.Unmarshal([]byte(nodeDir), &dir1)
	//db.Create(&NodeDetails{ClientId: cid, SubscriptionId: sid, NodeData: nodes, DirectionData: nodeDir})
	var getNode NodeDetails
	//db.Find(&getNode, "id=?", cid)

	db.Where(&NodeDetails{ClientId: cid}).Find(&getNode)
	fmt.Println("check:", getNode.NodeData)

	//dbs, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = dbs.Ping()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("connected")
	//}
	//res, err := dbs.Exec("create table DataLineage(clientId varchar(100), subscriptionId varchar(100),nodeData text,directionData text)")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//rs, _ := res.RowsAffected()
	//ls, _ := res.LastInsertId()
	//
	//fmt.Println("Result:", rs, ls)
	//cid := "c75694ad-2861-42fc-8e24-ec6e04d93eab"
	//sid := "33d5c9e6-809b-4dc3-adc0-102402363283"
	//ldata := "{\n\"sql_server\": {\n\"Name\": \"SQLTable1\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"\",\n\"Destination\": [\n{\n\"Name\": \"SD1\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"SQLTable1\",\n\"Destination\": [\n{\n\"Name\": \"DD1\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"SD1\",\n\"Destination\": [\n{\n\"Name\": \"testparquet\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"DD1\",\n\"Destination\": \"\"\n}\n]\n}\n]\n},\n{\n\"Name\": \"SD2\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"SQLTable1\",\n\"Destination\": [\n{\n\"Name\": \"DD1\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"SD2\",\n\"Destination\": [\n{\n\"Name\": \"SQLPool1\",\n\"Type\": \"\",\n\"Location\": \"\",\n\"Source\": \"DD1\",\n\"Destination\": \"\"\n}\n]\n}\n]\n}\n]\n}\n}"
	//res1, err := dbs.Exec("insert into DataLineage(clientId ,subscriptionId ,lineagedata) values ($1,$2,$3)", cid, sid, ldata)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//l1, err := res1.RowsAffected()
	//fmt.Println("rs:", l1)
}
