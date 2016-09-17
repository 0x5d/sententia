# sententia

Generate random adjective-noun combinations, madlibs-style, in Golang. Inspired by
[kylestetz](https://github.com/kylestetz)'s [Sentencer](https://github.com/kylestetz/Sentencer)
(node.js module).

```go
"Oh look, {{ an_adjective }} {{ noun }}!"
```

becomes something like (it's random!):

```go
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

```go
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
```go
sententia.Make("{{ noun }}")
// => "avenue", "knot", "show", "narcissus"
```

**`{{ a_noun }}`**

Picks a random noun, precedes it with an article ("a" or "an"), and replaces it.
```go
sententia.Make("{{ a_noun }}")
// => "a romanian", "an archer", "a tyvek"
```

**`{{ nouns }}`**

Picks a random noun and pluralizes it.
```go
sententia.Make("{{ noun }}")
// => "harbors", "pumps", "overcoats", "gongs"
```

**`{{ adjective }}`**

Is replaced in the template by a random adjective.
```go
sententia.Make("{{ adjective }}")
// => "sprightful", "naif", "glowing", "surfy"
```

**`{{ an_adjective }}`**

Picks a random adjective and precedes it with an article.
```go
sententia.Make("{{ an_adjective }}")
// => "a sphygmic", "a clubby", "an uncocked", "a bumbling"
```

**These all** can be mixed to provide beautiful nonsensic sentences:
```go
sententia.Make(
  "Once I had {{ an_adjective }} {{ noun }} full of {{ nouns }} but they flew into {{ a_noun }}.",
)
// => "Once I had a spleeny wallet full of toilets but they flew into an orchestra."
```

## Extending sententia

**Adjectives and nouns**

sententia can be extended by adding custom nouns or adjectives, or overwriting the built-in ones.
```go
customNouns := []string{"potato", "carrot", "letuce"}
sententia.AddNouns(nouns) // Apepnds the new nouns to the existent ones.
sententia.SetNouns(nouns) // Overwrites the existent nouns.

customAdjectives := []string{"delicious", "earthy", "healthy"}
sententia.AddAdjectives(customAdjectives) // Apepnds the new adjectives to the existent ones.
sententia.SetAdjectives(customAdjectives) // Overwrites existent adjectives.
```

**Actions**

Because sententia is based on Go's `text/template` package, we can also compose actions by feeding
their output into other ones. For example, we could write an action to turn a noun into title case,
where the first letter is upper case:

```go
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
