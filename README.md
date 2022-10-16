# PlanetX invasion!

## My notes :)

Hello reviewers! 

## Algorithm

This problem can be modeled as a simple graph where cities are nodes and roads between them represent directed edges between the cities (see Assumptions why those are directed).
To move the aliens on the planet a `Randomizer` interface is used to help to determine which way to move (in either completely pseudo-random way (can be repeatable) or with the cryptorandom method).
The algorithm is quite easy. We spawn M aliens (where M<=N, where N is number of cities) randomly on a map, and move them on every iteration. If two end up in the same city, we remove the city from the graph by removing roads from the cities that lead to the target city + we also remove the two aliens that fought.

### Time & Space complexity

Let's define variables:
- N number of cities on a planet
- M number of aliens
- K max iteration

Given that M's upper bound is set with the N, which is the number of cities, we can say that `M == N`.

#### Time

Time complexity is: `O(K * M) ~> *O(M)* == *O(N)*` which means that it linearly depends on the number of aliens spawned. Size of a planet and number of roads do not determine the complexity. Given that K is a constant, we can ignore it, thus we come to a `O(M)` as time complexity.


#### Space

Space complexity is: `O(4*N + O(M)) ~> O(N+M) -> given that M can get as large as N, then ~> O(N+N) ~> O(2N) -> *O(N)*`. For every city in a graph we are saving at most 4 locations (N, S, W, E) + we are also saving the location of every alien on a map.


### A possible alternative

One of the alternative solutions would be to use goroutines, but it just adds a bit more complexity as there would be a lot of syncing required between the goroutines (each Alien would be its own goroutine), and we would not gain any performance boosts, rather it will become slower and more difficult to maintain. For large enough N (where N is number of cities => max number of Aliens) then we might spawn quite a lot of goroutines, which is not a problem in itself as goroutines are quite light-weight, but rather it becomes impractical.

### Possible improvements

- Given that aliens can get "stranded" quite easily when city that happens to be "hub" between two other sub-graphs gets destroyed. Then those alies can't come to the other city. Improvement that I see is to detect when two graphs _split_ when destroying a city so that those aliens can be ignored.
- Given that this is a directed graph, an Alien can end up on the "leaf" of the graph and can't go anywhere else, then we can simply stop tracking that alien for its movements.

None of those improvements have been implemented because those would not improve the Big-Oh of the already implemented algorithm (+ it's Sunday for me ;) ).



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
