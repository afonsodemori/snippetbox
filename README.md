# Snippetbox

A full-stack web application for creating and sharing text snippets, built with Go using only the standard library — no frameworks.

This is a learning project built section-by-section through [_Let's Go_ by Alex Edwards](https://lets-go.alexedwards.net/). Each commit maps 1:1 to a book section.

> \[!NOTE]
> This repository is actively under development. I am currently working through the book, and features are being added incrementally. See [PROGRESS.md](PROGRESS.md) for the current completion status and full chapter map.

## Tech Stack

- **Go** — HTTP server, routing, templating, and middleware via `net/http` and `html/template`
- **MySQL** — relational data storage with a hand-rolled database model
- **HTML/CSS** — server-rendered templates with a custom inheritance layout system
- **TLS** — HTTPS with self-signed certificates

## Features

- Create, view, and share snippets with expiry dates
- User accounts with signup, login, and session management
- CSRF protection and secure cookie handling
- Structured logging, panic recovery middleware, and centralized error handling

## Getting Started

To run the application locally at its current stage of development:

```bash
git clone https://github.com/afonsodemori/snippetbox.git
cd snippetbox
go run ./cmd/web
```

## Licence

[MIT](LICENSE)
