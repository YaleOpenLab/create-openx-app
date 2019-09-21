package core

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"

	edb "github.com/Varunram/essentials/database"
	utils "github.com/Varunram/essentials/utils"
	xlm "github.com/YaleOpenLab/openx/chains/xlm"
	wallet "github.com/YaleOpenLab/openx/chains/xlm/wallet"
	openx "github.com/YaleOpenLab/openx/database"

	consts "github.com/org/plat/consts"
	notif "github.com/org/plat/notif"
)

// Entity defines a common structure for contractors, developers and originators
type Entity struct {
	// U is the base User class inherited from openx
	U *openx.User
	// EntityType can be used to define different entities in the system
	EntityType bool
}

// RetrieveAllEntitiesWithoutRole retrieves all the entities from the database
func RetrieveAllEntitiesWithoutRole() ([]Entity, error) {
	var users []Entity
	x, err := edb.RetrieveAllKeys(consts.DbDir+consts.DbName, EntityBucket)
	if err != nil {
		return users, errors.Wrap(err, "error while retrieving all keys")
	}

	for _, value := range x {
		var temp Entity
		err = json.Unmarshal(value, &temp)
		if err != nil {
			return users, errors.New("could not unmarshal json")
		}
		users = append(users, temp)
	}

	return users, nil
}

// RetrieveAllEntities gets all the proposed contracts associated with a particular entity
func RetrieveAllEntities(role string) ([]Entity, error) {
	var entities []Entity

	x, err := edb.RetrieveAllKeys(consts.DbDir+consts.DbName, EntityBucket)
	if err != nil {
		return entities, errors.Wrap(err, "error while retrieving all keys")
	}

	for _, value := range x {
		var entity Entity
		err = json.Unmarshal(value, &entity)
		if err != nil {
			return entities, errors.New("could not unmarshal entity")
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// RetrieveEntityHelper is a helper associated with the RetrieveEntity function
func RetrieveEntityHelper(key int) (Entity, error) {
	var entity Entity
	x, err := edb.Retrieve(consts.DbDir+consts.DbName, EntityBucket, key)
	if err != nil {
		return entity, errors.Wrap(err, "error while retrieving key from bucket")
	}

	err = json.Unmarshal(x, &entity)
	return entity, err
}

// RetrieveEntity retrieves an entity from the database
func RetrieveEntity(key int) (Entity, error) {
	var entity Entity
	user, err := RetrieveUser(key)
	if err != nil {
		return entity, err
	}

	entity, err = RetrieveEntityHelper(key)
	if err != nil {
		return entity, err
	}

	entity.U = &user
	return entity, entity.Save()
}

// newEntity creates a new entity based on the role passed
func newEntity(uname string, pwd string, seedpwd string, Name string, Address string, Description string, role string) (Entity, error) {
	var a Entity
	var err error
	user, err := NewUser(uname, utils.SHA3hash(pwd), seedpwd, Name)
	if err != nil {
		return a, errors.Wrap(err, "couldn't retrieve new user from db")
	}

	user.Address = Address
	user.Description = Description

	err = user.Save()
	if err != nil {
		return a, err
	}

	switch role {
	case "entity":
		a.EntityType = true
		return a, errors.New("invalid entity type passed!")
	}

	a.U = &user
	err = a.Save()
	return a, err
}

// TopReputationEntitiesWithoutRole returns the list of all the entities in descending order of reputation
func TopReputationEntitiesWithoutRole() ([]Entity, error) {
	allEntities, err := RetrieveAllEntitiesWithoutRole()
	if err != nil {
		return allEntities, errors.Wrap(err, "couldn't retrieve all entities without role")
	}
	for i := range allEntities {
		for j := range allEntities {
			if allEntities[i].U.Reputation < allEntities[j].U.Reputation {
				tmp := allEntities[i]
				allEntities[i] = allEntities[j]
				allEntities[j] = tmp
			}
		}
	}
	return allEntities, nil
}

// TopReputationEntities returns the list of all the entities belonging to a specific role in descending order of reputation
func TopReputationEntities(role string) ([]Entity, error) {
	allEntities, err := RetrieveAllEntities(role)
	if err != nil {
		return allEntities, errors.Wrap(err, "couldn't retrieve all entities")
	}
	for i := range allEntities {
		for j := range allEntities {
			if allEntities[i].U.Reputation < allEntities[j].U.Reputation {
				tmp := allEntities[i]
				allEntities[i] = allEntities[j]
				allEntities[j] = tmp
			}
		}
	}
	return allEntities, nil
}

// ValidateEntity validates the username and pwhash of the entity
func ValidateEntity(name string, token string) (Entity, error) {
	var rec Entity
	user, err := ValidateUser(name, token)
	if err != nil {
		return rec, errors.Wrap(err, "couldn't validate user")
	}
	return RetrieveEntity(user.Index)
}

// AgreeToContractConditions agrees to some specified contract conditions
func AgreeToContractConditions(contractHash string, projIndex string,
	debtAssetCode string, entityIndex int, seedpwd string) error {
	// we need to display this on the frontend and once the user presses agree, commit
	// a tx to the blockchain with the outcome

	message := "I agree to the terms and conditions specified in contract " + contractHash +
		"and by signing this message to the blockchain agree that I accept the investment in project " + projIndex +
		"whose debt asset is: " + debtAssetCode

	// hash the message and transmit the message in 5 parts due to stellar's memo field limit
	// eg.
	// CONTRACTHASH9a768ace36ff3d17
	// 71d5c145a544de3d68343b2e7609
	// 3cb7b2a8ea89ac7f1a20c852e6fc
	// 1d71275b43abffefac381c5b906f
	// 55c3bcff4225353d02f1d3498758

	user, err := RetrieveUser(entityIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve user from db")
	}

	seed, err := wallet.DecryptSeed(user.StellarWallet.EncryptedSeed, seedpwd)
	if err != nil {
		return errors.Wrap(err, "couldn't decrypt seed")
	}

	messageHash := "CONTRACTHASH" + strings.ToUpper(utils.SHA3hash(message))
	firstPart := messageHash[:28] // higher limit is not included in the slice
	secondPart := messageHash[28:56]
	thirdPart := messageHash[56:84]
	fourthPart := messageHash[84:112]
	fifthPart := messageHash[112:140]

	timestamp := float64(utils.Unix())

	_, firstHash, err := xlm.SendXLM(user.StellarWallet.PublicKey, timestamp, seed, firstPart)
	if err != nil {
		return errors.Wrap(err, "couldn't send tx 1")
	}

	_, secondHash, err := xlm.SendXLM(user.StellarWallet.PublicKey, timestamp, seed, secondPart)
	if err != nil {
		return errors.Wrap(err, "couldn't send tx 2")
	}

	_, thirdHash, err := xlm.SendXLM(user.StellarWallet.PublicKey, timestamp, seed, thirdPart)
	if err != nil {
		return errors.Wrap(err, "couldn't send tx 3")
	}

	_, fourthHash, err := xlm.SendXLM(user.StellarWallet.PublicKey, timestamp, seed, fourthPart)
	if err != nil {
		return errors.Wrap(err, "couldn't send tx 4")
	}

	_, fifthHash, err := xlm.SendXLM(user.StellarWallet.PublicKey, timestamp, seed, fifthPart)
	if err != nil {
		return errors.Wrap(err, "couldn't send tx 5")
	}

	if user.Notification {
		notif.SendContractNotification(firstHash, secondHash, thirdHash, fourthHash, fifthHash, user.Email)
	}

	return nil
}
