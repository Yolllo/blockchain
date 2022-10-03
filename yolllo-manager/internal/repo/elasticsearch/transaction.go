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

func (r *Repository) GetLastTransactions(pageSize int64) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	ctx := context.Background()
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	query := strings.NewReader(
		`{
			"size": ` + pageSizeStr + `,
			"query": {
				"match_all": {}
			},
			"sort": [
				{"timestamp": "desc"},
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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

func (r *Repository) GetNextTransactions(pageSize int64, timestampAfter float64, searchOrderAfter float64) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	ctx := context.Background()
	sizeStr := strconv.FormatInt(pageSize, 10)
	timestampAfterStr := fmt.Sprintf("%f", timestampAfter)
	searchOrderAfterStr := fmt.Sprintf("%f", searchOrderAfter)
	query := strings.NewReader(
		`{
			"size": ` + sizeStr + `,
			"query": {
				"match_all": {}
			},
			"search_after": [` + timestampAfterStr + `, ` + searchOrderAfterStr + `],
			"sort": [
				{"timestamp": "desc"},
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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

func (r *Repository) GetLastTransactionsByAddr(pageSize int64, walletAddr string) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	ctx := context.Background()
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	query := strings.NewReader(
		`{
			"size": ` + pageSizeStr + `,
			"query": {
				"multi_match": {
					"query" : "` + walletAddr + `"
					, "fields": ["receiver","sender"]
				}
			},
			"sort": [
				{"timestamp": "desc"},
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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

func (r *Repository) GetNextTransactionsByAddr(pageSize int64, walletAddr string, timestampAfter float64, searchOrderAfter float64) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	ctx := context.Background()
	sizeStr := strconv.FormatInt(pageSize, 10)
	timestampAfterStr := fmt.Sprintf("%f", timestampAfter)
	searchOrderAfterStr := fmt.Sprintf("%f", searchOrderAfter)
	query := strings.NewReader(
		`{
			"size": ` + sizeStr + `,
			"query": {
				"multi_match": {
					"query" : "` + walletAddr + `"
					, "fields": ["receiver","sender"]
				}
			},
			"search_after": [` + timestampAfterStr + `, ` + searchOrderAfterStr + `],
			"sort": [
				{"timestamp": "desc"},
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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

func (r *Repository) GetTransactionByHash(trxHash string) (esResp models.ElasticAPIQueryGetTransactionResp, err error) {
	ctx := context.Background()
	query := strings.NewReader(
		`{
			"query": {
			  "match": {
				"_id": {
				  "query": "` + trxHash + `"
				}
			  }
			}
		}`)
	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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

func (r *Repository) GetRangeTransactions(pageSize int64, pageFrom int64, timestampFrom int64, timestampTo int64) (esResp models.ElasticAPIQueryGetLastTransactionsResp, err error) {
	timestampFrom = timestampFrom * 1000
	timestampTo = timestampTo * 1000
	pageSizeStr := strconv.FormatInt(pageSize, 10)
	pageFromStr := strconv.FormatInt(pageFrom, 10)
	timestampFromStr := strconv.FormatInt(timestampFrom, 10)
	timestampToStr := strconv.FormatInt(timestampTo, 10)
	ctx := context.Background()
	query := strings.NewReader(
		`{
			"query": {
				"range": {
					"timestamp": {
						"gte": ` + timestampFromStr + `,
						"lte": ` + timestampToStr + `
					}
				}
			},
			"size": ` + pageSizeStr + `,
			"from": ` + pageFromStr + `,
			"sort": [
				{"timestamp": "desc"},
				{"searchOrder": "desc"}
			]
		}`)

	res, err := r.Conn.Search(
		r.Conn.Search.WithContext(ctx),
		r.Conn.Search.WithIndex("transactions-000001"),
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
