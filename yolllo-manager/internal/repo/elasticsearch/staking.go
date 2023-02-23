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

func (r *Repository) GetLastClaimedRewardByAddr(addr string, timestampFrom int64) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	timestampFrom = timestampFrom * 1000
	timestampFromStr := strconv.FormatInt(timestampFrom, 10)
	ctx := context.Background()
	query := strings.NewReader(
		`{
			"query": {
				"bool": {
					"must": [
						{
							"range": {
								"timestamp": {
									"gt": ` + timestampFromStr + `
								}
							}
						},
						{
							"match": {
								"receiver": "` + addr + `"
							}
						},
						{
							"match": {
								"originalSender": "` + addr + `"
							}
						},
						{
							"match": {
								"sender": "yol1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqhllllsj2afxf"
							}
						}
					]
				}
			},
			"from": 0,
			"size": 1,
			"sort": [
				{
					"timestamp": {
						"order": "asc"
					}
				}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("operations-000001"),
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
	fmt.Println(string(body))
	err = json.Unmarshal(body, &esResp)
	if err != nil {

		return
	}

	return
}
