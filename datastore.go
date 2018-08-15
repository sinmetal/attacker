package main

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

// LightEntity is Single Property Indexを貼ってないEntity
type LightEntity struct {
	I          int       `datastore:",noindex"`
	J          int       `datastore:",noindex"`
	K          float64   `datastore:",noindex"`
	Order1     int64     `datastore:",noindex"`
	Order2     int64     `datastore:",noindex"`
	Order3     int64     `datastore:",noindex"`
	Order4     int64     `datastore:",noindex"`
	Order5     int64     `datastore:",noindex"`
	CreatedAt1 time.Time `datastore:",noindex"`
	CreatedAt2 time.Time `datastore:",noindex"`
	CreatedAt3 time.Time `datastore:",noindex"`
	CreatedAt4 time.Time `datastore:",noindex"`
	CreatedAt5 time.Time `datastore:",noindex"`
}

// HeavyEntity is 単調増加するSingle Property Indexをがっつり貼ってるEntity
type HeavyEntity struct {
	I          int
	J          int
	K          float64
	Order1     int64
	Order2     int64
	Order3     int64
	Order4     int64
	Order5     int64
	CreatedAt1 time.Time
	CreatedAt2 time.Time
	CreatedAt3 time.Time
	CreatedAt4 time.Time
	CreatedAt5 time.Time
}

// Datastore is Datastore Client struct
type Datastore struct {
	DS *datastore.Client
}

// NewDatastore is Datastore Client 作成
func NewDatastore(projectID string) (*Datastore, error) {
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Datastore{
		DS: dsClient,
	}, nil
}

// PutHeavyEntity is HeavyEntity put to Datastore
func (ds *Datastore) PutHeavyEntity(ctx context.Context, i int, j int, k float64) error {
	ctx, span := trace.StartSpan(ctx, "/PutHeavyEntity")
	defer span.End()
	id := uuid.New().String()

	now := time.Now()
	e := &HeavyEntity{
		I:          i,
		J:          j,
		K:          k,
		Order1:     now.Unix(),
		Order2:     now.Unix(),
		Order3:     now.Unix(),
		Order4:     now.Unix(),
		Order5:     now.Unix(),
		CreatedAt1: now,
		CreatedAt2: now,
		CreatedAt3: now,
		CreatedAt4: now,
		CreatedAt5: now,
	}

	_, err := ds.DS.Put(ctx, datastore.NameKey("Heavy3", id, nil), e)
	return err
}

// PutLightEntity is LightEntity put to Datastore
func (ds *Datastore) PutLightEntity(ctx context.Context, i int, j int, k float64) error {
	ctx, span := trace.StartSpan(ctx, "/PutLightEntity")
	defer span.End()
	id := uuid.New().String()

	now := time.Now()
	e := &LightEntity{
		I:          i,
		J:          j,
		K:          k,
		Order1:     now.Unix(),
		Order2:     now.Unix(),
		Order3:     now.Unix(),
		Order4:     now.Unix(),
		Order5:     now.Unix(),
		CreatedAt1: now,
		CreatedAt2: now,
		CreatedAt3: now,
		CreatedAt4: now,
		CreatedAt5: now,
	}

	_, err := ds.DS.Put(ctx, datastore.NameKey("Light3", id, nil), e)
	return err
}
