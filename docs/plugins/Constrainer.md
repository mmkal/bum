# Constrainer

There may be some cases where you want to prevent some combination of Chaos
Bum terminations, but the [configuration options](../Configuring-behavior-via-spinnaker) aren't flexible
enough for your use case.

You can define a custom constrainer to do this.

As an example, let's say you wanted to disallow any terminations for apps
that contain "foo" as a substring.

```go
package constrainer

import (
	"github.com/Netflix/chaosbum/deps"
	"github.com/Netflix/chaosbum/config"
	"github.com/Netflix/chaosbum/schedule"
    "strings"
)

func init() {
    deps.GetConstrainer = getConstrainer()
}

type noFoo struct {}

func getConstrainer(cfg *config.Bum) (schedule.Constrainer, error) {
    return noFoo{}, nil
}

func (n noFoo) Filter(s schedule.Schedule) schedule.Schedule {
	result := schedule.New()
	for _, entry := range s.Entries() {
        if !strings.Contains(entry.Group.App(), "foo") {
            result.Add(entry.Time, entry.Group)
        }
    }
    return result
}

```

See the [Plugins](index.md) page for info on how to build a custom version of
Chaos Bum with your plugin.
