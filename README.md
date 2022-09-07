# Person API

This API allows for person-related operations, namely:

- CRUD operations
- retrieving the name of syllables for fornavn, etternavn, and fornavn + etternavn
- checking if fornavn, etternavn, or fornavn + etternavn are a palindrome


## Running

To run the application, you will need Docker and Docker-compose. From the root folder of the project, run:

```
$ sudo docker-compose build
$ sudo docker-compose up
```
N.B.: You may not use `sudo` if you have configured your user to be part of the `docker` group.

The API will be available at `localhost:8080`


## Tests

To run unit tests (assuming you are in the project root folder):

```
$ cd server && go test
```


## Endpoints
### CRUD Person
#### GET `localhost:8080/person/:brukerid`

Send `brukerid`, receive person

```
{
  brukerid: int
  fornavn: string
  etternavn: string
}
```


#### POST `localhost:8080/person`

Send person and receive `brukerid` (int)

```
{
  fornavn: string
  etternavn: string
}
```


#### PUT `localhost:8080/person`

Send new person (`fornavn` and `etternavn` may be different from original), receive 1

```
{
  brukerid: int
  fornavn: string
  etternavn: string
}
```


#### DELETE `localhost:8080/person/:brukerid`

Send `brukerid`, receive 1


### Palindrome checker
#### GET `localhost:8080/palindrom/:brukerid`

Send `brukerid`, receive whether full name is palindrome (bool)

Spaces and apostrophes are dismissed. "Mo M" and "M'am" count as palindromes.


#### GET `localhost:8080/palindrom/:brukerid/fornavn`

Send `brukerid`, receive whether `fornavn` is palindrome (bool)


#### GET `localhost:8080/palindrom/:brukerid/etternavn`

Send `brukerid`, receive whether `etternavn` is palindrome (bool)


### Syllable Counter
#### GET `localhost:8080/stavelser/:brukerid`

Send `brukerid`, receive number of syllables in full name (int)


#### GET `localhost:8080/stavelser/:brukerid/fornavn`

Send `brukerid`, receive number of syllables in `fornavn` (int)


#### GET `localhost:8080/stavelser/:brukerid/etternavn`

Send `brukerid`, receive number of syllables in `etternavn` (int)
