package main

import (
    "flag"
    "fmt"
    "hash/crc64"
    "hash/fnv"
    "io"
    "log"
    "os"
    "path/filepath"
    "runtime"
    "strings"
    "sync"
    "time"
)

type Runner struct {
    sync.RWMutex
    checksums  map[string][]string
    collisions map[string]bool
}

// a simple workaround for "too many open" files problem
const (
    NumberOfGorutines = 128
)

var (
    wg                  sync.WaitGroup
    fileExtension       = flag.String("extension", "", "File fileExtension like JPG, MOV etc...")
    NumberOfBytesToRead = flag.Int("kbytes", 16, "Number of bytes to read")
    NumberOfFiles       int
    NumberOfCollisions  int
)

func (r *Runner) calculateHash(filename string) {
    // get the CRC64 checksum from first NumberOfBytesToRead bytes
    checksum := calculateCRC64Hash(filename)

    r.RLock()
    conflicted_files, ok := r.checksums[checksum]
    r.RUnlock()

    if ok {
        // calculate whole sum if checksum in checksums map
        if calculateFNVHash(filename) == calculateFNVHash(conflicted_files[0]) {
            r.Lock()
            r.checksums[checksum] = append(conflicted_files, filename)
            r.Unlock()

            r.RLock()
            _, ok := r.collisions[checksum]
            r.RUnlock()

            if !ok {
                r.Lock()
                r.collisions[checksum] = true
                r.Unlock()
            }
        }
    } else {
        r.Lock()
        r.checksums[checksum] = append(r.checksums[checksum], filename)
        r.Unlock()
    }
    wg.Done()
}

func calculateFNVHash(filename string) string {
    f, err := os.Open(filename)
    defer f.Close()
    if err != nil {
        log.Panic(err)
    }

    h := fnv.New64()
    io.Copy(h, f)
    return fmt.Sprintf("%x", h.Sum(nil))
}

func calculateCRC64Hash(filename string) string {
    var buffer []byte = make([]byte, 1024**NumberOfBytesToRead)

    f, err := os.Open(filename)
    defer f.Close()
    if err != nil {
        log.Panic(err)
    }

    _, err = f.Read(buffer)
    if err != nil {
        log.Panic(err)
    }

    return fmt.Sprintf("%d", crc64.Checksum(buffer, crc64.MakeTable(crc64.ISO)))
}

func init() {
    flag.Parse()
    if flag.NArg() < 1 {
        flag.Usage()
        os.Exit(1)
    }
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    r := &Runner{checksums: make(map[string][]string), collisions: make(map[string]bool)}

    path, err := filepath.Abs(flag.Arg(0))
    if err != nil {
        log.Panic(err)
    }

    _, err = os.Stat(path)
    if os.IsNotExist(err) {
        fmt.Println("Directory does not exist...")
        os.Exit(1)
    }

    fmt.Printf("Started to walk over %s and its sub-directories to find duplicated files\n\n", path)

    start := time.Now()

    filepath.Walk(path, func(path string, info os.FileInfo, _ error) error {
        if _, file := filepath.Split(path); file != "" {
            // Skip various temporary or "hidden" files or directories.
            if file[0] == '.' || file[0] == '#' || file[0] == '~' || file[len(file)-1] == '~' {
                if info.IsDir() {
                    return filepath.SkipDir
                }
                return nil
            }
        }
        if info.Size() > 0 && info.Mode()&os.ModeType == 0 && strings.HasSuffix(strings.ToLower(path), strings.ToLower(*fileExtension)) {
            wg.Add(1)
            go r.calculateHash(path)
            NumberOfFiles++
            // wait NumberOfGorutines to finish
            if NumberOfFiles%NumberOfGorutines == 0 {
                wg.Wait()
            }
        }
        return nil
    })
    // wait remaining gorutines
    wg.Wait()

    elapsedTime := time.Since(start).Seconds()
    for i := range r.collisions {
        NumberOfCollisions++
        checksums := r.checksums[i]
        fmt.Printf("%d) %d items\n", NumberOfCollisions, len(checksums))
        for _, j := range checksums {
            fmt.Printf("\t%s\n", j)
        }
        fmt.Println()
    }
    fmt.Printf("Checked %d files, found %d different duplicated set of files (took %.2f sec)\n", NumberOfFiles, NumberOfCollisions, elapsedTime)
}
