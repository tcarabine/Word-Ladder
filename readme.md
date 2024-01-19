# Word Ladder

Given 2 words of equal length, starting with one and ending with the other, build a chain of words. Successive entries in the chain must all be real words (use the dictionary provided), and each can differ from the previous word by just one letter. For example, you can get from “cat” to “dog” using the following chain.

- cat
- cot
- cog
- dog

## Running me

### To Build

```sh
docker build -t wordchain:latest .
```

### With defaults

```sh
docker run wordchain
```

### With custom search

```sh
docker run -e start=cat -e end=dog wordchain
```