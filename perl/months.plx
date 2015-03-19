#!/usr/bin/perl
# months.plx
use warnings;
use strict;
my $month = 3;

print qw(
	January February March
	April May June
	July August September
	October November December
)[$month];
