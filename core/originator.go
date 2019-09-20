package core

import (
	"github.com/pkg/errors"
)

// RepOriginatedProject adds reputation to an originator on successful origination of a contract
func RepOriginatedProject(origIndex int, projIndex int) error {
	originator, err := RetrieveEntity(origIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve entity from db")
	}
	project, err := RetrieveProject(projIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve project from db")
	}
	return originator.U.ChangeReputation(project.TotalValue * OriginatorWeight)
}
