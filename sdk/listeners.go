package sdk

import "github.com/launchboxio/cloudscale/internal/api"

type Listeners struct {
	*Client
}

type ListenersList struct {
	Listeners []api.Listener `json:"listeners"`
}

type ListenerResponse struct {
	Listener api.Listener `json:"listener"`
}

type CreateListenerInput struct {
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Port      uint16 `json:"port"`
	Type      string `json:"type,omitempty"`
	Protocol  string `json:"protocol,omitempty"`
}

func (l *Listeners) List() (ListenersList, error) {
	var listenerList ListenersList
	_, err := l.http.R().
		SetResult(&listenerList).
		Get("/listeners")
	return listenerList, err
}

func (l *Listeners) Get(listenerId string) (ListenerResponse, error) {
	var result ListenerResponse
	_, err := l.http.R().
		SetResult(&result).
		Get("/listeners/" + listenerId)
	return result, err
}

func (l *Listeners) Delete(listenerId string) error {
	_, err := l.http.R().Delete("/listeners/" + listenerId)
	return err
}

func (l *Listeners) Create(input *CreateListenerInput) (ListenerResponse, error) {
	var result ListenerResponse
	_, err := l.http.R().
		SetResult(&result).
		SetBody(input).
		Post("/listeners")
	return result, err
}
