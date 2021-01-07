package source

import (
	"acnh/data/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Load provides
func (s *Source) Load() error {
	api := config.Config.API()
	models := config.Config.Models()
	all := make([]Item, 0)

	for _, model := range models {
		res, e := s.fromAPI(api, model)
		// api error
		if e != nil {
			return e
		}
		// unmarshal from byte slice to []Item
		data, err := s.unmarshal(res)
		if err != nil {
			return err
		}
		cnfErr := s.AddConfigToData(model, &data)
		// error setting config data
		if cnfErr != nil {
			return cnfErr
		}

		all = append(all, data...)
	}

	s.All = all

	return nil
}

// AddConfigToData iterates over each data item and appends the relevant Config details
// so each Item knows its original source type etc
// NOTE: unsure if needs to be a pointer here - check later
func (s *Source) AddConfigToData(configName string, data *[]Item) error {
	cnf, ok := config.Config.ModelConfigs[configName]
	if ok == false {
		return fmt.Errorf("Failed to find [%v] in model config", configName)
	}

	for i := range *data {
		(*data)[i].AddConfig(cnf)
	}

	return nil

}

// unmarshal handles the unmarshaling of the byte array from the api call
func (s *Source) unmarshal(res []byte) (data []Item, err error) {
	data = make([]Item, 0)
	err = json.Unmarshal(res, &data)
	return
}

// fromAPI makes the http client call and returns byte array
func (s *Source) fromAPI(
	api string,
	model string) ([]byte, error) {

	url := fmt.Sprintf("%s%s", api, model)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
