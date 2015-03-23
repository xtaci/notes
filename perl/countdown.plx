#!/usr/bin/perl
# countdown.plx
use warnings;
use strict;

my @count = (1..5);

for (reverse(@count)) {
	print "$_...\n";
	sleep 1;
}

print "BLAST OFF!\n";
