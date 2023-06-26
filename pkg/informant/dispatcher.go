package informant

import (
	"context"
	"fmt"
	"time"

	"github.com/neondatabase/autoscaling/pkg/util"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type MonitorResult struct {
	Result       DownscaleResult
	Confirmation struct{}
}

// TODO: do we need synchronization?
// The Dispatcher is the main object managing the websocket connection to the
// monitor. For more information on the protocol, see TODO.
type Dispatcher struct {
	Conn *websocket.Conn
	ctx  context.Context

	// When someone sends a message, the dispatcher will attach a transaction id
	// to it so that it knows when a response is back. When it receives a packet
	// with the same transaction id, it knows that that is the repsonse to the original
	// message and will send it down the oneshot so the original sender can use it.
	waiters map[uint64]util.OneshotSender[MonitorResult]

	// A message is sent along this channel when an upscale is requested.
	// When the informant.NewState is called, a goroutine will be spawned that
	// just tries to receive off the channel and then request the upscale.
	// A different way to do this would be to keep a backpointer to the parent
	// `State` and then just call a method on it when an upscale is requested.
	notifier chan<- struct{}

	// This counter represents the current transaction id. When we need a new one
	// we simply bump it and take the new number.
	//
	// Only we care about this number. The other side will just send back packets
	// with the id of the request, but they never do anything with the number.
	counter uint64

	logger *zap.Logger
}

// Create a new Dispatcher. Note that this does not immediately start the Dispatcher.
// Call Run() to start it.
func NewDispatcher(addr string, logger *zap.Logger, notifier chan<- struct{}) (disp Dispatcher, _ error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		return disp, fmt.Errorf("Error creating dispatcher: %v", err)
	}
	disp = Dispatcher{
		Conn:     c,
		ctx:      ctx,
		notifier: notifier,
		waiters:  make(map[uint64]util.OneshotSender[MonitorResult]),
		counter:  0,
		logger:   logger.Named("dispatcher"),
	}
	return disp, nil
}

// Send a packet down the connection.
func (disp *Dispatcher) send(p Packet) error {
	return wsjson.Write(disp.ctx, disp.Conn, p)
}

// Try to receive a packet off the connection.
func (disp *Dispatcher) recv() (*Packet, error) {
	var p Packet
	err := wsjson.Read(disp.ctx, disp.Conn, &p)
	if err != nil {
		return &p, nil
	}
	return &p, nil
}

// Make a request to the monitor. The dispatcher will handle returning a response
// on the provided oneshot.
//
// *Note*: sending a RequestUpscale to the monitor is incorrect. The monitor does
// not (and should) not know how to handle this and will panic. Likewise, we panic
// upon receiving a TryDownscale or NotifyUpscale request.
func (disp *Dispatcher) Call(req Request, sender util.OneshotSender[MonitorResult]) {
	id := disp.counter
	disp.counter += 1
	disp.send(Packet{Stage: Stage{
		Request:  &req,
		Response: nil,
		Done:     nil,
	}, SeqNum: id})
	disp.waiters[id] = sender
}

// Long running function that performs all orchestrates all requests/responses.
func (disp *Dispatcher) run() {
	for {
		packet, err := disp.recv()
		if err != nil {
			disp.logger.Warn("Error receiving from ws connection. Continuing.", zap.Error(err))
			continue
		}
		stage := packet.Stage
		id := packet.SeqNum
		switch {
		case stage.Request != nil:
			{
				req := stage.Request
				switch {
				case req.RequestUpscale != nil:
					{
						// The goroutine listening on the other side will make the
						// request
						disp.notifier <- struct{}{}
					}
				case req.NotifyUpscale != nil:
					{
						panic("informant should never receive a NotifyUpscale request from monitor")
					}
				case req.TryDownscale != nil:
					{
						panic("informant should never receive a TryDownscale request from monitor")
					}
				default:
					{
						panic("all fields nil")
					}
				}
			}
		case stage.Response != nil:
			{
				res := stage.Response
				switch {
				case res.DownscaleResult != nil:
					{
						// Loop up the waiter and send back the result
						sender, ok := disp.waiters[id]
						if ok {
							sender.Send(MonitorResult{Result: *res.DownscaleResult})
							// Don't forget to delete the waiter
							delete(disp.waiters, id)
							disp.send(Done())
						} else {
							panic("Received response for id without a registered sender")
						}
					}
				case res.ResourceConfirmation != nil:
					{
						// Loop up the waiter and send back the result
						sender, ok := disp.waiters[id]
						if ok {
							sender.Send(MonitorResult{Result: *res.DownscaleResult})
							// Don't forget to delete the waiter
							delete(disp.waiters, id)
							disp.send(Done())
						} else {
							panic("Received response for id without a registered sender")
						}
					}
				case res.UpscaleResult != nil:
					{
						panic("informant should never receive an UpscaleResult response from monitor")
					}
				default:
					{
						// This is a serialization error
						panic("all fields nil")
					}
				}
			}
		case stage.Done != nil:
			{
				// yay! :)
			}
		}
	}
}