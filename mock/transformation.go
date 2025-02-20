package mock

import (
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/execute/table"
)

type Transformation struct {
	ProcessFn func(id execute.DatasetID, tbl flux.Table) error
	FinishFn  func(id execute.DatasetID, err error)
}

func (t *Transformation) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return nil
}

func (t *Transformation) Process(id execute.DatasetID, tbl flux.Table) error {
	return t.ProcessFn(id, tbl)
}

func (t *Transformation) UpdateWatermark(id execute.DatasetID, ts execute.Time) error {
	return nil
}

func (t *Transformation) UpdateProcessingTime(id execute.DatasetID, ts execute.Time) error {
	return nil
}

func (t *Transformation) Finish(id execute.DatasetID, err error) {
	t.FinishFn(id, err)
}

type NarrowTransformation struct {
	ProcessFn func(chunk table.Chunk, d *execute.TransportDataset, mem memory.Allocator) error
}

func (n *NarrowTransformation) Process(chunk table.Chunk, d *execute.TransportDataset, mem memory.Allocator) error {
	return n.ProcessFn(chunk, d, mem)
}
