#!/usr/bin/env python3

import sys
import json
import yaml
from rich.console import Console
from rich.panel import Panel
from rich.text import Text

def main():
    """Display the intro card."""
    # Check for flags
    is_help_mode = '--help' in sys.argv or '-h' in sys.argv
    is_json_mode = '--json' in sys.argv
    is_yaml_mode = '--yaml' in sys.argv
    is_vcard_mode = '--vcard' in sys.argv

    # Profile data
    profile = {
        'name': 'Niels Swimberghe',
        'nickname': 'Swimburger',
        'title': 'C# and TypeScript SDK generator engineer, at Fern',
        'location': 'NYC',
        'timezone': 'Eastern Time',
        'bio': 'Niels Swimberghe is a Belgian American software engineer and Microsoft MVP at Fern where he owns the C# and TypeScript SDK generators.\nGet in touch with Niels on Twitter @RealSwimburger and follow Niels\' personal blog on .NET, Azure, and web development at swimburger.net.',
        'website': 'https://swimburger.net',
        'github': 'https://github.com/Swimburger',
        'twitter': 'https://twitter.com/RealSwimburger',
        'linkedin': 'https://linkedin.com/in/nielsswimberghe',
        'youtube': 'https://youtube.com/@RealSwimburger',
        'facebook': 'https://facebook.com/SwimburgerDotNet',
        'stackoverflow': 'https://stackoverflow.com/users/2919731/swimburger',
        'mvp': 'https://mvp.microsoft.com (Microsoft MVP)',
        'bluesky': 'https://bsky.app/profile/swimburger.bsky.social',
        'mastodon': '@swimburger@dotnet.social',
        'sponsor_paypal': 'https://www.paypal.com/paypalme/swimburger',
        'sponsor_github': 'https://github.com/sponsors/Swimburger'
    }

    # If --help flag is provided, display usage and exit
    if is_help_mode:
        print("""
\033[1m\033[38;5;161mswimburger\033[0m - Personal intro card for Niels Swimberghe (Swimburger)

\033[1mUSAGE:\033[0m
  swimburger [OPTIONS]

\033[1mOPTIONS:\033[0m
  \033[36m--help, -h\033[0m    Show this help message
  \033[36m--json\033[0m        Output profile data as JSON
  \033[36m--yaml\033[0m        Output profile data as YAML
  \033[36m--vcard\033[0m       Output profile data as vCard (VCF format)

\033[1mEXAMPLES:\033[0m
  swimburger              Display visual profile card
  swimburger --json       Output as JSON
  swimburger --yaml       Output as YAML
  swimburger --vcard      Output as vCard

\033[1mMORE INFO:\033[0m
  Website: https://swimburger.net
  GitHub:  https://github.com/Swimburger
""")
        return

    # If --json flag is provided, output JSON and exit
    if is_json_mode:
        print(json.dumps(profile, indent=2))
        return

    # If --yaml flag is provided, output YAML and exit
    if is_yaml_mode:
        print(yaml.dump(profile, default_flow_style=False, sort_keys=False))
        return

    # If --vcard flag is provided, output vCard and exit
    if is_vcard_mode:
        vcard = f"""BEGIN:VCARD
VERSION:3.0
FN:{profile['name']}
NICKNAME:{profile['nickname']}
TITLE:{profile['title']}
NOTE:{profile['bio'].replace(chr(10), '\\n')}
URL:{profile['website']}
URL;type=github:{profile['github']}
URL;type=twitter:{profile['twitter']}
URL;type=linkedin:{profile['linkedin']}
URL;type=youtube:{profile['youtube']}
URL;type=facebook:{profile['facebook']}
URL;type=stackoverflow:{profile['stackoverflow']}
URL;type=mvp:{profile['mvp']}
URL;type=bluesky:{profile['bluesky']}
X-SOCIALPROFILE;type=mastodon:{profile['mastodon']}
URL;type=sponsor:{profile['sponsor_paypal']}
URL;type=sponsor:{profile['sponsor_github']}
ADR;TYPE=WORK:;;{profile['location']};;;{profile['timezone']};
END:VCARD"""
        print(vcard)
        return

    # Build the card content
    text = Text()
    text.append(f"🍔 {profile['name']} ({profile['nickname']})\n", style="bold #a02c56")
    text.append(f"🌎 {profile['location']} | {profile['timezone']}\n\n", style="dim")
    text.append(f"{profile['bio']}\n\n", style="italic")
    text.append("━━━ Connect ━━━\n", style="bold #a02c56")
    text.append("🌐 ", style="")
    text.append("Website:        ", style="bold")
    text.append(f"{profile['website']}\n", style=f"dim link {profile['website']}")
    text.append("💻 ", style="")
    text.append("GitHub:         ", style="bold")
    text.append(f"{profile['github']}\n", style=f"dim link {profile['github']}")
    text.append("🐦 ", style="")
    text.append("Twitter:        ", style="bold")
    text.append(f"{profile['twitter']}\n", style=f"dim link {profile['twitter']}")
    text.append("💼 ", style="")
    text.append("LinkedIn:       ", style="bold")
    text.append(f"{profile['linkedin']}\n", style=f"dim link {profile['linkedin']}")
    text.append("🎥 ", style="")
    text.append("YouTube:        ", style="bold")
    text.append(f"{profile['youtube']}\n", style=f"dim link {profile['youtube']}")
    text.append("👥 ", style="")
    text.append("Facebook:       ", style="bold")
    text.append(f"{profile['facebook']}\n", style=f"dim link {profile['facebook']}")
    text.append("💬 ", style="")
    text.append("Stack Overflow: ", style="bold")
    text.append(f"{profile['stackoverflow']}\n", style=f"dim link {profile['stackoverflow']}")
    text.append("🦋 ", style="")
    text.append("Bluesky:        ", style="bold")
    text.append(f"{profile['bluesky']}\n", style=f"dim link {profile['bluesky']}")
    text.append("🐘 ", style="")
    text.append("Mastodon:       ", style="bold")
    text.append(f"{profile['mastodon']}\n\n", style="dim")
    text.append("━━━ Support ━━━\n", style="bold #a02c56")
    text.append("💖 ", style="")
    text.append("PayPal:         ", style="bold")
    text.append(f"{profile['sponsor_paypal']}\n", style=f"dim link {profile['sponsor_paypal']}")
    text.append("🎁 ", style="")
    text.append("GitHub Sponsor: ", style="bold")
    text.append(profile['sponsor_github'], style=f"dim link {profile['sponsor_github']}")

    # Create and display the panel
    console = Console()

    panel = Panel(
        text,
        border_style="#a02c56",
        padding=(1, 2)
    )

    console.print()
    console.print(panel)
    console.print()

if __name__ == '__main__':
    main()
