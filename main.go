package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}

	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)

	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Didn't understand")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
		You are in a large chamber,
		deep underground with hunners of gobbos.

		To the south there's a passage
		to the east another passage
		to the west nuhin
		to the north also nuhin
	`}

	darkRoom := storyNode{text: "It is black"}

	darkRoomLit := storyNode{text: "The room is now lit and you can continue north"}

	grue := storyNode{text: "While wlaking about in the dark you are eaten by a grue"}

	trap := storyNode{text: "Oh shit you fell down a hole and have died on a horrible spike"}

	treasure := storyNode{text: "You found all that tasty gold bro!"}

	start.addChoice("n", "Go north", &darkRoom)
	start.addChoice("s", "Go south", &darkRoom)
	start.addChoice("e", "Go east", &trap)

	darkRoom.addChoice("s", "Try to go back south", &grue)
	darkRoom.addChoice("o", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("n", "Go north", &treasure)
	darkRoomLit.addChoice("s", "Go south", &start)

	start.play()

	fmt.Println()
	fmt.Println("The end")

}
