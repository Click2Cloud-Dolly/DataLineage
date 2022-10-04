package main

import (
	"fmt"
	//"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBDetails struct {
	//gorm.Model
	ClientId       string `gorm:"primaryKey"`
	SubscriptionId string `json:"subscriptionId"`
	NodeData       string `json:"nodeData"`
	DirectionData  string `json:"directionData"`
}

func main() {

	//DB_HOST := os.Getenv("DB_HOST")
	//fmt.Printf("DB_HOST:%v \n", DB_HOST)
	//
	//DB_PORT := os.Getenv("DB_PORT")
	//fmt.Printf("DB_PORT:%v \n", DB_PORT)
	//
	//MYSQL_USER := os.Getenv("MYSQL_USER")
	//fmt.Printf("MYSQL_USER:%v \n", MYSQL_USER)
	//
	//MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	//fmt.Printf("MYSQL_PASSWORD:%v \n", MYSQL_PASSWORD)
	//
	//MYSQL_DB := os.Getenv("MYSQL_DB")
	//fmt.Printf("MYSQL_DB:%v \n", MYSQL_DB)

	DNS := "root" + ":" + "ROOT#123" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "datalineagedb"
	fmt.Println("DNS", DNS)

	var DB *gorm.DB
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic("Database not connected")
	}

	DB.AutoMigrate(&DBDetails{})

	//cid := "c75694ad-2861-42fc-8e24-ec6e04d93eab"
	//var getNode DBDetails
	//
	//DB.Where(&DBDetails{ClientId: cid}).Find(&getNode)
	//fmt.Println("check:", getNode.NodeData)
}
