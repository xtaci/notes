#!/usr/bin/perl

sub sum_of_fred_and_barney {
	print "Hey, you called the sum_of_fred_and_barney subroutine!\n";
	$fred + $barney; # That's the return value
}

$fred = 3;
$barney = 4;
$wilma = &sum_of_fred_and_barney; # $wilma gets 7
print "\$wilma is $wilma.\n";
$betty = 3 * &sum_of_fred_and_barney; # $betty gets 21
print "\$betty is $betty.\n";
