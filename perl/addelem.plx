#!/usr/bin/perl
# addelem.plx
use warnings;
use strict;

my @array1 = (1, 2, 3);
my @array2;
@array2 = (@array1, 4, 5, 6);

print "@array2\n";

@array2 = (3, 5, 7, 9);
@array2 = (1, @array2, 11);

print "@array2\n";
