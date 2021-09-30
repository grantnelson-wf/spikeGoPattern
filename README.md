# spikeGoPattern

This is a test of different Go implementations for creating a complex facade pattern with testing tools in Go.

## Problem

- We have some large Go packages used for complex data structures and management around those structures.
- These packages tend to expose internals making it difficult to do maintain, debug, and update.
- The code also has a lot of repeat testing tools in each package and consumer.
- The few places we have a testing tools exposed publicly, they can not be used for testing the package
  they are designed to test without being inside that package. So we either have a circular dependency,
  a huge package with lots in it, or a tool for self-testing and another for the consumers to use when testing.
- We would like our unit-tests to be in the same package as the code being tested since
  coverage works best for us if it is per package. Having the unit-test outside of the package would
  require the tests to only run on public methods which limits self-testing.
- Whatever pattern is created it needs to be easily to understand and work with and easy to consume.
- It would be nice if the implementation could have sub-packages in it to isolate complexity needed
  for implementing the facade. This may include the facade being broken up into smaller interfaces.

TL;DR: We need a pattern which helps hide internals, is easy to use, and exposes public testing tools
       which are also used when self-testing this package.

## Pattern Trials

- [Trial 1](./trial1/): A triangle of dependencies with a header package (like C++ headers but with interfaces),
                        shared testing tools, and an implementations package.
