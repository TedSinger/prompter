package main

type DirType int
const (
    Normal DirType = iota
    Link DirType = iota
    CrossDeviceLink DirType = iota
)

func getDirType(p string) DirType {
    return Normal
}

type PermType int
const (
    ClosedRead PermType = iota
    OpenReadClosedWrite PermType = iota
    OpenWrite PermType = iota
)

func getPermissions(p string) PermType {
    return OpenWrite
}

type Ownership int
const (
    You Ownership = iota
    Other Ownership = iota
)

func getOwnership(p string) Ownership {
    return You
}