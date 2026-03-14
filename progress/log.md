# Progress Log

Track daily sessions here. Update after each session.

**Format**: Copy the template below for each day.

---

## How to Use This Log

After each session, add an entry below. The tutor will read this at the start of each session to know where you are. Be honest about what was confusing — that's the most useful signal.

---

## Template

### Day X — YYYY-MM-DD
**Week**: X | **Topic**: ...
**Session type**: (guided exercise / code review / pair programming / etc.)
**What I did**:
- ...

**What clicked**:
- ...

**What was confusing**:
- ...

**Code written**: (link to file in exercises/ or projects/ if applicable)

**Tutor notes**: (tutor fills this in)

---

## Entries

(Start logging below this line)

### Day 7 — 2026-03-14
**Week**: 2 | **Topic**: Loops & range — data transformations
**Session type**: Guided exercise
**What we did**:
- Covered all forms of `for`: classic, while-style, infinite, range over slice, range over map
- Implemented `filterAbove`, `applyTax`, `cartTotal` — Go equivalents of `.filter()`, `.map()`, `.reduce()`
- Chained all three together for a realistic pipeline
- Reorganized week-02 exercises into per-day subdirectories to fix `main redeclared` conflict

**What clicked**:
- `for _, value := range` — used correctly this time (improvement from day 6)
- Building new slices with `append` for filter/map patterns
- Function composition as Go's alternative to method chaining
- Integer math for currency (`price * 110 / 100`)

**What needs reinforcement**:
- No issues today — all three functions written correctly on first attempt

**Code written**: exercises/week-02/day07/main.go

**Tutor notes**: All exercises done independently, clean code, correct results. Used idiomatic two-variable range throughout. Good understanding that Go has no built-in filter/map — explicit loops are the Go way. Ready for Day 8 (pointers).

---

### Day 6 — 2026-03-13
**Week**: 2 | **Topic**: Arrays, slices, and maps
**Session type**: Compare & contrast + exercise
**What we did**:
- Predicted output of 4 snippets covering: array copy vs slice reference, sub-slice mutation, map zero values with comma-ok idiom
- Implemented `cartTotal(items map[string]int) int` — Go equivalent of `Object.values().reduce()`
- Initialized Go module (`go mod init`)

**What clicked**:
- Arrays are value types (copy on assign), slices are reference types — got all predictions right
- Sub-slices share underlying array — predicted the mutation correctly
- Comma-ok idiom for map lookups — predicted `0, false` for missing key
- `make(map[...])` requirement for map initialization

**What needs reinforcement**:
- Idiomatic `for _, value := range map` — used `for key := range` + manual lookup instead (works but extra step)

**Code written**: exercises/week-02/day06_collections.go

**Tutor notes**: Perfect 4/4 on predictions — the value-vs-reference distinction clicked immediately (likely from JS experience with objects vs primitives). cartTotal implementation was correct, just not yet using the two-variable range form. Quick session, strong start to Week 2. Next: Day 7 (loops & range, data transformation).

---

### Day 5 — 2026-03-12
**Week**: 1 | **Topic**: Review — read 3 Go snippets, predict output
**Session type**: Quiz + review
**What we did**:
- Predicted output of 3 snippets covering: multiple returns, `_` discard, zero value traps, error wrapping
- Got 2/3 fully correct, missed that `double("5")` returns `n * 2` = 10, not 5

**What clicked**:
- Zero values in structs — correctly predicted `Rectangle{Width: 3}` has `Height` = 0
- `_` to discard unwanted return values
- Error wrapping output format

**What needs reinforcement**:
- Read function bodies carefully — missed the `* 2` in `double`

**Tutor notes**: Strong Week 1 finish. Concepts are landing well — error handling, structs, interfaces, zero values all solid. Ready for Week 2 (slices, maps, pointers, packages).

---

### Day 4 — 2026-03-12
**Week**: 1 | **Topic**: Structs vs interfaces vs TS classes
**Session type**: Code reading + exercise
**What we did**:
- Compared TS classes to Go structs + methods
- Learned method receivers `(p Payment)` syntax
- Saw implicit interface satisfaction — no `implements` keyword
- Built a `Charge` struct that satisfies `Summarizable` interface
- Introduced value vs pointer receivers (preview for Week 2)

**What clicked**:
- Struct creation with `Type{field: value}` — no `new` keyword
- Implicit interface satisfaction — just add the method
- Method receiver syntax felt natural

**Code written**: exercises/week-01/day04_structs.go

**Tutor notes**: Completed exercise independently and quickly (<10 min). Moved straight to Day 5 in the same session. Structs and interfaces clicked immediately.

---

### Day 3 — 2026-03-11
**Week**: 1 | **Topic**: Functions, multiple returns, error handling
**Session type**: Exercise — rewrite a TS function in Go
**What we did**:
- Rewrote a TypeScript `parseAmount("$49.99") → 4999` function in Go
- Used `strings.HasPrefix`, `strconv.ParseFloat`, `math.Round`
- Practiced returning `(value, error)` pairs instead of throwing
- Applied happy path on the left idiom — both error checks return early
- Learned about error wrapping with `fmt.Errorf` and `%w`

**What clicked**:
- Happy path on the left — applied it naturally this time without prompting
- Multiple return values for error handling feels intuitive now
- String slicing `input[1:]` same as JS

**What needs reinforcement**:
- Error wrapping (`%w`) vs plain `errors.New` — when to use which
- Haven't written custom error types yet

**Code written**: exercises/week-01/day03_functions.go

**Tutor notes**: Wrote the function independently with no hints needed — clean code on first attempt. Used `errors.New` initially, then added `fmt.Errorf` wrapping after discussion. Happy path idiom is landing. Ready for Day 4 (structs vs interfaces vs TS classes).

---

### Day 2 — 2026-03-11
**Week**: 1 | **Topic**: Types, variables, zero values — Go vs TypeScript
**Session type**: Compare & contrast
**What we did**:
- Compared Go's type system to TypeScript's side by side
- Covered the 3 ways to declare variables (`var` with type, `var` with inference, `:=`)
- Explored zero values, explicit type conversion, and constants
- Fixed 3 type-error exercises: `float64` not `float`, explicit `float64()` conversion, `fmt.Sprintf` for string+int

**What clicked**:
- Explicit type conversion — understood why `float64(meters) / 1000` is needed
- `fmt.Sprintf` for string formatting (chose it over `strconv.Itoa`)

**What needs reinforcement**:
- `string(int)` trap (converts to Unicode codepoint, not string representation) — worth revisiting
- Sized numeric types (`int8`/`int16`/`int32`/`int64`) — only touched briefly
- Happy path idiom from Day 1 still needs practice (didn't come up today)

**Code written**: exercises/week-01/day02_types.go

**Tutor notes**: All 3 fixes done correctly after guidance. Hit the `string(42)` trap which is a great learning moment — now knows to use `fmt.Sprintf` or `strconv.Itoa`. Ready for Day 3 (functions, multiple returns, rewrite a TS function in Go).

---

### Day 1 — 2026-03-10
**Week**: 1 | **Topic**: Setup + Hello World + go run / go build
**Session type**: Guided exercise
**What we did**:
- Ran first Go program with `go run`
- Explored zero values (int, string, bool, float64)
- Used `:=` short variable declaration
- Called a function with multiple returns, handled errors with `if err != nil`
- Saw that Go refuses to compile with unused variables

**What clicked**:
- Error handling with multiple returns — got the `if err != nil` pattern right away
- Zero values concept landed (no undefined/null for basic types)

**What needs reinforcement**:
- Idiomatic style: early return instead of if/else after error check ("happy path on the left")
- Haven't explored `go build` yet (only used `go run`)

**Code written**: exercises/week-01/day01_hello.go

**Tutor notes**: Learner jumped right into uncommenting and running code — good instinct. Handled error return naturally. Introduced "happy path left" idiom; worth reinforcing tomorrow. Ready to move to Day 2 (types compare & contrast). Suggested optional exercise: write a `greet` function to practice function signatures.
