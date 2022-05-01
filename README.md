# Igloo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/miniscruff/igloo)](https://pkg.go.dev/github.com/miniscruff/igloo)
[![codecov](https://codecov.io/gh/miniscruff/igloo/branch/main/graph/badge.svg?token=1tn4p0EOAC)](https://codecov.io/gh/miniscruff/igloo/)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/miniscruff/igloo/unit%20test%20and%20coverage)](https://github.com/miniscruff/igloo/actions?query=workflow%3A"unit+test+and+coverage")

Extension framework to [ebiten](https://github.com/hajimehoshi/ebiten)

**Very much in progress and not production ready**

```go
func main() {
    // same regardless
    exScene := NewExampleScene()

    // static game wins
    igloo.Push(exScene) // make sure to push our starting scene
    igloo.Pop() // calls dispose and sets next active scene

    // error catch the same either way
	if err := igloo.Run(); err != nil {
		fmt.Printf("run game: %s", err.Error())
	}
}

type ExampleScene struct {
    img *ebiten.Image
}

func NewExampleScene() *ExampleScene {
    return &ExampleScene{
    }
}

// FS should not be hard coded
func (s *ExampleScene) LoadContent() {
    s.img = content.LoadImage(fs, "blah.png")
    s.img = igloo.LoadImage("blah.png") // magically has the content fs
}

func (s *ExampleScene) Unload() {
    s.img.Dispose()
}

func (s *ExampleScene) Update() {
    // can load another scene
    igloo.Push(newScene())
    // or pop ourselves off
    igloo.Pop(s)
}

func (s *ExampleScene) Draw(dest *ebiten.Image) {
}

```

## Design Notes
* game: high level game
* scene: stack of scenes ( only the top gets input, for now )
* content: load different types of content
    * ContentCache: tbd way of loading and caching content across scenes
    * image
    * fonts
    * nine-slice data
* mathf: complex math
    * vec2
    * transform
    * tweens
    * easing
    * math for ui
* ui: control and organize elements for an interface
    * window
    * grid
    * horizontal layout
    * vertical layout
    * slider
    * checkbox
    * text input
* graphics: render different visual elements
    * sprite
    * nine-slice
    * label

## Notes
### Animations
1. Repeat animations, not repeat, ping-pong
1. Translate a little bit
1. Transition from one animation to another, based on some condition
1. Rotation
1. Easing
1. Number of frames
1. Thing with multiple animations [sprites, translates, rotates, scales, etc]
1. Total amount of time for the animation
1. Alter some animation based on game values

### Particles
1. TBD
