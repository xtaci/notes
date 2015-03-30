##!/usr/bin/perl
$maximum = &max(3, 5, 10, 4, 6);
sub max {
	my($max_so_far) = shift @_; # the first one is the largest yet seen
	foreach (@_) { # look at the remaining arguments
		if ($_ > $max_so_far) { # could this one be bigger yet?
			$max_so_far = $_;
		}
	}
	$max_so_far;
}
