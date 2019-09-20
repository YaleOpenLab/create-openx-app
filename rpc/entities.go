package rpc

import (
	"github.com/pkg/errors"
	"log"
	"net/http"

	erpc "github.com/Varunram/essentials/rpc"
	core "github.com/test/blah/core"
)

// setupEntityRPCs sets up all entity related endpoints
func setupEntityRPCs() {
	validateEntity()
	getStage0Contracts()
	getStage1Contracts()
	getStage2Contracts()
}

var EntityRpc = map[int][]string{
	1: []string{"/entity/validate"}, // GET
	2: []string{"/entity/stage0"},   // GET
	3: []string{"/entity/stage1"},   // GET
	4: []string{"/entity/stage2"},   // GET
}

// EntityValidateHelper is a helper that helps validate an entity
func EntityValidateHelper(w http.ResponseWriter, r *http.Request) (core.Entity, error) {
	var prepEntity core.Entity
	if r.Method == "GET" {
		if r.URL.Query() == nil || r.URL.Query()["username"] == nil ||
			len(r.URL.Query()["token"][0]) != 32 {
			return prepEntity, errors.New("Invalid params passed")
		}

		prepEntity, err := core.ValidateEntity(r.URL.Query()["username"][0], r.URL.Query()["token"][0])
		if err != nil {
			return prepEntity, errors.Wrap(err, "Error while validating entity")
		}

		return prepEntity, nil
	} else if r.Method == "POST" {
		if r.FormValue("username") == "" || r.FormValue("password") == "" {
			return prepEntity, errors.New("Invalid params passed")
		}

		prepEntity, err := core.ValidateEntity(r.FormValue("username"), r.FormValue("token"))
		if err != nil {
			return prepEntity, errors.Wrap(err, "Error while validating entity")
		}
		return prepEntity, nil
	}
	return prepEntity, errors.New("invalid method type")
}

// validateEntity is an endpoint that vlaidates is a specific entity is registered on the platform
func validateEntity() {
	http.HandleFunc(EntityRpc[1][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}
		erpc.MarshalSend(w, prepEntity)
	})
}

// getStage0Contracts gets a list of all the pre origianted contracts on the platform
func getStage0Contracts() {
	http.HandleFunc(EntityRpc[2][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveOriginatorProjects(core.Stage0.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving originator project", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}

// getStage1Contracts gets a list of all the originated contracts on the platform
func getStage1Contracts() {
	http.HandleFunc(EntityRpc[3][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveOriginatorProjects(core.Stage1.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving originator projects", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}

// getStage2Contracts gets a list of all the proposed contracts on the platform
func getStage2Contracts() {
	http.HandleFunc(EntityRpc[4][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveContractorProjects(core.Stage2.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving contractor projects", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}
