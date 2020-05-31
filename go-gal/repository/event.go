package repository

import (
	"github.com/erybz/go-gal-analytics/go-gal/internal/counter"
	"github.com/erybz/go-gal-analytics/go-gal/model"
)

type Stats string

const (
	LocationCountry Stats = "country"
	LocationCity          = "city"
	DeviceType            = "deviceType"
	DevicePlatform        = "devicePlatform"
	DeviceOS              = "os"
	DeviceBrowser         = "browser"
)

type EventRepository struct {
	locationCountry *counter.Counter
	locationCity    *counter.Counter
	deviceType      *counter.Counter
	devicePlatform  *counter.Counter
	deviceOS        *counter.Counter
	deviceBrowser   *counter.Counter
}

func NewEventRepository() *EventRepository {
	return &EventRepository{
		locationCountry: counter.NewCounter(),
		locationCity:    counter.NewCounter(),
		deviceType:      counter.NewCounter(),
		devicePlatform:  counter.NewCounter(),
		deviceOS:        counter.NewCounter(),
		deviceBrowser:   counter.NewCounter(),
	}
}

func (tr *EventRepository) AddEvent(ev *model.Event) {
	tr.locationCountry.Incr(ev.Location.Country)
	tr.locationCity.Incr(ev.Location.City)
	tr.deviceType.Incr(ev.Device.Type)
	tr.devicePlatform.Incr(ev.Device.Platform)
	tr.deviceOS.Incr(ev.Device.OS)
	tr.deviceBrowser.Incr(ev.Device.Browser)
}

func (tr *EventRepository) Events(d Stats) map[string]uint64 {
	m := make(map[string]uint64)
	switch d {
	case LocationCountry:
		m = tr.locationCountry.Items()
	case LocationCity:
		m = tr.locationCity.Items()
	case DeviceType:
		m = tr.deviceType.Items()
	case DevicePlatform:
		m = tr.devicePlatform.Items()
	case DeviceOS:
		m = tr.deviceOS.Items()
	case DeviceBrowser:
		m = tr.deviceBrowser.Items()
	}
	return m
}
