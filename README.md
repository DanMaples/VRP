# VRP

## Notes from the author
This code sample was written to solve a Vehicle Routing Problem(VRP).
This code was written and tested on macOS 12.7.6.
No AI tools were used.
Google's OR-Tools were not used.
No external sources were referenced or used.
All functions have as much unit test coverage as reasonably possible.

## To Run
From the top level directory run:
> go run main.go {path_to_problem}

## To Run The Unit Tests
From the top level directory run:
> go test ./...

## To Run Inside Of evaluateShared.py
From the top level directory run:
>  python3 evaluateShared.py --cmd "go run main.go" --problemDir problems