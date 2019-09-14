package skilltrees

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//ParseSkills ...
func ParseSkills(Type string) *[]Skill {

	if Type == "frontend" {
		frontFile, err := ioutil.ReadFile("skilltrees/skillTree_front.json")

		if err != nil {
			fmt.Printf("Error with config path: %v %v\n", "skilltrees/skillTree_front.json", err)
		}

		Skills := []Skill{}

		_ = json.Unmarshal(frontFile, &Skills)

		return &Skills
	}

	backFile, err := ioutil.ReadFile("skilltrees/skillTree_back.json")

	if err != nil {
		fmt.Printf("Error with config path: %v\n", "skilltrees/skillTree_back.json")
	}

	Skills := []Skill{}

	_ = json.Unmarshal(backFile, &Skills)

	return &Skills

}
