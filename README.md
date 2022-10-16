

## Assumptions

- Roads are not bi-directional, unless explicitly defined in the _planetx_ file. Example input: `X north=Y` does not mean that I can get to X from Y if I go south, unless there is specific line in the file that says `Y {direction}=X`.
- Input `A north=B`,`B north=A` and similar paths are considered to be valid, even though they don't make sense in real life unless both are located on north pole of the planet X ;). Using the same analogy `A west=B`, 'B west=A' also makes sense. As an example, if we start in London, and go west (by some imaginary road) we will end up in New York. Then going west and around the globe, we will end up back in London. Given that those restrictions are not specified in the task description, I have to make that assumption.
