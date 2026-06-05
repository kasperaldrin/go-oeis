# OEIS

A Go client for the [OEIS](https://oeis.org) (Online Encyclopedia of Integer Sequences). Look up sequences by ID from a local database dump, or (planned) query the live site.

## Installation

1. Clone the entire sequence database from [https://github.com/oeis/oeisdata](https://github.com/oeis/oeisdata)

## Usage
Get() - Will get a sequence by it's ID.
Search() - Get a list of sequences matching the sequence.
Filter() - Get a list of sequences matching the set of filters applied to all sequences.
Sequences() - Get an itterator over all sequences in the database.

## Iterate all sequences (offline)

Sequence files live at `{OfflinePath}/seq/{first4OfID}/{ID}.seq`. Use `NewSequenceIterator` or `for seq, err := range client.Sequences()` on an `OfflineClient` to walk every sequence in sorted order.
