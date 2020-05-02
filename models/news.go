package models

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const separator = `,`

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
	Newspaper    string `json:"newspaper"`
	Section      string `json:"section"`
}

func GetFilteredNews(db *sql.DB, id int64, filter string) []News {
	return getNews(db, id, filter)
}

func GetNewsBySection(db *sql.DB, id int64) []News {
	return getNews(db, id, ``)
}

func GetMultipleNews(db *sql.DB, sections string, filter string) []News {
	results := []News{}

	sectionList := strings.Split(sections, separator)
	for _, section := range sectionList {
		sectionID, _ := strconv.ParseInt(section, 10, 64)
		news := getNews(db, sectionID, filter)
		results = append(results, news...)
	}

	return results
}

func getMultipleNews(db *sql.DB, sections []int64, filter string) []News {
	results := []News{}

	for _, section := range sections {
		news := getNews(db, section, filter)
		results = append(results, news...)
	}

	return results
}

func GetNews(db *sql.DB, userID int64) []News {
	filters := getFiltersByUser(db, userID)
	sections := getSectionsByUser(db, userID)

	results := []News{}
	for _, section := range sections.Sections {
		document, err := retrieveNews(section.RSS)
		if err != nil {
			return nil
		}

		for _, channelItem := range document.Channel.Item {
			if !match(channelItem, filters) {
				n := News{
					Title:        channelItem.Title,
					Descripition: channelItem.Description,
					Link:         channelItem.Link,
					Newspaper:    section.Newspaper,
					Section:      section.Name,
				}

				results = append(results, n)
			}
		}
	}

	return results
}

func getNews(db *sql.DB, id int64, filter string) []News {
	row := db.QueryRow(NewsByID, id)

	var uri string
	var newspaper string
	var section string

	err := row.Scan(&uri, &section, &newspaper)

	if err != nil {
		panic(err)
	}

	results := []News{}

	document, err := retrieveNews(uri)
	if err != nil {
		return nil
	}

	noFilter := filter == ``

	for _, channelItem := range document.Channel.Item {
		m := false
		if !noFilter {
			m = match(channelItem, strings.Split(filter, separator))
		}

		if !m {
			n := News{
				Title:        channelItem.Title,
				Descripition: channelItem.Description,
				Link:         channelItem.Link,
				Newspaper:    newspaper,
				Section:      section,
			}

			results = append(results, n)
		}
	}

	return results
}

func match(i item, filters []string) bool {
	for _, filter := range filters {
		filter = strings.ToLower(strings.TrimSpace(filter))
		matched, err := regexp.MatchString(filter, strings.ToLower(i.Title))

		if err != nil {
			panic(err)
		}

		if matched {
			return true
		}

		matched, err = regexp.MatchString(filter, strings.ToLower(i.Description))

		if err != nil {
			panic(err)
		}

		if matched {
			return true
		}
	}

	return false
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
