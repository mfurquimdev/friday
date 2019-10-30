package main

import (
    "fmt"
)

type Time struct {
    Hour int
    Minute int
}

func NewTime(hour, minute int) *Time {
    return &Time {
        Hour: hour,
        Minute: minute,
    }
}

func (t *Time) TotalMinutes() int {
    return t.Hour*60 + t.Minute
}

func (t *Time) Sum(t2 *Time) {
    var minutes int = t.TotalMinutes() + t2.TotalMinutes()
    t.Hour, t.Minute = minutes/60, minutes%60
}

func (t *Time) Sub(t2 *Time) {
    var minutes int = t2.TotalMinutes() - t.TotalMinutes()
    t.Hour, t.Minute = minutes/60, minutes%60
}

type Elapsed struct {
    In *Time
    Out *Time
}

func (e *Elapsed) ElapsedTime() (int, int) {
    var minutes = e.Out.TotalMinutes() - e.In.TotalMinutes()
    return minutes/60, minutes%60
}

func NewElapsed(in_hour, in_minute, out_hour, out_minute int) *Elapsed {
    return &Elapsed {
        In: NewTime(in_hour, in_minute),
        Out: NewTime(out_hour, out_minute),
    }
}

func main() {
    week37 := make([]*Elapsed, 0, 4)

    week37 = append(week37, NewElapsed(8,54,18,2))
    week37 = append(week37, NewElapsed(9,17,18,19))
    week37 = append(week37, NewElapsed(9,23,19,20))
    week37 = append(week37, NewElapsed(9,12,18,54))
    friday := NewTime(9,33)

    totalTime := new(Time)
    for _, value := range week37 {
        fmt.Printf("%dh%d~%dh%d\t",
            value.In.Hour, value.In.Minute, value.Out.Hour, value.Out.Minute)

        h, m := value.ElapsedTime()
        fmt.Printf("%2d Hours %2d Minutes\n", h, m);
        totalTime.Sum(NewTime(h,m))
    }

    fmt.Printf("totalTime:\t%2d Hours %2d Minutes\n", totalTime.Hour, totalTime.Minute)

    fmt.Printf("%dh%d~", friday.Hour, friday.Minute)
    totalTime.Sub(NewTime(9*(len(week37)+1),0))

    friday.Sum(totalTime)
    fmt.Printf("%dh%d\n", friday.Hour, friday.Minute)
}

