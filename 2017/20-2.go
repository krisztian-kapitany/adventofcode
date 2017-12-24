package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"util"
)

type vector struct {
	x, y, z int
}

type particle struct {
	id            int
	pos, vel, acc vector
}

func getVectorCoords(s string) (int, int, int) {
	var ax, ay, az int
	for i, part := range strings.Split(s, ",") {

		var ai, _ = strconv.Atoi(part)

		switch i {
		case 0:
			ax = ai
		case 1:
			ay = ai
		case 2:
			az = ai
		}
	}
	return ax, ay, az
}

func posEquals(a vector, b vector) bool {
	if a.x != b.x || a.y != b.y || a.z != b.z {
		return false
	}
	return true
}

func remove(slice []*particle, s int) []*particle {
	return append(slice[:s], slice[s+1:]...)
}

func animate(ps []*particle) {
	for _, p := range ps {
		(*p).vel.x += p.acc.x
		(*p).vel.y += p.acc.y
		(*p).vel.z += p.acc.z

		(*p).pos.x += p.vel.x
		(*p).pos.y += p.vel.y
		(*p).pos.z += p.vel.z
	}
}

func main() {
	lines, _ := util.ReadLines("input20.txt")

	var particles []*particle

	for i, line := range lines {
		var pos, vel, acc vector
		for j, part := range strings.Split(line, ", ") {
			coords := part[strings.Index(part, "=<")+2 : len(part)-1]

			switch j {
			case 0:
				x, y, z := getVectorCoords(coords)
				pos = vector{x, y, z}
			case 1:
				x, y, z := getVectorCoords(coords)
				vel = vector{x, y, z}
			case 2:
				x, y, z := getVectorCoords(coords)
				acc = vector{x, y, z}
			}
		}
		particles = append(particles, &particle{i, pos, vel, acc})
	}

	for tick := 0; tick < 100; tick++ {
		collisionIds := make(map[int]struct{})

		for i := 0; i < len(particles); i++ {
			for j := 0; j < len(particles); j++ {
				if i == j {
					continue
				}
				if posEquals(particles[i].pos, particles[j].pos) {
					var empty struct{}
					collisionIds[i] = empty
					collisionIds[j] = empty
				}
			}
		}

		var toRemove []int
		for i := range collisionIds {
			toRemove = append(toRemove, i)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(toRemove)))

		for _, r := range toRemove {
			particles = remove(particles, r)
		}

		animate(particles)

	}
	fmt.Println(len(particles))
}
