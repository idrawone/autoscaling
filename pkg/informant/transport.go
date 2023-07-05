package informant

import (
	"github.com/neondatabase/autoscaling/pkg/api"
)

// Defines types that are used to communicate with the monitor over websocket
// connection
type Packet struct {
	Stage Stage  `json:"stage"`
	Id    uint64 `json:"id"`
}

type Stage struct {
	Request  *Request  `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
	Done     *struct{} `json:"done,omitempty"`
}

type Request struct {
	RequestUpscale *struct{}  `json:"requestUpscale,omitempty"`
	NotifyUpscale  *Resources `json:"notifyUpscale,omitempty"`
	TryDownscale   *Resources `json:"tryDownscale,omitempty"`
}

type Response struct {
	UpscaleResult        *Resources       `json:"upscaleResult,omitempty"`
	ResourceConfirmation *struct{}        `json:"resourceConfirmation,omitempty"`
	DownscaleResult      *DownscaleResult `json:"downscaleResult,omitempty"`
}

type Resources struct {
	Cpu uint64 `json:"cpu"`
	Mem uint64 `json:"mem"`
}

type DownscaleResult struct {
	Ok     bool   `json:"ok"`
	Status string `json:"status"`
}

// Convert into api.DownscaleResult.
//
// The reason for having two types is to prevent us from having to keep track/control
// how api.DownscaleResult is serialized
func (res *DownscaleResult) Into() *api.DownscaleResult {
	return &api.DownscaleResult{
		Ok:     res.Ok,
		Status: res.Status,
	}
}

func Done(id uint64) Packet {
	return Packet{
		Stage: Stage{
			Request:  nil,
			Response: nil,
			Done:     &struct{}{},
		},
		Id: id,
	}
}