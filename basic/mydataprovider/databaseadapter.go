package mydataprovider

import (
	"io"
	"github.com/rihtim/core/log"
	"github.com/rihtim/core/utils"
)

type Provider struct {
}

func (da *Provider) Init() (err *utils.Error) {

	log.Info("Database initializing...")
	return
}

func (da *Provider) Connect() (err *utils.Error){

	log.Info("Database connecting...")
	return
}

func (da *Provider) Create(collection string, data map[string]interface{}) (response map[string]interface{}, err *utils.Error){

	log.Info("Creating object")

	response = map[string]interface{}{
		"_id": "57b36760e63bce1eeb000004",
		"createdAt": 123456789,
		"updatedAt": 123456789,
	}

	return
}

func (da *Provider) Get(collection string, id string) (response map[string]interface{}, err *utils.Error){

	log.Info("Getting object")

	response = map[string]interface{}{
		"_id": "57b36760e63bce1eeb000004",
		"firstName": "Ahmet",
		"lastName": "Öztürk",
		"email": "ahmet@ozturk.com",
		"createdAt": 1597536000,
		"updatedAt": 1597536000,
	}

	return
}

func (da *Provider) Query(collection string, parameters map[string][]string) (response map[string]interface{}, err *utils.Error){

	log.Info("Querying collection")

	results := make([]map[string]interface{}, 1)
	results[0] = map[string]interface{}{
		"_id": "57b36760e63bce1eeb000004",
		"firstName": "Ahmet",
		"lastName": "Öztürk",
		"email": "ahmet@ozturk.com",
		"createdAt": 1597536000,
		"updatedAt": 1597536000,
	}

	response = map[string]interface{}{
		"results": results,
	}

	return
}

func (da *Provider) Update(collection string, id string, data map[string]interface{}) (response map[string]interface{}, err *utils.Error){

	log.Info("Updating object")

	response = map[string]interface{}{
		"updatedAt": 1568160000,
	}

	return
}

func (da *Provider) Delete(collection string, id string) (response map[string]interface{}, err *utils.Error){

	log.Info("Deleting object")
	return
}

func (da *Provider) CreateFile(data io.ReadCloser) (response map[string]interface{}, err *utils.Error){

	log.Info("Creating file")
	return
}

func (da *Provider) GetFile(id string) (response []byte, err *utils.Error){

	log.Info("Getting file")
	return
}
