# Sententia

Generate random adjective-noun combinations, madlibs-style, in Golang. Inspired by
[kylestetz](https://github.com/kylestetz)'s [Sentencer](https://github.com/kylestetz/Sentencer)
(node.js module).

```golang
"Oh look, {{ an_adjective }} {{ noun }}!"
```

becomes something like (it's random!):

```golang
"Oh look, a tameless recorder!"
```

## Actions

Built in actions are:

`{{ noun }}`, `{{ a_noun }}`, `{{ nouns }}`, `{{ adjective }}` and `{{ an_adjective }}`.