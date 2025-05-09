package sse

import (
    "fmt"
    "net/http"
    "sync"
    "github.com/uchimann/air_pollution_project/notifier/internal/model"
    "encoding/json"
)

type EventServer struct {
    clients map[chan string]bool
    mu      sync.Mutex
}

func NewEventServer() *EventServer {
    return &EventServer{clients: make(map[chan string]bool)}
}

func (es *EventServer) AddClient() chan string {
    ch := make(chan string, 10)
    es.mu.Lock()
    es.clients[ch] = true
    es.mu.Unlock()
    return ch
}

func (es *EventServer) RemoveClient(ch chan string) {
    es.mu.Lock()
    delete(es.clients, ch)
    close(ch)
    es.mu.Unlock()
}

func (es *EventServer) Broadcast(analysis model.PollutionAnalysis) {
    data, _ := json.Marshal(analysis)
    msg := fmt.Sprintf("data: %s\n\n", data)
    es.mu.Lock()
    for ch := range es.clients {
        select {
        case ch <- msg:
        default:
        }
    }
    es.mu.Unlock()
}

func (es *EventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    ch := es.AddClient()
    defer es.RemoveClient(ch)

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
        return
    }

    for {
        select {
        case msg := <-ch:
            fmt.Fprint(w, msg)
            flusher.Flush()
        case <-r.Context().Done():
            return
        }
    }
}