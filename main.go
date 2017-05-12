package main

import "fmt"

type section struct {
    host string
    port int
    user string
    pass string
    next_section *section
}

type road struct {
    sections int
    next_section section
}

func main() {
    r1s2 := section{host: "192.168.100.102", port: 22, user: "root", pass: "root", next_section: nil}
    r1s1 := section{host: "192.168.100.101", port: 22, user: "root", pass: "root", next_section: &r1s2}
    r1 := road{sections: 2, next_section: r1s1}

    fmt.Println(r1.sections)
    fmt.Println(r1.next_section.host)
    fmt.Println(r1.next_section.next_section.host)

    r2s1 := section{host: "192.168.200.101", port: 22, user: "root", pass: "root", next_section: nil}
    r2 := road{sections: 1, next_section: r2s1}

    fmt.Println(r2.sections)
    fmt.Println(r2.next_section.host)

    roads := [2]road{r1, r2}
    fmt.Println(roads)
    fmt.Println(roads[0])
    fmt.Println(roads[1])
}