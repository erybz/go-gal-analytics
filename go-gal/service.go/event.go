package service

import (
	"log"
	"net"
	"net/http"

	"github.com/avct/uasurfer"
	"github.com/erybz/go-gal-analytics/go-gal/model"
	"github.com/erybz/go-gal-analytics/go-gal/repository"
	"github.com/oschwald/geoip2-golang"
)

type EventService struct {
	eventRepo   *repository.EventRepository
	geoIPReader *geoip2.Reader
}

func NewEventService() *EventService {
	return &EventService{
		eventRepo:   repository.NewEventRepository(),
		geoIPReader: initGeoIPReader("go-gal/assets/GeoLite2-City.mmdb"),
	}
}

func (ts *EventService) BuildEvent(r *http.Request) (*model.Event, error) {
	// clientIP := net.ParseIP(realip.FromRequest(r))
	clientIP := net.ParseIP("111.119.248.36")
	userAgent := uasurfer.Parse(r.UserAgent())

	geoData, err := ts.geoIPReader.City(clientIP)
	if err != nil {
		return nil, err
	}

	if userAgent.IsBot() {
		return nil, nil
	}

	event := &model.Event{
		Location: model.EventLocation{
			Country: geoData.Country.Names["en"],
			City:    geoData.City.Names["en"],
		},
		Device: model.EventDevice{
			Type:     userAgent.DeviceType.StringTrimPrefix(),
			Platform: userAgent.OS.Platform.StringTrimPrefix(),
			OS:       userAgent.OS.Name.StringTrimPrefix(),
			Browser:  userAgent.Browser.Name.StringTrimPrefix(),
		},
	}
	return event, nil
}

func (ts *EventService) LogEvent(event *model.Event) {
	ts.eventRepo.AddEvent(event)
}

func (ts *EventService) Stats(dim repository.Stats) []map[string]interface{} {
	allStats := make([]map[string]interface{}, 0, 1)
	for k, v := range ts.eventRepo.Events(dim) {
		stat := map[string]interface{}{
			string(dim): k,
			"pageViews": v,
		}
		allStats = append(allStats, stat)
	}
	return allStats
}

func initGeoIPReader(path string) *geoip2.Reader {
	db, err := geoip2.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
