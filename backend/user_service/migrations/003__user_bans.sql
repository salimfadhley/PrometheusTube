-- +goose Up
ALTER TABLE users ADD COLUMN banned bool DEFAULT false;