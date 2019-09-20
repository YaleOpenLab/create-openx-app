package core

import (
	"github.com/pkg/errors"
)

// Slash slashes the contractor's reputation in the event of bad behaviour.
func (contractor *Entity) Slash(contractValue float64) error {
	// slash an entity's reputation score if it reneges on an agreed contract
	contractor.U.Reputation -= contractValue * 0.1
	return contractor.Save()
}

// RepInstalledProject adds reputatuon to the contractor on completion of installation of a project. By default,
// we add reputation to the entity. In case the recipient wants to dispute this claim, we review and
// change the reputation accordingly
func RepInstalledProject(contrIndex int, projIndex int) error {
	contractor, err := RetrieveEntity(contrIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve all entities from db")
	}

	project, err := RetrieveProject(projIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve project from db")
	}

	err = project.SetStage(5)
	if err != nil {
		return errors.Wrap(err, "couldn't set installed project's stage")
	}

	contractor.U.Reputation += project.TotalValue * ContractorWeight
	return contractor.Save()
}
