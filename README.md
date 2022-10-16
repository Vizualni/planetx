# PlanetX invasion!

## How to use?


### Run tests

To run test you can use standard "out-of-the box" tools provided by go toolset:
```
go test ./...
```

e.g.
```sh
$ go test ./...
ok      github.com/vizualni/planetx     0.334s
?       github.com/vizualni/planetx/cmd/planetx [no test files]
```

or you can use the `make test`.


###  Build the app

```
go build -o build/planetx ./cmd/planetx
```

or `make build` to build it.


### Running the application

```
$ go run ./cmd/planetx simulate-invasion --path a.planetx --aliens 2 --seed 1
City B destroyed by aliens Alien 2 and Alien 1 fighting.
========================
New planet looks like:
A
C east=A
```

where `a.planetx` is:

```
A north=B
B north=C
C south=B east=A
```

## Assumptions and decisions I had to make when making this

- Roads are not bi-directional. Example input: `X north=Y` does not mean that I can get to X from Y if I go south, unless there is specific line in the file that says `Y {direction}=X`.
- Input `A north=B`,`B north=A` and similar paths are considered to be valid, even though they don't make sense in real life unless both are located on north pole of the planet X ;). Using the same analogy `A west=B`, 'B west=A' also makes sense. As an example, if we start in London, and go west (by some imaginary road) we will end up in New York. Then going west and around the globe, we will end up back in London. Given that those restrictions are not specified in the task description, I have to make that assumption.
- City names are case sensitive. A city name `A` is different than the city name `a` and those are two different cities.
- City names can be made out of letters (both uppercase and lowercase) and can contain `-`. City name must start with a letter.
- Input such as `A north=B west=B` is valid. If B is exactly on the same latitude, but on the opposide side of the globe, then there could exist a road going west which would take us to B and there could also exist a road going north that could take us to B.
- When Aliens first arrive, they can't arrive at the same city at the same time. This doesn't actually play any role, but rather a decision I had to make.
- Inputs:
```
A north=B
```

and

```
A north=B
B
```
are equivalent.

- I added the option for the alien not to move at all. If we have `A north=B` and `B north=A` and we dispatch two aliens, then they would never meet because they would always travel at the same time.
