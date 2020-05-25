package models

// Sections Queries

// SectionsByTags const
const SectionsByTags = `SELECT s.id, s.name, s.rss, s.failed, n.name FROM section s
INNER JOIN newspaper n ON s.newspaper_id = n.id
WHERE UPPER(s.name) LIKE '%' || $1 || '%'`

// SectionsByNewspaper is used for return a list of sections by newspaper.id
const SectionsByNewspaper = `SELECT s.id, s.name, s.rss, s.failed, "",
CASE WHEN subs.user_id IS NULL THEN false ELSE true END
FROM section s
LEFT JOIN subscription subs ON (subs.section_id = s.id AND subs.user_id=$2 )
WHERE s.newspaper_id=$1`

// SectionsByName is used for return a list of sections filtering by section.name
const SectionsByName = `SELECT s.id, s.name, s.rss, s.failed, n.name,
CASE WHEN subs.user_id IS NULL THEN false ELSE true END
FROM section s
INNER JOIN newspaper n ON s.newspaper_id = n.id
LEFT JOIN subscription subs ON (subs.section_id = s.id AND subs.user_id=$2 )
WHERE UPPER(s.name) LIKE '%' || $1 || '%'`

const getUserSections = `SELECT s.id, s.name, s.rss, s.failed, n.name, true FROM section s
INNER JOIN newspaper n ON s.newspaper_id = n.id
INNER JOIN subscription subs ON subs.section_id = s.id
WHERE subs.user_id=?`

const subscribeUser = `INSERT INTO subscription(section_id, user_id) VALUES(?, ?)`
const unsubscribeUser = `DELETE FROM subscription WHERE section_id=? AND user_id=?`

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

// Users Queries
const newUser = `INSERT INTO user(email, username, password) VALUES(?, ?, ?)`
const loginUser = `SELECT id, username, password FROM user WHERE username=$1`
const attachFilter = `INSERT INTO user_filter(user_id, filter_id) VALUES(?, ?)`
const detachFilter = `DELETE FROM user_filter WHERE user_id=? AND filter_id=?`
const getUserFilter = `SELECT id FROM user_filter WHERE user_id=($1) AND filter_id=($2)`

// Filter Queries
const searchFilter = `SELECT id FROM filter WHERE UPPER(filter) = UPPER($1)`
const createFilter = `INSERT INTO filter(filter) VALUES(?)`
const getFilters = `SELECT id, filter FROM filter`
const getUserFilters = `SELECT f.id, f.filter FROM filter f INNER JOIN user_filter uf ON (f.id = uf.filter_id) WHERE uf.user_id=?`
