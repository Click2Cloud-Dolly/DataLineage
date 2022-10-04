package controllers

import (
	"azureDataLineage/synapse/datasets"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	imageLocation = "../../assets/image/chart-img/"
)

type Nodes struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
	Title string `json:"title"`
	Image string `json:"image"`
}
type Direction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Label string `json:"label"`
	Color string `json:"color"`
}

func NodesLineage(dataList []datasets.Mappings) ([]byte, []byte, error) {
	var directions []Direction

	var finalnodes []Nodes
	for _, data := range dataList {
		finalnodes = append(finalnodes, Nodes{
			ID:    strings.Join([]string{data.Location, data.Name}, ";"),
			Label: data.Name,
			Color: "#ffc3bf",
			Title: fmt.Sprintf("Location: %s \n Type: %s", data.Location, data.Type),
			Image: strings.Join([]string{imageLocation, data.Type, ".png"}, ""),
		})
		if data.TDestination != nil {
			for _, destination := range data.TDestination {
				directions = append(directions, Direction{
					From:  strings.Join([]string{data.Location, data.Name}, ";"),
					To:    destination,
					Color: "#000",
					Label: "Sends to",
				})
			}
		}
	}

	var occurred = make(map[string]bool)
	result := []Nodes{}
	for _, fnode := range finalnodes {
		if occurred[fnode.ID] != true {
			occurred[fnode.ID] = true
			result = append(result, fnode)

		}

	}
	nodeJson, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return nil, nil, err
	}
	nodeDirJson, err := json.MarshalIndent(directions, "", " ")
	if err != nil {
		return nil, nil, err
	}
	return nodeJson, nodeDirJson, nil
}
