#!/usr/bin/perl
#Add up some odd numbers
$n = 1;
while ($n < 10) {
	$sum += $n;
	$n += 2; # On to the next odd number
}
print "The total was $sum.\n";
