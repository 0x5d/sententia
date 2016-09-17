# sententia

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

sententia uses Go's `text/template` to identify the places where a new noun or adjective should be
generated.

## Usage

Download the lib.
```sh
go get github.com/castillobgr/sententia
```

```golang
package main

import (
	"fmt"

	"github.com/castillobgr/sententia"
)

func main() {
	sentence, err := sententia.Make("Aw yis, {{ adjective }} {{ nouns }}.")
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence)
}
```

## Actions

Built in actions are:

**`{{ noun }}`**

Picks a random noun and replaces it in the input.
```golang
sententia.Make("{{ noun }}")
// => "avenue", "knot", "show", "narcissus"
```

**`{{ a_noun }}`**

Picks a random noun, precedes it with an article ("a" or "an"), and replaces it.
```golang
sententia.Make("{{ a_noun }}")
// => "a romanian", "an archer", "a tyvek"
```

**`{{ nouns }}`**

Picks a random noun and pluralizes it.
```golang
sententia.Make("{{ noun }}")
// => "harbors", "pumps", "overcoats", "gongs"
```

**`{{ adjective }}`**

Is replaced in the template by a random adjective.
```golang
sententia.Make("{{ adjective }}")
// => "sprightful", "naif", "glowing", "surfy"
```

**`{{ an_adjective }}`**

Picks a random adjective and precedes it with an article.
```golang
sententia.Make("{{ an_adjective }}")
// => "a sphygmic", "a clubby", "an uncocked", "a bumbling"
```

**These all** can be mixed to provide beautiful nonsensic sentences:
```golang
sententia.Make(
  "Once I had {{ an_adjective }} {{ noun }} full of {{ nouns }} but they flew into {{ a_noun }}."
)
// => "Once I had a spleeny wallet full of toilets but they flew into an orchestra."
```

## Use cases

- Creating random IDs a la Heroku
- Generating lorem ipsums
- Fun
- ??
- Profit
