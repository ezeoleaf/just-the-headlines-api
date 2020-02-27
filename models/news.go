package models

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          string   `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type News struct {
	Title        string `json:"title"`
	Descripition string `json:"description"`
	Link         string `json:"link"`
}

func GetFilteredNews(db *sql.DB, id int, filter string) []News {
	return getNews(db, id, filter)
}

func GetNews(db *sql.DB, id int) []News {
	return getNews(db, id, ``)
}

func getNews(db *sql.DB, id int, filter string) []News {
	sql := "SELECT rss FROM section WHERE id=$1"

	row := db.QueryRow(sql, id)

	var uri string

	err := row.Scan(&uri)

	if err != nil {
		panic(err)
	}

	results := []News{}

	document, err := retrieveNews(uri)
	if err != nil {
		return nil
	}

	for _, channelItem := range document.Channel.Item {
		if !match(channelItem, strings.ToLower(filter)) {
			n := News{
				Title:        channelItem.Title,
				Descripition: channelItem.Description,
				Link:         channelItem.Link,
			}

			results = append(results, n)
		}
	}

	return results
}

func match(i item, filter string) bool {
	if filter == `` {
		return false
	}

	var matched bool
	var err error
	matched, err = regexp.MatchString(filter, strings.ToLower(i.Title))

	if err != nil {
		panic(err)
	}

	if !matched {
		matched, err = regexp.MatchString(filter, strings.ToLower(i.Description))

		if err != nil {
			panic(err)
		}
	}

	return matched
}

func retrieveNews(uri string) (*rssDocument, error) {
	if uri == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response error %d", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err
}
