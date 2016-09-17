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


## Use cases

- Creating random IDs a la Heroku
- Generating lorem ipsums
- Fun
- ??
- Profit


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
  "Once I had {{ an_adjective }} {{ noun }} full of {{ nouns }} but they flew into {{ a_noun }}.",
)
// => "Once I had a spleeny wallet full of toilets but they flew into an orchestra."
```

## Extending actions

sententia can be extended by adding custom actions. Let's say we want to append 'er' to every noun:
We could write something like
```golang
package main

import (
	"fmt"
	"math/rand"

	"github.com/castillobg/sententia"
)

func main() {
	customActions := map[string]interface{}{
		"rating": func() string {
			ratings := []string{"OK", "bad", "meh", "super"}
			return ratings[rand.Intn(len(ratings))]
		},
	}
	sententia.AddActions(customActions)
	sentence, err := sententia.Make(
		"Dinner was {{ rating }}.",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence)
}
```
and maybe we'd get `"Dinner was super."`

Because sententia is based on Go's `text/template` package, we can also compose actions by feeding
their output into other ones. For example, we could write an action to turn a noun into title case,
where the first letter is upper case:

```golang
package main

import (
	"fmt"
	"strings"

	"github.com/castillobg/sententia"
)

func main() {
	customActions := map[string]interface{}{
		"capitalize": strings.Title,
	}
	sententia.AddActions(customActions)
	sentence, err := sententia.Make(
		"She wrote a book called '{{ capitalize an_adjective }} {{ capitalize noun }}'",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(sentence)
}
```
Then, we could get something like `"She wrote a book called 'A Subdued Duckling'"`.
If you want a refresher on how Go templates handle arguments,
[check this out](https://golang.org/pkg/text/template/#hdr-Arguments)!
