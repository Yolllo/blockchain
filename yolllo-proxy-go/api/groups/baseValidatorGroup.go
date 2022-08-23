package groups

import (
	"net/http"

	"github.com/ElrondNetwork/elrond-proxy-go/api/shared"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
	"github.com/gin-gonic/gin"
)

type validatorGroup struct {
	facade ValidatorFacadeHandler
	*baseGroup
}

// NewValidatorGroup returns a new instance of validatorGroup
func NewValidatorGroup(facadeHandler data.FacadeHandler) (*validatorGroup, error) {
	facade, ok := facadeHandler.(ValidatorFacadeHandler)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}

	vg := &validatorGroup{
		facade:    facade,
		baseGroup: &baseGroup{},
	}

	baseRoutesHandlers := []*data.EndpointHandlerData{
		{Path: "/statistics", Handler: vg.statistics, Method: http.MethodGet},
	}
	vg.baseGroup.endpoints = baseRoutesHandlers

	return vg, nil
}

// statistics returns the validator statistics
func (group *validatorGroup) statistics(c *gin.Context) {
	validatorStatistics, err := group.facade.ValidatorStatistics()
	if err != nil {
		shared.RespondWith(c, http.StatusBadRequest, nil, err.Error(), data.ReturnCodeRequestError)
		return
	}

	shared.RespondWith(c, http.StatusOK, gin.H{"statistics": validatorStatistics}, "", data.ReturnCodeSuccess)
}
