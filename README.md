# gfn

Generate fantasy names for games and stories. It uses the fine
[fantasyname module](https://github.com/s0rg/fantasyname) by
[s0rg](https://github.com/s0rg/), which implements the code created by
the [rinkworks fantasy name
generator](http://rinkworks.com/namegen/). The code itself is
[outlined here](http://rinkworks.com/namegen/instr.shtml), or take a
quick look at the [reference
guide](http://rinkworks.com/namegen/reference.shtml).

In case the site vanishes some day, a copy of those documents is
contained here in the repository.

# Install

Execute

```shell
% go build
% cp gfn $HOME/bin
```

# Usage

There are a bunch of pre compiled fantasy name codes builtin, you can
get a list with: 

```shell
% gfn -l
```

To use one of them and limit the number of
words generated:

```shell
% gfn JapaneseNamesDiverse -c 24
yumufuchi  afuchin    keyu       amorekin   ekimuwo    ashihewani rosa       chireki
oterun     ruwahi     uwamine    emiyumu    temimon    yuwa       awayason   fuki
emiwa      nushiron   achihora   yomichi    saniyutan  kewaritsu  saroru     uhashi 
```

You can also write a code yourself:

```shell
gfn '!sVm' -c 24
Quaoobunker    Emeemoopsie    Angeepookie    Osousmoosh     Umuisweetie    Ustoesnookum   Sulealover     Imopookie
Skelaesnoogle  Echiapookie    Cereepoochie   Gariwuddle     Echaewookie    Tiaieschmoopie Queaubooble    Athesmoosh 
Undousnuggy    Urnuigooble    Mosoesnuggy    Eldoegooble    Denoipoochie   Mosoosmooch    Shyucuddly     Tiaeylovey
```

A short outline of the code will be printed if you add the `-h`
parameter:
```shell
This is gfn, a fantasy name generator cli.

Usage: gfn [-vlc] <name|code>

Options:
-c --count    How many names to generate
-l --list     List pre-compiled shortcuts
-v --version  Show program version

pattern  syntax The  letters s,  v, V,  c, B,  C, i,  m, M,  D, and  d
represent different types of random replacements:

s - generic syllable
v - vowel
V - vowel or vowel combination
c - consonant
B - consonant or consonant combination suitable for beginning a word
C - consonant or consonant combination suitable anywhere in a word
i - insult
m - mushy name
M - mushy name ending
D - consonant suited for a stupid person's name
d - syllable suited for a stupid person's name (begins with a vowel)
Everything else is emitted literally.

All  characters  between parenthesis  ()  are  emitted literally.  For
example, the pattern s(dim), emits  a random generic syllable followed
by dim.

Characters  between angle  brackets <>  emit patterns  from the  table
above. Imagine the entire pattern is wrapped in one of these.

In  both  types of  groupings,  a  vertical  bar  | denotes  a  random
choice. Empty groups are allowed.  For example, (foo|bar) emits either
foo or bar. The pattern <c|v|>  emits a constant, vowel, or nothing at
all.

An exclamation point ! means  to capitalize the component that follows
it. For  example, !(foo) will emit  Foo and v!s will  emit a lowercase
vowel followed by a capitalized syllable, like eRod.
```

# Report bugs

[Please open an issue](https://github.com/TLINDEN/gfn/issues). Thanks!

# License

This work is licensed under the terms of the General Public Licens
version 3.

# Author

Copyleft (c) 2024 Thomas von Dein
