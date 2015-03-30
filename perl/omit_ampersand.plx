#!/usr/bin/perl
sub division {
	$_[0] / $_[1]; # Divide first param by second
}
my $quotient = division 355, 113; # Uses &division
print $quotient;
