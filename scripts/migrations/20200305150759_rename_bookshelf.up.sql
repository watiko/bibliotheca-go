ALTER TABLE book_shelves RENAME TO bookshelves;
ALTER TABLE bookshelves RENAME COLUMN book_shelf_id TO bookshelf_id;
ALTER TABLE books RENAME COLUMN book_shelf_id TO bookshelf_id;
