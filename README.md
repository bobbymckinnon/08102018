
# Patient Data

You have been asked to assist in the creation of an internal API and Search page for a national healthcare provider. This provider has a set of inpatient prospective payment systems providers.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/bobbymckinnon/08102018
$ cd $GOPATH/src/github.com/bobbymckinnon/08102018
$ heroku local
```
You should also install [govendor](https://github.com/kardianos/govendor) if you are going to add any dependencies to this app.

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Endpoint Examples
Available URL parameters (all are optional, however at least one is required)

```sh
https://quiet-refuge-38556.herokuapp.com/providers?state=AL
https://quiet-refuge-38556.herokuapp.com/providers?min_discharges=50&max_discharges=75
```

 - state
 - max_discharges
 - min_discharges
 - max_average_covered_charges
 - min_average_covered_charges
 - max_average_medicare_payments
 - min_average_medicare_payments


## Client Search Interface
```sh
https://quiet-refuge-38556.herokuapp.com/
```

## CLI Testing
The following command runs go's internal testing framework for the API, and jest/enzyme tests for the client code.
```sh
make test
```

```sh
go test ./...
npm test
```

## Providers Data Import
The following endpoint can be used to trigger an import from the providers.csv file into the providers table.
(Please be aware that running the importer over the API endpoint will take time)

```sh
https://quiet-refuge-38556.herokuapp.com/importProviders?importProviders=yes
```
create index on providers (upper(state));

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
