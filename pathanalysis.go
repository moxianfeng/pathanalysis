package main

import (
    "flag"
    "log"
    "io/ioutil"
    "os"
    "fmt"
)

var (
    depth int = 4;
    root string = "";
    pathSize map[string]int64;
)

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
        pathSize[path] = _size;
    }
    return _size;
}

func main() {
    flag.Parse();
    if len(root) == 0 {
        flag.PrintDefaults();
        return;
    }

    statPath(root, 0);

    for p, s := range(pathSize) {
        fmt.Println(p, s);
    }
}

func init() {
    flag.IntVar(&depth, "depth", 4, "report result depth, base on the root path");
    flag.StringVar(&root, "root", "", "root path to analysys");

    pathSize = make(map[string]int64);
}

