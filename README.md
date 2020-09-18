# Dewey - Your personal library manager

Dewey is a command line program that allows you to manage your library(ies) of books.

## Basic usage

To create a new library run `dewey new library <name>` (dewey will prompt you for a name if none is provided.).

You will now have an empty library. You can add books with `dewey add <library>` which will start the add prompt
for the specified library. If you only have one library `dewey add` will suffice.

`dewey show` will print out all the books in your library sorted alphabetically. Future version will allow more
advanced filtering.

## Future Plans

    - Advanced filtering features (all of the below combineable) 
      for example: `dewey search +author JRR Tolkien +read yes` and `dewey search -read -tag horror`.
      Where `+` means "include" and `-` "filter out".
      - Filter by author
      - Filter by tag
      - Filter by read
      - Fuzzy search?
      - etc.
    - Create collections in a library (for example a collection of "Law Litterature")
      `dewey new collection Kennedy +subject Kennedy`?
    - Emacs integration
