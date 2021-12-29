package drawflow

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Setup() *mux.Router {
	seed()

	r := mux.NewRouter()
	r.HandleFunc("/drawflow", GetDrawflowsController).Methods("GET", "OPTIONS")
	r.HandleFunc("/drawflow", GetDrawflowsController).Methods("GET", "OPTIONS").Queries("name", "{[.]]*?}")
	r.HandleFunc("/drawflow/{id}", GetDrawflowController).Methods("GET", "OPTIONS")
	r.HandleFunc("/drawflow/save/update", SaveDrawflowsController).Methods("POST", "OPTIONS").Headers("Content-Type", "application/json")
	r.HandleFunc("/drawflow/save/update", SaveDrawflowsController).Methods("OPTIONS")
	r.HandleFunc("/drawflow/save/code", SavePythonCodeController).Methods("POST", "OPTIONS").Headers("Content-Type", "application/json")
	r.HandleFunc("/drawflow/save/code", SavePythonCodeController).Methods("OPTIONS")
	r.Use(mux.CORSMethodMiddleware(r))

	return r
}

func seed() {
	dr := DrawflowData{
		Name:  "Example sumatoria",
		Data:  `{"drawflow":{"Home":{"data":{"1":{"id":1,"name":"NodeNumber","data":{"type":"input","value":"5","color":"success","pythonStr":"nodo_1=5"},"class":"node","html":"NodeNumber","typenode":"vue","inputs":{},"outputs":{"output_1":{"connections":[{"node":"5","output":"input_1"}]}},"pos_x":148,"pos_y":152},"2":{"id":2,"name":"NodeNumber","data":{"type":"input","value":"5","color":"success","pythonStr":"nodo_2=5"},"class":"node","html":"NodeNumber","typenode":"vue","inputs":{},"outputs":{"output_1":{"connections":[{"node":"5","output":"input_1"}]}},"pos_x":386,"pos_y":153},"3":{"id":3,"name":"NodeIfElse","data":{"type":"operation","valueIf":"'sumatoria correcta!'","valueElse":"'sumatoria incorrecta!'","value":"'sumatoria correcta!'","color":"dark","nodeFrom":"5","valueCompare":"10","pythonStr":"if nodo_5 == 10: nodo_3='sumatoria correcta!';\nelse: nodo_3='sumatoria incorrecta!';"},"class":"node-code","html":"NodeIfElse","typenode":"vue","inputs":{"input_1":{"connections":[{"node":"5","input":"output_1"}]}},"outputs":{"output_1":{"connections":[{"node":"4","output":"input_1"}]}},"pos_x":352,"pos_y":304},"4":{"id":4,"name":"NodeCodeBlock","data":{"type":"bloque","value":"print(nodo_3)","color":"dark","pythonStr":"print(nodo_3)","pythoStr":"print(nodo_3)"},"class":"node-code","html":"NodeCodeBlock","typenode":"vue","inputs":{"input_1":{"connections":[{"node":"3","input":"output_1"}]}},"outputs":{"output_1":{"connections":[]}},"pos_x":807,"pos_y":324},"5":{"id":5,"name":"NodeAdd","data":{"type":"operation","value":10,"color":"warning","operation":"add","pythonStr":"nodo_5=nodo_1 + nodo_2"},"class":"node","html":"NodeAdd","typenode":"vue","inputs":{"input_1":{"connections":[{"node":"1","input":"output_1"},{"node":"2","input":"output_1"}]}},"outputs":{"output_1":{"connections":[{"node":"3","output":"input_1"}]}},"pos_x":150,"pos_y":300}}}}}`,
		DType: []string{"DrawflowData"}}
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
		log.Fatal(err)
	}

	mu.SetJson = pb

	response, err := client.NewTxn().Mutate(ctx, mu)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", response.Json)

}

type CancelFunc func()

func newClient() (*dgo.Dgraph, CancelFunc) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	conn, err := grpc.Dial("dgraph:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
			api.NewDgraphClient(conn),
		), func() {
			if err := conn.Close(); err != nil {
				log.Printf("Error while closing connection:%v", err)
			}
		}
}

func newSchema() string {
	return (`name: string @index(fulltext,term) .
			data: string .
			type DrawflowData { 
				name
				data 
			}

			`)
}
