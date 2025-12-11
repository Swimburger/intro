#!/usr/bin/env node

import chalk from 'chalk';
import boxen from 'boxen';
import YAML from 'yaml';

// Check for flags
const isHelpMode = process.argv.includes('--help') || process.argv.includes('-h');
const isJsonMode = process.argv.includes('--json');
const isYamlMode = process.argv.includes('--yaml');
const isVcardMode = process.argv.includes('--vcard');

// Profile data
const profile = {
  name: 'Niels Swimberghe',
  nickname: 'Swimburger',
  title: 'C# and TypeScript SDK generator engineer, at Fern',
  location: 'NYC',
  timezone: 'Eastern Time',
  bio: 'Niels Swimberghe is a Belgian American software engineer and Microsoft MVP at Fern where he owns the C# and TypeScript SDK generators.\nGet in touch with Niels on Twitter @RealSwimburger and follow Niels\' personal blog on .NET, Azure, and web development at swimburger.net.',
  website: 'https://swimburger.net',
  github: 'https://github.com/Swimburger',
  twitter: 'https://twitter.com/RealSwimburger',
  linkedin: 'https://linkedin.com/in/nielsswimberghe',
  youtube: 'https://youtube.com/@RealSwimburger',
  facebook: 'https://facebook.com/SwimburgerDotNet',
  stackoverflow: 'https://stackoverflow.com/users/2919731/swimburger',
  mvp: 'https://mvp.microsoft.com (Microsoft MVP)',
  bluesky: 'https://bsky.app/profile/swimburger.bsky.social',
  mastodon: '@swimburger@dotnet.social',
  sponsor_paypal: 'https://www.paypal.com/paypalme/swimburger',
  sponsor_github: 'https://github.com/sponsors/Swimburger'
};

// If --help flag is provided, display usage and exit
if (isHelpMode) {
  console.log(`
${chalk.bold.hex('#a02c56')('swimburger')} - Personal intro card for Niels Swimberghe (Swimburger)

${chalk.bold('USAGE:')}
  swimburger [OPTIONS]

${chalk.bold('OPTIONS:')}
  ${chalk.cyan('--help, -h')}    Show this help message
  ${chalk.cyan('--json')}        Output profile data as JSON
  ${chalk.cyan('--yaml')}        Output profile data as YAML
  ${chalk.cyan('--vcard')}       Output profile data as vCard (VCF format)

${chalk.bold('EXAMPLES:')}
  swimburger              Display visual profile card
  swimburger --json       Output as JSON
  swimburger --yaml       Output as YAML
  swimburger --vcard      Output as vCard

${chalk.bold('MORE INFO:')}
  Website: https://swimburger.net
  GitHub:  https://github.com/Swimburger
`);
  process.exit(0);
}

// If --json flag is provided, output JSON and exit
if (isJsonMode) {
  console.log(JSON.stringify(profile, null, 2));
  process.exit(0);
}

// If --yaml flag is provided, output YAML and exit
if (isYamlMode) {
  console.log(YAML.stringify(profile));
  process.exit(0);
}

// If --vcard flag is provided, output vCard and exit
if (isVcardMode) {
  const vcard = `BEGIN:VCARD
VERSION:3.0
FN:${profile.name}
NICKNAME:${profile.nickname}
TITLE:${profile.title}
NOTE:${profile.bio.replace(/\n/g, '\\n')}
URL:${profile.website}
URL;type=github:${profile.github}
URL;type=twitter:${profile.twitter}
URL;type=linkedin:${profile.linkedin}
URL;type=youtube:${profile.youtube}
URL;type=facebook:${profile.facebook}
URL;type=stackoverflow:${profile.stackoverflow}
URL;type=mvp:${profile.mvp}
URL;type=bluesky:${profile.bluesky}
X-SOCIALPROFILE;type=mastodon:${profile.mastodon}
URL;type=sponsor:${profile.sponsor_paypal}
URL;type=sponsor:${profile.sponsor_github}
ADR;TYPE=WORK:;;${profile.location};;;${profile.timezone};
END:VCARD`;
  console.log(vcard);
  process.exit(0);
}

// Helper function to create clickable links using OSC 8
const makeLink = (url, text) => {
  return `\u001b]8;;${url}\u001b\\${chalk.gray(text)}\u001b]8;;\u001b\\`;
};

// Build the card content
const card = `
${chalk.bold.hex('#a02c56')(`🍔 ${profile.name} (${profile.nickname})`)}
${chalk.dim('🌎 ' + profile.location + ' | ' + profile.timezone)}

${chalk.italic(profile.bio)}

${chalk.bold.hex('#a02c56')('━━━ Connect ━━━')}
🌐 ${chalk.bold('Website:')}        ${makeLink(profile.website, profile.website)}
💻 ${chalk.bold('GitHub:')}         ${makeLink(profile.github, profile.github)}
🐦 ${chalk.bold('Twitter:')}        ${makeLink(profile.twitter, profile.twitter)}
💼 ${chalk.bold('LinkedIn:')}       ${makeLink(profile.linkedin, profile.linkedin)}
🎥 ${chalk.bold('YouTube:')}        ${makeLink(profile.youtube, profile.youtube)}
👥 ${chalk.bold('Facebook:')}       ${makeLink(profile.facebook, profile.facebook)}
💬 ${chalk.bold('Stack Overflow:')} ${makeLink(profile.stackoverflow, profile.stackoverflow)}
🦋 ${chalk.bold('Bluesky:')}        ${makeLink(profile.bluesky, profile.bluesky)}
🐘 ${chalk.bold('Mastodon:')}       ${chalk.gray(profile.mastodon)}

${chalk.bold.hex('#a02c56')('━━━ Support ━━━')}
💖 ${chalk.bold('PayPal:')}         ${makeLink(profile.sponsor_paypal, profile.sponsor_paypal)}
🎁 ${chalk.bold('GitHub Sponsor:')} ${makeLink(profile.sponsor_github, profile.sponsor_github)}
`;

// Display the card in a box
console.log(
  boxen(card.trim(), {
    padding: 1,
    margin: 1,
    borderStyle: 'round',
    borderColor: '#a02c56'
  })
);
