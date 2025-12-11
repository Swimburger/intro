using Spectre.Console;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Linq;
using YamlDotNet.Serialization;

// Check for flags
var isHelpMode = args.Length > 0 && args.Any(arg => arg == "--help" || arg == "-h");
var isJsonMode = args.Length > 0 && args.Any(arg => arg == "--json");
var isYamlMode = args.Length > 0 && args.Any(arg => arg == "--yaml");
var isVcardMode = args.Length > 0 && args.Any(arg => arg == "--vcard");

// Profile data
var profile = new Profile
{
    Name = "Niels Swimberghe",
    Nickname = "Swimburger",
    Title = "C# and TypeScript SDK generator engineer, at Fern",
    Location = "NYC",
    Timezone = "Eastern Time",
    Bio = "Niels Swimberghe is a Belgian American software engineer and Microsoft MVP at Fern where he owns the C# and TypeScript SDK generators.\nGet in touch with Niels on Twitter @RealSwimburger and follow Niels' personal blog on .NET, Azure, and web development at swimburger.net.",
    Website = "https://swimburger.net",
    GitHub = "https://github.com/Swimburger",
    Twitter = "https://twitter.com/RealSwimburger",
    LinkedIn = "https://linkedin.com/in/nielsswimberghe",
    YouTube = "https://youtube.com/@RealSwimburger",
    Facebook = "https://facebook.com/SwimburgerDotNet",
    StackOverflow = "https://stackoverflow.com/users/2919731/swimburger",
    MVP = "https://mvp.microsoft.com (Microsoft MVP)",
    Bluesky = "https://bsky.app/profile/swimburger.bsky.social",
    Mastodon = "@swimburger@dotnet.social",
    SponsorPayPal = "https://www.paypal.com/paypalme/swimburger",
    SponsorGitHub = "https://github.com/sponsors/Swimburger"
};

// If --help flag is provided, display usage and exit
if (isHelpMode)
{
    System.Console.WriteLine(@"
[bold #a02c56]swimburger[/] - Personal intro card for Niels Swimberghe (Swimburger)

[bold]USAGE:[/]
  swimburger [OPTIONS]

[bold]OPTIONS:[/]
  [cyan]--help, -h[/]    Show this help message
  [cyan]--json[/]        Output profile data as JSON
  [cyan]--yaml[/]        Output profile data as YAML
  [cyan]--vcard[/]       Output profile data as vCard (VCF format)

[bold]EXAMPLES:[/]
  swimburger              Display visual profile card
  swimburger --json       Output as JSON
  swimburger --yaml       Output as YAML
  swimburger --vcard      Output as vCard

[bold]MORE INFO:[/]
  Website: https://swimburger.net
  GitHub:  https://github.com/Swimburger
"
        .Replace("[bold #a02c56]", "\x1b[1;38;5;161m")
        .Replace("[bold]", "\x1b[1m")
        .Replace("[cyan]", "\x1b[36m")
        .Replace("[/]", "\x1b[0m"));
    return;
}

// If --json flag is provided, output JSON and exit
if (isJsonMode)
{
    var json = JsonSerializer.Serialize(profile, new JsonSerializerOptions { WriteIndented = true });
    System.Console.WriteLine(json);
    return;
}

// If --yaml flag is provided, output YAML and exit
if (isYamlMode)
{
    var serializer = new SerializerBuilder().Build();
    var yaml = serializer.Serialize(profile);
    System.Console.WriteLine(yaml);
    return;
}

// If --vcard flag is provided, output vCard and exit
if (isVcardMode)
{
    var vcard = $@"BEGIN:VCARD
VERSION:3.0
FN:{profile.Name}
NICKNAME:{profile.Nickname}
TITLE:{profile.Title}
NOTE:{profile.Bio.Replace("\n", "\\n")}
URL:{profile.Website}
URL;type=github:{profile.GitHub}
URL;type=twitter:{profile.Twitter}
URL;type=linkedin:{profile.LinkedIn}
URL;type=youtube:{profile.YouTube}
URL;type=facebook:{profile.Facebook}
URL;type=stackoverflow:{profile.StackOverflow}
URL;type=mvp:{profile.MVP}
URL;type=bluesky:{profile.Bluesky}
X-SOCIALPROFILE;type=mastodon:{profile.Mastodon}
URL;type=sponsor:{profile.SponsorPayPal}
URL;type=sponsor:{profile.SponsorGitHub}
ADR;TYPE=WORK:;;{profile.Location};;;{profile.Timezone};
END:VCARD";
    System.Console.WriteLine(vcard);
    return;
}

// Create a panel with the profile information
var panel = new Panel(
    Align.Left(
        new Markup($@"[bold #a02c56]🍔 {profile.Name} ({profile.Nickname})[/]
[dim]🌎 {profile.Location} | {profile.Timezone}[/]

[italic]{profile.Bio}[/]

[bold #a02c56]━━━ Connect ━━━[/]
🌐 [bold]Website:[/]        [grey link={profile.Website}]{profile.Website}[/]
💻 [bold]GitHub:[/]         [grey link={profile.GitHub}]{profile.GitHub}[/]
🐦 [bold]Twitter:[/]        [grey link={profile.Twitter}]{profile.Twitter}[/]
💼 [bold]LinkedIn:[/]       [grey link={profile.LinkedIn}]{profile.LinkedIn}[/]
🎥 [bold]YouTube:[/]        [grey link={profile.YouTube}]{profile.YouTube}[/]
👥 [bold]Facebook:[/]       [grey link={profile.Facebook}]{profile.Facebook}[/]
💬 [bold]Stack Overflow:[/] [grey link={profile.StackOverflow}]{profile.StackOverflow}[/]
🦋 [bold]Bluesky:[/]        [grey link={profile.Bluesky}]{profile.Bluesky}[/]
🐘 [bold]Mastodon:[/]       [grey]{profile.Mastodon}[/]

[bold #a02c56]━━━ Support ━━━[/]
💖 [bold]PayPal:[/]         [grey link={profile.SponsorPayPal}]{profile.SponsorPayPal}[/]
🎁 [bold]GitHub Sponsor:[/] [grey link={profile.SponsorGitHub}]{profile.SponsorGitHub}[/]")
    )
)
{
    Border = BoxBorder.Rounded,
    BorderStyle = new Style(new Color(160, 44, 86)),
    Padding = new Padding(1, 0, 1, 0)
};

AnsiConsole.Write(panel);

class Profile
{
    [JsonPropertyName("name")]
    [YamlMember(Alias = "name")]
    public string Name { get; set; }

    [JsonPropertyName("nickname")]
    [YamlMember(Alias = "nickname")]
    public string Nickname { get; set; }

    [JsonPropertyName("title")]
    [YamlMember(Alias = "title")]
    public string Title { get; set; }

    [JsonPropertyName("location")]
    [YamlMember(Alias = "location")]
    public string Location { get; set; }

    [JsonPropertyName("timezone")]
    [YamlMember(Alias = "timezone")]
    public string Timezone { get; set; }

    [JsonPropertyName("bio")]
    [YamlMember(Alias = "bio")]
    public string Bio { get; set; }

    [JsonPropertyName("website")]
    [YamlMember(Alias = "website")]
    public string Website { get; set; }

    [JsonPropertyName("github")]
    [YamlMember(Alias = "github")]
    public string GitHub { get; set; }

    [JsonPropertyName("twitter")]
    [YamlMember(Alias = "twitter")]
    public string Twitter { get; set; }

    [JsonPropertyName("linkedin")]
    [YamlMember(Alias = "linkedin")]
    public string LinkedIn { get; set; }

    [JsonPropertyName("youtube")]
    [YamlMember(Alias = "youtube")]
    public string YouTube { get; set; }

    [JsonPropertyName("facebook")]
    [YamlMember(Alias = "facebook")]
    public string Facebook { get; set; }

    [JsonPropertyName("stackoverflow")]
    [YamlMember(Alias = "stackoverflow")]
    public string StackOverflow { get; set; }

    [JsonPropertyName("mvp")]
    [YamlMember(Alias = "mvp")]
    public string MVP { get; set; }

    [JsonPropertyName("bluesky")]
    [YamlMember(Alias = "bluesky")]
    public string Bluesky { get; set; }

    [JsonPropertyName("mastodon")]
    [YamlMember(Alias = "mastodon")]
    public string Mastodon { get; set; }

    [JsonPropertyName("sponsor_paypal")]
    [YamlMember(Alias = "sponsor_paypal")]
    public string SponsorPayPal { get; set; }

    [JsonPropertyName("sponsor_github")]
    [YamlMember(Alias = "sponsor_github")]
    public string SponsorGitHub { get; set; }
}
