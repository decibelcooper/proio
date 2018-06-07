package proio_test

import (
	"bytes"
	"fmt"

	"github.com/decibelcooper/proio/go-proio"
	"github.com/decibelcooper/proio/go-proio/model/eic"
)

func Example_scan() {
	buffer := &bytes.Buffer{}
	writer := proio.NewWriter(buffer)

	for i := 0; i < 5; i++ {
		event := proio.NewEvent()
		p := &eic.Particle{
			Pdg:    443,
			Charge: float32(i + 1),
		}
		event.AddEntry("Particle", p)
		writer.Push(event)
	}
	writer.Flush()

	reader := proio.NewReader(buffer)

	for event := range reader.ScanEvents() {
		fmt.Print(event)
	}

	// Output:
	// ---------- TAG: Particle ----------
	// ID: 1
	// Entry type: proio.model.eic.Particle
	// pdg: 443
	// charge: 1
	//
	// ---------- TAG: Particle ----------
	// ID: 1
	// Entry type: proio.model.eic.Particle
	// pdg: 443
	// charge: 2
	//
	// ---------- TAG: Particle ----------
	// ID: 1
	// Entry type: proio.model.eic.Particle
	// pdg: 443
	// charge: 3
	//
	// ---------- TAG: Particle ----------
	// ID: 1
	// Entry type: proio.model.eic.Particle
	// pdg: 443
	// charge: 4
	//
	// ---------- TAG: Particle ----------
	// ID: 1
	// Entry type: proio.model.eic.Particle
	// pdg: 443
	// charge: 5
}
