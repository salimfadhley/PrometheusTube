-- +goose Up
ALTER TABLE videos DROP COLUMN IF EXISTS download_id;