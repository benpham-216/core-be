package main

import (
 "encoding/json"
 "log"
 "net/http"
 "os"
 "strings"
 "github.com/benpham-216/core-be/internal/application"
 "github.com/benpham-216/core-be/internal/domain"
 "github.com/benpham-216/core-be/internal/eventstore"
)
func main(){s:=application.NewService(eventstore.NewMemoryStore());mux:=http.NewServeMux();mux.HandleFunc("/healthz",func(w http.ResponseWriter,r *http.Request){w.WriteHeader(http.StatusNoContent)});mux.HandleFunc("/v1/work-orders",func(w http.ResponseWriter,r *http.Request){if r.Method!="POST"{w.WriteHeader(405);return};var in struct{ID,Title,Description,CorrelationID string;Actor domain.Actor};if json.NewDecoder(r.Body).Decode(&in)!=nil{http.Error(w,"invalid JSON",400);return};wo,e:=s.Create(r.Context(),in.ID,in.Title,in.Description,in.CorrelationID,in.Actor);if e!=nil{http.Error(w,e.Error(),400);return};w.WriteHeader(201);json.NewEncoder(w).Encode(wo)});mux.HandleFunc("/v1/work-orders/",func(w http.ResponseWriter,r *http.Request){parts:=strings.Split(strings.TrimPrefix(r.URL.Path,"/v1/work-orders/"),"/");id:=parts[0];if len(parts)==1&&r.Method=="GET"{wo,e:=s.Get(r.Context(),id);if e!=nil{http.Error(w,e.Error(),404);return};json.NewEncoder(w).Encode(wo);return};if len(parts)==2&&parts[1]=="events"&&r.Method=="GET"{e,_:=s.Events(r.Context(),id);json.NewEncoder(w).Encode(e);return};if len(parts)==3&&parts[1]=="commands"&&parts[2]=="transition"&&r.Method=="POST"{var in struct{ExpectedVersion int `json:"expectedVersion"`;Next domain.Status `json:"next"`;CorrelationID string `json:"correlationId"`;Actor domain.Actor `json:"actor"`};if json.NewDecoder(r.Body).Decode(&in)!=nil{http.Error(w,"invalid JSON",400);return};wo,e:=s.Transition(r.Context(),id,in.ExpectedVersion,in.Next,in.CorrelationID,in.Actor);if e!=nil{http.Error(w,e.Error(),409);return};json.NewEncoder(w).Encode(wo);return};w.WriteHeader(404)});port:=os.Getenv("PORT");if port==""{port="8080"};log.Fatal(http.ListenAndServe(":"+port,mux))}
