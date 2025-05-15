package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	_ "github.com/glebarez/go-sqlite"
)

var (
	initFlag         = flag.Bool("init", false, "Initializes the database and stores SHA 256 hashes retrieved from \"path\"")
	checkFlag        = flag.Bool("check", false, "Checks the path for hash mismatch")
	checkFlagSimple  = flag.Bool("c", false, "Checks the path for hash mismatch (simplified output, no color)")
	updateFlag       = flag.Bool("update", false, "Updates the hash database for the directory")
	databaseLocation = flag.String("hash-db-loc", "./hashes.db", "The database location for the SHA 256 hashes")
)

func main() {
	flag.Parse()
	path := flag.Arg(0) // get the filepath

	if path == "" {
		fmt.Fprintln(os.Stderr, "Usage: ./file-integrity-checker [options] path/directory.")
		fmt.Fprintln(os.Stderr, "See ./file-integrity-checker -h for options")
		return
	}

	// Connect to the SQLite database
	db, err := sql.Open("sqlite", *databaseLocation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating database: %s\n", err)
		return
	}

	defer db.Close()

	_, err = create_table(db, "hashes")
	if err != nil {
		fmt.Println(err)
		return
	}

	// index 0 is get, index 1 is add, index 2 is update
	prepared_statements, err := get_prepared_statements(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting prepared statements: %s\n", err)
	}

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return process_file(path, prepared_statements)
		}
		return nil
	})
}

func process_file(file string, prepared_statements [3]*sql.Stmt) error {
	file_hash, err := hash_file(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open filename specified by %s: %v\n", file, err)
		return err
	}
	if *initFlag {
		err = modify_hash_in_table(prepared_statements[1], file, file_hash)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting %s into the database. The hash already exists\n", file)
			return err
		}

		fmt.Printf("Identified file %s and successfully stored its hash: %s\n", file, file_hash)

	} else if *updateFlag {
		// first, we make sure if the entry actually exists
		_, err = get_hash_from_filename(prepared_statements[0], file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating the database for %s. The hash does not exist\n", file)
			return err
		}

		err = modify_hash_in_table(prepared_statements[2], file, file_hash)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error modifying the hash for %s in the database: %v", file, err)
			return err
		}

		fmt.Printf("Hash for file [%s] updated successfully.\n", file)
	} else if *checkFlag {
		stored_hash, err := get_hash_from_filename(prepared_statements[0], file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating the database for %s. The hash does not exist\n", file)
			return err
		}

		if stored_hash != file_hash {
			fmt.Fprintf(os.Stderr, "File [%s] Status: Modified (", file)
			// print in color
			color.New(color.FgGreen).Print("Hash ")
			color.New(color.FgYellow).Print("mismatch")
			fmt.Fprintln(os.Stderr, ")")
		} else {
			fmt.Printf("File [%s] Status: Unmodified\n", file)
		}
	} else if *checkFlagSimple {
		stored_hash, err := get_hash_from_filename(prepared_statements[0], file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating the database for %s. The hash does not exist\n", file)
			return err
		}

		if stored_hash != file_hash {
			fmt.Fprintf(os.Stderr, "File [%s] Status: Modified (Hash mismatch)\n", file)
		} else {
			fmt.Printf("File [%s] Status: Unmodified\n", file)
		}
	}
	return nil
}
