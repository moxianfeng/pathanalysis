package main

import (
    "flag"
    "log"
    "io/ioutil"
    "os"
    "fmt"
    "sort"
)

var (
    depth int = 4;
    root string = "";
    pathSize map[string]int64;

    sizeArray []PathSize;
)

type PathSize struct {
    path string;
    size int64;
}

func statPath(path string, curdepth int) int64 {
    var _size int64 = 0;

    files, err := ioutil.ReadDir(path);
    if nil != err {
        log.Print(err);
        return 0;
    }

    for _, file := range(files) {
        if file.IsDir() {
            subpath := fmt.Sprintf("%s%c%s", path, os.PathSeparator, file.Name());
            _size += statPath(subpath, curdepth + 1);
        } else {
            _size += file.Size();
        }
    }


    if curdepth == depth && _size > 0 {
        sizeArray = append(sizeArray, PathSize{path: path, size: _size});
        // pathSize[path] = _size;
    }
    return _size;
}

func humanSize(size int64) string {
    unit := []string{"", "K", "M", "G", "T", "P", "E"};
    var i int = 0;
    for size > 1000 {
        size /= 1000;
        i++;
    }
    return fmt.Sprintf("%d%s", size, unit[i]);
}

func main() {
    flag.Parse();
    if len(root) == 0 {
        flag.PrintDefaults();
        return;
    }

    statPath(root, 0);

    sort.Slice(sizeArray, func(i, j int) bool {
        return sizeArray[i].size > sizeArray[j].size
    });
    for _, s := range(sizeArray) {
        fmt.Println(s.path, humanSize(s.size));
    }
}

func init() {
    flag.IntVar(&depth, "depth", 4, "report result depth, base on the root path");
    flag.StringVar(&root, "root", "", "root path to analysys");

    pathSize = make(map[string]int64);
    sizeArray = make([]PathSize, 0);
}

