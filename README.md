# SEQ Search

Search SEQ files in a directory for a specific hex sequence. Used when you have the
bytes of an SEQ file but don't know which SEQ file you're looking at.

## How to Use

Download the [latest release](https://github.com/NicholasMoser/SEQ-Search/releases),
extract it, and run the executable. Enter the byte sequence in hex and type in the
directory you wish to search.

Let's say we want to search for the byte sequence in this file:

![Example Bytes](example_bytes.png?raw=true "Example Bytes")

And here is an example of how we would use the program and the output:

![Input](seq_search_input.png?raw=true "Input")

![Output](seq_search_output.png?raw=true "Output")

Many files and offsets are returned, meaning that this is a common byte sequence.
We would want to use a larger byte sequence to narrow down exactly what file this
code is in.
