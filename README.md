Trains problem
--------------------------

### Usage

#### Using docker-compose
1. `docker-compose -f trains-problem.yml up --build`

#### Using docker by itself
1. `docker build -t trains-problem .`
2. `docker run -it --rm --name trains-problem trains-problem ./bin/trains-problem` and run the following commands from within the container

#### Directly on the host
1. `make`
2. `make run`

-------------

### Structure
- `digraph.go`: Logic behind the problems being solved
- `digraph_test.go`: Unit tests for core logic
- `main.go`: Consuming digraph functions to solve the ten problems

-------------

### Challenges faced
- Inability to use `.next`/`.prev` definitions caused problems when recursively traversing due to the restrictions of the chosen data structure
- Performing pathfinding optimizations on complex datastructure