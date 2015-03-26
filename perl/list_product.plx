#!/usr/bin/perl
@backwards = reverse qw/ yabba dabba doo /;
# gives doo, dabba, yabba
$backwards = reverse qw/ yabba dabba doo /;
# # gives oodabbadabbay
#
@fred = 6 * 7; # gets the one-element list (42)
@barney = "hello" . ' ' . "world";
@wilma = undef; # OOPS! Gets the one-element list (undef)
# which is not the same as this:
@betty = ( ); # A correct way to empty an array
print "@fred @barney @wilma @betty";
