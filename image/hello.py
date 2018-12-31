#!/usr/bin/env python3
import sys
import time


def main():
    name = input("Name:")

    for i in range(10):
        print("Hello, {}".format(name))
        time.sleep(1)


if __name__ == "__main__":
    sys.exit(main())
