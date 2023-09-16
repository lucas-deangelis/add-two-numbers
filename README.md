# go-testing-quick

Trying to use `testing/quick` on the "Add Two Numbers" Leetcode problem. One issue is that to check the result of `addTwoNumbers`, you can't use `addTwoNumbers` too, because then you're not testing anything. To avoid this issue, I have two functions on `ListNode`, `FromInt` and `ToInt`, and compare the result of `addTwoNumbers` to converting the two `ListNode` to int, adding them, and converting them back to a `ListNode`. This assumes that those functions work as expected.
