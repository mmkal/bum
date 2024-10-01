We currently don't have a streamlined process for deploying Chaos Bum. This
page describes the manual steps required to build and deploy. A great way to
contribute to this project would be to use Docker containers to make it easier
for other users to get up and running quickly.

## Prerequisites

* [Spinnaker]
* MySQL (5.6 or later)

To use this version of Chaos Bum, you must be using [Spinnaker] to manage your applications. Spinnaker is the
continuous delivery platform that we use at Netflix.

Chaos Bum also requires a MySQL-compatible database, version 5.6 or later.

[Spinnaker]: http://www.spinnaker.io/


## Build

To build Chaos Bum on your local machine (requires the Go
toolchain).

```
go get github.com/netflix/chaosbum/cmd/chaosbum
```

This will install a `chaosbum` binary in your `$GOBIN` directory.

## How Chaos Bum runs

Chaos Bum does not run as a service. Instead, you set up a cron job
that calls Chaos Bum once a weekday to create a schedule of terminations.

When Chaos Bum creates a schedule, it creates another cron job to schedule terminations
during the working hours of the day.

## Deploy overview

To deploy Chaos Bum, you need to:

1. Configure Spinnaker for Chaos Bum support
1. Set up the MySQL database
1. Write a configuration file (chaosbum.toml)
1. Set up a cron job that runs Chaos Bum daily schedule

## Configure Spinnaker for Chaos Bum support

Spinnaker's web interface is called *Deck*. You need to be running Deck version
v.2839.0 or greater for Chaos Bum support. Check which version of Deck you are
running by hitting the `/version.json` endpoint of your Spinnaker deployment.
(Note that this version information will not be present if you are running
Deck using a [Docker container hosted on Quay][quay]).

[quay]: https://quay.io/repository/spinnaker/deck

Deck has a config file named `/var/www/settings.js`. In this file there is a
"feature" object that contains a number of feature flags:

```
  feature: {
    pipelines: true,
    notifications: false,
    fastProperty: true,
    ...
```

Add the following flag:

```
chaosBum: true
```

If the feature was enabled successfully, when you create a new app with Spinnaker, you will see
a "Chaos Bum: Enabled" checkbox in the "New Application" modal dialog. If it
does not appear, you may need to deploy a more recent version of Spinnaker.

![new-app](new-app.png "new application dialog")

For more details, see [Additional configuration files][spinconfig] on the
Spinnaker website.

[spinconfig]: http://www.spinnaker.io/docs/custom-configuration#section-additional-configuration-files



## Create the MySQL database

Chaos Bum uses a MySQL database as a backend to record a daily termination
schedule and to enforce a minimum time between terminations. (By default, Chaos
Bum will not terminate more than one instance per day per group).

Log in to your MySQL deployment and create a database named `chaosbum`:

```
mysql> CREATE DATABASE chaosbum;
```

Note: Chaos Bum does not currently include a mechanism for purging old data.
Until this function exists, it is the operator's responsibility to remove old
data as needed.

## Write a configuration file (chaosbum.toml)

See [Configuration file format](Configuration-file-format) for the configuration file format.

## Create the database schema

Once you have created a `chaosbum` database and have populated the
configuration file with the database credentials, add the tables to the database
by doing:

```
chaosbum migrate
```


### Verifying Chaos Bum is configured properly

Chaos Bum supports a number of command-line arguments that are useful for
verifying that things are working properly.

#### Spinnaker

You can verify that Chaos Bum can reach Spinnaker by fetching the Chaos Bum
configuration for an app:

```
chaosbum config <appname>
```

If successful, you'll see output that looks like:

```
(*chaosbum.AppConfig)(0xc4202ec0c0)({
 Enabled: (bool) true,
 RegionsAreIndependent: (bool) true,
 MeanTimeBetweenKillsInWorkDays: (int) 2,
 MinTimeBetweenKillsInWorkDays: (int) 1,
 Grouping: (chaosbum.Group) cluster,
 Exceptions: ([]chaosbum.Exception) {
 }
})
```

If it fails, you'll see an error message.

#### Database

You can verify that Chaos Bum can reach the database by attempting to
retrieve the termination schedule for the day.

```
chaosbum fetch-schedule
```

If successful, you should see output like:

```
[69400] 2016/09/30 23:41:03 chaosbum fetch-schedule starting
[69400] 2016/09/30 23:41:03 Writing /etc/cron.d/chaosbum-daily-terminations
[69400] 2016/09/30 23:41:03 chaosbum fetch-schedule done
```

(Chaos Bum will write an empty file to
`/etc/cron.d/chaosbum-daily-terminations` since the database does not contain
any termination schedules yet).

If Chaos Bum cannot reach the database, you will see an error. For example:

```
[69668] 2016/09/30 23:43:50 chaosbum fetch-schedule starting
[69668] 2016/09/30 23:43:50 FATAL: could not fetch schedule: failed to retrieve schedule for 2016-09-30 23:43:50.953795019 -0700 PDT: dial tcp 127.0.0.1:3306: getsockopt: connection refused
```

#### Generate a termination schedule

You can manually invoke Chaos Bum to generate a schedule file. When testing,
you may want to specify `--no-record-schedule` so the schedule doesn't get
written to the database.

If you have many apps and you don't want to sit there while Chaos Bum
generates a complete schedule, you can limit the number of apps  using the
`--max-apps=<number>`. For example:

```
chaosbum schedule --no-record-schedule --max-apps=10
```

#### Terminate an instance

You can manually invoke Chaos Bum to terminate an instance. For example:

```
chaosbum terminate chaosguineapig test --cluster=chaosguineapig --region=us-east-1
```


### Optional: Dynamic properties (etcd, consul)

Chaos Bum supports changing the following configuration properties dynamically:

* chaosbum.enabled
* chaosbum.leashed
* chaosbum.schedule_enabled
* chaosbum.accounts

These are intended to allow an operator to make certain changes to Chaos
Bum's behavior without having to redeploy.

Note: the configuration file takes precedence over dynamic provider, so do
not specify these properties in the config file if you want to set them
dynamically.

To take advantage of dynamic properties, you need to keep those properties in
either [etcd] or [Consul] and add a `[dynamic]` section that contains the
endpoint for the service and a path that returns a JSON file that has each of
the properties you want to set dynamically.

Chaos Bum uses the [Viper][viper] library to implement dynamic configuration, see the
Viper [remote key/value store support][remote] docs for more details.


[etcd]: https://coreos.com/etcd/docs/latest/
[consul]: https://www.consul.io/
[viper]: https://github.com/spf13/viper
[remote]: https://github.com/spf13/viper#remote-keyvalue-store-support


## Set up a cron job that runs Chaos Bum daily schedule

### Create /apps/chaosbum/chaosbum-schedule.sh

For the remainder if the docs, we assume you have copied the chaosbum binary
to `/apps/chaosbum`, and will create the scripts described below there as
well. However, Chaos Bum makes no explicit assumptions about the location of
these files.


Create a file called `chaosbum-schedule.sh` that invokes `chaosbum
schedule` and writes the output to a logfile.

Note that because this will be invoked from cron, the PATH will likely not include the
location of the chaosbum binary so be sure to specify it explicitly.

/apps/chaosbum/chaosbum-schedule.sh:
```bash
#!/bin/bash
/apps/chaosbum/chaosbum schedule >> /var/log/chaosbum-schedule.log 2>&1
```

### Create /etc/cron.d/chaosbum-schedule

Once you have this script, create a cron job that invokes it once a day. Chaos
Bum starts terminating at `chaosbum.start_hour` in
`chaosbum.time_zone`, so it's best to pick a time earlier in the day.

The example below generates termination schedules each weekday at 12:00 system
time (which we assume is in UTC).

/etc/cron.d/chaosbum-schedule:
```bash
# Run the Chaos Bum scheduler at 5AM PDT (4AM PST) every weekday
# This corresponds to: 12:00 UTC
# Because system clock runs UTC, time change affects when job runs

# The scheduler must run as root because it needs root permissions to write
# to the file /etc/cron.d/chaosbum-daily-terminations

# min  hour  dom  month  day  user  command
    0    12    *      *  1-5  root  /apps/chaosbum/chaosbum-schedule.sh
```

### Create /apps/chaosbum/chaosbum-terminate.sh

When Chaos Bum schedules terminations, it will create cron jobs that call the
path specified by `chaosbum.term_path`, which defaults to /apps/chaosbum/chaosbum-terminate.sh

/apps/chaosbum/chaosbum-terminate.sh:
```
#!/bin/bash
/apps/chaosbum/chaosbum terminate "$@" >> /var/log/chaosbum-terminate.log 2>&1
```

