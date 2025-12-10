package year2025

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Day8Vector3D struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func (v *Day8Vector3D) GetRelativeDistance(o *Day8Vector3D) float64 {
	return math.Sqrt(math.Pow(float64(v.X-o.X), 2) + math.Pow(float64(v.Y-o.Y), 2) + math.Pow(float64(v.Z-o.Z), 2))
}

type Day8VectorDistance struct {
	V1       *Day8Vector3D `json:"v1"`
	V2       *Day8Vector3D `json:"v2"`
	Distance float64       `json:"distance"`
}

type Day8Circuit struct {
	Vectors map[*Day8Vector3D]bool `json:"vectors"`
	Count   int                    `json:"count"`
}

func (c *Day8Circuit) AddVector(v *Day8Vector3D) {
	if c.Vectors[v] {
		return // Already in circuit, don't double-count
	}
	c.Vectors[v] = true
	c.Count++
}

func (c *Day8Circuit) checkMembership(v *Day8Vector3D) bool {
	_, ok := c.Vectors[v]
	return ok
}

func (c *Day8Circuit) Print() {
	fmt.Println("[")
	for v := range c.Vectors {
		value := *v
		fmt.Printf("(%d, %d, %d)\n", value.X, value.Y, value.Z)
	}
	fmt.Println("]")
	fmt.Println("count: ", c.Count)
}

func createDay8Circuit(d Day8VectorDistance) *Day8Circuit {
	c := &Day8Circuit{
		Vectors: make(map[*Day8Vector3D]bool),
		Count:   0,
	}
	c.AddVector(d.V1)
	c.AddVector(d.V2)
	return c
}

func buildDay8Vector(x int, y int, z int) *Day8Vector3D {
	return &Day8Vector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func checkVectorMapMembership(key *Day8Vector3D, m map[*Day8Vector3D]int) bool {
	_, ok := m[key]
	return ok
}

func getDay8Part1Example() string {
	return `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`
}

func Day8(data string, cycleLimit int) {
	if data == "" {
		data = getDay8Part1Example()
	}
	// closestCount := 3
	vectors := make([]*Day8Vector3D, 0)
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x := convertStringToInt(parts[0])
		y := convertStringToInt(parts[1])
		z := convertStringToInt(parts[2])
		vectorPtr := buildDay8Vector(x, y, z)
		vectors = append(vectors, vectorPtr)
	}
	vectorDistances := make([]Day8VectorDistance, 0)
	for i, v1 := range vectors {
		for _, v2 := range vectors[i+1:] {
			if v1 == v2 {
				continue
			}
			distance := v1.GetRelativeDistance(v2)
			vectorDistances = append(vectorDistances, Day8VectorDistance{V1: v1, V2: v2, Distance: distance})
		}
	}
	slices.SortFunc(vectorDistances, func(a Day8VectorDistance, b Day8VectorDistance) int {
		if a.Distance <= b.Distance {
			return -1
		}
		return 1
	})
	// files.PrintJsonType(vectorDistances[0:5])

	circuits := make([]*Day8Circuit, 0)
	added := make(map[*Day8Vector3D]int)
	for i, v := range vectorDistances {
		if i == cycleLimit {
			break
		}
		// fmt.Printf("v: (%d, %d, %d) - (%d, %d, %d)\n", v.V1.X, v.V1.Y, v.V1.Z, v.V2.X, v.V2.Y, v.V2.Z)

		v1 := v.V1
		v2 := v.V2
		v1ok, v2ok := checkVectorMapMembership(v1, added), checkVectorMapMembership(v2, added)
		if !v1ok && !v2ok {
			fmt.Printf("adding new circuit: %d for v: (%d, %d, %d) - (%d, %d, %d)\n", len(circuits), v.V1.X, v.V1.Y, v.V1.Z, v.V2.X, v.V2.Y, v.V2.Z)
			circuits = append(circuits, createDay8Circuit(v))
			added[v1] = len(circuits) - 1
			added[v2] = len(circuits) - 1
		} else if v1ok && !v2ok {
			fmt.Printf("adding v2 (%d, %d, %d) to circuit: %d\n", v2.X, v2.Y, v2.Z, added[v1])
			circuits[added[v1]].AddVector(v2)
			added[v2] = added[v1]
		} else if !v1ok && v2ok {
			fmt.Printf("adding v1 (%d, %d, %d) to from distance (%d, %d, %d) - (%d, %d, %d) circuit: %d\n", v1.X, v1.Y, v1.Z, v2.X, v2.Y, v2.Z, v.V1.X, v.V1.Y, v.V1.Z, added[v2])
			circuits[added[v2]].AddVector(v1)
			added[v1] = added[v2]
		} else if v1ok && v2ok && added[v1] != added[v2] {
			circuit1Idx := added[v1]
			circuit2Idx := added[v2]
			circuit1 := circuits[circuit1Idx]
			circuit2 := circuits[circuit2Idx]

			// Merge smaller circuit into larger circuit for efficiency
			var targetIdx, sourceIdx int
			var targetCircuit, sourceCircuit *Day8Circuit
			if circuit1.Count >= circuit2.Count {
				targetIdx = circuit1Idx
				sourceIdx = circuit2Idx
				targetCircuit = circuit1
				sourceCircuit = circuit2
			} else {
				targetIdx = circuit2Idx
				sourceIdx = circuit1Idx
				targetCircuit = circuit2
				sourceCircuit = circuit1
			}

			fmt.Printf("merging circuits: %d (size %d) into %d (size %d)\n", sourceIdx, sourceCircuit.Count, targetIdx, targetCircuit.Count)

			// Add all vectors from source circuit to target circuit
			for vec := range sourceCircuit.Vectors {
				if !targetCircuit.checkMembership(vec) {
					targetCircuit.AddVector(vec)
				}
				// Update added map to point to target circuit
				added[vec] = targetIdx
			}

			// Clear the source circuit (optional, but keeps things clean)
			sourceCircuit.Vectors = make(map[*Day8Vector3D]bool)
			sourceCircuit.Count = 0
		}

		/*for i, circuit := range circuits {
			fmt.Printf("checking circuit: %d\n", i)
			if circuit.checkMembership(v1) {
				add = false
				if circuits[i].checkMembership(v2) {
					added[v2] = i
					// fmt.Printf("v2 (%d, %d, %d) already in circuit: %d\n", v2.X, v2.Y, v2.Z, i)
					break
				}
				fmt.Printf("adding v2 (%d, %d, %d) to circuit: %d\n", v2.X, v2.Y, v2.Z, i)
				circuits[i].AddVector(v2)
				break
			} else if circuit.checkMembership(v2) {
				add = false
				if circuits[i].checkMembership(v1) {
					added[v1] = true
					// fmt.Printf("v1 (%d, %d, %d) already in circuit: %d\n", v1.X, v1.Y, v1.Z, i)
					break
				}
				fmt.Printf("adding v1 (%d, %d, %d) to circuit: %d\n", v1.X, v1.Y, v1.Z, i)
				circuits[i].AddVector(v1)
				break
			}
		}
		if add {
			fmt.Printf("adding new circuit: %d for v: (%d, %d, %d) - (%d, %d, %d)\n", len(circuits), v.V1.X, v.V1.Y, v.V1.Z, v.V2.X, v.V2.Y, v.V2.Z)
			circuits = append(circuits, createDay8Circuit(v))
		}
		*/

	}

	slices.SortFunc(circuits, func(a *Day8Circuit, b *Day8Circuit) int {
		if a.Count >= b.Count {
			return -1
		}
		return 1
	})

	part1Result := 1
	for i := 0; i < 3; i++ {
		circuits[i].Print()
		part1Result *= circuits[i].Count
	}
	fmt.Println("part1 result: ", part1Result)
}
