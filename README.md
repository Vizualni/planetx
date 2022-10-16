

## Assumptions

- Roads are not bi-directional, unless explicitly defined in the _planetx_ file. Example input: `X north=Y` does not mean that I can get to X from Y if I go south, unless there is specific line in the file that says `Y {direction}=X`.
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
