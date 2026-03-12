# Go Learning Roadmap

**Duration**: 8 weeks (30 min/day, 5 days/week)
**Approach**: Learn concepts through building, leverage TypeScript/NestJS background
**Guiding principle**: Curation > generation — practice reading and evaluating Go code, not just writing it

---

## Phase 1: Foundations Through Contrast (Weeks 1–2)

### Week 1 — Go's Mental Model

**Goal**: Understand how Go thinks differently from TypeScript.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Setup + Hello World + `go run` / `go build` | Guided exercise |
| 2 | Types, variables, zero values — compared to TS | Compare & contrast |
| 3 | Functions, multiple returns, error handling | Exercise: rewrite a TS function in Go |
| 4 | Structs vs interfaces vs TS classes/interfaces | Code reading + discussion |
| 5 | Review: read 3 short Go snippets, predict output | Quiz + review |

**Key contrast**: Go has no classes, no exceptions, no generics (well, now it does but they're different). Error handling is explicit. This week is about rewiring instincts.

### Week 2 — Control Flow & Collections

**Goal**: Get comfortable with Go's data structures and flow control.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Arrays, slices, maps — compared to JS arrays/objects | Guided exercise |
| 2 | Loops (`for` is the only loop), range | Exercise: data transformation |
| 3 | Pointers — the concept TS devs don't have | Visual explanation + exercises |
| 4 | Packages, imports, visibility (uppercase = exported) | Restructure code into packages |
| 5 | Review: debug 3 Go programs with subtle bugs | Code review exercise |

**Key contrast**: Pointers exist. Slices are not arrays. Capitalization means something. These are the traps for TS developers.

---

## Phase 2: Building Something Real (Weeks 3–4)

### Week 3 — Project: CLI Tool

**Goal**: Build a command-line tool from scratch. Suggested: a payment webhook event logger.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Project setup, `cobra` or stdlib `flag`, basic CLI structure | Pair programming |
| 2 | File I/O — reading and writing JSON/files | Build feature |
| 3 | Error handling patterns in practice | Refactor for proper error handling |
| 4 | Testing with `go test` — table-driven tests | Write tests for existing code |
| 5 | Review: evaluate AI-generated Go code for your CLI | Curation exercise |

### Week 4 — Project: HTTP Server

**Goal**: Build a simple HTTP server. Suggested: a Stripe webhook receiver.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | `net/http` basics — handlers, mux, request/response | Pair programming |
| 2 | JSON marshaling/unmarshaling — compared to TS parsing | Build endpoint |
| 3 | Middleware pattern in Go vs NestJS middleware/guards | Build middleware |
| 4 | Project structure & dependency injection without a framework | Architecture discussion |
| 5 | Review: compare your Go server to equivalent NestJS code | Contrast & evaluate |

**Key insight**: Go doesn't have NestJS. There's no decorator magic. This is intentional. Understanding *why* Go is explicit where NestJS is implicit builds judgment.

---

## Phase 3: Go's Superpowers (Weeks 5–6)

### Week 5 — Concurrency

**Goal**: Learn goroutines and channels — Go's signature feature.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Goroutines — compared to JS async/await | Visual explanation + exercise |
| 2 | Channels — sending and receiving data between goroutines | Guided exercise |
| 3 | `select`, timeouts, context cancellation | Exercise: add timeouts to HTTP calls |
| 4 | Common concurrency patterns (fan-out, worker pools) | Build a concurrent webhook processor |
| 5 | Review: spot concurrency bugs in Go code | Code review exercise |

### Week 6 — Interfaces & Composition

**Goal**: Master Go's approach to polymorphism and code reuse.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Interfaces — implicit satisfaction vs TS explicit implements | Compare & contrast |
| 2 | The io.Reader/io.Writer pattern — why it's powerful | Exercise: build something composable |
| 3 | Embedding vs inheritance — struct composition | Refactor exercise |
| 4 | Interface design — small interfaces, big impact | Architecture discussion |
| 5 | Review: evaluate interface designs in open-source Go code | Curation exercise |

---

## Phase 4: Production Go (Weeks 7–8)

### Week 7 — Project: Payment Service

**Goal**: Build a small payment processing service combining everything learned.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Project architecture — design the service before coding | Architecture session |
| 2 | Database access with `sqlx` or `pgx` | Build feature |
| 3 | Graceful shutdown, health checks, configuration | Build feature |
| 4 | Structured logging, error wrapping | Refactor for production readiness |
| 5 | Integration testing patterns | Write integration tests |

### Week 8 — Polish & Reflect

**Goal**: Ship something you're proud of. Reflect on what you've learned.

| Day | Topic | Format |
|-----|-------|--------|
| 1 | Code review your own project end-to-end | Self-review + tutor feedback |
| 2 | Performance: benchmarking, profiling basics | Exercise |
| 3 | Linting, formatting, CI setup (`golangci-lint`) | Polish |
| 4 | Compare: how you'd build this differently in NestJS | Reflection |
| 5 | Retrospective: what worked, what's next | Planning session |

---

## Session Types Reference

Each day uses one of these formats:

- **Guided exercise**: Tutor introduces concept, you code with guidance
- **Compare & contrast**: See the same idea in TS/Go side by side, discuss tradeoffs
- **Pair programming**: Build features together, tutor writes some code, you write some
- **Code reading**: Read existing Go code and explain what it does
- **Code review exercise**: Evaluate Go code (often AI-generated) for correctness, style, performance
- **Curation exercise**: Given multiple implementations, choose the best and explain why
- **Architecture discussion**: Design systems on paper before coding
- **Quiz + review**: Predict outputs, find bugs, explain behavior

## Adapting the Roadmap

This roadmap is a starting point. The tutor will adjust based on:
- What clicks quickly (skip ahead)
- What needs more time (add practice days)
- What you find interesting (lean into it)
- Real work problems that come up (apply Go to them)
