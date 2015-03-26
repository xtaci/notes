#!/usr/bin/perl
use 5.010;
@array = qw( pebbles dino fred barney betty );
@removed = splice @array, 2; # remove everything after fred
# @removed is qw(fred barney betty)
# # @array is qw(pebbles dino)
say "removed:@removed array:@array";

@array = qw( pebbles dino fred barney betty );
@removed = splice @array, 1, 2; # remove dino, fred
# @removed is qw(dino fred)
# # @array is qw(pebbles barney betty)
say "removed:@removed array:@array";

@array = qw( pebbles dino fred barney betty );
@removed = splice @array, 1, 2, qw(wilma); # remove dino, fred
# @removed is qw(dino fred)
# # @array is qw(pebbles wilma
# # barney betty)
say "removed:@removed array:@array";

@array = qw( pebbles dino fred barney betty );
@removed = splice @array, 1, 0, qw(wilma); # remove nothing
# @removed is qw()
# # @array is qw(pebbles wilma dino
# # fred barney betty)
say "removed:@removed array:@array";
