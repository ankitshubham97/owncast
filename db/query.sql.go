// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const addFollower = `-- name: AddFollower :exec
INSERT INTO ap_followers(iri, inbox, name, username, image, approved_at) values($1, $2, $3, $4, $5, $6)
`

type AddFollowerParams struct {
	Iri        string
	Inbox      string
	Name       sql.NullString
	Username   string
	Image      sql.NullString
	ApprovedAt sql.NullTime
}

func (q *Queries) AddFollower(ctx context.Context, arg AddFollowerParams) error {
	_, err := q.db.ExecContext(ctx, addFollower,
		arg.Iri,
		arg.Inbox,
		arg.Name,
		arg.Username,
		arg.Image,
		arg.ApprovedAt,
	)
	return err
}

const addToOutbox = `-- name: AddToOutbox :exec
INSERT INTO ap_outbox(id, iri, value, type) values($1, $2, $3, $4)
`

type AddToOutboxParams struct {
	ID    string
	Iri   string
	Value []byte
	Type  string
}

func (q *Queries) AddToOutbox(ctx context.Context, arg AddToOutboxParams) error {
	_, err := q.db.ExecContext(ctx, addToOutbox,
		arg.ID,
		arg.Iri,
		arg.Value,
		arg.Type,
	)
	return err
}

const approveFederationFollower = `-- name: ApproveFederationFollower :exec
UPDATE ap_followers SET approved_at = $1 WHERE iri = $2
`

type ApproveFederationFollowerParams struct {
	ApprovedAt sql.NullTime
	Iri        string
}

func (q *Queries) ApproveFederationFollower(ctx context.Context, arg ApproveFederationFollowerParams) error {
	_, err := q.db.ExecContext(ctx, approveFederationFollower, arg.ApprovedAt, arg.Iri)
	return err
}

const getFederationFollowerApprovalRequests = `-- name: GetFederationFollowerApprovalRequests :many
SELECT iri, inbox, name, username, image FROM ap_followers WHERE approved_at = null
`

type GetFederationFollowerApprovalRequestsRow struct {
	Iri      string
	Inbox    string
	Name     sql.NullString
	Username string
	Image    sql.NullString
}

func (q *Queries) GetFederationFollowerApprovalRequests(ctx context.Context) ([]GetFederationFollowerApprovalRequestsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFederationFollowerApprovalRequests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFederationFollowerApprovalRequestsRow
	for rows.Next() {
		var i GetFederationFollowerApprovalRequestsRow
		if err := rows.Scan(
			&i.Iri,
			&i.Inbox,
			&i.Name,
			&i.Username,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFederationFollowersWithOffset = `-- name: GetFederationFollowersWithOffset :many
SELECT iri, inbox, name, username, image, created_at FROM ap_followers WHERE approved_at is not null LIMIT $1 OFFSET $2
`

type GetFederationFollowersWithOffsetParams struct {
	Limit  int32
	Offset int32
}

type GetFederationFollowersWithOffsetRow struct {
	Iri       string
	Inbox     string
	Name      sql.NullString
	Username  string
	Image     sql.NullString
	CreatedAt sql.NullTime
}

func (q *Queries) GetFederationFollowersWithOffset(ctx context.Context, arg GetFederationFollowersWithOffsetParams) ([]GetFederationFollowersWithOffsetRow, error) {
	rows, err := q.db.QueryContext(ctx, getFederationFollowersWithOffset, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFederationFollowersWithOffsetRow
	for rows.Next() {
		var i GetFederationFollowersWithOffsetRow
		if err := rows.Scan(
			&i.Iri,
			&i.Inbox,
			&i.Name,
			&i.Username,
			&i.Image,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowerCount = `-- name: GetFollowerCount :one


SElECT count(*) FROM ap_followers
`

// Queries added to query.sql must be compiled into Go code with sqlc. Read README.md for details.
// Federation related queries.
func (q *Queries) GetFollowerCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowerCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLocalPostCount = `-- name: GetLocalPostCount :one
SElECT count(*) FROM ap_outbox
`

func (q *Queries) GetLocalPostCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLocalPostCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getObjectFromOutboxByID = `-- name: GetObjectFromOutboxByID :one
SELECT value FROM ap_outbox WHERE id = $1
`

func (q *Queries) GetObjectFromOutboxByID(ctx context.Context, id string) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, getObjectFromOutboxByID, id)
	var value []byte
	err := row.Scan(&value)
	return value, err
}

const getObjectFromOutboxByIRI = `-- name: GetObjectFromOutboxByIRI :one
SELECT value FROM ap_outbox WHERE iri = $1
`

func (q *Queries) GetObjectFromOutboxByIRI(ctx context.Context, iri string) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, getObjectFromOutboxByIRI, iri)
	var value []byte
	err := row.Scan(&value)
	return value, err
}

const getOutbox = `-- name: GetOutbox :many
SELECT value FROM ap_outbox
`

func (q *Queries) GetOutbox(ctx context.Context) ([][]byte, error) {
	rows, err := q.db.QueryContext(ctx, getOutbox)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][]byte
	for rows.Next() {
		var value []byte
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		items = append(items, value)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeFollowerByIRI = `-- name: RemoveFollowerByIRI :exec
DELETE FROM ap_followers WHERE iri = $1
`

func (q *Queries) RemoveFollowerByIRI(ctx context.Context, iri string) error {
	_, err := q.db.ExecContext(ctx, removeFollowerByIRI, iri)
	return err
}

const updateFollowerByIRI = `-- name: UpdateFollowerByIRI :exec
UPDATE ap_followers SET inbox = $1, name = $2, username = $3, image = $4 WHERE iri = $5
`

type UpdateFollowerByIRIParams struct {
	Inbox    string
	Name     sql.NullString
	Username string
	Image    sql.NullString
	Iri      string
}

func (q *Queries) UpdateFollowerByIRI(ctx context.Context, arg UpdateFollowerByIRIParams) error {
	_, err := q.db.ExecContext(ctx, updateFollowerByIRI,
		arg.Inbox,
		arg.Name,
		arg.Username,
		arg.Image,
		arg.Iri,
	)
	return err
}
