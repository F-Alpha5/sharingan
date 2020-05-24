package cli

import (
	"github.com/urfave/cli/v2"
)

func GetGlobalFlags(s *settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "target-name",
			Aliases:     []string{"target"},
			Value:       "",
			Usage:       "Target domain",
			Destination: &s.Target,
		},
		&cli.IntFlag{
			Name:        "threads",
			Aliases:     []string{"t"},
			Value:       20,
			Usage:       "Max number of go routines",
			Destination: &s.Threads,
		},
		&cli.BoolFlag{
			Name:        "no-prompt",
			Value:       false,
			Usage:       "Disable prompts and continue without confirmation",
			Destination: &s.NoPrompt,
		},
	}
}

func GetNMapFlags() []cli.Flag {
	return []cli.Flag{}
}

func GetDNSFlags(s *DNSSettings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "wordlist",
			Aliases:     []string{"w", "dns-wordlist"},
			Value:       "",
			Usage:       "Wordlist for DNS bruteforcing",
			Destination: &s.DNSWordlist,
		},
		&cli.StringFlag{
			Name:        "domain",
			Aliases:     []string{"root-domain", "d"},
			Value:       "",
			Usage:       "Basis for subdomain scanning",
			Destination: &s.RootDomain,
		},
		&cli.BoolFlag{
			Name:        "skip-probe",
			Usage:       "Skips host-up nmap scan",
			Destination: &s.SkipProbe,
		},
		&cli.BoolFlag{
			Name:        "replace-subs",
			Usage:       "used with add subs to replace all subs for program",
			Destination: &s.ReplaceSubs,
		},
		&cli.BoolFlag{
			Name:        "rescan",
			Usage:       "Scans domain regardless of the existance of previous results",
			Destination: &s.Rescan,
		},
	}
}

// GetDirbFlags return flags to CLI command struct, add flags to settings struct
func GetDirbFlags(s *DirbSettings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "wordlist",
			Aliases:     []string{"w", "dirb-wordlist"},
			Value:       "",
			Usage:       "Wordlist for directory busting",
			Destination: &s.Wordlist,
		},
		&cli.StringFlag{
			Name:        "domain",
			Aliases:     []string{"d"},
			Value:       "",
			Usage:       "Domain + Path to start bruteforcing from",
			Destination: &s.Domain,
		},
	}
}
