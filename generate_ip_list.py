import random


def generate_random_octet(octets=4):
    ip_range = tuple(range(1, 256))
    for _ in range(octets):
        yield str(random.choice(ip_range))


def main():
    # Keep Asking until they give you a number!!!
    # Been Benchmarking with 500,000
    prompt = False
    if prompt:

        while True:
            try:
                total = int(input("How Many Unique IP's do you need generated?> "))
            except ValueError:
                print("Not a Number!")
            else:
                break
    else:
        total = 500000

    print("Generating Unique IP's")
    collection = set()
    while len(collection) < total:
        # Generate an octet, and add it to a set. Collection Length will only increase if it's not a duplicate value.
        octets = f"{'.'.join(generate_random_octet())}\n"
        collection.add(octets)

    with open("python_unique_ips.txt", 'w') as ip_file:
        ip_file.writelines(collection)

    print("Done.")


if __name__ == '__main__':
    main()
