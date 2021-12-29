package drawflow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

func GetAllDrawflows() ([]DrawflowData, error) {
	dg, cancel := newClient()
	defer cancel()
	ctx := context.Background()
	var err error
	var resp *api.Response
	q := `query Drawflow(){
		results(func: eq(dgraph.type,"DrawflowData")){
			uid
			name
			}}`

	resp, err = dg.NewTxn().Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var r Query
	err = json.Unmarshal(resp.Json, &r)
	fmt.Println(string(resp.Json))

	return r.Results, err
}
func GetDrawflows(name string) ([]DrawflowData, error) {
	dg, cancel := newClient()
	defer cancel()
	ctx := context.Background()
	vars := make(map[string]string)
	vars["$name"] = name
	q := `query Drawflow($name:string){
		results(func: anyoftext(name,$name)){
			uid
			name
			}}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, vars)
	if err != nil {
		log.Default().Print(err)
	}

	var r Query
	err = json.Unmarshal(resp.Json, &r)
	fmt.Println(string(resp.Json))

	return r.Results, err
}

func GetDrawflow(idFilter string) (DrawflowData, error) {

	dg, cancel := newClient()
	defer cancel()
	ctx := context.Background()
	vars := make(map[string]string)
	vars["$id"] = idFilter
	q := `query Drawflow($id:string){
		results(func: uid($id)){
			uid
			name
			data
			}}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, vars)
	if err != nil {
		log.Print(err)
	}

	var r Query
	err = json.Unmarshal(resp.Json, &r)
	fmt.Println(string(resp.Json))

	if len(r.Results) == 0 {
		err = errors.New("no results")
	}
	return r.Results[0], err
}
func SaveDrawflow(dr DrawflowData) (DrawflowData, error) {

	dr.DType = []string{"DrawflowData"}
	client, cancel := newClient()

	defer cancel()
	op := &api.Operation{
		Schema: newSchema(),
	}
	ctx := context.Background()

	if err := client.Alter(ctx, op); err != nil {
		log.Print(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(dr)

	if err != nil {
		log.Print(err)
	}

	mu.SetJson = pb

	response, err := client.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Println(err, response)
	}

	return DrawflowData{}, err
}
