package core

import (
	"github.com/pkg/errors"
	"log"
)

// SetOneTimeUnlock sets a one time seedpwd that can be used to automatically unlock the project once an investment comes in
func (a *Recipient) SetOneTimeUnlock(projIndex int, seedpwd string) error {
	log.Println("setting one time unlock for project with index: ", projIndex)
	project, err := RetrieveProject(projIndex)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve project")
	}

	recp, err := RetrieveRecipient(project.RecipientIndex)
	if err != nil {
		return errors.Wrap(err, "did not retrieve recipient belonging to project")
	}

	if recp.U.Index != a.U.Index {
		return errors.Wrap(err, "recipient index does not match with project recipient index")
	}

	project.OneTimeUnlock = seedpwd
	return project.Save()
}
