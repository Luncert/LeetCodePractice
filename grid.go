package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y int
}

func (p Point) compare(t Point) int {
	if p.X == t.X && p.Y == t.Y {
		return 0
	} else if p.X >= t.X && p.Y >= t.Y {
		return 1
	} else {
		return -1
	}
}

func (p Point) adjacentTo(t Point) bool {
	return math.Abs(float64(p.X-t.X)) <= 1 ||
		math.Abs(float64(p.Y-t.Y)) <= 1
}

func (p *Point) set(x, y int) {
	p.X, p.Y = x, y
}

func (p *Point) setX(x int) {
	p.X = x
}

func (p *Point) setY(y int) {
	p.Y = y
}

type Rect struct {
	sz         int
	start, end *Point
}

func newRect(x1, y1, x2, y2 int) *Rect {
	if x2 > x1 && y2 > y1 {
		return &Rect{
			sz:    (x2 - x1) * (y2 - y1),
			start: &Point{x1, y1},
			end:   &Point{x2, y2},
		}
	}
	return nil
}

func (r *Rect) containsPoint(p Point) bool {
	return r.start.compare(p) < 0 && r.end.compare(p) > 0
}

type RectList struct {
	rects []*Rect
}

func newRectList(w, h int) (rl *RectList) {
	rl = &RectList{rects: make([]*Rect, 1)}
	rl.rects[0] = newRect(0, 0, w-1, h-1)
	return
}

func (rl *RectList) Len() int { return len(rl.rects) }

func (rl *RectList) Swap(i, j int) { rl.rects[i], rl.rects[j] = rl.rects[j], rl.rects[i] }

func (rl *RectList) Less(i, j int) bool { return rl.rects[i].sz < rl.rects[j].sz }

func (rl *RectList) autoCut(p Point) {
	for idx, rect := range rl.rects {
		if rect.containsPoint(p) {
			rl.rects = append(rl.rects[:idx], rl.rects[idx+1:]...)
			// top
			if r := newRect(rect.start.X, rect.start.Y, rect.end.X, p.Y-1); r != nil {
				rl.rects = append(rl.rects, r)
			}
			// left
			if r := newRect(rect.start.X, rect.start.Y, p.X-1, rect.end.Y); r != nil {
				rl.rects = append(rl.rects, r)
			}
			// right
			if r := newRect(p.X+1, rect.start.Y, rect.end.X, rect.end.Y); r != nil {
				rl.rects = append(rl.rects, r)
			}
			// bottom
			if r := newRect(p.X-1, p.Y+1, rect.end.X, rect.end.Y); r != nil {
				rl.rects = append(rl.rects, r)
			}
			break
		}
	}
}

func kMarsh(grid []string) {
	row := len(grid)
	col := len(grid[0])

	rl := newRectList(col, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 'x' {
				rl.autoCut(Point{j, i})
			}
		}
	}
	if len(rl.rects) == 0 {
		fmt.Println("impossible")
	} else {
		sort.Sort(rl)

	}
}
