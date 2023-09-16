# Add two numbers

Some experiments around the "Add Two Numbers" Leetcode problem.

## `testing/quick`

One issue is that to check the result of `addTwoNumbers`, you can't use `addTwoNumbers` too, because then you're not testing anything. To avoid this issue, I have two functions on `ListNode`, `FromInt` and `ToInt`, and compare the result of `addTwoNumbers` to converting the two `ListNode` to int, adding them, and converting them back to a `ListNode`. This assumes that those functions work as expected.

To be able to use `quick.Check()` you need to implement the `Generator` interface.

## Benchmark

I benchmarked different solutions. Mine, one suggested by Codeium with the "make this code faster and more efficient" refactor, one by asking GPT-4 to improve the solution, and one from Leetcode.
