package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	db_path string
}

var (
	config *Config = &Config{
		db_path: "/tmp/goleveldb-data",
	}
)

func init() {
	flag.StringVar(&config.db_path, "dbpath", config.db_path, "go leveldb path")
}

func main() {
	flag.Parse()

	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, err := leveldb.OpenFile(config.db_path, nil)
	if err != nil {
		fmt.Println("Failed to open db:", err.Error())
		return
	}
	defer db.Close()

	err = db.Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		panic(err)
	}

	// Remember that the contents of the returned slice should not be modified.
	data, err := db.Get([]byte("key"), nil)
	if err != nil {
		fmt.Println("Failed to get data:", err.Error())
	}
	fmt.Printf("Data: %+v\n", data)

	buf := make([]byte, 0, len(data))
	copy(buf, data)
	buf = append(buf, '!')
	err = db.Put([]byte("key"), buf, nil)
	if err != nil {
		fmt.Println("Failed to update:", err.Error())
	}

	err = db.Delete([]byte("key"), nil)
	if err != nil {
		fmt.Println("Failed to delete:", err.Error())
	}

	samples := [5]string{"these", "are", "test", "data", "!"}
	for i := 0; i < len(samples); i++ {
		key := "key" + strconv.Itoa(i)
		err = db.Put([]byte(key), []byte(samples[i]), nil)
		if err != nil {
			panic(err)
		}
	}

	idx := 0
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		idx++
		fmt.Println("index:", idx, "key:", string(key), "value:", string(value))
	}
	iter.Release()
	err = iter.Error()
	fmt.Println("Iter satus:", err)

	//Seek-then-Iterate:

	key := []byte("key3")
	iter = db.NewIterator(nil, nil)
	for ok := iter.Seek(key); ok; ok = iter.Next() {
		// Use key/value.
		key := iter.Key()
		value := iter.Value()
		fmt.Println("index:", idx, "key:", string(key), "value:", string(value))
	}
	iter.Release()
	err = iter.Error()
	fmt.Println("Iter satus:", err)

	/*
		Iterate over subset of database content:

		iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
		for iter.Next() {
			// Use key/value.
			...
		}
		iter.Release()
		err = iter.Error()
		...
		Iterate over subset of database content with a particular prefix:

		iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
		for iter.Next() {
			// Use key/value.
			...
		}
		iter.Release()
		err = iter.Error()
		...
		Batch writes:

		batch := new(leveldb.Batch)
		batch.Put([]byte("foo"), []byte("value"))
		batch.Put([]byte("bar"), []byte("another value"))
		batch.Delete([]byte("baz"))
		err = db.Write(batch, nil)
		...
		Use bloom filter:

		o := &opt.Options{
			Filter: filter.NewBloomFilter(10),
		}
		db, err := leveldb.OpenFile("path/to/db", o)
		...
		defer db.Close()
		...
	*/
}
