# nob for Neocities 🐱
nob is a neocities-oriented blog manager and static site generator. It is
really easy to use and fast. It is intended to be used together with
[Neocities CLI](https://neocities.org/cli).

## Installation
Grab a binary for your operating system from the
[releases page](https://github.com/nshebang/nob/releases/)
and move it to a convenient directory (for example, `~/.local/bin` if you are
using Linux)

Finally, just run `nob` from your terminal and that's it!

## Features
* Neocities-oriented
* Cross-platform
* Simple templating system
* You can write your posts in Markdown
* Support for RSS
* No theming system. You just use any CSS stylesheet you like and works.

## How to use nob
The [wiki](https://github.com/nshebang/nob/wiki) has more information
about this. 

## Q&A 

### Is this project still in development?

As long as there is people pushing commits, it's active. As for me, the main
maintainer (Olivia), I'm often too busy with my life and just maintain this
whenever I can.

### Will you ever finish this generator?

It's finished, for now.

### Does nob support themes?

Yes, you can use any CSS stylesheet / theme you want by simply `<link>`ing
it to the HTML templates of your blog.

### Does nob support Markdown and LaTeX?

Yes, Markdown is fully supported. The Markdown parser detects LaTeX blocks too,
but you need a javascript typesetting library such as
[KaTeX](https://katex.org/) or [Mathjax](https://www.mathjax.org/) 
for those blocks to be properly rendered as nob doesn't provide a
LaTeX parser nor a stylesheet for mathematical notation.

### How can I upload my blog to Neocities?

You can either do it manually (a hassle) or use
[Neocities CLI](https://neocities.org/cli) which automatically updates your
blog entries. Remember to add the `.nob/` directory of your blog to a
`.gitignore` file to prevent Neocities CLI from uploading your Markdown
drafts and HTML templates. You shouldn't upload any files in `.nob/` if you
actually want your Markdown drafts to remain private.

### This generator is missing a feature that Hugo already has!

Hugo is a more complex program that manages web content in general, not just
blogs. The goal of nob is not trying to be another Hugo.
The goal of this project is to provide a simple and straightforward blog 
manager and generator for Neocities users who know how to use a CLI interface.

Of course some features are missing, but some are not really necessary
(such as weird yaml config files). However, if you think nob is missing an 
_essential_ feature, you can always [open a PR](https://github.com/nshebang/nob/pulls).

### I need a tagging system!

Soon.

