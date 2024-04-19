## Solution

The problem is solved using simple math. The difference between the min and max
bid is calculated. This value, divided by the increment and rounded down gives us the
number of times an increment can happen. This value is then used to calculate the
`ActualMax`. Finally, all of these `ActualMax` values are compared against each other
to come up with the winning bid.

## Assumptions

1. It is assumed that a bid is tied to a user, but that information doesn't
belong in the `Bid` struct.

2. There is probably some time management for an auction, but that is not being 
considered for this problem. The assumption here is that we are doing the calculation
of the winning bid once this time has run out.

3. All the calculations are being done in cents. This is to avoid any rounding errors
and to do all the calculations using integers. Conversion from dollars to cents are
handled elsewhere.

4. It is assumed that there will always be at least 1 bid. Otherwise, 
the system will crash. In addition, the maximum number of bids is limited by the
available memory in the computer.