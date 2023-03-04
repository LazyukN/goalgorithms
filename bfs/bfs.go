package main

import "fmt"

type Queue struct {
	queue []Person
}

type Person struct {
	name string
	sex  string
}

func main() {
	bob := Person{"bob", "female"}
	alisa := Person{"alisa", "male"}
	clar := Person{"clar", "male"}
	peggy := Person{"peggy", "male"}
	anudg := Person{"anudg", "male"}
	tom := Person{"tom", "male"}
	johny := Person{"johny", "female"}

	graph := map[string][]Person{
		"nikita": []Person{bob, clar, alisa},
		"bob":    []Person{anudg, peggy},
		"clar":   []Person{tom, johny},
		"alisa":  []Person{peggy},
	}

	femaleName := findFemale(graph, "nikita")
	fmt.Println(femaleName)

}

func (q *Queue) enqueue(element Person) {
	q.queue = append(q.queue, element)
}

func (q *Queue) dequeue() Person {
	element := q.queue[0]
	q.queue = q.queue[1:]
	return element
}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func (p *Person) isFemale() bool {
	return p.sex == "female"
}

func findFemale(graph map[string][]Person, vertex string) string {
	q := Queue{
		graph[vertex],
	}
	cache := make(map[string]int)
	for !q.isEmpty() {
		p := q.dequeue()
		if p.isFemale() {
			return p.name
		} else {
			if _, ok := cache[p.name]; ok {
				continue
			} else {
				for _, friend := range graph[p.name] {
					q.enqueue(friend)
				}
				cache[p.name] = 1
			}
		}
	}

	return "sausage party"
}
