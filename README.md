# Go Tutor — Daily Practice System

A structured learning system for learning Go through daily 30-minute sessions, guided by a Claude Project tutor that tracks your progress and adapts to your level.

## How It Works

1. **Check in daily** with your Claude Project tutor
2. The tutor reviews your progress log and assigns today's session
3. You work through exercises or project tasks — the tutor works alongside you
4. Log what you did and what you learned

## Your Background

- **Strong in**: TypeScript/JavaScript, React Native (2yr), NestJS (3yr), Next.js (6mo)
- **Domain**: Payment integrations (Stripe, Checkout.com)
- **Go level**: Read some Go, haven't built anything yet
- **Daily commitment**: 30 minutes

## Repo Structure

```
go-tutor/
├── README.md                  # You're here
├── curriculum/
│   └── roadmap.md             # Full learning roadmap (8 weeks)
├── progress/
│   └── log.md                 # Daily progress log (you + tutor update this)
├── exercises/
│   ├── week-01/               # Exercises organized by week
│   ├── week-02/
│   ├── week-03/
│   └── week-04/
├── projects/
│   └── (project folders appear as you progress)
├── .claude/
│   └── project-instructions.md  # Copy this into your Claude Project
└── go.mod                     # (created when you init your first Go module)
```

## Getting Started

### 1. Clone this repo

```bash
git clone <your-repo-url>
cd go-tutor
```

### 2. Install Go

```bash
# macOS
brew install go

# Verify
go version
```

### 3. Set up the Claude Project

1. Go to [claude.ai](https://claude.ai) → Projects → New Project
2. Copy the contents of `.claude/project-instructions.md` into the project's custom instructions
3. Upload the `curriculum/roadmap.md` and `progress/log.md` files as project knowledge

### 4. Start your first session

Open the Claude Project and say:

> "I'm ready for my first session."

## Philosophy

This system is built on the idea that **learning a programming language with an LLM should be fundamentally different** from traditional learning:

- **The tutor codes alongside you** — not just explains, but pair-programs
- **Concepts are learned through building** — not isolated exercises
- **Your existing knowledge is leveraged** — Go is taught through the lens of TypeScript/NestJS patterns you already know
- **Curation over generation** — you'll practice reading, evaluating, and improving code as much as writing it
