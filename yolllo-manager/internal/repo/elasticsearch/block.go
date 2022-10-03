package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"yolllo-manager/models"
)

func (r *Repository) GetLastBlocks(pageSize int64) (esResp models.ElasticAPIQueryGetLastBlocksResp, err error) {
	ctx := context.Background()
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	query := strings.NewReader(
		`{
			"size": ` + pageSizeStr + `,
			"query": {
				"match_all": {}
			},
			"sort": [
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("blocks-000001"),
		r.Conn.Search.WithBody(query),
		r.Conn.Search.WithTrackTotalHits(true),
	)
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return
	}
	err = json.Unmarshal(body, &esResp)
	if err != nil {

		return
	}

	return
}

func (r *Repository) GetNextBlocks(pageSize int64, nextPageOffset float64) (esResp models.ElasticAPIQueryGetLastBlocksResp, err error) {
	ctx := context.Background()
	sizeStr := strconv.FormatInt(pageSize, 10)
	offsetStr := fmt.Sprintf("%f", nextPageOffset)
	query := strings.NewReader(
		`{
			"size": ` + sizeStr + `,
			"query": {
				"match_all": {}
			},
			"search_after": [` + offsetStr + `],
			"sort": [
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("blocks-000001"),
		r.Conn.Search.WithBody(query),
		r.Conn.Search.WithTrackTotalHits(true),
	)
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return
	}
	err = json.Unmarshal(body, &esResp)
	if err != nil {

		return
	}

	return
}
