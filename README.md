# hypnos

A sleepy cron

Ever write little scripts like to do things on some interval?

```bash
#!/bin/bash
while [ 1 ]; do
  echo "running some tasks..."
  # ...
  sleep 3600
done
```

Want it to be more like cron? And start repeatedly at specific times?

```bash
#!/bin/bash
while [ 1 ]; do
  hypnos "0 * * * *"
  echo "running some tasks..."
  # ...
done
```
```
Next run at 2017-01-31 19:00:00 -0800 PST Sleeping for 33m25.234920265s
```

## Installation

Grab a tarball from [Releases](https://github.com/transitorykris/hypnos/releases)
and install in your path.

Or build from source

```
go get github.com/transitorykris/hypnos
```

## Examples

Hypnos follows [cron](https://en.wikipedia.org/wiki/Cron)

Run once per minute

```
hypnos "* * * * *"
```

Run once per week

```
hypnos @weekly
```

Run on Fridays at 5:30pm

```
hypnos "30 17 * * FRI"
```
