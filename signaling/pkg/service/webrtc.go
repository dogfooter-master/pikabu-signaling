package service

import (
	"bytes"
	"fmt"
	"github.com/pion/webrtc"
	"net/http"
	"pikabu-signaling/signaling/pkg/signal"
	"time"
)

type WebRTC struct {
	PeerConnection  *webrtc.PeerConnection
	DataChannel *webrtc.DataChannel
	Offer string
}

var WebRTCConfig webrtc.Configuration
var WebRTCMap map[string]WebRTC

func init() {
	WebRTCConfig = webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
}

func (o *WebRTC) CreateDataChannel() (err error){
	o.PeerConnection, err = webrtc.NewPeerConnection(WebRTCConfig)
	if err != nil {
		panic(err)
	}

	o.PeerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		payload := []byte(c.ToJSON().Candidate)
		resp, onICECandidateErr := http.Post(fmt.Sprintf("http://%s/candidate", *answerAddr), "application/json; charset=utf-8", bytes.NewReader(payload))
		if onICECandidateErr != nil {
			panic(onICECandidateErr)
		} else if closeErr := resp.Body.Close(); closeErr != nil {
			panic(closeErr)
		}
	})

	o.DataChannel, err = o.PeerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}
	o.PeerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	o.DataChannel.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", o.DataChannel.Label(), o.DataChannel.ID())

		for range time.NewTicker(5 * time.Second).C {
			message := signal.RandSeq(15)
			fmt.Printf("Sending '%s'\n", message)

			// Send the message as text
			sendErr := o.DataChannel.SendText(message)
			if sendErr != nil {
				panic(sendErr)
			}
		}
	})
	o.DataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", o.DataChannel.Label(), string(msg.Data))
	})
	return
}

func (o *WebRTC) CreateOffer() (err error) {
	// Create an offer to send to the browser
	offer, err := o.PeerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	err = o.PeerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	o.Offer = signal.Encode(offer)

	return
}

func (o *WebRTC) ReceiveAnswer(rs string) (err error) {
	answer := webrtc.SessionDescription{}
	signal.Decode(rs, &answer)

	// Apply the answer as the remote description
	err = o.PeerConnection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}
	return
}

func (o *WebRTC) AddCandidate(rs string) (err error) {
	var candidate string
	signal.Decode(rs, &candidate)
	if err = o.PeerConnection.AddICECandidate(webrtc.ICECandidateInit{Candidate: candidate}); err != nil {
		panic(err)
	}
	return
}