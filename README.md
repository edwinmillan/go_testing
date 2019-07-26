# Go Testing

This Project serves as a test bed for me to explore Go.

I am using it as a way to get familiar with Go using my experience in Python to help drive progress.

## Baseline Tests

### Requirements:
- 500,000 Unique IP Addresses must be generated randomly.
- Random IP Addresses should be written to a file.

### Initial Baseline

Go 1.12.5:
`go run generate_ip_list.go  7.02s user 0.32s system 191% cpu 3.830 total`

Python 3.7:
`python generate_ip_list.py  3.64s user 0.11s system 98% cpu 3.798 total`


### Refactored Baseline
Go 1.12.5:
`go run generate_ip_list.go  0.52s user 0.25s system 114% cpu 0.673 total`
