package uchiwa

import (
    "github.com/sensu/uchiwa/uchiwa/logger"
    "github.com/sensu/uchiwa/uchiwa/structs"
)

// ClearSilenced send a POST request to the /stashes endpoint in order to create a stash
func (u *Uchiwa) ClearSilenced(data structs.Silence) error {
	api, err := getAPI(u.Datacenters, data.Dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	_, err = api.ClearSilenced(data)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}

// PostSilence send a POST request to the /stashes endpoint in order to create a stash
func (u *Uchiwa) PostSilence(data structs.Silence) error {
	api, err := getAPI(u.Datacenters, data.Dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	_, err = api.Silence(data)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}
