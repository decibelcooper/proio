package proio

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/decibelcooper/proio/go-proio/model/eic"
	"go-hep.org/x/hep/fwk"
	"go-hep.org/x/hep/fwk/job"
)

func TestInputStream1(t *testing.T) {
	buffer := &bytes.Buffer{}

	wrt := NewWriter(buffer)
	event := NewEvent()
	event.AddEntry("test", &eic.SimHit{})
	wrt.Push(event)
	wrt.Push(event)
	wrt.Push(event)
	wrt.Push(event)
	wrt.Close()

	app := job.New(job.P{
		"EvtMax":   int64(4),
		"NProcs":   2,
		"MsgLevel": job.MsgLevel("ERROR"),
	})

	app.Create(job.C{
		Type: "github.com/decibelcooper/proio/go-proio.PrintEvents",
		Name: "eventprinter",
		Props: job.P{
			"Input":  "toprint",
			"Output": "printed",
		},
	})

	app.Create(job.C{
		Type: "go-hep.org/x/hep/fwk.InputStream",
		Name: "input",
		Props: job.P{
			"Ports": []fwk.Port{
				{
					Name: "toprint",
					Type: reflect.TypeOf(&Event{}),
				},
			},
			"Streamer": &InputStream{R: buffer},
		},
	})

	app.Run()
}

type PrintEvents struct {
	fwk.TaskBase

	input  string
	output string
}

func (tsk *PrintEvents) Configure(ctx fwk.Context) error {
	if err := tsk.DeclInPort(tsk.input, reflect.TypeOf(&Event{})); err != nil {
		return err
	}

	if err := tsk.DeclOutPort(tsk.output, reflect.TypeOf("")); err != nil {
		return err
	}

	return nil
}

func (tsk *PrintEvents) StartTask(ctx fwk.Context) error {
	return nil
}

func (tsk *PrintEvents) Process(ctx fwk.Context) error {
	store := ctx.Store()

	v, err := store.Get(tsk.input)
	if err != nil {
		return err
	}
	event := v.(*Event)

	evtString := event.String()
	err = store.Put(tsk.output, evtString)
	if err != nil {
		return err
	}

	return nil
}

func (tsk *PrintEvents) StopTask(ctx fwk.Context) error {
	return nil
}

func init() {
	fwk.Register(
		reflect.TypeOf(PrintEvents{}),
		func(typ, name string, mgr fwk.App) (fwk.Component, error) {
			tsk := &PrintEvents{
				TaskBase: fwk.NewTask(typ, name, mgr),
				input:    "printevents",
				output:   "printedevents",
			}

			if err := tsk.DeclProp("Input", &tsk.input); err != nil {
				return nil, err
			}

			if err := tsk.DeclProp("Output", &tsk.output); err != nil {
				return nil, err
			}

			return tsk, nil
		},
	)
}
