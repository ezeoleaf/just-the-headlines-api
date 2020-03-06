package models

// Sections Queries

// SectionsByTags const
const SectionsByTags = `SELECT s.id, s.name, s.rss, s.failed, n.name FROM section s
INNER JOIN newspaper n ON s.newspaper_id = n.id
WHERE UPPER(s.name) LIKE '%' || $1 || '%'`

// SectionsByNewspaper is used for return a list of sections by newspaper.id
const SectionsByNewspaper = `SELECT id, name, rss, failed, "" FROM section WHERE newspaper_id=$1`

// SectionsByName is used for return a list of sections filtering by section.name
const SectionsByName = `SELECT s.id, s.name, s.rss, s.failed, n.name FROM section s
INNER JOIN newspaper n ON s.newspaper_id = n.id
WHERE UPPER(s.name) LIKE '%' || $1 || '%'`

// News Queries

// NewsByID returns a RSS url for a section.id
const NewsByID = `SELECT s.rss, s.name, n.name FROM section s INNER JOIN newspaper n ON(s.newspaper_id = n.id) WHERE s.id=$1`

// Newspapers Queries

// NewspapersByName returns a list of newspapers filtering by newspaper.name
const NewspapersByName = `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE UPPER(n.name) LIKE '%' || $1 || '%'`

// NewspapersByCountry returns a list of newspapers filtering by country.code
const NewspapersByCountry = `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE UPPER(c.code) LIKE $1 || '%'`

// NewspapersAll returns a list of newspapers for all the newspapers in the database
const NewspapersAll = `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id)`

// NewspaperByID returns a single newspaper filtering by newspaper.id
const NewspaperByID = `SELECT n.id, n.name, c.name, c.code FROM newspaper n INNER JOIN country c ON(n.country_id = c.id) WHERE n.id=$1`

// Tags Queries

// TagsIDs returns the category ids for a list of tags
const TagsIDs = `SELECT id FROM category WHERE UPPER(name) in ($1)`
