#!/usr/bin/perl
#sort3.plx
use warnings;
use strict;

my @unsorted = (1, 2, 11, 24, 3, 36, 40, 4);

my @string = sort { $a cmp $b } @unsorted;
print "String sort: @string\n";

my @number = sort { $a <=> $b } @unsorted;
print "Numeric sort: @number\n";
