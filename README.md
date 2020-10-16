# simple-host-lookup
Simple host lookup tool written in Go

## Usage examples:

To lookup host ns :
```bash
  go run ./cmd/cli/cli.go ns [HOSTNAME]
  ```

To lookup host ip :
```bash
  go run ./cmd/cli/cli.go ip [HOSTNAME]
  ```

To lookup host cname :
```bash
  go run ./cmd/cli/cli.go cname [HOSTNAME]
  ```

To lookup host mx :
```bash
  go run ./cmd/cli/cli.go mx [HOSTNAME]
  ```

To full lookup host :
```bash
  go run ./cmd/cli/cli.go all [HOSTNAME]
  ```

To install :
```bash
  go install
  ```



