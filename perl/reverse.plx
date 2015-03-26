#!/usr/bin/perl
use 5.010;
@fred = 6..10;  
say @fred;
@barney = reverse(@fred); # gets 10, 9, 8, 7, 6
say @barney;
@wilma = reverse 6..10; # gets the same thing, without the other array
say @wilma;
@fred = reverse @fred; # puts the result back into the original array
say @fred;
