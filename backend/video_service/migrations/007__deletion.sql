-- +goose Up
ALTER TABLE videos ADD COLUMN is_deleted bool DEFAULT FALSE;
