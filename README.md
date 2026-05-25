# invisible-characters

A static website about [invisible Unicode characters](https://invisible-characters.com/).

## What Visitors Can Do

Invisible Characters helps users find, understand, and copy Unicode characters that are hard or impossible to see.

Typical visitor goals:

1. Find a specific invisible character and copy it.
2. Learn what a character does and where it is used.
3. Use practical guides for apps like TikTok, WhatsApp, Instagram, or X/Twitter.
4. Decode text to reveal hidden or non-printable characters.

## Page Structure (End User Perspective)

### 1. Home page (`/`)

The home page is the starting point for most users. It contains:

- A short introduction to invisible Unicode characters.
- Practical examples (empty messages, invisible usernames, etc.).
- A quick link to the decoder tool.
- A browsable list of individual characters.
- Links to larger Unicode character groups.

For visitors, this page answers: "What is this?" and "Where should I click next?"

### 2. Character detail pages (`/<CODE>-<NAME>.html`)

Each character has its own page, for example `U+3164 HANGUL FILLER`.

These pages provide:

- Character name and Unicode codepoint.
- Optional group information.
- Optional description with links to related characters.
- A copy button to place the character into the clipboard.

For visitors, this is the main "copy and use" page.

### 3. Group pages (`/block-tags.html`, `/block-variation-selectors.html`)

Group pages collect related characters in one place.

They provide:

- A short explanation of the Unicode block/group.
- A list of all characters in that group.
- Copy buttons for each listed character.

For visitors, these pages are useful when exploring families of similar characters.

### 4. Decoder tool page (`/view.html`)

The decoder lets users paste text and reveal hidden characters.

It shows:

- Annotated output where invisible characters are highlighted.
- Unicode codepoints and names for detected characters.
- Basic per-character statistics.

For visitors, this page answers: "What hidden characters are inside this text?"

### 5. Practical guide pages

These pages are task-focused and give quick copy-and-paste instructions:

- `/empty-tweet.html`
- `/empty-whatsapp.html`
- `/invisible-tiktok-name.html`
- `/empty-instagram-comment.html`

For visitors, these are "how-to" pages with immediate action.

### 6. Legal and fallback pages

- `/legal.html`: Privacy policy and legal information.
- `/404.html`: Not-found page for invalid URLs.

## Visitor Journey

Most users follow one of these paths:

1. Search engine -> specific character page -> copy character.
2. Home page -> practical guide page -> copy character -> use in target app.
3. Home page -> decoder tool -> inspect text for hidden characters.
4. Home page -> group page -> compare/copy related characters.

## Technical Note

All pages are statically generated from templates and character data. The output is fast, CDN-friendly, and simple to host.

