#!/usr/bin/perl
my @names = qw/ fred barney betty dino wilma pebbles bamm-bamm /;
my $result = &which_element_is("dino", @names);
sub which_element_is {
	my($what, @array) = @_;
	foreach (0..$#array) { # indices of @array's elements
		if ($what eq $array[$_]) {
			return $_; # return early once found
		}
	}
	-1; # element not found (return is optional here)
}

print $result;
