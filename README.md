# Fractional Knapsack

> Striver SDE Sheet 046 · Day 8 Greedy — sort by value density and pour; the only split happens at the final item. Go, stdlib only, tested.

## Problem

A knapsack holds a given weight; each item has value and weight, and items may be taken fractionally. Maximize the carried value. (The divisibility is what makes greed *provably* optimal here — the indivisible 0/1 variant on Day 25 needs DP.)

## Approach

Take items in descending value-per-weight. Whole items go in while they fit; the first that doesn't is split to exactly fill the remaining capacity, and everything after it is ignored. The exchange argument: any plan carrying lower-density weight while higher-density material is left over improves by swapping, so the density ordering is optimal. One implementation nicety: the sort compares `v1·w2 > v2·w1` (cross-multiplication) instead of dividing — no float ratios, no precision traps in the ordering itself.

**Complexity:** O(n log n) time, O(n) space for the sorted copy.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
