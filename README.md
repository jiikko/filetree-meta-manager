# README

## Ruby version

`cat .ruby-version`

## System dependencies

## Configuration

## Setup

## How to run the test suite

## Services (job queues, cache servers, search engines, etc.)

## Deployment

## デバッグ

- 開発環境で filetree を取得する
  - cd client; go run cmd/dump-filetree/main.go tmp
- 開発環境で filetree を登録する
  - `curl -X POST 'http://localhost:3000/api/v1/filetrees' -H 'Accept: application/json' -H 'Content-Type: application/json' -H 'Authorization: 018fd3b6-5645-7ad0-8b7b-3660c43e98f2' -d '{"filetree": {"key1": "value1", "key2": "value2"}, "device": "8TB" }'`
