package groups_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ElrondNetwork/elrond-proxy-go/api/groups"
	"github.com/ElrondNetwork/elrond-proxy-go/api/mock"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const nodePath = "/node"

func TestNewNodeGroup_WrongFacadeShouldErr(t *testing.T) {
	wrongFacade := &mock.WrongFacade{}
	group, err := groups.NewNodeGroup(wrongFacade)
	require.Nil(t, group)
	require.Equal(t, groups.ErrWrongTypeAssertion, err)
}

func TestHeartbeat_GetHeartbeatDataReturnsStatusOk(t *testing.T) {
	t.Parallel()

	facade := &mock.Facade{
		GetHeartbeatDataHandler: func() (*data.HeartbeatResponse, error) {
			return &data.HeartbeatResponse{Heartbeats: []data.PubKeyHeartbeat{}}, nil
		},
	}

	nodeGroup, err := groups.NewNodeGroup(facade)
	require.NoError(t, err)
	ws := startProxyServer(nodeGroup, nodePath)

	req, _ := http.NewRequest("GET", "/node/heartbeatstatus", nil)
	resp := httptest.NewRecorder()
	ws.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHeartbeat_GetHeartbeatDataReturnsOkResults(t *testing.T) {
	t.Parallel()

	name1, identity1 := "name1", "identity1"
	name2, identity2 := "name2", "identity2"

	facade := &mock.Facade{
		GetHeartbeatDataHandler: func() (*data.HeartbeatResponse, error) {
			return &data.HeartbeatResponse{
				Heartbeats: []data.PubKeyHeartbeat{
					{
						NodeDisplayName: name1,
						Identity:        identity1,
					},
					{
						NodeDisplayName: name2,
						Identity:        identity2,
					},
				},
			}, nil
		},
	}
	nodeGroup, err := groups.NewNodeGroup(facade)
	require.NoError(t, err)
	ws := startProxyServer(nodeGroup, nodePath)

	req, _ := http.NewRequest("GET", "/node/heartbeatstatus", nil)
	resp := httptest.NewRecorder()
	ws.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var result data.HeartbeatApiResponse
	loadResponse(resp.Body, &result)
	assert.Equal(t, name1, result.Data.Heartbeats[0].NodeDisplayName)
	assert.Equal(t, name2, result.Data.Heartbeats[1].NodeDisplayName)
	assert.Equal(t, identity1, result.Data.Heartbeats[0].Identity)
	assert.Equal(t, identity2, result.Data.Heartbeats[1].Identity)
}

func TestHeartbeat_GetHeartbeatBadRequestShouldErr(t *testing.T) {
	t.Parallel()

	facade := &mock.Facade{
		GetHeartbeatDataHandler: func() (*data.HeartbeatResponse, error) {
			return nil, errors.New("bad request")
		},
	}
	nodeGroup, err := groups.NewNodeGroup(facade)
	require.NoError(t, err)
	ws := startProxyServer(nodeGroup, nodePath)

	req, _ := http.NewRequest("GET", "/node/heartbeatstatus", nil)
	resp := httptest.NewRecorder()
	ws.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
