package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var blockChain *BlockChain

func run() {

	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/add", blockChainAddHandler)

	_ = http.ListenAndServe("localhost:8080", nil)

}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
}
func blockChainAddHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)

	blockChainGetHandler(w, r)

}

func main() {
	blockChain = NewBlockChain()
	run()
}
