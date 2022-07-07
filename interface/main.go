package main

import "fmt"

type WeightFeedCalculater interface {
	CalculateWeightFeed() int
	GetName() string
	GetWeight() int
}

type Dog struct {
	weight int
}

type Cat struct {
	weight int
}

type Cow struct {
	weight int
}

func (d Dog) CalculateWeightFeed() int {
	const feedPerKilogram = 2
	return feedPerKilogram * d.weight
}

func (d Dog) GetName() string {
	const animalName = "Dog"
	return animalName
}
func (d Dog) GetWeight() int {
	return d.weight
}

func (c Cat) CalculateWeightFeed() int {
	const feedPerKilogram = 7
	return feedPerKilogram * c.weight
}

func (c Cat) GetName() string {
	const animalName = "Cat"

	return animalName
}

func (c Cat) GetWeight() int {
	return c.weight
}

func (cw Cow) CalculateWeightFeed() int {
	const feedPerKilogram = 25
	return feedPerKilogram * cw.weight
}

func (cw Cow) GetName() string {
	const animalName = "Cow"
	return animalName
}

func (cw Cow) GetWeight() int {
	return cw.weight
}

func main() {

	arr := []WeightFeedCalculater{
		Cow{
			weight: 100,
		},
		Dog{
			weight: 24,
		},
		Cat{
			weight: 10,
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
