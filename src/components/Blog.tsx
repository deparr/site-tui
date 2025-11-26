import { createSignal, For, Match, Show, Switch } from "solid-js";
import { tui, colors, syntaxColors } from "../lib/theme";
import {
  type BulletList,
  type Delete,
  parse, type Heading, type Inline, type Para, type CodeBlock, type OrderedList, type RawBlock, type Str, type Verbatim, type Block, type Section, type Emph, type SmartPunctuation, type Link, type DoubleQuoted,
  type Strong
} from "@djot/djot";
import { TextAttributes } from "@opentui/core";

interface PostMetadata {
  title?: string;
  description?: string;
  date?: string;
  tags?: string[];
};

type Size = {
  width: number;
  height: number;
};

const blogPageStr = `
\`\`\`=meta
title = Building Static Sites with Neovim
description = One step closer to becoming emacs
alt_title = Did you know Neovim could do THIS!?
date = April 9, 2025
tags = blog-post
template = post,base
\`\`\`

I enjoy having a personal website.

I don't enjoy Javascript front-end frameworks, especially for a site that's mostly static.
I previously used Sveltekit for this site, and while it was nice, I've been wanting to move to something simpler.

A couple of blogs I follow[^djot-blogs] have mentioned they use custom site
generators built around [Djot](https://djot.net), which got me interested in creating my own.
At first I wanted to write my own Djot parser, to make highlighting code
blocks easier (and better), but I dropped it a few hours in after it became unfun.

[^djot-blogs]: The wonderful blogs of [matklad](https://matklad.github.io) and [natecraddock](https://nathancraddock.com). Check them out!

A few weeks later, while reading about the new \`:TOhtml\` in nvim 0.10 news, I had a {-terrible-} fantastic idea: _turning Neovim into a site generator_.

- I want a site with highlighted code blocks
- Djot has a lua implementation
- Neovim has treesitter and \`:TOhtml\`

What's to stop me from wiring \`djot\` up to a janky neovim plugin?

## The Plugin

Enter [jolt.nvim](https://github.com/deparr/jolt.nvim).

The core plugin actions are exposed as vim user commands, which means they
can be run from the command line as well as interactively.

From the command line:

\`\`\`bash
$ nvim --headless +Jolt
$ nvim --headless "+Jolt watch"
$ nvim --headless "+Jolt clean"
\`\`\`

or inside a Neovim session:

\`\`\`
:Jolt
:Jolt watch
:Jolt clean
\`\`\`

### The Build "System"

It's relatively straight forward:

1. a user specified content directory is scanned:

    - \`.dj\` files are pages to be rendered with \`djot\`
    - \`.html\` files are templates to be filled with page content
    - any other file is static content to be copied to the output directory

2. in a filtering pass, code blocks are highlighted (more on this below)

3. results are written to the output directory

4. results are cached, so future, single-file updates are quick to rebuild

Djot does most of the real work. My code mainly does highlighting and book keeping.

### Highlighting

My options for highlighting were [\`vim.tohtml\`](https://neovim.io/doc/user/lua.html#vim.tohtml) and treesitter.
Figuring \`vim.tohtml\` would be easier, I tried it first, extracting the style sheet and code block from the resulting HTML document.
As you could probably guess, this was pretty janky.

For starters, \`vim.tohtml\` only operates on valid *winids*; which means I have to keep a window around to use for highlighting.
This is fine in headless usage, but opening and closing tons of windows sometimes resulted in ui layout shifts;
kind of annoying when you're trying to work.

Layout issues could probably be mitigated with better window management, but there were also issues with the extracted content.
I was being lazy and taking each line between \`<code></code>\` tags as a line of code, which wasn't always the case.
I'd have to properly parse the HTML to get what I want out of it, and at that point I might as well just generate the HTML myself.

That and wanting more control over the highlight style sheet pushed me to treesitter.

Neovim's treesitter interface is quite nice actually.
In just a few lines you can get an iterator over a query set's captures:

\`\`\`lua
local parser = vim.treesitter.get_string_parser(source, lang, {})
local queries = vim.treesitter.query.get(lang, "highlights")
local trees = parser:parse()
local root = trees[1]:root() -- returns a list in the case there's injections
for id, node in queries:iter_captures(root, lang) do
    -- ... do highlight stuff
end
\`\`\`

Using treesitter is much better than \`tohtml\`.
I can reliably generate the HTML I want, and I have more control over which highlight groups end up in the final style sheet.
However, it's not without its issues, as multi-line and nested captures can be tricky[^bad-hl].

Most multi-line nodes are easy to handle (e.g. multi-line strings), but some (e.g. rust doc comments) are not.
The annoying thing is this is parser dependent; rust doc comment nodes always span multiple lines, even when they're only a single line! 

\`\`\`rust
/// computes euclidean distance between two entities
pub fn dist_to(a: &Entity, b: &Entity) -> f32 {
}
\`\`\`
{.note}
[pic of bad hl if above is fixed](/images/rust_doc_bad_hl.png)

whereas the zig parser gives the proper ranges

\`\`\`zig
/// computes euclidean distance between two entities
pub fn dist_to(a: *const Entity, b: *const Entity) f32 {
}
\`\`\`

Properly handling nested captures also makes the highlighting code more complex for, honestly, very little gain.
Having nice highlights for string escapes, shell substitutions, and complex type defs is quite nice in your editor.
Though in my opinion, they're not as necessary in blog-style code blocks where the code context window is small.

\`\`\`bash
target=$HOME/.local/bin # no nesting
target="$HOME/.local/bin" # nested in @string

path=\${path##*/} # no nesting
path="\${path##*/}" # nested in @string
\`\`\`

Plus I like the lighter syntax highlighting that skipping the nested nodes gives, so I've opted to skip them for now.

For longer examples, see [here](/blog/hl-test).

[^bad-hl]: any inaccurate highlighting is result of my own half-baked query parsing lmao

### Watch Mode

Vite's dev mode with file watching and hot reloading is _really nice_; I like it a lot.
I wanted something similar for jolt.nvim.
Sure I could use something like [watchexec](https://github.com/watchexec/watchexec) to rerun \`nvim --headless +Jolt\`.
But where's the fun in that? If we're going to {-ab-}use Neovim this way, we might as well go all in!

Luckily Neovim exposes libuv to lua land, making this quite easy.

We just need to listen on a \`uv.fs_event\`, and, after filtering and debouncing for real changes, send
the list of changed files to the builder for re-rendering.
Since the builder caches the rendered pages, template changes are also supported!
Changing a template in watch mode will trigger a re-render for all pages that use it.

## Closing Thoughts

Even though it started out as a meme idea: _"What if I used nvim to build my website"_, I'm actually really enjoying the workflow I've got setup.

    * All the tooling is inside my editor, so I can easily create keybinds for custom actions.
    * The code highlighting is accurate because it uses treesitter[^bad-hl] and not a bunch of regexes.
    * It's cool to be able to highlight code on my site _exactly_ the same as my editor.
    * I don't have to deal with arcane templating languages; everything is either Djot or HTML.
    * And it just feels good to make and use my own tools.

*It turns out making things for fun is, well, fun!*

Source code for [this website](https://github.com/deparr/site) and [jolt.nvim](https://github.com/deparr/jolt.nvim).

Thanks for reading!

`;

const superscript = ["⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"]
function toSuperscript(n: number): string {
  const digits = [];
  while (n > 0) {
    digits.push(superscript[n % 10]);
    n = (n / 10) | 0;
  }

  return digits.reverse().join("");
}

// todo might not need this to be a signal
const [getFootnotes, setFootnotes] = createSignal<string[]>([]);

const punctuation = {
  left_single_quote: "'",
  right_single_quote: "'",
  left_double_quote: '"',
  right_double_quote: '"',
  ellipses: "…",
  em_dash: "—",
  en_dash: "–",
};
function InlineNode(props: { inline: Inline }) {
  const { inline } = props;
  let fnRefCount = 0;
  if (inline.tag == "footnote_reference") {
    const footnotes = getFootnotes();
    const idx = footnotes.findIndex((fn) => fn === inline.text);
    if (idx === -1) {
      fnRefCount = footnotes.push(inline.text);
      // setFootnotes(footnotes)
    } else {
      fnRefCount = idx + 1;
    }
  }
  return (
    <Switch>
      <Match when={inline.tag === "str"} >
        {(inline as Str).text}
      </Match>
      <Match when={inline.tag === "verbatim"} >
        <span style={{ bg: "#323232", fg: syntaxColors.dark.green, bold: true }}
        >{" " + (inline as Verbatim).text + " "}</span>
      </Match>
      <Match when={inline.tag === "link"}>
        {/* todo: these aren't working, dont feel like debugging them */}
        <span style={{ underline: true }}><InlineNodes inline={(inline as Link).children} /></span>
      </Match>
      <Match when={inline.tag === "smart_punctuation"}>
        {punctuation[(inline as SmartPunctuation).type]}
      </Match>
      <Match when={inline.tag === "hard_break"}>
        {"\n"}
      </Match>
      <Match when={inline.tag === "soft_break"}>
        {" "}
      </Match>
      <Match when={inline.tag === "emph"}>
        <em><InlineNodes inline={(inline as Emph).children} /></em>
      </Match>
      <Match when={inline.tag === "strong"}>
        <strong><InlineNodes inline={(inline as Strong).children} /></strong>
      </Match>
      <Match when={inline.tag === "delete"}>
        <span style={{ strikethrough: true }}
        ><InlineNodes inline={(inline as Delete).children} /></span>
      </Match>
      <Match when={inline.tag === "double_quoted"}>
        {'"'}<InlineNodes inline={(inline as DoubleQuoted).children} />{'"'}
      </Match>
      <Match when={inline.tag === "single_quoted"}>
        {"'"}<InlineNodes inline={(inline as DoubleQuoted).children} />{"'"}
      </Match>
      <Match when={inline.tag === "footnote_reference"}>
        {toSuperscript(fnRefCount)}
      </Match>
      <Match when={true}>
        <span style={{ bg: "#ff0000" }}>{inline.tag}</span>
      </Match>
    </Switch>
  );
}

function InlineNodes(props: {
  inline: Inline[]
}) {
  return (
    <For each={props.inline}>{(s) => (
      <InlineNode inline={s} />
    )}
    </For>
  );

}

function BlockNodes(props: { blocks: Block[], dim: Size }) {
  return (
    <For each={props.blocks} fallback={<text>BAD BLOCKS</text>}>{(b) => (
      <BlockNode block={b} dim={props.dim} />
    )}
    </For>
  );
}

function Hr(props: { width: number }) {
  return (<text fg={colors.dark.border} marginTop={1} alignSelf="center"
  >{"─".repeat((props.width * 0.86) | 0)}</text>);
}

function ListNode(props: {
  list: BulletList | OrderedList,
  kind: "-" | "1",
  listNest: boolean,
}) {
  const { list, kind, listNest: nested } = props;
  return (
    <box marginTop={1} marginBottom={nested ? 1 : 0} paddingLeft={2} paddingRight={2}>
      <For each={list.children}>{(item, i) => (
        <For each={item.children}>{(listBlock) => {
          return (<Switch>
            <Match when={listBlock.tag === "para"}>
              <text>{kind == "1" ? `${i() + ((list as OrderedList).start ?? 1)}. ` : '- '}
                <InlineNodes inline={(listBlock as Para).children} />
              </text>
            </Match>
            <Match when={listBlock.tag === "bullet_list" || listBlock.tag === "ordered_list"} >
              <ListNode
                list={listBlock as BulletList | OrderedList}
                kind={listBlock.tag === "bullet_list" ? "-" : "1"}
                listNest />
            </Match>
          </Switch>);
        }}
        </For>
      )}
      </For>
    </box>);
}

function BlockNode(props: { block: Block, dim: Size }) {
  const { block, dim } = props;
  return (
    <Switch>
      <Match when={block.tag === "section"}>
        <box marginTop={2}><BlockNodes blocks={(block as Section).children} dim={dim} /></box>
      </Match>
      <Match when={block.tag === "heading"}>
        <text fg={colors.dark.accent} attributes={TextAttributes.BOLD} marginBottom={1}
        >{"▎▋█ "}
          <InlineNodes inline={(block as Heading).children} />
        </text>
      </Match>
      <Match when={block.tag === "para"}>
        <text marginTop={1}>
          <InlineNodes inline={(block as Para).children} />
        </text>
      </Match>
      <Match when={block.tag === "thematic_break"}>
        <Hr width={dim.width} />
      </Match>
      <Match when={block.tag === "code_block"}>
        <box
          marginTop={2}
          marginBottom={1}
          paddingTop={1}
          paddingLeft={2}
          paddingRight={2}
          backgroundColor="#323232">
          <code
            fg={syntaxColors.dark.white}
            content={(block as CodeBlock).text}
            filetype={(block as CodeBlock).lang}
            syntaxStyle={tui.syntax.dark}
          />
        </box>
      </Match>
      <Match when={block.tag === "bullet_list"}>
        <ListNode list={block as BulletList} kind="-" listNest={false} />
      </Match>
      <Match when={block.tag === "ordered_list"}>
        <ListNode list={block as OrderedList} kind="1" listNest={false} />
      </Match>
      <Match when={block.tag === "raw_block"}>
        <Show when={(block as RawBlock).format === "tui"}>
          <text fg="#000000" bg="#00ff00">{(block as RawBlock).text}</text>
        </Show>
      </Match>
      <Match when={true}>
        <box><text bg="#0000ff">{block.tag}</text></box>
      </Match>
    </Switch >
  );
}

export default function Blog(props: { width: number, height: number }) {
  const { width, height } = props;
  const dim: Size = { width, height };
  const ast = parse(blogPageStr);
  return (
    <box paddingTop={2} paddingBottom={1} paddingLeft={3} paddingRight={3}>
      <box rowGap={1}>
        <ascii_font color={syntaxColors.dark.off_white} text="Building Static Sites" font="tiny" />
        <ascii_font color={syntaxColors.dark.off_white} text="with Neovim" font="tiny" />
        <text fg={colors.dark.dim} attributes={TextAttributes.ITALIC}
        >April 9th 2025</text>
      </box>
      <BlockNodes blocks={ast.children} dim={dim} />
      <Hr width={dim.width} />
      {/* just assume fns only contain a single para */}
      <For each={getFootnotes()}>{(fn, idx) => {
        const para = ast.footnotes[fn]?.children[0] as Para;
        const fnRef = idx() + 1;
        return (
          <text id={`fn${fnRef}`} marginTop={1}>{`${fnRef}. `}<InlineNodes inline={para.children} /></text>
        );
      }}
      </For>
    </box>
  );
}
