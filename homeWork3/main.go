package main

import "fmt"

type WeightFeedCalculater interface {
	CalculateWeightFeed() int
	GetName() string
	GetWeight() int
}

type Dog struct {
	weight          int
	feedPerKilogram int
	amount          int
	name            string
}

type Cat struct {
	weight          int
	feedPerKilogram int
	amount          int
	name            string
}

type Cow struct {
	weight          int
	feedPerKilogram int
	amount          int
	name            string
}

func (d Dog) CalculateWeightFeed() int {
	return d.feedPerKilogram * d.weight * d.amount
}

func (d Dog) GetName() string {
	return d.name
}
func (d Dog) GetWeight() int {
	return d.weight
}

func (c Cat) CalculateWeightFeed() int {
	return c.feedPerKilogram * c.weight * c.amount
}

func (c Cat) GetName() string {
	return c.name
}

func (c Cat) GetWeight() int {
	return c.weight
}

func (cw Cow) CalculateWeightFeed() int {
	return cw.feedPerKilogram * cw.weight * cw.amount
}

func (cw Cow) GetName() string {
	return cw.name
}

func (cw Cow) GetWeight() int {
	return cw.weight
}

func main() {

	arr := []WeightFeedCalculater{
		Cow{
			weight:          100,
			feedPerKilogram: 25,
			amount:          1,
			name:            "Cow",
		},
		Dog{
			weight:          24,
			feedPerKilogram: 2,
			amount:          1,
			name:            "Dog",
		},
		Cat{
			weight:          10,
			feedPerKilogram: 7,
			amount:          1,
			name:            "Cat",
		},
	}

	sum := getAllFeed(arr)
	fmt.Println("Feed for the whole farm", sum)
}

func getAllFeed(farm []WeightFeedCalculater) int {

	var sum int
	for _, animal := range farm {

		sum += animal.CalculateWeightFeed()
		fmt.Printf("Animal name: %s\nAnimal weight: %d\nAnimal feed: %d\n", animal.GetName(), animal.GetWeight(), animal.CalculateWeightFeed())
	}

	return sum

}
