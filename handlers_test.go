package mongoshake

import "testing"

var BaseURL = "http://192.168.124.65:9100"

func TestClient_GetConfig(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetConfig()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetExecutor(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetExecutor()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetRepl(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetRepl()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetQueue(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetQueue()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetWorker(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetWorker()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetSentinel(t *testing.T) {
	client := NewClient(BaseURL)
	_, err := client.GetSentinel()
	if err != nil {
		t.Error(err)
	}
}
