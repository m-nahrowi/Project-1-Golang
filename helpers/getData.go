package helpers

import (
	"encoding/json"
	"io/ioutil"
)

type Teman struct {
	Nama      string `json:"Nama"`
	Alamat    string `json:"Alamat"`
	Pekerjaan string `json:"Pekerjaan"`
	Alasan    string `json:"Alasan"`
}

func Data() ([]Teman, error) {
	// Baca file data.json
	file, err := ioutil.ReadFile("helpers/data.json")
	if err != nil {
		return nil, err
	}

	// Dekode data JSON menjadi slice Teman
	var teman []Teman
	err = json.Unmarshal(file, &teman)
	if err != nil {
		return nil, err
	}

	return teman, nil
}
