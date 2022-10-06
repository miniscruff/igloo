# Igloo

[![PkgGoDev](https://pkg.go.dev/badge/github.com/miniscruff/igloo)](https://pkg.go.dev/github.com/miniscruff/igloo)
[![codecov](https://codecov.io/gh/miniscruff/igloo/branch/main/graph/badge.svg?token=1tn4p0EOAC)](https://codecov.io/gh/miniscruff/igloo/)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/miniscruff/igloo/unit%20test%20and%20coverage)](https://github.com/miniscruff/igloo/actions?query=workflow%3A"unit+test+and+coverage")

Extension framework to [ebiten](https://github.com/hajimehoshi/ebiten)

**Very much in progress and not production ready**

## Examples

Coming soon, the go package docs is probably the easiest to read for now.


## Tech notes

SceneTree -> Tree of all visual objects
XVisual -> visual representation of something
	-> sprite
	-> label
	-> nine slice ( just 9 anchored child visuals )
	-> animated sprite (???)
	-> rich text (???)

If you want to have like a follow camera just modify the root view transform

Scene tree is auto generated from the editor:
1. Load all content
1. In dispose unload all content
1. For each object build out the parent/child relationship
1. Each object should be of type XVisual

Example:
Space escape scene:
```json
{
    "metadata": {
        "name": "Game", // this determines the type names
        "contentPath": "content/"
    },
    "assets": [{
        "name": "door",
        "type": "image",
        "path": "door_wall.png",
    }, {
        "name": "whatever",
        "type": "opentype",
        "path": "file.ttf",
    }],
    "content": [{
        "name": "Door",
        "type": "sprite",
        "sprite": {
            "image": "door", // ref name above
            "color": "white",
        },
	}, {
		"name": "SliderBackground",
		"type": "SlicedSprite",
		"slicedSprite": {
			"image": "slider_background",
			"borders": {
				"left": 5,
				"right": 5,
				"top": 5,
				"bottom": 5,
			},
			"color": "white",
		}
    }, {
        "name": "Regular20",
        "type": "font",
        "font": {
            "opentype": "whatever",
            "size": 20,
            "dpi": 72,
        },
    }],
    "visuals": [{
        "name": "world",
        "visual": { // optional
        },
        "transform": { // transform stuff
            "x": 0,
            "y": 0,
        }
        "children": [{
            "name": "door",
            "type": "sprite",
            "sprite": "door",
            "transform": {
                "X": 0, // etc
                "Y": 0,
            },
            "children": [/* repeat if needed */],
        }],
    }, {
        // same as world
        "name": "ui",
    }]
}
```
