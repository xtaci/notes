#!/usr/bin/perl
# badprefix.plx
# The prime rule is this: the prefix represents what you want to get, not what you've got.
#
# So @ represents a list of values, and $ represents a single scalar. Hence, when we're
#
# getting a single scalar from an array, we never prefix the variable with @ – that would
#
# mean a list. A single scalar is always prefixed with a $.

use warnings;
use strict;

my @array = (1, 3, 5, 7, 9);
print @array[1];
