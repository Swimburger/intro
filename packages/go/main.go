package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

// Profile data
type Profile struct {
	Name          string `json:"name" yaml:"name"`
	Nickname      string `json:"nickname" yaml:"nickname"`
	Title         string `json:"title" yaml:"title"`
	Location      string `json:"location" yaml:"location"`
	Timezone      string `json:"timezone" yaml:"timezone"`
	Bio           string `json:"bio" yaml:"bio"`
	Website       string `json:"website" yaml:"website"`
	GitHub        string `json:"github" yaml:"github"`
	Twitter       string `json:"twitter" yaml:"twitter"`
	LinkedIn      string `json:"linkedin" yaml:"linkedin"`
	YouTube       string `json:"youtube" yaml:"youtube"`
	Facebook      string `json:"facebook" yaml:"facebook"`
	StackOverflow string `json:"stackoverflow" yaml:"stackoverflow"`
	MVP           string `json:"mvp" yaml:"mvp"`
	Bluesky       string `json:"bluesky" yaml:"bluesky"`
	Mastodon      string `json:"mastodon" yaml:"mastodon"`
	SponsorPayPal string `json:"sponsor_paypal" yaml:"sponsor_paypal"`
	SponsorGitHub string `json:"sponsor_github" yaml:"sponsor_github"`
}

func main() {
	// Check for flags
	isHelpMode := false
	isJsonMode := false
	isYamlMode := false
	isVcardMode := false
	for _, arg := range os.Args[1:] {
		if arg == "--help" || arg == "-h" {
			isHelpMode = true
		}
		if arg == "--json" {
			isJsonMode = true
		}
		if arg == "--yaml" {
			isYamlMode = true
		}
		if arg == "--vcard" {
			isVcardMode = true
		}
	}

	profile := Profile{
		Name:          "Niels Swimberghe",
		Nickname:      "Swimburger",
		Title:         "C# and TypeScript SDK generator engineer, at Fern",
		Location:      "NYC",
		Timezone:      "Eastern Time",
		Bio:           "Niels Swimberghe is a Belgian American software engineer and Microsoft MVP at Fern where he owns the C# and TypeScript SDK generators.\nGet in touch with Niels on Twitter @RealSwimburger and follow Niels' personal blog on .NET, Azure, and web development at swimburger.net.",
		Website:       "https://swimburger.net",
		GitHub:        "https://github.com/Swimburger",
		Twitter:       "https://twitter.com/RealSwimburger",
		LinkedIn:      "https://linkedin.com/in/nielsswimberghe",
		YouTube:       "https://youtube.com/@RealSwimburger",
		Facebook:      "https://facebook.com/SwimburgerDotNet",
		StackOverflow: "https://stackoverflow.com/users/2919731/swimburger",
		MVP:           "https://mvp.microsoft.com (Microsoft MVP)",
		Bluesky:       "https://bsky.app/profile/swimburger.bsky.social",
		Mastodon:      "@swimburger@dotnet.social",
		SponsorPayPal: "https://www.paypal.com/paypalme/swimburger",
		SponsorGitHub: "https://github.com/sponsors/Swimburger",
	}

	// If --help flag is provided, display usage and exit
	if isHelpMode {
		fmt.Printf("\n\033[1m\033[38;5;161mswimburger\033[0m - Personal intro card for Niels Swimberghe (Swimburger)\n\n")
		fmt.Printf("\033[1mUSAGE:\033[0m\n")
		fmt.Printf("  swimburger [OPTIONS]\n\n")
		fmt.Printf("\033[1mOPTIONS:\033[0m\n")
		fmt.Printf("  \033[36m--help, -h\033[0m    Show this help message\n")
		fmt.Printf("  \033[36m--json\033[0m        Output profile data as JSON\n")
		fmt.Printf("  \033[36m--yaml\033[0m        Output profile data as YAML\n")
		fmt.Printf("  \033[36m--vcard\033[0m       Output profile data as vCard (VCF format)\n\n")
		fmt.Printf("\033[1mEXAMPLES:\033[0m\n")
		fmt.Printf("  swimburger              Display visual profile card\n")
		fmt.Printf("  swimburger --json       Output as JSON\n")
		fmt.Printf("  swimburger --yaml       Output as YAML\n")
		fmt.Printf("  swimburger --vcard      Output as vCard\n\n")
		fmt.Printf("\033[1mMORE INFO:\033[0m\n")
		fmt.Printf("  Website: https://swimburger.net\n")
		fmt.Printf("  GitHub:  https://github.com/Swimburger\n\n")
		return
	}

	// If --json flag is provided, output JSON and exit
	if isJsonMode {
		jsonData, err := json.MarshalIndent(profile, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
		return
	}

	// If --yaml flag is provided, output YAML and exit
	if isYamlMode {
		yamlData, err := yaml.Marshal(profile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling YAML: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(yamlData))
		return
	}

	// If --vcard flag is provided, output vCard and exit
	if isVcardMode {
		vcard := fmt.Sprintf(`BEGIN:VCARD
VERSION:3.0
FN:%s
NICKNAME:%s
TITLE:%s
NOTE:%s
URL:%s
URL;type=github:%s
URL;type=twitter:%s
URL;type=linkedin:%s
URL;type=youtube:%s
URL;type=facebook:%s
URL;type=stackoverflow:%s
URL;type=mvp:%s
URL;type=bluesky:%s
X-SOCIALPROFILE;type=mastodon:%s
URL;type=sponsor:%s
URL;type=sponsor:%s
ADR;TYPE=WORK:;;%s;;;%s;
END:VCARD`,
			profile.Name,
			profile.Nickname,
			profile.Title,
			profile.Bio,
			profile.Website,
			profile.GitHub,
			profile.Twitter,
			profile.LinkedIn,
			profile.YouTube,
			profile.Facebook,
			profile.StackOverflow,
			profile.MVP,
			profile.Bluesky,
			profile.Mastodon,
			profile.SponsorPayPal,
			profile.SponsorGitHub,
			profile.Location,
			profile.Timezone,
		)
		fmt.Println(vcard)
		return
	}

	// Define styles
	nameStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#a02c56"))

	titleStyle := lipgloss.NewStyle().
		Faint(true)

	bioStyle := lipgloss.NewStyle().
		Italic(true)

	labelStyle := lipgloss.NewStyle().
		Bold(true)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#a02c56"))

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#a02c56")).
		Padding(1, 2).
		Width(145).
		MarginTop(1).
		MarginBottom(1)

	// Build the card content
	content := fmt.Sprintf("%s\n%s\n\n%s\n\n%s\n🌐 %s %s\n💻 %s %s\n🐦 %s %s\n💼 %s %s\n🎥 %s %s\n👥 %s %s\n💬 %s %s\n🦋 %s %s\n🐘 %s %s\n\n%s\n💖 %s %s\n🎁 %s %s",
		nameStyle.Render(fmt.Sprintf("🍔 %s (%s)", profile.Name, profile.Nickname)),
		titleStyle.Render("🌎 "+profile.Location+" | "+profile.Timezone),
		bioStyle.Render(profile.Bio),
		sectionStyle.Render("━━━ Connect ━━━"),
		labelStyle.Render("Website:       "),
		valueStyle.Render(profile.Website),
		labelStyle.Render("GitHub:        "),
		valueStyle.Render(profile.GitHub),
		labelStyle.Render("Twitter:       "),
		valueStyle.Render(profile.Twitter),
		labelStyle.Render("LinkedIn:      "),
		valueStyle.Render(profile.LinkedIn),
		labelStyle.Render("YouTube:       "),
		valueStyle.Render(profile.YouTube),
		labelStyle.Render("Facebook:      "),
		valueStyle.Render(profile.Facebook),
		labelStyle.Render("Stack Overflow:"),
		valueStyle.Render(profile.StackOverflow),
		labelStyle.Render("Bluesky:       "),
		valueStyle.Render(profile.Bluesky),
		labelStyle.Render("Mastodon:      "),
		valueStyle.Render(profile.Mastodon),
		sectionStyle.Render("━━━ Support ━━━"),
		labelStyle.Render("PayPal:         "),
		valueStyle.Render(profile.SponsorPayPal),
		labelStyle.Render("GitHub Sponsor:"),
		valueStyle.Render(profile.SponsorGitHub),
	)

	// Display the card
	fmt.Println(boxStyle.Render(content))
}
