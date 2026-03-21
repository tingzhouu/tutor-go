package main

import (
	"flag"
	"fmt"
	"os"
)

// What we're building this week:
// A CLI tool that logs and queries payment webhook events.
//
// Usage:
//   webhook-logger log --type payment_intent.succeeded --amount 4999
//   webhook-logger list
//   webhook-logger list --type payment_intent.failed
//   webhook-logger summary
//
// Today (Day 11): CLI skeleton with flag parsing
// Day 12: File I/O — persist events to JSON
// Day 13: Error handling — make it robust
// Day 14: Tests — table-driven tests
// Day 15: Review AI-generated code for the tool

func main() {
	// In NestJS you'd use a CLI framework like commander or yargs.
	// In Go, the stdlib `flag` package handles simple cases.
	// For complex CLIs, people use cobra — but flag is enough for us.

	// Subcommand approach: check os.Args[1] for the command name,
	// then use flag.NewFlagSet for each subcommand's flags.

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "log":
		if err := handleLog(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error %s\n", err)
			os.Exit(1)
		}
	case "list":
		if err := handleList(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "error %s\n", err)
			os.Exit(1)
		}
	case "summary":
		if err := handleSummary("events.json"); err != nil {
			fmt.Fprintf(os.Stderr, "error %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: webhook-logger <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  log      Log a new webhook event")
	fmt.Println("  list     List logged events")
	fmt.Println("  summary  Show event summary")
}

// TODO 1: Implement handleLog
// Parse these flags:
//   --type   string  (required) e.g. "payment_intent.succeeded"
//   --amount int     (required) amount in cents
//   --id     string  (optional) event ID, default "evt_" + some placeholder for now
//
// Use flag.NewFlagSet("log", flag.ExitOnError) to create a subcommand flag set.
// Then call flagSet.StringVar / flagSet.IntVar to define flags.
// Then flagSet.Parse(args) to parse.
//
// For now, just print the parsed values. We'll save to file tomorrow.
//
// TS equivalent of flag parsing:
//   const { type, amount, id } = yargs(process.argv).option('type', { ... })

func handleLog(args []string) error {
	// Your code here
	fs := flag.NewFlagSet("log", flag.ExitOnError)
	var eventType string
	var amount int
	var id string
	fs.StringVar(&eventType, "type", "", "event type")
	fs.StringVar(&id, "id", "evt_XXX", "event id")
	fs.IntVar(&amount, "amount", 0, "amount in cents")
	fs.Parse(args)

	if eventType == "" {
		return fmt.Errorf("--type is required")
	}

	if amount <= 0 {
		return fmt.Errorf("--amount is required and must be positive")
	}

	fmt.Printf("eventType %s, amount %d, id %s\n", eventType, amount, id)
	events, err := loadEvents("events.json")
	if err != nil {
		return fmt.Errorf("loading events: %w", err)
	}
	event := newEvent(id, eventType, amount)
	events = append(events, event)
	err = saveEvents("events.json", events)
	if err != nil {
		return fmt.Errorf("error saving events: %w", err)
	}
	fmt.Printf("Successfully saved event of id %s\n", id)
	return nil
}

// TODO 2: Implement handleList
// Parse one optional flag:
//   --type string  (optional) filter by event type
//
// For now, just print "Listing events..." and the filter if provided.

func handleList(args []string) error {
	// Your code here
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	var eventType string
	fs.StringVar(&eventType, "type", "", "filter by event type")
	fs.Parse(args)

	fmt.Printf("Listing events... %s\n", eventType)

	events, err := loadEvents("events.json")
	if err != nil {
		return fmt.Errorf("encountered err: %w", err)
	}
	for _, e := range events {
		if eventType != "" && eventType != e.Type {
			continue
		}
		fmt.Printf("eventType %s, amount %d, id %s\n", e.Type, e.Amount, e.ID)
	}
	return nil
}

// TODO 3: Implement handleSummary
// No flags needed. Just print "Summary coming soon..." for now.

func handleSummary(path string) error {
	// Your code here
	events, err := loadEvents(path)
	if err != nil {
		return fmt.Errorf("Unexpected error when loading events: %w", err)
	}

	eventAmount := make(map[string]int)
	eventCount := make(map[string]int)
	totalAmount := 0
	for _, event := range events {
		eventAmount[event.Type] += event.Amount
		eventCount[event.Type]++
		totalAmount += event.Amount
	}

	for eventType, amount := range eventAmount {
		fmt.Printf("%s\t%d\tevents\t$%.2f\n", eventType, eventCount[eventType], float64(amount)/100)
	}
	fmt.Printf("Total %d events\t$%.2f\n", len(events), float64(totalAmount)/100)
	return nil
}
