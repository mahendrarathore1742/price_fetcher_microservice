package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/mahendrarathore1742/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func makeHttpHanderFunc(apifn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {

		if err := apifn(ctx, w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func NewJsonAPiServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {

	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {

	http.HandleFunc("/", makeHttpHanderFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)

	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return WriteJSON(w, http.StatusOK, &priceResp)
}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)

	return json.NewEncoder(w).Encode(v)
}
