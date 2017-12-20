package main

// TODO: DRY, optimise

import (
	"os"
	"bufio"
	"regexp"
	"aoc2017"
	"fmt"
)

type vector3 struct {
	x int
	y int
	z int
}

type particle struct {
	pos vector3
	vel vector3
	acc vector3
}

func distOf(ptl particle) int {
	return aoc2017.Abs(ptl.pos.x) + aoc2017.Abs(ptl.pos.y) + aoc2017.Abs(ptl.pos.z)
}

func getSwarm() map[int]*particle {
	dat, _ := os.Open("./day20_input.txt")
	scanner := bufio.NewScanner(bufio.NewReader(dat))
	re, _ := regexp.Compile("p=<(-?\\d*),(-?\\d*),(-?\\d*)>, v=<(-?\\d*),(-?\\d*),(-?\\d*)>, a=<(-?\\d*),(-?\\d*),(-?\\d*)>")

	var particles = map[int]*particle{}

	i := 0

	for scanner.Scan() {
		nums := aoc2017.StringsToInts(re.FindStringSubmatch(scanner.Text())[1:])

		particles[i] = &particle{
			vector3{nums[0], nums[1], nums[2]},
			vector3{nums[3], nums[4], nums[5]},
			vector3{nums[6], nums[7], nums[8]},
		}

		i++
	}

	return particles
}

func main() {
	part1()
	part2()
}

func part2() {
	particles := getSwarm()

	for i := 0; i < 1000; i++ {
		for _, ptl := range particles {
			ptl.vel.x += ptl.acc.x
			ptl.vel.y += ptl.acc.y
			ptl.vel.z += ptl.acc.z
			ptl.pos.x += ptl.vel.x
			ptl.pos.y += ptl.vel.y
			ptl.pos.z += ptl.vel.z
		}

		sweep := map[int]bool{}

		for i1, ptl1 := range particles {
			for i2, ptl2 := range particles {
				if i1 != i2 && ptl1.pos == ptl2.pos {
					sweep[i1] = true
					sweep[i2] = true
					fmt.Printf("Boom! %d and %d\n", i1, i2)
				}
			}
		}

		for i, _ := range sweep {
			fmt.Printf("Sweeping %d\n", i)
			delete(particles, i)
		}
	}

	fmt.Printf("Part 2: %d remaining particles\n", len(particles))
}

func part1() {
	particles := getSwarm()

	for i := 0; i < 1000; i++ {
		for _, ptl := range particles {
			ptl.vel.x += ptl.acc.x
			ptl.vel.y += ptl.acc.y
			ptl.vel.z += ptl.acc.z
			ptl.pos.x += ptl.vel.x
			ptl.pos.y += ptl.vel.y
			ptl.pos.z += ptl.vel.z
		}
	}

	closest := 0
	closestValue := distOf(*particles[0])

	for i, ptl := range particles {
		dst := distOf(*ptl)

		if dst < closestValue {
			closest = i
			closestValue = dst
		}
	}

	fmt.Printf("Part 1: Closest is particle %d with dist %d\n", closest, closestValue)
}
