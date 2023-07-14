package words

import (
	"encoding/json"
	"io/ioutil"
)

type Words struct {
	Words []string `json:"forbidden_words"`
}

func ReadWordsFromFile(filepath string) ([]string, error) {
	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var words Words
	err = json.Unmarshal(fileData, &words)
	if err != nil {
		return nil, err
	}

	return words.Words, nil
}
