#!/usr/bin/perl
sub list_from_fred_to_barney {
	if ($fred < $barney) {
# Count upwards from $fred to $barney
		$fred..$barney;
	} else {
# Count downwards from $fred to $barney
		reverse $barney..$fred;
	}
}
$fred = 11;
$barney = 6;
@c = &list_from_fred_to_barney; # @c gets (11, 10, 9, 8, 7, 6)
