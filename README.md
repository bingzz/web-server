# Golang Web Server tutorial sample project

This is the basic overview of using Golang as a Web Server using Gin Framework


## Install packages:

Use go get `<package_url>`
> This is similar to installing packages in NodeJs

| Package | Uses |
| - | - |
| `github.com/gofor-little/env` | `.env` reader | 
| `github.com/gin-gonic/gin` | Golang API Web Framework |


## Setting up Database (Postgres)

| Command | Description |
| - | - |
| `systemctl status postgresql` | Check RDBMS status |
| `psql -U postgres -d personal_db -h 127.0.0.1 -W` | Start Connection |
| `sudo ufw status` | Check port status |
| `sudo ufw allow 5432/tcp` | Allow port 5432 |