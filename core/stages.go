package core

import (
	"github.com/pkg/errors"
	"log"
)

// StageXtoY promtoes a contract from  stage X.Number to stage Y.Number
func StageXtoY(index int) error {
	// check for out of bound errors
	// retrieve the project
	project, err := RetrieveProject(index)
	if err != nil {
		return errors.Wrap(err, "couldn't retrieve project")
	}

	if project.Stage < 0 || project.Stage > 8 {
		log.Println("project stage number out of bounds, quitting")
		return errors.New("stage number out of bounds or not eligible for stage updation")
	}

	if project.StageChecklist == nil || project.StageData == nil {
		log.Println("stage checklist or stage data is nil, quitting")
		return errors.New("stage checklist or stage data is nil, quitting")
	}

	var baseStage Stage
	var finalStage Stage

	switch project.Stage {
	case 0:
		baseStage = Stage0
		finalStage = Stage1
	case 1:
		baseStage = Stage1
		finalStage = Stage2
	case 2:
		baseStage = Stage2
		finalStage = Stage3
	case 3:
		baseStage = Stage3
		finalStage = Stage4
	case 4:
		baseStage = Stage4
		finalStage = Stage5
	case 5:
		baseStage = Stage5
		finalStage = Stage6
	case 6:
		baseStage = Stage6
		finalStage = Stage7
	default:
		// shouldn't come here? in case it does, error out.
		return errors.New("base stage doesn't match with predefined stages, quitting")
	}

	if len(project.StageChecklist[baseStage.Number]) != len(baseStage.Activities) {
		log.Println("length of checklists don't match, quitting")
		return errors.New("length of checklists don't match, quitting")
	}

	if len(project.StageData[baseStage.Number]) == 0 {
		log.Println("baseStage data is empty, can't upgrade stages!")
		return errors.New("baseStage data is empty, can't upgrade stages")
	}

	// go through the checklist and see if something's wrong
	for _, check := range project.StageChecklist[baseStage.Number] {
		if !check {
			log.Println("checklist not satisfied, quitting")
			return errors.New("checklist not satisfied, quitting")
		}
	}

	// everything in the checklist is set to true, so we can upgrade from stage 0 to 1 safely
	log.Println("Upgrading: ", project.Index, " from stage: ", baseStage.Number, " to stage: ", finalStage.Number)
	return project.SetStage(finalStage.Number)
}

// Stage0 is a predefined transition stage
var Stage0 = Stage{
	Number:       0,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage1 is a predefined transition stage
var Stage1 = Stage{
	Number:       1,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage2 is a predefined transition stage
var Stage2 = Stage{
	Number:       2,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage3 is a predefined transition stage
var Stage3 = Stage{
	Number:       3,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage4 is a predefined transition stage
var Stage4 = Stage{
	Number:       4,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage5 is a predefined transition stage
var Stage5 = Stage{
	Number:       5,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage6 is a predefined transition stage
var Stage6 = Stage{
	Number:       6,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// Stage7 is a predefined transition stage
var Stage7 = Stage{
	Number:       7,
	FriendlyName: "",
	Name:         "",
	Activities:   []string{},
	StateTrigger: []string{},
}

// SetStage sets the stage of a project
func (a *Project) SetStage(number int) error {
	switch number {
	case 3:
		a.Reputation = a.TotalValue // upgrade reputation since totalValue might have changed from the originated contract
		err := a.Save()
		if err != nil {
			log.Println("Error while saving project", err)
			return err
		}
	case 5:
		for _, i := range a.InvestorIndices {
			elem, err := RetrieveInvestor(i)
			if err != nil {
				log.Println("Error while retrieving investor", err)
				return err
			}
			err = elem.U.ChangeReputation(a.TotalValue * InvestorWeight)
			if err != nil {
				log.Println("Couldn't change investor reputation", err)
				return err
			}
		}
	case 6:
		recp, err := RetrieveRecipient(a.RecipientIndex)
		if err != nil {
			return err
		}
		err = recp.U.ChangeReputation(a.TotalValue * RecipientWeight) // modify recipient reputation now that the system had begun power generation
		if err != nil {
			log.Println("Error while changing recipient reputation", err)
			return err
		}
	default:
		log.Println("default")
	}
	a.Stage = number
	return a.Save()
}
